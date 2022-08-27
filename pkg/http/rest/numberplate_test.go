package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/fake"
)

func TestPutCondition(t *testing.T) {
	r := gin.Default()
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)
	handler := NewNumberPlateHandler(s)
	handler.Route(r)

	t.Run("PutCondition success", func(t *testing.T) {
		condition, _ := json.Marshal(numberplate.PutCondition{
			HowMany: 10,
			HowLong: 60,
			Remind:  10,
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodPut,
			"/eshop_api/number_plate/condition",
			bytes.NewBuffer(condition),
		)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
