package handler

// UserResource is the REST layer to the User domain
type UserResource struct {
    // normally one would use DAO (data access object)
    users map[string]User
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
        Metadata(restfulspec.KeyOpenAPITags,tags).
        Writes([]User{}).
        Returns(200, "OK", []User{}))

    ws.Route(ws.PUT("/{user-id}").To(u.upsertUser).
        // docs
        Doc("update a user").
        Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
        Metadata(restfulspec.KeyOpenAPITags,tags).
        Reads(User{})) // from the request




    return ws
}

