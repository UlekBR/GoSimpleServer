package server

import (
	"net/http"
	"net/url"
)

type ResponseStringWriterFunc func(string)
type ResponseBytesWriterFunc func([]byte)
type ResponseJsonWriterFunc func(interface{})

type ResponseManager struct {
	stringWriteFunc ResponseStringWriterFunc
	bytesWriteFunc  ResponseBytesWriterFunc
	jsonWriteFunc   ResponseJsonWriterFunc
	URl             *url.URL
	Method          string
	Body            string
	Header          http.Header
}

func (rw *ResponseManager) WriteString(str string) {
	rw.stringWriteFunc(str)
}
func (rw *ResponseManager) WriteBytes(bytes []byte) {
	rw.bytesWriteFunc(bytes)
}
func (rw *ResponseManager) WriteJson(json interface{}) {
	rw.jsonWriteFunc(json)
}
