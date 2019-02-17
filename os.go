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

func expand(p string) (string, error) {
	p, err := homedir.Expand(p)
	if err != nil {
		return "", err
	}
	p, err = filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return p, nil
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
	p, err := expand(filepath.Join(fs.basepath, path))
	if err != nil {
		return nil, err
	}
	return os.Open(p)
}
