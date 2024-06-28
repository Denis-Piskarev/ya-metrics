package compression

import (
	"compress/gzip"
	"net/http"
)

type compressWriter struct {
	w  http.ResponseWriter
	gz *gzip.Writer
}

func NewCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		gz: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Write(b []byte) (int, error) {
	return c.gz.Write(b)
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) WriteHeader(statusCode int) {
	c.w.WriteHeader(statusCode)
}

func (c *compressWriter) Close() error {
	return c.gz.Close()
}
