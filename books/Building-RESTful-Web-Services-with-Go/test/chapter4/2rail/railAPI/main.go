package main

import (
	"database/sql"
	"dbutils"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"

	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int    `json:"ID"`
	DriverName      string `json:"driverName"`
	OperatingStatus bool   `json:"operatingStatus"`
}

// StationResource holds information about locations
type StationResource struct {
	ID          string    `json:"ID"`
	Name        string    `json:"name"`
	OpeningTime time.Time `json:"openingTime"`
	ClosingTime time.Time `json:"closingTime"`
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int       `json:"ID"`
	TrainID     int       `json:"trainID"`
	StationID   int       `json:"stationID"`
	ArrivalTime time.Time `json:"arrivalTime"`
}

// Register adds paths and routes to a new service instance
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))

	ws.Route(ws.POST("").To(t.createTrain))

	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))

	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t *TrainResource) getTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train-id")

	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).
		Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		resp.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
// {"driverName": "Veronica", "operatingStatus": true}
func (t *TrainResource) createTrain(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(t)
	if err != nil {
		log.Println(err)
	}
	log.Println(t.DriverName, t.OperatingStatus)

	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?,?)")
	result, err := statement.Exec(t.DriverName, t.OperatingStatus)
	if err != nil {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		newID, _ := result.LastInsertId()
		t.ID = int(newID)
		resp.WriteHeaderAndEntity(http.StatusCreated, t)
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("train-id")
	stmt, _ := DB.Prepare("delete from train where id=?")
	_, err := stmt.Exec(id)
	if err != nil {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		resp.WriteHeader(http.StatusOK)
	}
}

func main() {
	// Connect to Database
	var err error
	DB, err = sql.Open("sqlite3", "railapi.db") // 始化全局变量, 使用 =
	if err != nil {
		log.Println("Driver creation failed!")
	}

	// Create Tables(train/station/schedule)
	dbutils.Initialize(DB)

	nc := restful.NewContainer()
	nc.Router(restful.CurlyRouter{}) // 使用花括号 {} 作为路径参数的标记

	t := &TrainResource{}
	t.Register(nc)

	log.Printf("start listening on localhost:8000")
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      nc,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
