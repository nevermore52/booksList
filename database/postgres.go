package database

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host 		string
	Port	 	string
	Username 	string
	Password	string
	DBName		string
	SSLMode 	string
}

type Postgres struct {
	db 		*sqlx.DB
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", config)
		if err != nil {
			fmt.Println("error to open postgres db")
			return nil, err
		}

		err = db.Ping()
		if err != nil{
			return nil, err
		}

		sql := (` CREATE TABLE IF not exists books(
		id INT primary KEY, 
		title VARCHAR(50) NOT NULL, 
		author VARCHAR(50) NOT NULL, 
		readed BOOLEAN DEFAULT FALSE,
		timeADD TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`)
		if _, err := db.Exec(sql); err != nil {
			fmt.Println(err)
			return nil, err
		}


		return db, nil
}

func (p *Postgres) DBInsertBooks(title string, author string, pages int) {
	if _, err := p.db.Exec(`
	INSERT INTO books (title,author,pages)
	VALUES(хуй, пизда, головка)
	`); err != nil {
		fmt.Print("error insert to books", err)
	}
	
}
