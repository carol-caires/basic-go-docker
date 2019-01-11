package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var connectionString = fmt.Sprintf(
	"%s:%s@tcp(%s)/%s",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASS"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_NAME"))

type Student struct {
	ID   int    `json:"id"`
	NOME string `json:"name"`
}

type Employee struct {
	ID   int    `json:"id"`
	NOME string `json:"name"`
}

func main() {

	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/employees", getEmployees)

	http.ListenAndServe(":3000", nil)
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", connectionString)
	defer db.Close()

	result, err := db.Query("SELECT * FROM students;")
	defer result.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	student := Student{}
	studentRes := []Student{}

	for result.Next() {
		var id int
		var name string

		err = result.Scan(&id, &name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		student.ID = id
		student.NOME = name

		studentRes = append(studentRes, student)
	}

	b, err := json.Marshal(studentRes)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()

	result, err := db.Query("SELECT * FROM employees;")
	defer result.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	employee := Employee{}
	employeeRes := []Employee{}

	for result.Next() {
		var id int
		var name string

		err = result.Scan(&id, &name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		employee.ID = id
		employee.NOME = name

		employeeRes = append(employeeRes, employee)
	}

	b, err := json.Marshal(employeeRes)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b))
}
