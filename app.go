package main

	

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "admin"
    password = "admin"
    dbname   = "api_golang"
)
// Book Struct(Model)
// type Book struct {
// 	ID 		string	`json:"id"`
// 	Isbn	string	`json:"isbn"`
// 	Title	string	`json:"title"`
// 	// Author	*Author `json:"author"`
// }

// Article Struct(Model)
type Article struct {
	Id		int  `json:"id"`
	Author	string	`json:"author"`
	Title	string	`json:"title"`
	Body	string 	`json:"body"`
	Created string 	`json:"created"`
}

// //Author struct
// type Author struct {
// 	Firstname	string	`json:"firstname"`
// 	Lastname	string	`json:"lastname"`
// }

// Init books var as a slice book struct
// var books []Book

//get All Books
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		
	// open database
	db, err := sql.Open("postgres", psqlconn)

	rows, err := db.Query(`SELECT "id", "author", "title", "body", "created" FROM "article"`)
	CheckError(err)
	
	var articles []Article

	defer rows.Close()
	for rows.Next() {
		var id int
		var author string
		var title string
		var body string
		var created string

	
		err = rows.Scan(&id, &author, &title, &body, &created)
		CheckError(err)
		articles = append(articles, Article{Id: id, Author: author, Title: title, Body: body, Created: created})
	}
	
	CheckError(err)
	json.NewEncoder(w).Encode(articles)
}
//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	//
}
//create new book
func createArticles(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		
	// open database
	db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        panic(err)
    }

	author, au_ok := r.URL.Query()["author"]
	title, ti_ok := r.URL.Query()["title"]
	body, bo_ok := r.URL.Query()["body"]

	if !au_ok || !ti_ok || !bo_ok || len(author[0]) < 1 || len(title[0]) < 1 || len(body[0]) < 1 {
			fmt.Println("error in keys %s", author)
			return
		}
	// insert
	data_author := author[0]
	data_title := title[0]
	data_body := body[0]
    // hardcoded
	insertStmt:= `insert into article (author, title, body, created) values($1, $2, $3, current_timestamp)`
    _, e := db.Exec(insertStmt, data_author, data_title, data_body)
    CheckError(e)
}
//get All Books
func updateBook(w http.ResponseWriter, r *http.Request) {
	//
}
//get All Books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	//init router
	r:= mux.NewRouter()

	//mock Data - @todo - implement DB
	// books = append(books, Book{ID: "1", Isbn: "436455", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "436444", Title: "Book two", Author: &Author{Firstname: "Samuel", Lastname: "Etoo"}})

        // connection string
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)
		
			// close database
		defer db.Close()
	
			// check db
		err = db.Ping()
		CheckError(err)
	
		fmt.Println("Connected!")
	// db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/api_golang")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// getarticle, err := db.Query("SELECT * FROM article")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer getarticle.Close()

	// for getarticle.Next() {
	// 	var article Article

	// 	err = getarticle.Scan(&article.Author)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(article.Author)
	// }
	//Route Handlers
	r.HandleFunc("/api/articles", getArticles).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/articles", createArticles).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
