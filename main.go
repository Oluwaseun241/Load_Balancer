package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

/* Load balancer */
func NewLoadBalancer(targetServers []string) (*httputil.ReverseProxy, error) {

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
