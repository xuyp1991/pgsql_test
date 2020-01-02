package main

import (
	db "github.com/xuyp1991/pgsql_test/db"
	query "github.com/xuyp1991/pgsql_test/query"
	"github.com/cihub/seelog"
	// "time"
)



func main() {
	defer seelog.Flush()
	db.Init()
	query.Query_action_from_object("XIANG4")
	query.Query_alllink("SKU5")
	db.CloseDB()
}