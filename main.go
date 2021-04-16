package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/golang/glog"
)

var (
	userMap = make(map[string]string)
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/loginSubmit", handleLoginSubmit)
	mux.HandleFunc("/changePassword", handleChangePassword)
	mux.HandleFunc("/deleteUser", handleDeleteUser)

	server := http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		glog.Error(err)
	}
}

// The handler for the root path.
func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "html") // application/json
	fmt.Fprintf(writer, "Hello, World")
}

// handle login
func handleLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	t, _ := template.ParseFiles("login.gtpl")
	log.Println(t.Execute(writer, nil))

}

// handle login submit
func handleLoginSubmit(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	request.ParseForm()
	fmt.Println("username:", request.Form["username"])
	fmt.Println("password:", request.Form["password"])

	userMap[request.Form["username"][0]] = request.Form["password"][0]

	fmt.Println("userMap", userMap)

	fmt.Fprintf(writer, "done")
}

// handleChangePassword...
func handleChangePassword(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	request.ParseForm()
	fmt.Println("username:", request.Form["username"])
	fmt.Println("password:", request.Form["password"])

	userMap[request.Form["username"][0]] = request.Form["password"][0]

	fmt.Println("userMap", userMap)

	fmt.Fprintf(writer, "done")
}

// handleDeleteUser...
func handleDeleteUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method)
	request.ParseForm()
	fmt.Println("deleted username:", request.Form["username"])

	delete(userMap, request.Form["username"][0])

	fmt.Println("userMap", userMap)

	fmt.Fprintf(writer, "done")
}
