package main

// User is just a sample type
type User struct {
    ID string `xml:"id" json:"id" description:"identifier of the user"`
    Name string `xml:"name" json:"name" description:"name of the user" default:"john"`
    Age int `xml:"age" json:"age" description:"age of the user" default:"21"`
}

// UserResource is the REST layer to the User domain
// User Domain
type UserResource struct {
    // normally one would use DAO (data access object)
    users map[string]User
}

// WebService creates a new service that can handle REST requests for User resources.


