package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

	db, err := sql.Open("mysql", "root:example@tcp(172.20.0.2:3306)/dbmp")
	defer db.Close()

	result, err := db.Query("SELECT * FROM students;")
	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	student := Student{}
	studentRes := []Student{}

	for result.Next() {
		var id int
		var name string

		err = result.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
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
	db, err := sql.Open("mysql", "root:example@tcp(172.20.0.2:3306)/dbmp")
	defer db.Close()

	result, err := db.Query("SELECT * FROM employees;")
	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}
	employeeRes := []Employee{}

	for result.Next() {
		var id int
		var name string

		err = result.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
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
