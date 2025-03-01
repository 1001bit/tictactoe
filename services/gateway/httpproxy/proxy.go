package httpproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	url *url.URL
}

func New(addr string) (*Proxy, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &Proxy{url: url}, nil
}

func (p *Proxy) ProxyHandler(stripPrefix string) http.Handler {
	return http.StripPrefix(stripPrefix, httputil.NewSingleHostReverseProxy(p.url))
}
