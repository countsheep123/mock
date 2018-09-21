package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
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
			e, err := normalize(endpoint.Endpoint)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("===> [%d] %6s %s\n", http.StatusInternalServerError, r.Method, r.URL)
				return
			}
			reqEndpoint, err := normalize(r.URL.String())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("===> [%d] %6s %s\n", http.StatusInternalServerError, r.Method, r.URL)
				return
			}

			if e == reqEndpoint &&
				r.Method == method.Method {

				jsonBytes, err := json.Marshal(method.Response)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Printf("===> [%d] %6s %s\n", http.StatusInternalServerError, r.Method, r.URL)
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set("Access-Control-Allow-Origin", "*")
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

func normalize(endpoint string) (string, error) {
	oldURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	q := oldURL.Query()

	newQuery := url.Values{}
	for _, k := range keys(q) {
		vs := q[k]
		sort.Strings(vs)
		for _, v := range vs {
			newQuery.Add(k, v)
		}
	}

	newURL, err := url.Parse(oldURL.Path)
	if err != nil {
		return "", err
	}
	newURL.RawQuery = newQuery.Encode()
	return newURL.String(), nil
}

func keys(v url.Values) []string {
	keys := []string{}
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}
