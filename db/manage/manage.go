package manage

import (
	"BOOT-BOT/db/timers"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

const CreateTableUsersReq = `CREATE TABLE IF NOT EXISTS USERS(ID BIGINT PRIMARY KEY, AUTOPAYMENT BOOLEAN, AI CHARACTER(30), AMOUNT_REQUESTS BIGINT, SUBSCRIBE_TIME BIGINT);`
const DefaultReqsPerWeek = 3
const getParam = "select %s from users where id = %d;"
const insertUser = `INSERT INTO USERS(ID,
                  AUTOPAYMENT,
                  AI,
                  AMOUNT_REQUESTS,
                  SUBSCRIBE_TIME) VALUES(%v, %v, '%v', %v, %v);`
const selectUser = `SELECT * FROM USERS WHERE ID = %d;`
const updateUserParam = "update users set %s = '%v' where id = %d;"

const (
	GoogleAI = "GoogleAI"
	GPT35    = "GPT-3.5"
	GPT4     = "GPT-4"
)

func AddUser(db *sqlx.DB, id int64) error {
	nextMonday := timers.StartWeekUpdate()

	_, err := db.Query(fmt.Sprintf(insertUser, id, false, GoogleAI, DefaultReqsPerWeek, nextMonday))
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

func GetParam[T comparable](db *sqlx.DB, id int64, reqParam string) (T, error) {
	var param T
	err := db.Get(&param, fmt.Sprintf(getParam, reqParam, id))
	if err != nil {
		return param, err
	}

	return param, nil
}

func UpdateParam(db *sqlx.DB, id int64, reqParam string, newMeaning any) error {
	if reflect.TypeOf(newMeaning) == reflect.TypeOf("") {

	}

	_, err := db.Query(fmt.Sprintf(updateUserParam, reqParam, newMeaning, id))
	if err != nil {
		return err
	}
	return nil
}
