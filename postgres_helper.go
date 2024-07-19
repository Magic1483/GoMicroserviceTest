package main

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

type datascore struct {
	db *sql.DB
}

type Message struct {
	Id 			int
	Text 		string
	IsChecked	bool
}



func CreateTable(db *sql.DB)  {
	_,err := db.Exec("CREATE TABLE IF NOT EXISTS messages"+
	" (id SERIAL PRIMARY KEY,text TEXT NOT NULL, isChecked BOOL NOT NULL)")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("[postgres helper] table created\n")
	}
}

func (d *datascore) InsertMessage(text string) error  {

	exec := fmt.Sprintf("INSERT INTO messages (text,isChecked) VALUES ('%s',false);",text)
	_,err := d.db.Exec(exec)
	return err
}

func (d *datascore) GetMessages() ([]Message , error)  {
	var m []Message
	var (
		id int
		text string
		isChecked bool
	)
	
	rows,err := d.db.Query("select id,text,ischecked from messages;")

	if err != nil {
		return nil,fmt.Errorf("failed to get messages")
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id,&text,&isChecked)
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, Message{id,text,isChecked})
	}
	return m,nil

	
}

// https://pkg.go.dev/github.com/lib/pq
func (s *server) setupDB() {
	var db *sql.DB

	config,_ := toml.LoadFile("conf.toml")
	pg_pass := config.Get("postgres_pass").(string)

	connStr := "user=postgres password="+pg_pass+" dbname=db sslmode=disable host=db"
	db,err := sql.Open("postgres",connStr)

	if err != nil {
		log.Fatal(err)
	}
	CreateTable(db)

	s.db = &datascore{
		db: db,
	}

}