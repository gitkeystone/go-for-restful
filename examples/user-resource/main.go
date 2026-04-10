package main


func main() {
    u := UserResource{
        map[string]User{}
    }

    restful.DefaultContainer.Add(u.WebService())


}