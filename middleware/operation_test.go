package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casualjim/go-swagger"
	"github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/internal/testing/petstore"
	"github.com/stretchr/testify/assert"
)

func TestOperationExecutor(t *testing.T) {
	spec, api := petstore.NewAPI(t)
	api.RegisterOperation("getAllPets", swagger.OperationHandlerFunc(func(params interface{}) (interface{}, error) {
		return []interface{}{
			map[string]interface{}{"id": 1, "name": "a dog"},
		}, nil
	}))

	context := NewContext(spec, api, nil)
	mw := newOperationExecutor(context)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/pets", nil)
	request.Header.Add("Accept", "application/json")
	request.SetBasicAuth("admin", "admin")
	mw.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, `[{"id":1,"name":"a dog"}]`+"\n", recorder.Body.String())

	spec, api = petstore.NewAPI(t)
	api.RegisterOperation("getAllPets", swagger.OperationHandlerFunc(func(params interface{}) (interface{}, error) {
		return nil, errors.New(422, "expected")
	}))

	context = NewContext(spec, api, nil)
	mw = newOperationExecutor(context)

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/pets", nil)
	request.Header.Add("Accept", "application/json")
	request.SetBasicAuth("admin", "admin")
	mw.ServeHTTP(recorder, request)
	assert.Equal(t, 422, recorder.Code)
	assert.Equal(t, `{"code":422,"message":"expected"}`, recorder.Body.String())
}
