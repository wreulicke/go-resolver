package resolver

import (
	"io"
	"net/http"
)

type httpFsResolver struct {
	http.FileSystem
}

// NewHTTPFsResolver provides resolver from net/http file system
func NewHTTPFsResolver(httpFs http.FileSystem) Resolver {
	return &httpFsResolver{
		FileSystem: httpFs,
	}
}

func (fs *httpFsResolver) Resolve(path string) (io.ReadCloser, error) {
	return fs.FileSystem.Open(path)
}
