package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ShowTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var message string

	if r.Method == "GET" {
		rows, err := db.Query("SELECT *  FROM posts")
		if err != nil {
			fmt.Println("db Query SELECT FROM posts failed")
			fmt.Println(err)
			return
		}
		defer rows.Close()
		var posts []Post
		post := new(Post)
		for rows.Next() {
			var bodyResponse sql.NullString
			err = rows.Scan(&post.Post_id, &post.Title, &bodyResponse, &post.User_id, &post.Completed, &post.Trashed)
			if err != nil {
				fmt.Println("scanning error: ", err)
			}
			if bodyResponse.Valid {
				post.Body = bodyResponse.String
			} else {
				post.Body = ""
			}
			posts = append(posts, *post)
		}
		fmt.Println(posts)
		json.NewEncoder(w).Encode(posts)

	} else {
		message = "POST/ not allowed"

	}

	w.Write([]byte(message))
}

func Completed(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var message string
	if r.Method == "GET" {
		message = "GET/ not allowed"

	} else {
		message = "POST/ not allowed"
		//Capital to maintain consistency of field naming
		type response struct {
			Post_id   uint64
			Completed bool
		}
		var data response
		err := json.NewDecoder(r.Body).Decode(&data)
		fmt.Println("post_id:", data.Post_id, "completed:", data.Completed)
		if err != nil {
			fmt.Println("Update prepare failed: ", err)
			http.Error(w, "Invalid Data", http.StatusBadRequest)
			return
		}
		stmt, err := db.Prepare("UPDATE posts SET completed=?  WHERE post_id=?")
		_, err = stmt.Exec(data.Completed, data.Post_id)
		if err != nil {
			fmt.Println("Update exec failed: ", err)
			http.Error(w, "data update failed", http.StatusInternalServerError)
			return
		}

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

func AddTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var message string
	if r.Method == "GET" {
		message = "GET/ add not allowed"
	} else {
		message = "POST/ add new task"

		var data Post
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		title := data.Title     //r.FormValue("title")
		body := data.Body       //r.FormValue("body")
		user_id := data.User_id //r.FormValue("user_id")
		completed := false
		trashed := false

		stmt, err := db.Prepare("insert into posts(title, body, user_id, completed, trashed) values(?,?,?,?,?)")
		if err != nil {
			fmt.Println("can't prepare insert")
			log.Fatal(err)
		}
		println(title, body, user_id, completed, trashed)
		_, err = stmt.Exec(title, body, user_id, completed, trashed)
		if err != nil {
			fmt.Println("can't insert prepared statement")
			log.Fatal(err)
		}
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
