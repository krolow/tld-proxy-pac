package main

import (
	"io"
	"fmt"
	"flag"
	"net/http"
	"strings"
)



var proxyTemplate = `function FindProxyForURL (url, host) {
  if (dnsDomainIs(host, '.%tld%')) {
    return 'PROXY %host%:%port%';
  }
  return 'DIRECT';
}`

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	tld := flag.String("tld", "dev", "top level domain to proxy e.g: .dev")
	forwardHost := flag.String("forward-host", "127.0.0.1", "host ip to forward")
	forwardPort := flag.String("forward-port", "80", "host port to forward")
	listenPort := flag.String("listen-port", "8040", "proxy listen port")

	flag.Parse()

	r := strings.NewReplacer(
		"%tld%", *tld,
		"%host%", *forwardHost,
		"%port%", *forwardPort,
	)

	content := r.Replace(proxyTemplate)

	fmt.Println("tld-proxy-pac running... http://127.0.0.1:" + *listenPort + "/")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, content)
	})

	http.ListenAndServe(":" + *listenPort, nil)
}
