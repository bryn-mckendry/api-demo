package database

import (
	"apidemo/internal/domain"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var (
		err      error
		host     = os.Getenv("POSTGRES_HOST")
		port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
}

func GetAnimals() ([]domain.Animal, error) {
	rows, err := db.Query("SELECT id, name, species FROM animals")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	animals := make([]domain.Animal, 0)
	for rows.Next() {
		var a domain.Animal
		if err := rows.Scan(&a.Id, &a.Name, &a.Species); err != nil {
			return nil, err
		}
		animals = append(animals, a)
	}

	return animals, nil
}

func CreateAnimal(a domain.Animal) error {
	if _, err := db.Exec("INSERT INTO animals (name, species) VALUES ($1, $2)", a.Name, a.Species); err != nil {
		return err
	}
	return nil
}
