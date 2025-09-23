package database

import (
	"fmt"

	"libraryes/struct"

	"github.com/jmoiron/sqlx"
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
	DB 		*sqlx.DB

}
/*
type Book struct {
	id 			int    `db:"id"`
	title 		string `db:"title"`
	author		string `db:"author"`
	pages 		int	   `db:"pages"`
	readed 		bool   `db:"readed"`
	
	timeadd 	time.Time	`db:"timeadd"`
	timeread	time.Time	`db:"timeread"`
} 
*/

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
		id SERIAL, 
		title VARCHAR(50) NOT NULL PRIMARY KEY, 
		author VARCHAR(50) NOT NULL, 
		pages VARCHAR(50) NOT NULL,
		readed BOOLEAN DEFAULT FALSE,
		timeADD TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		timeREAD TIMESTAMP)`)
		if _, err := db.Exec(sql); err != nil {
			fmt.Println(err)
			return nil, err
		}


		return db, nil
}

func (p *Postgres) DBInsertBooks(title string, author string, pages int) {
	if _, err := p.DB.Exec(`
	INSERT INTO books (title,author,pages)
	VALUES($1, $2, $3)
	`,title, author, pages); err != nil {
		fmt.Print("error insert to books: ", err)
	}
	
}

func (p *Postgres) DBReadBook(title string) {

	
	
	if _, err := p.DB.Exec(`
	UPDATE books SET 
	readed = $2,
	timeread = CURRENT_TIMESTAMP
	WHERE title = $1
	`,title, true); err != nil {
		fmt.Println("error update readed book: ", err)
	}

}

func (p *Postgres) DBExportBooks() (map[string]str.Book, error){
	books := []str.Book{}
	err := p.DB.Select(&books, "SELECT * FROM books")
	tempMap := make(map[string]str.Book)
	tmp := str.Book{}
	for i, v := range books {
		tmp = books[i] 
		tempMap[tmp.Title] = v
	}
	return tempMap, err
  
	
}