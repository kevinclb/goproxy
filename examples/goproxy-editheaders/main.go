package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		fmt.Println("rec request to : ", req.Host, "from : ", req.RemoteAddr)
		return req, nil
	})
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnResponse().DoFunc(respond)
	log.Fatal(http.ListenAndServe(":8080", proxy))

	//log.Fatal(http.ListenAndServeTLS(":8081", "ca.pem", "key.pem", proxy))
}

func respond(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	fmt.Println(ctx.Req.Header.Get("X-GoPoxy"))
	response := goproxy.NewResponse(ctx.Req,
		goproxy.ContentTypeText, http.StatusOK,
		"I intercepted my own request!")
	response.Header.Add("HelloHeader", "JustSayingHi")
	response.Header.Add("Cookie", "CustomCookie")
	return response
}
