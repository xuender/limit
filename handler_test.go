package limit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestLimitFuncHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := limit.FuncHandler(1, time.Second, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	start := time.Now()

	go func() {
		handler.ServeHTTP(httptest.NewRecorder(), req)
	}()

	go func() {
		errRec := httptest.NewRecorder()

		time.Sleep(time.Millisecond * 300)
		handler.ServeHTTP(errRec, req)
		assert.NotEqual(t, errRec.Code, http.StatusOK)
	}()

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.GreaterOrEqual(t, time.Since(start), time.Second)
}
