package resolver

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type httpResolver struct {
	base   *url.URL
	client *http.Client
}

type ServerError struct {
	*http.Response
}

type ClientError struct {
	*http.Response
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("Failed to connect %s, status;%d", e.Request.URL, e.StatusCode)
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("Failed to connect %s, status;%d", e.Request.URL, e.StatusCode)
}

// NewHTTPFsResolver provides resolver from net/http file system
func NewHTTPResolver(base *url.URL, client *http.Client) Resolver {
	return &httpResolver{
		base:   base,
		client: client,
	}
}

func (r *httpResolver) Resolve(p string) (io.ReadCloser, error) {
	u := *r.base // copy
	u.Path = path.Join(r.base.Path, p)
	resp, err := r.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.Body, nil
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		// not implement redirect.
	} else if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		return nil, &ClientError{
			Response: resp,
		}
	} else if resp.StatusCode >= 500 && resp.StatusCode <= 599 {
		return nil, &ServerError{
			Response: resp,
		}
	}
	return nil, fmt.Errorf("Unknow status: %d, url:%s", resp.StatusCode, resp.Request.URL)
}
