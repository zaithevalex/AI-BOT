package manage

import (
	"BOOT-BOT/db/timers"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const CreateTableUsersSubscriptions = `CREATE TABLE IF NOT EXISTS SUBSCRIPTIONS(ID BIGINT, AUTOPAYMENT BOOLEAN, SUBSCRIPTION CHARACTER(30), AMOUNT_REQUESTS BIGINT, SUBSCRIBE_TIME BIGINT);`

const CreateTableUsersReq = `CREATE TABLE IF NOT EXISTS USERS(ID BIGINT PRIMARY KEY, AI CHARACTER(30));`

const insertUser = `INSERT INTO USERS(ID, AI) VALUES(%v, '%v');`

const insertSubscription = `INSERT INTO SUBSCRIPTIONS(ID,
                          AUTOPAYMENT,
                          SUBSCRIPTION,
                          AMOUNT_REQUESTS,
                          SUBSCRIBE_TIME) VALUES(%v, %v, '%v', %v, %v);`

const selectUser = `SELECT * FROM USERS WHERE ID = %d;`

const GetUserParam = `select %s from users where id = %d;`

const GetSubscriptionParam = `select %s from subscriptions where id = %d;`

const UpdateUserParam = `update USERS set %s = '%v' where id = %d;`

const UpdateUserSubscriptionParam = `update SUBSCRIPTIONS set %s = '%v' where id = %d;`

const (
	GoogleAI = "GoogleAI"
	GPT35    = "GPT-3.5"
	GPT4     = "GPT-4"
)

const (
	PayloadDefault = "sub_default"
	Payload2Weeks  = "sub_2weeks"
	Payload1Month  = "sub_1month"
	Payload1Year   = "sub_1year"
)

const DefaultReqsPerWeek = 3

func AddUser(db *sqlx.DB, id int64) error {
	_, err := db.Query(fmt.Sprintf(insertUser, id, GoogleAI))
	if err != nil {
		return err
	}

	nextMonday := timers.StartWeekUpdate()
	_, err = db.Query(fmt.Sprintf(insertSubscription, id, false, PayloadDefault, DefaultReqsPerWeek, nextMonday))
	if err != nil {
		return err
	}

	return nil
}

func CheckUser(db *sqlx.DB, id int64) (bool, error) {
	users, err := db.Exec(fmt.Sprintf(selectUser, id))
	if err != nil {
		return false, err
	}

	iUser, err := users.RowsAffected()
	if err != nil {
		return false, err
	}
	if iUser > 0 {
		return true, nil
	}
	return false, nil
}

func GetParam[T comparable](db *sqlx.DB, query string, id int64, reqParam string) (T, error) {
	var param T
	err := db.Get(&param, fmt.Sprintf(query, reqParam, id))
	if err != nil {
		return param, err
	}

	return param, nil
}

func UpdateParam(db *sqlx.DB, query string, id int64, reqParam string, newMeaning any) error {
	_, err := db.Query(fmt.Sprintf(query, reqParam, newMeaning, id))
	if err != nil {
		return err
	}
	return nil
}
