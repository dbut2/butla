package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type VanityConfig struct {
	Host       string
	Repository string
	Module     string
}

const tmpl = `<!DOCTYPE html>
<html>
<head>
    <meta name="go-import" content="{{.Host}}/{{.Module}} git {{.Repository}}/{{.Module}}">
    <meta name="go-source" content="{{.Host}}/{{.Module}} {{.Repository}}/{{.Module}} {{.Repository}}/{{.Module}}/tree/master{/dir} {{.Repository}}/{{.Module}}/blob/master{/dir}/{file}#L{line}">
    <meta http-equiv="refresh" content="0; url={{.Repository}}/{{.Module}}">
    <link rel="canonical" href="{{.Repository}}/{{.Module}}">
    <title>Redirecting to {{.Repository}}/{{.Module}}</title>
</head>
</html>`

func main() {
	host := getEnvOrDefault("VANITY_HOST", "dbut.dev")
	githubUser := getEnvOrDefault("GITHUB_USER", "dbut2")
	port := getEnvOrDefault("PORT", "8080")

	t, err := template.New("vanity").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Remove leading slash and any query parameters
		path := strings.TrimPrefix(r.URL.Path, "/")
		path = strings.Split(path, "?")[0]

		// Skip serving for empty path or favicon
		if path == "" || path == "favicon.ico" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		config := VanityConfig{
			Host:       host,
			Repository: fmt.Sprintf("https://github.com/%s", githubUser),
			Module:     path,
		}

		w.Header().Set("Content-Type", "text/html")
		err := t.Execute(w, config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("Starting server on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
