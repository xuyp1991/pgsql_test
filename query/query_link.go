package db

import (
   db "github.com/xuyp1991/pgsql_test/db"
	"github.com/cihub/seelog"
	"time"
)
// table(action_id bigint,action_type integer,action_name text,node_id text,ship_num text,object_id_result text,record_time TIMESTAMP,action_hash text,action_hash_info text)
type object_link struct {
	Action_id		int
	Action_type		int
	Action_name 	string
	Node_id 			string
	Ship_num			string
	Object_id		string
	Record_time		time.Time
	Action_hash		string
	Action_hash_info string
}

func Query_alllink(object_id string) {
	rows, err := db.GetDB().Query("select * from query_all_link($1)", object_id)

	if err != nil {
		seelog.Errorf("select from t_actions failed, err: %v", err)
		return 
	}

	defer rows.Close()

	for rows.Next() {
		var link_inf object_link
		err := rows.Scan(&link_inf.Action_id,&link_inf.Action_type,&link_inf.Action_name,&link_inf.Node_id,&link_inf.Ship_num,&link_inf.Object_id,&link_inf.Record_time,
			&link_inf.Action_hash,&link_inf.Action_hash_info)

		if err != nil {
			seelog.Errorf("scan to action struct error, err: %v", err)
			return 
			//return nil, errors.Wrap(err, "scan from t_actions failed")
		}
		seelog.Tracef("link info : %v", link_inf)
		//objList = append(objList, objId)
	}

	err = rows.Err()

	if err != nil {
		seelog.Errorf("rows err in query_all_link, err: %v", err)
		return 
	}
}
