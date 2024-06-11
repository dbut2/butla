package main

import (
	"embed"
	"html/template"
	"net"
	"net/http"
	"slices"
)

//go:embed templates/*.tmpl.html
var content embed.FS

type PageData struct {
	Title       string
	Message     string
	Explanation template.HTML
}

func main() {
	cfHosts, _ := net.LookupHost("tunnel")
	tmpl, err := template.New("html").ParseFS(content, "templates/*.tmpl.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ips := r.Header.Values("X-Forwarded-For")
		data := PageData{
			Title:       "Am I Home? Yes",
			Message:     "Hello, Local!",
			Explanation: `<p>Congratulations! It appears your connection is being routed through the local network.</p>`,
		}
		if HasAny(cfHosts, ips) {
			data.Title = "Am I Home? No"
			data.Message = "Hello, World!"
			data.Explanation = `<p>You are seeing this message because your connection is being routed through the public internet. This may be due to one of the following reasons:</p>
                <ol>
                    <li>You are not connected to the network.</li>
                    <li>You might be using a VPN, causing your traffic to be routed outside of the network.</li>
                    <li>Your DNS settings may not be pointing directly to the network gateway.</li>
                    <li>Your DNS cache might still be holding the global address.</li>
                </ol>
                <p>To resolve this issue, ensure you are connected to the <code>dbut2</code> network and that your local DNS settings are pointing to the network gateway at <code>192.168.2.1</code>.</p>`
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Lookup("index.tmpl.html").Execute(w, data)
	})

	err = http.ListenAndServe(":8080", nil)
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
