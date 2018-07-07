package resolver

import "io"

type Resolver interface {
	Resolve(path string) (io.ReadCloser, error)
}
