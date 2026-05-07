package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"helper"
	"io"
	"log"
	"net/http"
)

type DBClient struct {
	db *gorm.DB
}

type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{
		db: db,
	}
	r := mux.NewRouter()
	// POST http://localhost:8000/v1/package
	// {}
	r.HandleFunc("/v1/package", dbclient.PostPackage).Methods("POST")

	// GET http://localhost:8000/v1/package/id
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods("GET")

	// GET http://localhost:8000/v1/package/id?weight=100
	r.HandleFunc("/v1/package", dbclient.GetPackageByWeight).Methods("GET")

	srv := &http.Server{
		Addr:         "localhost:8000",
		Handler:      r,
		ReadTimeout:  15e9,
		WriteTimeout: 15e9,
	}
	log.Fatal(srv.ListenAndServe())
}

// PostPackage saves the package information
// POST http://localhost:8000/v1/package
// {}
func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var pkg = helper.Package{}
	postBody, _ := io.ReadAll(r.Body)
	pkg.Data = string(postBody)

	driver.db.Save(&pkg)
	responseMap := map[string]any{
		"id": pkg.ID, // 执行 Save 函数后，ID 自动回填，因此，使用带地址的 &pkg
	}

	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

// GetPackage fetches the original URL for the given
// encoded(short) string
// GET http://localhost:8000/v1/package/id
func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var pkg = helper.Package{}
	vars := mux.Vars(r)

	driver.db.First(&pkg, "ID = ?", vars["id"])

	var pkgData any
	json.Unmarshal([]byte(pkg.Data), &pkgData)

	response := PackageResponse{
		Package: pkg,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(response)
	w.Write(jsonData)
}

// GetPackagesbyWeight fetches all packages with given weight
// GET http://localhost:8000/v1/package?weight=n
func (driver DBClient) GetPackageByWeight(w http.ResponseWriter, r *http.Request) {
	var pkgs []helper.Package

	weight := r.FormValue("weight")

	// Handle response details
	driver.db.Find(&pkgs, "data->>'weight' = ?", weight)
	//driver.db.Raw("SELECT * FROM \"Package\" WHERE data->>'weight' = ?", weight).Scan(&pkgs)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	respJSON, _ := json.Marshal(pkgs)
	w.Write(respJSON)
}
