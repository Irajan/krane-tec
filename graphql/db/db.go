package db

import (
	"database/sql"

	stringUtils "github.com/Irajan/krane-stack/utils/string"

	logger "github.com/Irajan/krane-stack/utils/logger"

	errorUtils "github.com/Irajan/krane-stack/utils/error"

	"github.com/Irajan/krane-stack/configs"

	_ "github.com/lib/pq"
)

var qb *sql.DB

func connect() *sql.DB {
	// Connect to the database

	logger.Info("Connecting to database . . . .")

	appConfig := configs.Config

	psqlInfo := stringUtils.CreateString("host=%s port=%s user=%s password=%s dbname=%s",
		appConfig.DbHost, appConfig.DbPort, appConfig.DbUser, appConfig.DbPassword, appConfig.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	errorUtils.HandleError(err, "Error connecting to database")

	// Check the connection
	err = db.Ping()

	errorUtils.HandleError(err, "Error pinging on database")

	return db
}

func Initialize() {
	qb = connect()
}

func Query(query string, args ...any) *sql.Rows {

	rows, err := qb.Query(query, args...)

	errorUtils.HandleError(err, "Error on query ")

	return rows
}

func Insert(query string, values []any, returns []any) {
	err := qb.QueryRow(query, values...).Scan(returns...)

	errorUtils.HandleError(err, "Error on insert query ")

}

func Close() {
	qb.Close()
}
