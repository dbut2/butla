package main

import (
	"fmt"
	"net"
	"net/http"
	"slices"
)

func main() {
	cfHosts, _ := net.LookupHost("tunnel")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ips := r.Header.Values("X-Forwarded-For")
		if HasAny(cfHosts, ips) {
			w.Write([]byte("Hello, World!\n"))
		} else {
			w.Write([]byte("Hello, Local!\n"))
		}
		w.Write([]byte(fmt.Sprintf("Tunnel IP is %v\n", cfHosts)))
		w.Write([]byte(fmt.Sprintf("Your IP is %v\n", ips)))
		w.Write([]byte(fmt.Sprintf("Headers are %v\n", r.Header)))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func HasAny[S ~[]E, E comparable](a, b S) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}
