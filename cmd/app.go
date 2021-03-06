/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-30 17:13:53
 * @LastEditTime    : 2022-07-15 18:03:55
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/app.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	clsgo "github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
)

var l = log.ClsLog()

func simpleHTTPServer() {

	l.Info("ClsGO application ", clsgo.Version)
	l.Info(config.Get("database.default.0.link"))

	// HTTP simple web server
	apis := make(http.APIS)
	apis["/"] = func(r *http.Request) {
		r.Response.Write("Home")
	}
	apis["/hello"] = func(r *http.Request) {
		r.Response.Write("Hello World!")
	}
	http.App("0.0.0.0", 8080, "v1", &apis)

}

func tcpServer() {
	net.TcpServer("0.0.0.0", 8081, &HMProtocol{})
}

func App() {
	go simpleHTTPServer()
	go tcpServer()
}
