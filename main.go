package main

import (
	_ "embed"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config map[string]map[string]string

func main() {
	var con Config
	sanic(yaml.Unmarshal(must(os.ReadFile("config.yaml")), &con))
	sanic(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimRight(must(url.ParseRequestURI(r.RequestURI)).Path, "/")
		w.Header().Set("Location", lengthen(con, r.Host, path))
		w.WriteHeader(http.StatusTemporaryRedirect)
	})))
}

func lengthen(con Config, host, path string) string {
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

func must[T any](v T, err error) T { sanic(err); return v }
