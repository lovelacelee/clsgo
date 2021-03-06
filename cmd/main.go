/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 08:59:04
 * @LastEditTime    : 2022-07-08 14:02:32
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/main.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
	"sync"
	"time"
)

func main() {

	var workGroup sync.WaitGroup
	workGroup.Add(1)

	cfg, err := config.ClsConfig("config", "clsgo", true)
	if err != nil {
		l.Error("Config load failed!")
	}
	// Init logger
	go func() {
		for {
			log.Update(cfg.Sub("log"))
			time.Sleep(time.Second)
		}
	}()

	App()
	workGroup.Wait()
}
