/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-05 14:26:51
 * @LastEditTime    : 2022-07-05 16:13:38
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/config/gfcfg.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 * Config implement use goframe, config automatic search supported
 * Only implement Get
 */
package config

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

// Get retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// It returns a default value specified by `def` if value for `pattern` is not found.
func Get(pattern string, def ...interface{}) (x *gvar.Var) {
	var ctx = gctx.New()
	result, err := gcfg.Instance().Get(ctx, pattern)
	if err != nil {
		if len(def) > 0 {
			return gvar.New(def[0])
		} else {
			return gvar.New(nil)
		}
	} else {
		return result
	}
}
