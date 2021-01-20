package main

import (
	"flag"
	"os"
)

var bucketURL = "http://storage.googleapis.com/parenting-app-ui-master1"

func main() {
	targetURL := flag.String("target", bucketURL, "Target URL")
	port := flag.Int("port", 3000, "Port to run proxy on")
	if len(os.Args) < 2 {
		flag.Usage()
	}
	flag.Parse()
	proxy := AngularProxy{
		ProxyToURL: *targetURL,
		Port:       *port,
	}
	proxy.start()
}
