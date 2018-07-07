package resolver

import (
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type s3Resolver struct {
	s3iface.S3API
	bucket  string
	basekey string
}

// NewS3Resolver provides resolver from s3
func NewS3Resolver(s3API s3iface.S3API, bucket string) Resolver {
	return &s3Resolver{
		S3API:   s3API,
		bucket:  bucket,
		basekey: "",
	}
}

// NewS3ResolverWithBaseKey provides resolver from s3 with basekey
func NewS3ResolverWithBaseKey(s3API s3iface.S3API, bucket string, basekey string) Resolver {
	if strings.HasPrefix(basekey, "/") {
		basekey = basekey[1:]
	}
	if basekey != "" && !strings.HasSuffix(basekey, "/") {
		basekey = basekey + "/"
	}
	return &s3Resolver{
		S3API:   s3API,
		bucket:  bucket,
		basekey: basekey,
	}
}

func (resolver *s3Resolver) Resolve(path string) (io.ReadCloser, error) {
	base := resolver.basekey
	if strings.HasPrefix(path, "/") {
		path = base + path[1:]
	} else {
		path = base + path
	}

	req := &s3.GetObjectInput{
		Bucket: aws.String(resolver.bucket),
		Key:    aws.String(path),
	}
	res, err := resolver.GetObject(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
