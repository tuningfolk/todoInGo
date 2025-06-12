package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Post_id   int64  `json:"post_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	User_id   int64  `json:"user_id"`
	Completed bool   `json:"completed"`
	Trashed   bool   `json:"trashed"`
} //to accommodate fields that can be null in the database. See: https://go.dev/doc/database/querying#nullable_columns

type Env struct {
	db *sql.DB
}

func (env *Env) handleShowTasks(w http.ResponseWriter, r *http.Request) {
	ShowTasks(w, r, env.db)
}

func (env *Env) handleAddTask(w http.ResponseWriter, r *http.Request) {
	AddTask(w, r, env.db)
}

func (env *Env) handleCompleted(w http.ResponseWriter, r *http.Request) {
	Completed(w, r, env.db)
}

func accessDB() *sql.DB {

	// pointer to an sql.DB struct, which represents access to a specific database
	var db *sql.DB

	// Capture connection properties
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "todo"

	//Get a database handle

	var err error
	//initialize db variable
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	//to confirm connecting works
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("no ping")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}

func main() {
	db := accessDB()
	env := &Env{db: db}
	http.HandleFunc("/api/completed/", env.handleCompleted)
	http.HandleFunc("/trash/", Trash)
	http.HandleFunc("/edit/", Edit)
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/changepassword/", ChangePassword)
	http.HandleFunc("/api/add/", env.handleAddTask)
	http.HandleFunc("/delete/", Delete)
	http.HandleFunc("/api/posts", env.handleShowTasks)
	//listing directory, we'll think about this later
	http.Handle("/", http.FileServer(http.Dir("../static")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
