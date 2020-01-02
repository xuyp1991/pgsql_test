package db

import (
   db "github.com/xuyp1991/pgsql_test/db"
	"github.com/cihub/seelog"
	"github.com/pkg/errors"
)
// table(action_id bigint,action_type integer,action_name text,node_id text,ship_num text,object_id_result text,record_time TIMESTAMP,action_hash text,action_hash_info text)
type repeat_accept struct {
	Object_id		string
	Count			 	int
}

func Query_repeat_accept() (repeat_accept_info []repeat_accept, err error) {
	rows, err := db.GetDB().Query("select * from sku_repeat_accept()")

	if err != nil {
		seelog.Errorf("select from sku_repeat_accept failed, err: %v", err)
		return nil , errors.Wrap(err, "select from sku_repeat_accept failed")
	}

	defer rows.Close()

	for rows.Next() {
		var repeat_accept_inf repeat_accept
		err := rows.Scan(&repeat_accept_inf.Object_id,&repeat_accept_inf.Count)

		if err != nil {
			seelog.Errorf("scan to repeat_accept struct error, err: %v", err)
			return nil, errors.Wrap(err, "scan to repeat_accept failed")
		}
		seelog.Tracef("repeat accept info : %v", repeat_accept_inf)
		repeat_accept_info = append(repeat_accept_info, repeat_accept_inf)
	}

	err = rows.Err()

	if err != nil {
		seelog.Errorf("rows err in Query_repeat_accept, err: %v", err)
		return nil , errors.Wrap(err, "rows err in Query_repeat_accept")
	}

	return repeat_accept_info , nil
}
