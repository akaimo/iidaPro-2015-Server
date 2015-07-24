package controllers

import (
	"database/sql"
	"iidaPro/app/models"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap // このデータベースマッパーからSQLを流す
)

func InitDB() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic(err.Error())
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// ここで好きにテーブルを定義する
	t := Dbm.AddTable(models.User{}).SetKeys(true, "Id")
	t.ColMap("Password").Transient = true
	t.ColMap("Name").MaxSize = 20
	t.ColMap("Name").SetUnique(true)

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	// bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.DefaultCost)
	// demoUser := &models.User{0, "demo", "demo", bcryptPassword}
	// if err := Dbm.Insert(demoUser); err != nil {
	// 	panic(err)
	// }
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
