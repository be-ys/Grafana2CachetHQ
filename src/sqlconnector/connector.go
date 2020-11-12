package sqlconnector

import (
	"database/sql"
	"helpers"
	"strconv"
)

func connector() *sql.DB{
	configuration := helpers.GetConfiguration()

	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", configuration.Sql.Username+":"+configuration.Sql.Password+"@tcp("+configuration.Sql.Host+":"+strconv.Itoa(configuration.Sql.Port)+")/"+configuration.Sql.Database)

	return db
}