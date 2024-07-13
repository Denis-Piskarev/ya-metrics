// Package for operate with compressing data

package compress

import (
	"bytes"
	"compress/gzip"
)

// GetGZip returns compressed data
func GetGZip(metrics []byte) (*bytes.Buffer, error) {
	// Creating a new gzip writer
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	defer gz.Close()

	// Writing the body to the gzip writer
	if _, err := gz.Write(metrics); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	return &buf, nil
}
