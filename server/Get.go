package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) Get(path string, fn func(*ResponseManager)) {
	http.HandleFunc("GET "+path, func(w http.ResponseWriter, r *http.Request) {

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
			Header:     w.Header(),
			Host:       host,
			Port:       port,
			RequestURI: r.RequestURI,
			RemoteAddr: r.RemoteAddr,
		})

	})
}
