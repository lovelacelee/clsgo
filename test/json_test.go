/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-13 15:25:49
 * @LastEditTime    : 2022-06-24 18:10:18
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /test/json_test.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package clsgo

import (
	"github.com/lovelacelee/clsgo/pkg/log"
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg/config"
)

type JsonModel struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	Github  string `json:"github"`
}

func TestJson(t *testing.T) {
	var json config.ClsJson
	//Load
	json.Load("test.json")

	var m JsonModel
	err := json.Parse(&m)
	log.Info("%v %v %v %v", err, m.Author, m.Version, m.Github)
	//Read check
	if !reflect.DeepEqual(m.Author, "Lovelace") {
		t.Errorf("Json value not match: %s != %s\n", m.Author, "Lovelace")
		t.Logf(json.String())
	}
	//Modification
	m.Author = "Lovelace"
	json.Save(m, "test.json")
	log.Info(json.String())
}

func BenchmarkJson(b *testing.B) {
	var json config.ClsJson
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Load("test.json")
		var m JsonModel
		json.Parse(&m)
		json.Save(m, "test.json")
	}
}
