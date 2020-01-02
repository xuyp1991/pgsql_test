package db

import (
   db "github.com/xuyp1991/pgsql_test/db"
	"github.com/cihub/seelog"
	"time"
	"github.com/pkg/errors"
)

type action struct {
	Action_id		int
	Action_type		int
	Node_serial 	int
	Ship_num 		string
	Destination		string
	Object_id		string
	Action_info		string
	Action_status	int
	Record_time		time.Time
	Action_hash		string
}

func Query_action_from_object(object_id string) (action_info []action, err error){
	rows, err := db.GetDB().Query("select * from t_actions where object_id=$1", object_id)

	if err != nil {
		seelog.Errorf("select from t_actions failed, err: %v", err)
		return nil,errors.Wrap(err, "select from t_actions failed")
	}

	defer rows.Close()

	for rows.Next() {
		var action_inf action
		err := rows.Scan(&action_inf.Action_id,&action_inf.Action_type,&action_inf.Node_serial,&action_inf.Ship_num,&action_inf.Destination,
			&action_inf.Object_id,&action_inf.Action_info,&action_inf.Action_status,&action_inf.Record_time,&action_inf.Action_hash)

		if err != nil {
			seelog.Errorf("scan to action struct error, err: %v", err)
			return nil,errors.Wrap(err, "scan to action struct error")
			//return nil, errors.Wrap(err, "scan from t_actions failed")
		}

		seelog.Tracef("actions : %d,\t%d,\t%d,\t%s,\t%s,\t%s,\t%s,\t%d,\t%s,\t%s", action_inf.Action_id,action_inf.Action_type,action_inf.Node_serial,action_inf.Ship_num,
			action_inf.Destination,action_inf.Object_id,action_inf.Action_info,action_inf.Action_status,action_inf.Record_time,action_inf.Action_hash)
		action_info = append(action_info, action_inf)
	}

	err = rows.Err()

	if err != nil {
		seelog.Errorf("rows err in t_actions, err: %v", err)
		return nil,errors.Wrap(err, "rows err in t_actions")
	}

	return action_info,nil
}
