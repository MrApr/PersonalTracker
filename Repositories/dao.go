package Repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

//dbConnection determines and specifies database connection
type dbConnection struct {
	user   string
	pass   string
	host   string
	port   int
	dbName string
}

func ConnectToDb() *gorm.DB {
	conn := dbConnection{
		user:   "heroes",
		pass:   "22510070",
		host:   "localhost",
		port:   3306,
		dbName: "process_tracker",
	} //Todo make it to read from conf file
	//Todo check whether gorm creates connection pool or not
	return conn.mysqlConnection()
}

//makeMysqlConnection to mysql server
func (dbConn *dbConnection) mysqlConnection() *gorm.DB {
	dsn := dbConn.makeConnString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//makeConnString to connect to database
func (dbConn dbConnection) makeConnString() string {
	port := strconv.Itoa(dbConn.port)
	dsn := dbConn.user + ":" + dbConn.pass + "@tcp(" + dbConn.host + ":" + port + ")/" + dbConn.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}
