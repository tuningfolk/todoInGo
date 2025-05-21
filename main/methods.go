package main

import "net/http"

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ all pending tasks"
	} else {
		message = "POST/ all pending tasks"
	}
	w.Write([]byte(message))
}

func ShowCompleted(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ show completed tasks"
	} else {
		message = "POST/ not allowed"
	}
	w.Write([]byte(message))
}

func Trash(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ show trashed tasks"
	} else {
		message = "POST/ trash post"
	}
	w.Write([]byte(message))
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ show edit page"
	} else {
		message = "POST/ submit edited task"
	}
	w.Write([]byte(message))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ show  login page"
	} else {
		message = "POST/ log in user"
	}
	w.Write([]byte(message))
}

func Register(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ show register page"
	} else {
		message = "POST/ register user"
	}
	w.Write([]byte(message))
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ changepassword page"
	} else {
		message = "POST/ changepassword page"
	}
	w.Write([]byte(message))
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ add not allowed"
	} else {
		message = "POST/ add new task"
	}
	w.Write([]byte(message))
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ complete not allowed"
	} else {
		message = "POST/ mark task as complete"
	}
	w.Write([]byte(message))

}

func Delete(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "GET/ delete not allowed"
	} else {
		message = "POST/ permanently delete task"
	}
	w.Write([]byte(message))
}
