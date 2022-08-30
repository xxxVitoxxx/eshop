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
	"github.com/xxxVitoxxx/eshop/pkg/storage/fake"
)

func TestPutConditionByStoreName(t *testing.T) {
	r := gin.Default()
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)
	handler := NewNumberPlateHandler(s)
	handler.Route(r)

	t.Run("PutConditionByStoreName success", func(t *testing.T) {
		condition, _ := json.Marshal(numberplate.PutCondition{
			HowMany: 10,
			HowLong: 60,
			Remind:  10,
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodPut,
			"/eshop_api/number_plate/condition/eshop",
			bytes.NewBuffer(condition),
		)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestCreateCondition(t *testing.T) {
	r := gin.Default()
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)
	handler := NewNumberPlateHandler(s)
	handler.Route(r)

	t.Run("CreateCondition success", func(t *testing.T) {
		condition, _ := json.Marshal(numberplate.Condition{
			StoreName: "eshop",
			HowMany:   20,
			HowLong:   60,
			Remind:    5,
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodPost,
			"/eshop_api/number_plate/condition",
			bytes.NewBuffer(condition),
		)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("CreateCondition missing required key", func(t *testing.T) {
		condition, _ := json.Marshal(numberplate.Condition{
			StoreName: "two circle",
			HowLong:   120,
			Remind:    5,
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodPost,
			"/eshop_api/number_plate/condition",
			bytes.NewBuffer(condition),
		)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(
			t,
			`{"message":"Key: 'Condition.HowMany' Error:Field validation for 'HowMany' failed on the 'required' tag"}`,
			w.Body.String(),
		)
	})
}

func TestDeleteConditionByStoreName(t *testing.T) {
	r := gin.Default()
	fakeRepo := fake.NewConditionRepo()
	handler := NewNumberPlateHandler(numberplate.NewService(fakeRepo))
	handler.Route(r)

	t.Run("DeleteConditionByStoreName success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			http.MethodDelete,
			"/eshop_api/number_plate/condition/eshop",
			nil,
		)

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
