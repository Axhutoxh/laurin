package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	
)

// Person type
type Person struct	{
	id int;
	name string;
}

func main()	{
	var p1 []Person = []Person	{
		{id: 3, name: "ashutosh"},
		{id: 4, name: "akash"},
		{id: 5, name: "sharad"},
	}

	var tempPerson Person;

	db, _ := sql.Open("sqlite3", "./person.db")
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS PERSON (ID INTEGER PRIMARY KEY, NAME TEXT)")
	stmt.Exec()

	for _, p := range p1	{
		stmt, _ = db.Prepare("INSERT INTO PERSON (NAME) VALUES (?)")
		stmt.Exec(p.name)
	}

	rows, _ := db.Query("SELECT ID, NAME FROM PERSON")

	for rows.Next()	{
		rows.Scan(&tempPerson.id, &tempPerson.name)

		fmt.Println(tempPerson);
	}
}
