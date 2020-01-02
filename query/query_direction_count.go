package db

import (
   db "github.com/xuyp1991/pgsql_test/db"
	"github.com/cihub/seelog"
	"github.com/pkg/errors"
)
// table(action_id bigint,action_type integer,action_name text,node_id text,ship_num text,object_id_result text,record_time TIMESTAMP,action_hash text,action_hash_info text)
type object_direcion struct {
	Object_id		string
	Direction		string
	Count			 	int
}

func Query_direction_count(object_id string) (direction_info []object_direcion, err error) {
	rows, err := db.GetDB().Query("select * from action_diction_count($1)", object_id)

	if err != nil {
		seelog.Errorf("select from action_diction_count failed, err: %v", err)
		return nil , errors.Wrap(err, "select from action_diction_count failed")
	}

	defer rows.Close()

	for rows.Next() {
		var direction_inf object_direcion
		err := rows.Scan(&direction_inf.Direction,&direction_inf.Count)
		direction_inf.Object_id = object_id;

		if err != nil {
			seelog.Errorf("scan to object_direcion struct error, err: %v", err)
			return nil, errors.Wrap(err, "scan to object_direcion failed")
		}
		seelog.Tracef("direction info : %v", direction_inf)
		direction_info = append(direction_info, direction_inf)
	}

	err = rows.Err()

	if err != nil {
		seelog.Errorf("rows err in action_diction_count, err: %v", err)
		return nil , errors.Wrap(err, "rows err in action_diction_count")
	}

	return direction_info , nil
}
