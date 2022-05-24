package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("Full name :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println()
}

const (
	HOST    = "localhost"
	PORT    = 5432
	USER    = "postgres"
	PASS    = "Indra19"
	DB_NAME = "sesi07"
)

var (
	db  *sql.DB
	err error
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	//Insert new employee
	// emp := Employee{
	// 	FullName: "Haby",
	// 	Email:    "haby@koinworks.com",
	// 	Age:      19,
	// 	Division: "Developer",
	// }

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	//Get all employees
	// employees, err := getAllEmployees(db)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// for _, e := range *employees {
	// 	e.Print()
	// }

	// //Get employee by Id
	// fmt.Println(strings.Repeat("=", 10), "Get Employee By ID 2", strings.Repeat("=", 10))
	// employee, err := getEmployeeById(db, 2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// employee.Print()

	//Update employee by Id
	emp := Employee{
		FullName: "Indra",
		Email:    "indra@koinworks.com",
		Age:      29,
		Division: "IT",
	}
	err = updateEmployeeById(db, 2, &emp)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	//Delete employee by Id
	// err = deleteEmployeeById(db, 1)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode= disable",
		HOST, PORT, USER, PASS, DB_NAME)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// fmt.Println("Successfully connected to database")
	// best practice, limit database connections or called connections pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
	SELECT id, full_name, email, age, division
	FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	employees := []Employee{}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	for rows.Next() {
		employee := Employee{}
		err := rows.Scan(
			&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return &employees, nil
}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `SELECT id, full_name, email, age, division
	FROM employees
	WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division)
	if err != nil {
		return nil, err
	}

	return &emp, nil

}

func updateEmployeeById(db *sql.DB, id int, emp *Employee) error {
	query := `UPDATE employees
	SET full_name=$2, email=$3, age=$4, division=$5
	WHERE id=$1`

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, emp.FullName, emp.Email, emp.Age, emp.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}
func deleteEmployeeById(db *sql.DB, id int) error {
	query := `DELETE from employees WHERE id=$1`

	res, err := db.Exec(query, 1)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Delete %d successfully\n", count)
	return nil
}
