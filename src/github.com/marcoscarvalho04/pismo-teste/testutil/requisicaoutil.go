package testutil

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
)

type handler func(http.ResponseWriter, *http.Request)

func FazerRequisicaoParaURL(method string, url string, body string, handler handler, params map[string]string) (statusCode int, responseString string) {
	var req *http.Request

	if len(body) == 0 {
		req = httptest.NewRequest(method, url, nil)
	} else {
		req = httptest.NewRequest(method, url, strings.NewReader(body))
	}
	if params != nil {
		req = mux.SetURLVars(req, params)
	}

	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	response, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		statusCode = 0
		responseString = "Erro"
		return statusCode, responseString
	}
	responseString = string(response)
	statusCode = resp.StatusCode
	return statusCode, responseString
}
