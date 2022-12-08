package rdb_test

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
	"github.com/xuender/limit/rdb"
	rdb_mock "github.com/xuender/limit/rdb/mock"
)

func TestNewDistributed(t *testing.T) {
	t.Parallel()

	limiter := rdb.NewDistributed(nil, "key", -1, time.Second)
	assert.NotNil(t, limiter.Wait())
}

func TestRdb_Wait(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limiter := rdb.NewDistributed(client, "key", 1000, time.Second)
	cmd1 := new(redis.IntCmd)
	cmd2 := new(redis.IntCmd)

	cmd1.SetVal(1)
	cmd2.SetVal(2)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd1).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd2).MinTimes(1).MaxTimes(20)

	assert.Nil(t, limiter.Wait())
	assert.Nil(t, limiter.Wait())
	time.Sleep(time.Millisecond * 4)
	assert.Nil(t, limiter.Wait())
}

func TestRdb_Wait_Error(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limiter := rdb.NewDistributed(client, "key", 1000, time.Second)
	cmd := new(redis.IntCmd)

	cmd.SetVal(3)
	cmd.SetErr(limit.ErrKey)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd).MinTimes(0).MaxTimes(3)

	assert.NotNil(t, limiter.Wait())
}

func TestRdb_Wait_Timeout(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limiter := rdb.NewDistributed(client, "key", 1, time.Millisecond)
	cmd := new(redis.IntCmd)
	cmd2 := new(redis.IntCmd)

	cmd.SetVal(1)
	cmd2.SetVal(2)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd2).MinTimes(1).MaxTimes(3)
	client.EXPECT().Decr(gomock.Any(), "key").Return(cmd).MinTimes(0).MaxTimes(3)

	assert.Nil(t, limiter.Wait())
	assert.NotNil(t, limiter.Wait())
}

func TestRdb_Wait_Timeout_error(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limiter := rdb.NewDistributed(client, "key", 1, time.Millisecond)
	cmd := new(redis.IntCmd)
	cmd2 := new(redis.IntCmd)
	cmd3 := new(redis.IntCmd)

	cmd.SetVal(1)
	cmd2.SetVal(2)
	cmd3.SetVal(3)
	cmd3.SetErr(limit.ErrKey)

	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "key").Return(cmd2).MinTimes(1).MaxTimes(3)
	client.EXPECT().Decr(gomock.Any(), "key").Return(cmd3).MinTimes(0).MaxTimes(3)

	assert.Nil(t, limiter.Wait())
	assert.NotNil(t, limiter.Wait())
}
