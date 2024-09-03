package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var c []byte

type Config map[string]map[string]string

func main() {
	con := Config{}
	sanic(yaml.Unmarshal(c, &con))
	fmt.Println(con)
	sanic(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimRight(must(url.ParseRequestURI(r.RequestURI)).Path, "/")
		w.Header().Set("Location", getCode(con, r.Host, path))
		w.WriteHeader(http.StatusTemporaryRedirect)
	})))
}

func getCode(con Config, host, path string) string {
	if _, ok := con[host]; !ok {
		host = "default"
	}
	if v, ok := con[host][path]; ok {
		return v
	}
	if v, ok := con[host]["default"]; ok {
		return v
	}
	if v, ok := con["default"][path]; ok {
		return v
	}
	return con["default"]["default"]
}

func sanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func must[T any](v T, err error) T {
	sanic(err)
	return v
}
