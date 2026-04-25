// TODO: 引入 httprouter
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := getCommandOutput("go", "version")
	_, _ = io.WriteString(w, resp)
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(w, getCommandOutput("cat", params.ByName("name")))
}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}
