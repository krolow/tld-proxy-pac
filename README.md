# tld-proxy-pac

The idea is to provide an easy way to setup custom TLD (Top Level Domains) for local development.


### Download

[Releases](https://github.com/krolow/tld-proxy-pac/releases)

**Or** you can use the docker image: [krolow/tld-proxy-pac:latest](https://hub.docker.com/r/krolow/tld-proxy-pac/)

### Options

- **-tld**: Top level domain to forward (default .local)
- **-forward-host**: IP address that will receive the requests TLD (default outbound ip)
- **-forward-port**: Port that proxy should redirect requests for specific TLD (default 80)
- **-listen-port**: Proxy server port (default 8040)


### Example

#### Running

```bash
tld-proxy-pac -tld=my-dev

tld-proxy-pac --forward-host=192.168.1.1 --forward-port=3000

tld-proxy-pac -tld=local --forward-port=3000 --listen-port=4040
```

### Enabling proxy

#### Ubuntu

`System Settings > Network > Network Proxy > Automatic`

#### OS X

`Network Preferences > Advanced > Proxies > Automatic Proxy Configuration`


#### Windows

`Settings > Network and Internet > Proxy > Use setup script`


## License

Licensed under <a href="http://krolow.mit-license.org/">The MIT License</a>
Redistributions of files must retain the above copyright notice.

## Author

Vinícius Krolow - krolow[at]gmail.com
