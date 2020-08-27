package Routes

import (
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

	//Handler for posting data into database
	//Currently works as 'GET' but not 'POST'
	sensor.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		s1 := Sensor{32.45,id, currentTime.Format("01-02-2006 15:04:05 Mon")}
		c.JSON(http.StatusOK, s1)

		//Calls func to open db connection
		db := DBConn()

		//Hardcoded, inserts values into db
		sqlStatement := `
			INSERT INTO sensordata (sensorid, date, airquality)
			VALUES ($1, $2, $3) RETURNING sensorid`

		//Executes insert statement, expects one row for return value
		err := db.QueryRow(sqlStatement, "a1", "01-02-2006 15:04:05 Mon", 2345).Scan(&id)

		fmt.Println("New record ID is:", id)

		if err != nil {
			panic(err)
		}

	//Calls func to close the db connection
		DBClose(db)
	})
}

