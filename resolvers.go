// experimental
package resolver

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type Resolvers struct {
	S3     s3iface.S3API
	SSM    ssmiface.SSMAPI
	client *http.Client
}

func (f *Resolvers) Resolve(u string) (io.ReadCloser, error) {
	if strings.HasPrefix(u, "file://") {
		t := strings.TrimPrefix(u, "file://")
		r, err := NewOsFsResolver("")
		if err != nil {
			return nil, err
		}
		return r.Resolve(t)
	} else if strings.HasPrefix(u, "s3://") {
		t := strings.TrimPrefix(u, "s3://")
		s := strings.SplitN(t, "/", 2)
		if len(s) < 2 {
			return nil, fmt.Errorf("Cannot create s3 resolver. Expected path is s3://bucket/your/object/key")
		}
		r := NewS3Resolver(f.S3, s[0])
		return r.Resolve(s[1])
	} else if strings.HasPrefix(u, "ssm://") {
		t := strings.TrimPrefix(u, "ssm://")
		return NewSSMResolver(f.SSM).Resolve(t)
	}
	return nil, errors.New("Not found resolver")
}

