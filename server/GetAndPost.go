package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (h *Handler) GetAndPost(path string, fn func(*ResponseManager)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}

			parts := strings.Split(r.Host, ":")
			host, port := parts[0], parts[1]

			fn(&ResponseManager{

				stringWriteFunc: func(str string) {
					_, err := w.Write([]byte(str))
					if err != nil {
						return
					}
				},
				bytesWriteFunc: func(bytes []byte) {
					_, err := w.Write(bytes)
					if err != nil {
						return
					}
				},
				jsonWriteFunc: func(jsonStruct interface{}) {
					jsonData, err := json.Marshal(jsonStruct)
					if err != nil {
						return
					}
					w.Header().Set("Content-Type", "application/json")

					_, err = w.Write(jsonData)
					if err != nil {
						return
					}
				},
				getPathValue: func(str string) string {
					return r.PathValue(str)
				},
				URL:        r.URL,
				Method:     r.Method,
				Body:       string(body),
				Header:     w.Header(),
				Host:       host,
				Port:       port,
				RequestURI: r.RequestURI,
				RemoteAddr: r.RemoteAddr,
			})

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not Allowed"))
		}

	})
}
