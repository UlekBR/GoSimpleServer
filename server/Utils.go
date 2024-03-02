package server

import (
	"net/http"
	"net/url"
)

type ResponseStringWriterFunc func(string)
type ResponseBytesWriterFunc func([]byte)
type ResponseJsonWriterFunc func(interface{})
type ResponseGetPathValueFunc func(string) string

type ResponseManager struct {
	stringWriteFunc ResponseStringWriterFunc
	bytesWriteFunc  ResponseBytesWriterFunc
	jsonWriteFunc   ResponseJsonWriterFunc
	URL             *url.URL
	Method          string
	Body            string
	Header          http.Header
	getPathValue    ResponseGetPathValueFunc
	Host            string
	Port            string
	RequestURI      string
	RemoteAddr      string
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
func (rw *ResponseManager) GetPathValue(name string) string {
	return rw.getPathValue(name)
}
