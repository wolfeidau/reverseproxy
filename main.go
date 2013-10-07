// Copyright 2013 Mark Wolfe. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE-MIT file.

// Simple HTTP reverse proxy which logs requests via Syslog

package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	proxyUrl   = flag.String("url", "https://api.tempo-db.com", "URL to proxy")
	proxyPort  = flag.String("port", ":8080", "Port to listen for connections")
	syslogHost = flag.String("syslog-host", "localhost", "Host to send UDP syslog messages")
	syslogPort = flag.String("syslog-port", "514", "Port to send UDP syslog messages")
)

func init() {
	flag.Parse()
}

func reverseProxy(target *url.URL, logger *syslog.Writer) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		logger.Info(fmt.Sprintf("%s", req.URL.Path))
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {

	u, err := url.Parse(*proxyUrl)
	if err != nil {
		log.Fatal(err)
	}

	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := syslog.Dial("udp", net.JoinHostPort(*syslogHost, *syslogPort), syslog.LOG_INFO, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started listening %s", *proxyPort)

	reverse_proxy := reverseProxy(u, logger)
	http.Handle("/", reverse_proxy)

	if err = http.ListenAndServe(*proxyPort, nil); err != nil {
		log.Fatal(err)
	}
}
