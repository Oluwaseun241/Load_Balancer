package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/* Load balancer */
func NewLoadBalancer(targetServers []string) (*httputil.ReverseProxy, error) {
  var targets []*url.URL
  for _, ts := range targetServers {
    u, err := url.Parse(ts)
    if err != nil {
      return nil, err
    }
    targets = append(targets, u)
  }

  director := func (req *http.Request) {
    targetURL := targets[req.RemoteAddr[7]%uint8(len(targets))]
    req.URL.Scheme = targetURL.Scheme
    req.URL.Host = targetURL.Host
    req.URL.Path = targetURL.Path + req.URL.Path
  }
  return &httputil.ReverseProxy{Director: director}, nil
}


/* Reverse proxy */
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    proxy.ServeHTTP(w,r)
  }
}

func main() {
  targetServers := []string {
    "http://127.0.0.1:5000",
    "http://127.0.0.1:5001",
    "http://127.0.0.1:5002",
  }

  proxy, err := NewLoadBalancer(targetServers)
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", ProxyRequestHandler(proxy))
  log.Fatal(http.ListenAndServe(":8080", nil))
}
