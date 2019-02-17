package resolver

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestResolve(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/base/path" {
			return
		}
		t.Errorf("Does not contains `Test` header. headers: %v", r.Header)
	}))
	url, _ := url.Parse(ts.URL + "/base")
	r := NewHTTPResolver(url, http.DefaultClient)

	_, err := r.Resolve("/path")
	if err != nil {
		t.Error(err)
	}
}
