package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ShowTasks)
	http.HandleFunc("/completed/", ShowCompleted)
	http.HandleFunc("/trash/", Trash)
	http.HandleFunc("/edit/", Edit)
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/changepassword/", ChangePassword)
	http.HandleFunc("/add/", AddTask)
	http.HandleFunc("/complete/", CompleteTask)
	http.HandleFunc("/delete/", Delete)

	http.Handle("/static/", http.Fileserver(http.Dir("public")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
