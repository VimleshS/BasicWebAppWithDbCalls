package main

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

type MainDb struct {
	dbConn mysql.Conn
}

func (mdb *MainDb) Init(dbname string) error {
	mdb.dbConn = mysql.New("tcp", "", "127.0.0.1:3306", "", "", dbname)
	err := mdb.dbConn.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (mdb *MainDb) ExecuteQuery(sql string) ([]mysql.Row, error) {
	rows, _, err := mdb.dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (mdb *MainDb) Dispose() {
	mdb.dbConn.Close()
}

var SQLConnection mysql.Conn

func init_(dbname string) error {
	SQLConnection = mysql.New("tcp", "", "127.0.0.1:3306", "", "", "test")
	err := SQLConnection.Connect()
	if err != nil {
		return err
	}
	return nil
}

func ExecuteQuery_(sql string) ([]mysql.Row, error) {
	rows, _, err := SQLConnection.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
