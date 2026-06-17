package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db, err := sql.Open("mysql", "root:12345@(127.0.0.1:3306)/go_simple?parseTime=true")

	if err != nil {
		fmt.Println(err)
		db.Ping()
	}

	query := `
	CREATE TABLE users (
    	id INT AUTO_INCREMENT,
    	username TEXT NOT NULL,
    	password TEXT NOT NULL,
    	created_at DATETIME,
    	PRIMARY KEY (id)
	);`

	_, err1 := db.Exec(query)
	if err1 != nil {
		fmt.Println(err1)
	}

	if false {

		username := "johndoe2"
		password := "secret123"
		createdAt := time.Now()
		
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(result)
		userID, err := result.LastInsertId()
		fmt.Println(userID)
	}
	
	if false {
	
	var (
		id int
		username string
		password string
		createdAt time.Time
	)

		query_get := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	
		err2 := db.QueryRow(query_get, 1).Scan(&id, &username, &password, &createdAt)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(id, username, password, createdAt)

	}

	type user struct {
		id int
		username string
		password string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u)
	}
	err2 := rows.Err()
	if err != nil {
		fmt.Println(err2)
	}

	_, err3 := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(users)



	r.HandleFunc("/books/{title}/page/{page}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// r.HandleFunc("/secure", SecureHandler).Schemes("https")
	// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// bookrouter := r.PathPrefix("/books").Subrouter()
	// bookrouter.HandleFunc("/", AllBooks)
	// bookrouter.HandleFunc("/{title}", GetBook)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))


	http.ListenAndServe(":80", r)

}