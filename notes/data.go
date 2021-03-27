package notes

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	conn *sql.DB
}

func (db *DataBase) Init() {
	var err error
	db.conn, err = sql.Open("mysql", "root:starfield@tcp(sql:3306)/postit")
	if err != nil {
		log.Fatal("Error connecting to DB. Error: ", err)
	}

	db.conn.SetMaxOpenConns(4)
	db.conn.SetMaxIdleConns(4)
	db.conn.SetConnMaxLifetime(60 * time.Second)
}

func (db *DataBase) List(o int, l int) ([]Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := DB.conn.QueryContext(ctx, `SELECT id,
	title,
	body,
	tags
	FROM notes
	LIMIT ?, ?`, o*l, l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Close()

	ns := make([]Note, 0, l)
	for res.Next() {
		var n Note
		res.Scan(&n.ID, &n.Title, &n.Body, &n.Tags)
		ns = append(ns, n)
	}

	return ns, nil
}

func (db *DataBase) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res := DB.conn.QueryRowContext(ctx, `SELECT COUNT(*) FROM notes`)

	var c int
	err := res.Scan(&c)
	if err != nil {
		return 0, err
	}

	return c, nil
}

func (db *DataBase) Add(n Note) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := DB.conn.ExecContext(ctx, `INSERT INTO notes (
	title,
	body,
	tags) 
	VALUES (?, ?, ?)`, n.Title, n.Body, n.Tags)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *DataBase) Get(id int) (Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res := DB.conn.QueryRowContext(ctx, `SELECT id,
	title,
	body,
	tags
	FROM notes
	WHERE id = ?`, id)

	var n Note
	err := res.Scan(&n.ID, &n.Title, &n.Body, &n.Tags)
	if err == sql.ErrNoRows {
		return n, errors.New("404 Not Found")
	} else if err != nil {
		return n, err
	}

	return n, nil
}

func (db *DataBase) Update(id int, n Note) error {
	_, err := db.Get(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = DB.conn.ExecContext(ctx, `UPDATE notes SET
	title = ?,
	body = ?,
	tags = ?
	WHERE id = ?`, n.Title, n.Body, n.Tags, id)

	return err
}

func (db *DataBase) Delete(id int) error {
	_, err := db.Get(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = DB.conn.QueryContext(ctx, `DELETE FROM notes WHERE id = ?`, id)
	return err
}

var DB DataBase

func init() {
	log.Println("Connecting to Database...")
	DB.Init()
}
