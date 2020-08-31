package Routes

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//Sensor struct holding information
type Sensor struct {
	airquality float64
	id string
	timestamp string
}

//Handles route for posting data into postgres
func sensorPutPath(rg *gin.RouterGroup) {
	sensor := rg.Group("/sensor")
	//currentTime := time.Now()

	//Handler for getting row from db based on requested id
	//Move into own function
	sensor.GET("/get/:id", func(c *gin.Context) {
		//Holds requested id
		id := c.Param("id")
		//Holds returned row data

		var s1 Sensor
		//Calls func to open db connection
		db := DBConn()
		println("We returned GET")
		//Selects an entire row from db
		//$1 is replaced by requested id
		sqlGet := `SELECT * FROM sensordata WHERE sensorid = $1`
		println(sqlGet)
		fmt.Printf("%s\n", id)
		//Executes sql statement based on id
		row := db.QueryRow(sqlGet, id)
		//Copies returned row into individual variables
		err := row.Scan(&s1.id, &s1.timestamp, &s1.airquality)
		//Display data to terminala
		fmt.Println("<---Returned data--->")
		fmt.Printf("id: %s\n", s1.id)
		fmt.Printf("Date: %s\n", s1.timestamp)
		fmt.Printf("Air Quality: %f\n", s1.airquality)
		fmt.Println("<------------->")

		DBClose(db)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err)
			}
		}
	})
	
	//Handler for posting data into database
	sensor.POST("/post/:id/:timestamp/:aq", func(c *gin.Context) {
		var id string
		sid := c.Param("id")
		ts := c.Param("timestamp")
		aq := c.Param("aq")

		s, _ := strconv.ParseFloat(aq, 32)

		s1 := Sensor{s, sid, ts}
		//c.JSON(http.StatusOK, s1)

		//Calls func to open db connection
		db := DBConn()
		println("We returned POST")
		//Inserts values into db
		sqlPost := `
			INSERT INTO sensordata (sensorid, date, airquality)
			VALUES ($1, $2, $3) RETURNING sensorid`

		//Executes insert statement, expects one row for return value
		err := db.QueryRow(sqlPost, s1.id, s1.timestamp, s1.airquality).Scan(&id)

		fmt.Println("<-----Posted Data----->")
		fmt.Printf("ID is: %s\n", id)
		fmt.Printf("Air Quality: %f\n", s1.airquality)
		fmt.Println("<-------------------->")

		//Calls func to close the db connection
		DBClose(db)

		if err != nil {
			panic(err)
		}
	})
}

