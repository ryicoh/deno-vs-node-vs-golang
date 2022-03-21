package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const port = 8080

type (
	request struct {
		Hello string `json:"hello"`
	}

	response struct {
		Data *request `json:"data"`
	}
)

func main() {

	fmt.Printf("HTTP server listening on http://localhost:%d\n", port)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		req := new(request)
		if err := json.Unmarshal(body, req); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		resData, err := json.Marshal(response{Data: req})
		if err := json.Unmarshal(body, req); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Write(resData)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
