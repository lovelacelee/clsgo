/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 09:03:01
 * @LastEditTime    : 2022-07-08 18:42:07
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/net/tcp.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package net

import (
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

type Conn = gtcp.Conn
type Retry = gtcp.Retry

var l = log.ClsLog()

type TcpProtocol interface {
	Recv(conn *Conn) ([]byte, error)
	Send(conn *Conn, data []byte) error
}

func connectionClose(conn *gtcp.Conn) {
	err := conn.Close()
	if err != nil {
		l.Error(err)
	}
}

func TcpServer(host string, port int, proto TcpProtocol) {
	l.Warningf("TCP server listen %s:%d", host, port)
	gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		l.Info("New connection ", conn.RemoteAddr().String())
		defer connectionClose(conn)
		for {
			data, err := proto.Recv(conn)
			if err != nil {
				l.Error(err)
				break
			}
			err = proto.Send(conn, data)
			if err != nil {
				l.Error(err)
				break
			}
		}
		l.Error("connection closed ", conn.RemoteAddr().String())
	}).Run()
}
