package chrome

import "github.com/gopherjs/gopherjs/js"

type Proxy struct {
	o        js.Object
	Settings js.Object
}

func NewProxy(proxyObj js.Object) *Proxy {
	p := new(Proxy)
	p.o = proxyObj
	if proxyObj.String() != "undefined" {
		p.Settings = proxyObj.Get("settings")
	}
	return p
}

/*
* Types
 */

type ProxyServer struct {
	js.Object
	Scheme string `js:"scheme"`
	Host   string `js:"host"`
	Port   int    `js:"port"`
}

type ProxyRules struct {
	js.Object
	SingleProxy   ProxyServer `js:"singleProxy"`
	ProxyForHttp  ProxyServer `js:"proxyForHttp"`
	ProxyForHttps ProxyServer `js:"proxyForHttps"`
	ProxyForFtp   ProxyServer `js:"proxyForFtp"`
	FallbackProxy ProxyServer `js:"fallbackProxy"`
	BypassList    []string    `js:"bypassList"`
}

type PacScript struct {
	js.Object
	Url       string `js:"url"`
	Data      string `js:"data"`
	Mandatory bool   `js:"mandatory"`
}

type ProxyConfig struct {
	js.Object
	Rules     ProxyRules `js:"rules"`
	PacScript PacScript  `js:"pacScript"`
	Mode      string     `js:"mode"`
}

/*
* Events
 */

// OnProxyError notifies about proxy errors.
func (p *Proxy) OnProxyError(callback func(details map[string]interface{})) {
	p.o.Get("onProxyError").Call("addListener", callback)
}
