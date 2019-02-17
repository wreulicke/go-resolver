package resolver

import (
	"io"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type osFsResolver struct {
	basepath string
}

// NewOsFsResolver provides resolver from Os file system
func NewOsFsResolver(basepath string) (Resolver, error) {
	p, err := homedir.Expand(basepath)
	if err != nil {
		return nil, err
	}
	p, err = filepath.Abs(p)
	if err != nil {
		return nil, err
	}
	return &osFsResolver{
		basepath: p,
	}, nil
}

func (fs *osFsResolver) Resolve(path string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(fs.basepath, path))
}
