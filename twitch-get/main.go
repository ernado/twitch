package main

import (
	"flag"
	"net"
	"net/http"
	"time"
	"log"

	"github.com/cydev/twitch/downloader"
)

const (
	defaultHTTPTimeout        = 30 * time.Second
	defaultRequestTimeout     = 30 * time.Second
	defaultKeepAliveInterval  = 600 * time.Second
	defaultHTTPHeadersTimeout = defaultRequestTimeout
)

func getDefaultHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: defaultRequestTimeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   defaultHTTPTimeout,
				KeepAlive: defaultKeepAliveInterval,
			}).Dial,
			TLSHandshakeTimeout:   defaultHTTPTimeout,
			ResponseHeaderTimeout: defaultHTTPHeadersTimeout,
		},
	}
	return client
}

func main() {
	flag.Parse()
	client := getDefaultHTTPClient()
	if flag.NArg() < 1 {
		log.Fatalln("no stream name specified")
	}
	streamName := flag.Arg(0)
	log.Println("waiting for stream", streamName)
	d := downloader.New(streamName, client)
	d.Start()
}
