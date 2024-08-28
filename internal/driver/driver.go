package driver

import (
	"database/sql"
	"time"
	//"github.com/jackc/pgconn"
	_"github.com/jackc/pgx/v5/stdlib"
	_"github.com/jackc/pgx/v5"
)

type DB struct {
	SQL *sql.DB
}

var dbConn =&DB{}
const maxOpenDbConn=10
const maxIdleDbConn=5
const maxDbLifetime=5*time.Minute

//connectSQL creates database pool for postgres
func ConnectSQL(dsn string)(*DB,error){
	d,err:=NewDatabase(dsn)
	if err!=nil{
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL=d
	err = testDB(d)
	if err!=nil{
		return nil,err
	}
	return dbConn,err
}
//creates a new database for the applicataion
func NewDatabase(dsn string)(*sql.DB,error){
	db,err:=sql.Open("pgx",dsn)
	if err!=nil{
		return nil,err
	}
	if err = db.Ping(); err!=nil{
		return nil,err
	}
	return db,err
}
//tries to ping the database
func testDB(d *sql.DB)error{
	err := d.Ping()
	if err!=nil{
		return err
	}
	return nil
}