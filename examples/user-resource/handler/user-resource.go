package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/gitkeystone/user-resource/model"
)

// UserResource is the REST layer to the User domain
type UserResource struct {
	// normally one would use DAO (data access object)
	Users map[string]model.User
}

// WebService creates a new service that can handle REST requests for User resources.
func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
		// docs
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]model.User{}).
		Returns(200, "OK", []model.User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		// docs
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.User{}). // on the response
		Returns(200, "OK", model.User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{user-id}").To(u.upsertUser).
		// docs
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.User{})) // from the request

	ws.Route(ws.POST("").To(u.createUser).
		// docs
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(model.User{}))

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		// docs
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

// GET http://localhost:8080/users
func (u UserResource) findAllUsers(_ *restful.Request, response *restful.Response) {
	log.Println("findAllUsers")
	var list []model.User
	for _, each := range u.Users {
		list = append(list, each)
	}
	_ = response.WriteEntity(list)
}

// GET http://localhost:8080/users/1
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	log.Println("findUser")
	id := request.PathParameter("user-id")
	usr := u.Users[id]
	if len(usr.ID) == 0 {
		_ = response.WriteErrorString(http.StatusNotFound, "")
	} else {
		_ = response.WriteEntity(usr)
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u UserResource) upsertUser(request *restful.Request, response *restful.Response) {
	log.Println("upsertUser")
	usr := model.User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		_ = response.WriteError(http.StatusInternalServerError, err)
	} else {
		u.Users[usr.ID] = usr
		_ = response.WriteEntity(usr)
	}
}

// POST http://localhost:8080/users
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u UserResource) createUser(request *restful.Request, response *restful.Response) {
	log.Println("createUser")
	usr := model.User{ID: fmt.Sprintf("%d", time.Now().Unix())}
	err := request.ReadEntity(&usr)
	if err != nil {
		_ = response.WriteError(http.StatusInternalServerError, err)
	} else {
		u.Users[usr.ID] = usr
		_ = response.WriteEntity(usr)
	}
}

// DELETE http://localhost:8080/users/1
func (u UserResource) removeUser(request *restful.Request, _ *restful.Response) {
	log.Println("removeUser")
	id := request.PathParameter("user-id")
	delete(u.Users, id)
}
