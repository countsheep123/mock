package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Port      int64       `json:"port"`
	Endpoints []*Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Endpoint string    `json:"endpoint"`
	Methods  []*Method `json:"methods"`
}

type Method struct {
	Method   string      `json:"method"`
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, endpoint := range handler.Endpoints {
		for _, method := range endpoint.Methods {
			if endpoint.Endpoint == r.URL.Path &&
				r.Method == method.Method {

				jsonBytes, err := json.Marshal(method.Response)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Printf("===> [%d] %6s %s\n", http.StatusInternalServerError, r.Method, r.URL)
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(method.Status)
				w.Write(jsonBytes)
				fmt.Printf("===> [%d] %6s %s\n", http.StatusOK, r.Method, r.URL)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Printf("===> [%d] %6s %s\n", http.StatusNotFound, r.Method, r.URL)
}
