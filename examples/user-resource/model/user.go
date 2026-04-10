package model

// User is just a sample type
// User Domain
type User struct {
    ID string `xml:"id" json:"id" description:"identifier of the user"`
    Name string `xml:"name" json:"name" description:"name of the user" default:"john"`
    Age int `xml:"age" json:"age" description:"age of the user" default:"21"`
}