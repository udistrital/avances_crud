package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TablasAvances_20210301_160503 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TablasAvances_20210301_160503{}
	m.Created = "20210301_160503"

	migration.Register("TablasAvances_20210301_160503", m)
}

// Run the migrations
func (m *TablasAvances_20210301_160503) Up() {
	file, err := ioutil.ReadFile("../scripts/tablas_avances.up.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *TablasAvances_20210301_160503) Down() {
	file, err := ioutil.ReadFile("../scripts/tablas_avances.down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
