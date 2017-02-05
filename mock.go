package main

import (
	"fmt"
	"net/http"
	"sync"
)

func serve(handlers []*handler) {
	wg := &sync.WaitGroup{}
	fmt.Println("---------- Endpoints ----------")
	for i := range handlers {
		handler := handlers[i]
		addr := fmt.Sprintf(":%d", handler.Port)
		for j := range handler.Endpoints {
			endpoint := handler.Endpoints[j]
			for k := range endpoint.Methods {
				method := endpoint.Methods[k]
				fmt.Printf("%6s %8s %s\n", addr, fmt.Sprintf("[%s]", method.Method), endpoint.Endpoint)
			}
		}
		wg.Add(1)
		go func() {
			panic(http.ListenAndServe(addr, handler))
			wg.Done()
		}()
	}
	fmt.Println("-------------------------------")
	wg.Wait()
}
