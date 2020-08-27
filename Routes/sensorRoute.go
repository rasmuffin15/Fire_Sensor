package Routes

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Sensor struct holding information
type Sensor struct {
	airquality float32
	id string
	timestamp string
}

//Handles route for posting data into postgres
func sensorPutPath(rg *gin.RouterGroup) {
	sensor := rg.Group("/sensor")
	currentTime := time.Now()

	//Handler for getting row from db based on requested id
	sensor.GET("/get/:id", func(c *gin.Context) {
		//Holds requested id
		id := c.Param("id")
		//Holds returned row data
		var s1 Sensor
		//Calls func to open db connection
		db := DBConn()

		//Selects an entire row from db
		//$1 is replaced by requested id
		sqlGet := `SELECT * FROM sensordata WHERE sensorid = $1`
		//Executes sql statement based on id
		row := db.QueryRow(sqlGet, id)
		//Copies returned row into individual variables
		err := row.Scan(&s1.id, &s1.timestamp, &s1.airquality)
		//Display data to terminal
		fmt.Println("<---Returned data--->")
		fmt.Printf("id: %s\n", s1.id)
		fmt.Printf("Date: %s\n", s1.timestamp)
		fmt.Printf("Air Quality: %f\n", s1.airquality)
		fmt.Println("<------------->")

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err)
			}
		}
		DBClose(db)
	})
	
	//Handler for posting data into database
	sensor.POST("/post/:id", func(c *gin.Context) {
		id := c.Param("id")
		s1 := Sensor{32.45,id, currentTime.Format("01-02-2006 15:04:05 Mon")}
		c.JSON(http.StatusOK, s1)

		//Calls func to open db connection
		db := DBConn()

		//Hardcoded, inserts values into db
		sqlPost := `
			INSERT INTO sensordata (sensorid, date, airquality)
			VALUES ($1, $2, $3) RETURNING sensorid`

		//Executes insert statement, expects one row for return value
		err := db.QueryRow(sqlPost, "c1", "01-02-2006 15:04:05 Mon", 2345).Scan(&id)

		fmt.Println("New record ID is:", id)

		if err != nil {
			panic(err)
		}
		//Calls func to close the db connection
		DBClose(db)
	})
}

