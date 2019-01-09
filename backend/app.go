package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/employees", getEmployees)

	http.ListenAndServe(":3000", nil)
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:example@tcp(172.18.0.4:3306)/dbmp") // TODO :: ver porque não está conectando
	defer db.Close()

	students, err := db.Query("SELECT * FROM students")
	defer students.Close()

	if err != nil {
		fmt.Printf("Deu pau :: %v", err.Error())
		return
	}

	fmt.Fprintf(w, "Students :: %v!", students)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "mydb:senhamydb@/mydb")
	defer db.Close()

	if err != nil {
		fmt.Println("Deu pau na conexão")
		return
	}

	students, err := db.Query("SELECT * FROM students")
	defer students.Close()

	if err != nil {
		fmt.Println("Deu pau na query")
		return
	}

	fmt.Fprintf(w, "Students :: %v!", students)
}
