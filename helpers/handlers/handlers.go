package handlers

import (
	"io"
	"app/helpers/logger"
	"net/http"
	"net/url"
	"fmt"
)

// HttpProxyHandler implemets http proxy handler
func HttpProxyHandler(route, proxyUrl string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		req.URL, _ = url.Parse(proxyUrl)
		resp, err := http.DefaultTransport.RoundTrip(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		logger.Debug(route + " => " + proxyUrl)

		copyHeader(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

// HealthCheckHandler implemets /healthCheck
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\" : \"ok\"}")
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
