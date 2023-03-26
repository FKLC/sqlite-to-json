package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "modernc.org/sqlite"
)

func fatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Arg[1] = db path, Arg[2] = query
	if len(os.Args) != 3 {
		log.Println("Missing arguments! Please provide database's path and a query.")
		os.Exit(1)
	}

	db, err := sql.Open("sqlite", os.Args[1])
	fatalErr(err)
	defer db.Close()

	record, err := db.Query(os.Args[2])
	fatalErr(err)
	defer record.Close()

	columns, err := record.Columns()
	fatalErr(err)
	row := make([]interface{}, len(columns))
	for i := 0; i < len(columns); i++ {
		row[i] = new(sql.RawBytes)
	}
	rowString := make([]string, len(columns))

	for record.Next() {
		record.Scan(row...)
		fatalErr(err)

		for i := 0; i < len(columns); i++ {
			if col, ok := row[i].(*sql.RawBytes); ok {
				rowString[i] = string(*col)
			} else {
				log.Fatalln("Cannot convert column into string")
			}
		}

		j, err := json.Marshal(&rowString)
		fatalErr(err)
		fmt.Println(string(j))
	}

}
