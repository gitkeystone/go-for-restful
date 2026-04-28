package main

import (
	"database/sql"
	"dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

// DB Driver visible to whole program.
var DB *sql.DB

// StationResource holds information about locations.
// 两种数据的占位符:
//  1. 来自 HTTP 请求的 POST Body;
//  2. 从数据库查询的数据;
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// GetStation returns the station detail
func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station-id")

	// 提前在数据库层面把时间类型转换成字符串，确保 Go 代码能用 string 类型安全地接收，避免因类型转换引发错误
	err := DB.QueryRow("select ID,NAME,CAST(OPENING_TIME as char),CAST(CLOSING_TIME as char) from station where id=?", id).
		Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": station,
		})
	}
}

// CreateStation handles the POST
func CreateStation(c *gin.Context) {
	var station StationResource
	// Parse the body into our resource
	if err := c.BindJSON(&station); err == nil {
		// Format Time to Go time format
		statement, _ := DB.Prepare("insert into station (name, opening_time, closing_time) values (?,?,?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveStation handles the removing of resource
func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "railapigin.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)

	r := gin.Default()
	// Add routes to REST verbs
	// GET http://localhost:8000/v1/stations/1
	r.GET("/v1/stations/:station-id", GetStation)
	// POST http://localhost:8000/v1/stations
	// {"name": "Brooklyn", "opening_time":"8:12:00", "closing_time":"18:23:00"}
	r.POST("/v1/stations", CreateStation)
	// DELETE http://localhost:8000/v1/stations/1
	r.DELETE("v1/stations/:station-id", RemoveStation)
	// Default listen and serve on 0.0.0.0:8000
	err = r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
