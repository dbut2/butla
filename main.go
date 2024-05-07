package main

import (
	rp "github.com/dbut2/reverse-proxy"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	proxy := rp.New(
		rp.Select("http://192.168.1.226:3000", HostMatches("dbut.dev")),
		rp.Select("http://192.168.1.226:8081", rp.Always()),
	)

	err := http.ListenAndServe(":"+port, proxy)
	if err != nil {
		panic(err)
	}
}

func HostMatches(host string) rp.Rule {
	r := rp.HostMatches(host)
	r.Matcher = func(r *http.Request) bool {
		return r.Host == host
	}
	return r
}
