package server

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) Post(path string, fn func(*ResponseManager)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
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
				URl:    r.URL,
				Method: r.Method,
				Body:   string(body),
				Header: w.Header(),
			})
		}
	})
}
