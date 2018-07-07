package resolver

import "io"

// Resolver abstracts very tin file system
type Resolver interface {
	Resolve(path string) (io.ReadCloser, error)
}
