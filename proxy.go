package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// AngularProxy - Proxies to URL or to index.html page if route not found
type AngularProxy struct {
	ProxyToURL string
	Port       int
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {

	url, _ := url.Parse(target)

	fmt.Println("Target URL:", url)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func (a *AngularProxy) handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := fmt.Sprint(a.ProxyToURL, "/")

	if req.URL.Path == "/" {
		req.URL.Path = "/index.html"
	} else if !strings.Contains(req.URL.Path, ".") {
		headReq, err := http.Head(fmt.Sprint(a.ProxyToURL, req.URL.Path))
		if err != nil {
			fmt.Println("Head request error ", err)
		} else if headReq.StatusCode == 404 {
			req.URL.Path = "/index.html"
		}
	}

	serveReverseProxy(url, res, req)
}

func (a *AngularProxy) getListenAddress() string {
	return fmt.Sprint(":", a.Port)
}

func (a *AngularProxy) start() {
	fmt.Println("Listening on port ", a.Port)
	fmt.Println("Proxying requests to ", a.ProxyToURL)
	http.HandleFunc("/", a.handleRequestAndRedirect)
	if err := http.ListenAndServe(a.getListenAddress(), nil); err != nil {
		fmt.Println("Error starting server ", err)
		panic(err)
	}
}
