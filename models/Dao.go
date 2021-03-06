package models

import (
	"database/sql"
	"github.com/MrApr/PersonalTracker/Error"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

//dbConnection determines and specifies database connection
type dbConnection struct {
	user   string
	pass   string
	host   string
	port   int
	dbName string
}

//DB holds database instance
var DB *gorm.DB

//ConnectToDb creates a new connection to database
func ConnectToDb(maxConn int, maxOpen int) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	conn := dbConnection{
		user:   os.Getenv("USER"),
		pass:   os.Getenv("PASS"),
		host:   os.Getenv("HOST"),
		port:   port,
		dbName: os.Getenv("DBNAME"),
	}
	DB = conn.mysqlConnection()

	dbConfigurator, err := DB.DB()
	if err != nil {
		panic(Error.AdvanceError{
			Message: err.Error(),
			Line:    37,
			Type:    "Critical",
			File:    "DAO",
		})
	}

	conn.configDBConn(dbConfigurator, maxConn, maxOpen)
}

//makeMysqlConnection to mysql server
func (dbConn *dbConnection) mysqlConnection() *gorm.DB {
	//dsn := dbConn.makeConnString()
	db, err := gorm.Open(sqlite.Open("Pt.db"), &gorm.Config{})
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(Error.AdvanceError{
			Message: err.Error(),
			Line:    53,
			Type:    "Critical",
			File:    "DAO",
		})
	}
	return db
}

//makeConnString to connect to database
func (dbConn *dbConnection) makeConnString() string {
	port := strconv.Itoa(dbConn.port)
	dsn := dbConn.user + ":" + dbConn.pass + "@tcp(" + dbConn.host + ":" + port + ")/" + dbConn.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

//configure connected database
func (dbConn *dbConnection) configDBConn(db *sql.DB, maxIdleConn int, maxOpenConn int) {
	if maxIdleConn != 0 {
		db.SetMaxIdleConns(maxIdleConn)
	}

	if maxOpenConn != 0 {
		db.SetMaxOpenConns(maxOpenConn)
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetConnMaxIdleTime(time.Hour)
}
