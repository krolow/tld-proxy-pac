package main

import (
	"io"
	"fmt"
	"flag"
	"log"
	"net"
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
	outboundIP := GetOutboundIP()
	tld := flag.String("tld", "dev", "top level domain to proxy e.g: .dev")
	forwardHost := flag.String("forward-host", outboundIP, "host ip to forward")
	forwardPort := flag.String("forward-port", "80", "host port to forward")
	listenPort := flag.String("listen-port", "8040", "proxy listen port")

	flag.Parse()

	r := strings.NewReplacer(
		"%tld%", *tld,
		"%host%", *forwardHost,
		"%port%", *forwardPort,
	)

	content := r.Replace(proxyTemplate)

	fmt.Println("Forwarding " + *tld + " to: " + *forwardHost + ":" + *forwardPort)
	fmt.Println("Outbound IP: " + outboundIP)
	fmt.Println("tld-proxy-pac running... http://0.0.0.0:" + *listenPort + "/")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, content)
	})

	http.ListenAndServe(":" + *listenPort, nil)
}


func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
			log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}
