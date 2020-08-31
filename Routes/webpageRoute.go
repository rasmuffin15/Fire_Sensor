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
	println("Inside sensorDataHandler func")

	pd := SensorData {
		Airquality: 23.45,
		Id:         "k5",
		Timestamp:  "tomorrow",
	}

	t, err := template.ParseFiles("Pages/showData.html")

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, pd)

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func HtmlPage() {
	println("Inside HtmlPage func")
	http.HandleFunc("/", sensorDataHandler)
	println("Created Path")
	http.ListenAndServe(":8080", nil)
	println("Exiting func")
}
