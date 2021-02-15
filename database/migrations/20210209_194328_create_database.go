package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateDatabase_20210209_194328 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateDatabase_20210209_194328{}
	m.Created = "20210209_194328"

	migration.Register("CreateDatabase_20210209_194328", m)
}

// Run the migrations
func (m *CreateDatabase_20210209_194328) Up() {
	file, err := ioutil.ReadFile("../script/20210209_194328_create_database.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CreateDatabase_20210209_194328) Down() {
	file, err := ioutil.ReadFile("../script/20210209_194328_create_database_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
