package Routes

import (
	"html/template"
	"log"
	"net/http"
)

type SensorData struct {
	Airquality float64
	Id string
	Timestamp string
}

func sensorDataHandler(w http.ResponseWriter, r *http.Request) {

	db := DBConn()
	i := 0
	var S1 [4]SensorData

	sqlGet := `SELECT * FROM sensordata`
	rows, err := db.Query(sqlGet)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&S1[i].Id, &S1[i].Timestamp, &S1[i].Airquality)

		if err != nil {
			panic(err)
		}

		i++

	}

	t, err := template.ParseFiles("Pages/showData.html")

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, S1)

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	DBClose(db)
}

func HtmlPage() {
	http.HandleFunc("/", sensorDataHandler)
	http.ListenAndServe(":8000", nil)

}
