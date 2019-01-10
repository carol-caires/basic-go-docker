package main

import (
	"database/sql"
	"fmt"
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

	db, err := sql.Open("mysql", "root:example@tcp(172.19.0.4:3306)/dbmp")
	defer db.Close()

	result, err := db.Query("SELECT * FROM students;")
	defer result.Close()

	if err != nil {
		fmt.Printf("Deu pau :: %v", err.Error())
		return
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

	fmt.Fprintf(w, "Students :: %v!", studentRes)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:example@tcp(172.19.0.4:3306)/dbmp") // TODO :: ver porque não está conectando
	defer db.Close()

	result, err := db.Query("SELECT * FROM employees;")
	defer result.Close()

	if err != nil {
		fmt.Printf("Deu pau :: %v", err.Error())
		return
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

	fmt.Fprintf(w, "Employees :: %v!", employeeRes)
}
