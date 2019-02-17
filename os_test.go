package resolver

import (
	"os"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func unwrap(r Resolver) *osFsResolver {
	t, ok := r.(*osFsResolver)
	if !ok {
		return nil
	}
	return t
}

type d struct {
	in  string
	out string
}

var testdata []d

func init() {
	home, _ := homedir.Dir()
	wd, _ := os.Getwd()
	testdata = []d{
		{"~/", home},
		{"~/base", home + "/base"},
		{"./test", wd + "/test"},
		{"test", wd + "/test"},
		{"./test/../test", wd + "/test"},
	}
}

const format = `not expected. 
in: %s
actual: %s
out: %s
`

func TestNewOsFsResolver(t *testing.T) {
	for _, tt := range testdata {
		t.Run(tt.in, func(t *testing.T) {
			resolver, _ := NewOsFsResolver(tt.in)
			r := unwrap(resolver)
			if r.basepath != tt.out {
				t.Errorf(format, tt.in, r.basepath, tt.out)
			}
		})
	}
}

