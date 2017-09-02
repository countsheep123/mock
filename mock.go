package mock

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func Serve(handlers []*Handler) {
	wg := &sync.WaitGroup{}
	fmt.Fprintln(os.Stdout, "---------- Endpoints ----------")
	for _, handler := range handlers {
		addr := fmt.Sprintf(":%d", handler.Port)
		for _, endpoint := range handler.Endpoints {
			for _, method := range endpoint.Methods {
				fmt.Fprintf(os.Stdout, "%6s %8s %s\n", addr, fmt.Sprintf("[%s]", method.Method), endpoint.Endpoint)
			}
		}
		wg.Add(1)
		go func() {
			panic(http.ListenAndServe(addr, handler))
			wg.Done()
		}()
	}
	fmt.Fprintln(os.Stdout, "-------------------------------")
	wg.Wait()
}
