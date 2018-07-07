package resolver

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type osFsResolver struct {
	basepath string
}

// NewOsFsResolver provides resolver from Os file system
func NewOsFsResolver(basepath string) (Resolver, error) {
	if !strings.HasPrefix(basepath, "/") {
		ex, err := os.Executable()
		if err != nil {
			return nil, err
		}
		exPath := filepath.Dir(ex)
		basepath = filepath.Join(exPath, basepath)
	}
	return &osFsResolver{
		basepath: basepath,
	}, nil
}

func (fs *osFsResolver) Resolve(path string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(fs.basepath, path))
}
