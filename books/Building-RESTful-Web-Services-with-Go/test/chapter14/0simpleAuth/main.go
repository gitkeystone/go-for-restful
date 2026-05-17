package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
var users = map[string]string{
	"naren": "passme",
	"admin": "password",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if originalPassword, ok := users[username]; ok {
		// 开发环境必须设为 false！
		store.Options.Secure = false
		session, _ := store.Get(r, "session.id") // 创建 Session
		if originalPassword == password {
			session.Values["authenticated"] = true
			err := session.Save(r, w)
			if err != nil {
				http.Error(w, "Error saving session", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Logged In Successfully"))
}

// HealthcheckHandler returns the date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session.id") // 消费 Session
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["authenticated"] != nil && session.Values["authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

// LogoutHandler removes the session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//从 HTTP 请求中获取指定名称的会话（Session）
	session, _ := store.Get(r, "session.id") // 消费 Session
	session.Values["authenticated"] = false

	//将当前会话的状态保存到客户端（浏览器）
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(""))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/logout", LogoutHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15e9,
		ReadTimeout:  15e9,
	}
	log.Fatal(srv.ListenAndServe())
}
