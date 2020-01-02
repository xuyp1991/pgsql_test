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
	action_info,_ := query.Query_action_from_object("XIANG4")
	for _,actinfo := range action_info {
		seelog.Tracef("main get action info : %v", actinfo)
	}
	link_info,_ := query.Query_alllink("SKU5")
	for _,info := range link_info {
		seelog.Tracef("main get link info : %v", info)
	}

	query.Query_direction_count("SKU5")
	query.Query_repeat_accept()
	db.CloseDB()
}