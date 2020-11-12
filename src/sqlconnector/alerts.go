package sqlconnector

import "log"

func DeleteAlert(ruleId int) {
	db := connector()

	_, err := db.Query("DELETE FROM alerts where ruleId = ?;", ruleId)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()

}

func CountAlerts(name string) int {
	db := connector()

	var nb int
	err := db.QueryRow("SELECT COUNT(*) FROM alerts where nom = ?;", name).Scan(&nb)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()

	return nb
}

func ReturnMaxCriticality(name string) int {
	db := connector()

	var nb2 int
	err := db.QueryRow("SELECT MAX(criticality) FROM alerts where nom = ?", name).Scan(&nb2)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()

	return nb2
}

func InsertAlert(name string, ruleId int, criticality int){
	db := connector()

	_, err := db.Query("INSERT INTO alerts(nom, ruleId, criticality) values(?, ?, ?);", name, ruleId, criticality)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()
}



