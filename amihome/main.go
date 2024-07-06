package main

import (
	"embed"
	"html/template"
	"net"
	"net/http"
)

//go:embed templates/*.tmpl.html
var content embed.FS

type PageData struct {
	Title       string
	Message     string
	Explanation template.HTML
}

var (
	local, docker *net.IPNet
)

func init() {
	local = mustParseCIDR("192.168.0.0/16")
	docker = mustParseCIDR("172.0.0.0/8")
}

var (
	localData = PageData{
		Title:       "Am I Home? Yes",
		Message:     "Hello, Local!",
		Explanation: `<p>Congratulations! It appears your connection is being routed through the local network.</p>`,
	}

	dockerData = PageData{
		Title:   "Am I Home? No",
		Message: "Hello, World!",
		Explanation: `<p>You are seeing this message because your connection is being routed through the public internet. This may be due to one of the following reasons:</p>
                <ol>
                    <li>You are not connected to the network.</li>
                    <li>You might be using a VPN, causing your traffic to be routed outside of the network.</li>
                    <li>Your DNS settings may not be pointing directly to the network gateway.</li>
                    <li>Your DNS cache might still be holding the global address.</li>
                </ol>
                <p>To resolve this issue, ensure you are connected to the <code>dbut2</code> network and that your local DNS settings are pointing to the network gateway at <code>192.168.2.1</code>.</p>`,
	}
)

func mustParseCIDR(str string) *net.IPNet {
	_, n, err := net.ParseCIDR(str)
	if err != nil {
		panic(err.Error())
	}
	return n
}

func main() {
	tmpl, err := template.New("html").ParseFS(content, "templates/*.tmpl.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ips := MapSlice(r.Header.Values("X-Forwarded-For"), func(ip string) net.IP {
			return net.ParseIP(ip)
		})

		var data PageData

		if IsAll(ips, func(ip net.IP) bool {
			return local.Contains(ip)
		}) {
			data = localData
		}

		if IsAny(ips, func(ip net.IP) bool {
			return docker.Contains(ip)
		}) {
			data = dockerData
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Lookup("index.tmpl.html").Execute(w, data)
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func IsAny[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func IsAll[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func MapSlice[S ~[]E, E any, U any](s S, f func(E) U) []U {
	u := make([]U, len(s))
	for i, e := range s {
		u[i] = f(e)
	}
	return u
}
