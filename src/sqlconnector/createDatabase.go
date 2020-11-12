package sqlconnector

import (
	"log"
)

func CreateDatabase() {
	db := connector()

	// Connect and check the server version
	_, err := db.Query("CREATE TABLE IF NOT EXISTS alerts(id int primary key auto_increment, nom varchar(100), ruleId int, criticality int);")
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()
}