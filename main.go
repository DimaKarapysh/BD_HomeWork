package main

import (
	"Portfolio_Nodes/db"
	"Portfolio_Nodes/tool"
)

func main() {
	err := tool.InitEnv()
	if err != nil {
		panic("PANIC ENV")
	}

	_, err = db.InitDatabase()
	if err != nil {
		panic("PANIC DB")
	}

	err = db.RunMigrations()
	if err != nil {
		panic("PANIC MIGRATION")
	}

}
