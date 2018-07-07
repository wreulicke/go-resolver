package resolver

import (
	"errors"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type ssmResolver struct {
	ssmiface.SSMAPI
	prefix         string
	withDecription *bool
}

type dummyReadCloser struct {
	io.Reader
}

// NewSSMResolver provides resolver from SSM
func NewSSMResolver(ssmAPI ssmiface.SSMAPI) Resolver {
	return &ssmResolver{
		SSMAPI: ssmAPI,
		prefix: "",
	}
}

// NewSSMResolverWithPrefix provides resolver from SSM with prefix
func NewSSMResolverWithPrefix(ssmAPI ssmiface.SSMAPI, prefix string) Resolver {
	return &ssmResolver{
		SSMAPI: ssmAPI,
		prefix: prefix,
	}
}

func (resolver *ssmResolver) Resolve(path string) (io.ReadCloser, error) {
	r := &ssm.GetParameterInput{
		Name: aws.String(resolver.prefix + path),
	}
	res, err := resolver.GetParameter(r)
	if err != nil {
		return nil, err
	}
	if res.Parameter == nil {
		return nil, errors.New("parameter " + *r.Name + "is not found")
	} else if res.Parameter.Name == nil {
		return nil, errors.New("parameter " + *r.Name + "is not found")
	}
	return newDummyReadCloser(*res.Parameter.Value), nil
}

func newDummyReadCloser(v string) io.ReadCloser {
	return &dummyReadCloser{
		Reader: strings.NewReader(v),
	}
}

func (*dummyReadCloser) Close() error {
	return nil
}
