package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// what we learn is CRUD using sqlx or sql
// in db. We should use begin tran, rollback, commit
// 

type Users struct {
	Id           int       `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}


var db *sqlx.DB

// var db *sql.DB


func main() {
	var err error
	// db, err = sql.Open("mysql", "root:admin1234@tcp(127.0.0.1:3306)/sample_app?parseTime=true")
	db, err = sqlx.Open("mysql", "root:admin1234@tcp(127.0.0.1:3306)/sample_app?parseTime=true")
	if err != nil {
		panic(err)
	}

	// GetUsers()
	// users, err := GetUsersX()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	// GetUserById()
	// firstUser, err := GetUserById(1)
	// 	if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(firstUser)

	// AddUser()
	// user := Users{Username: "james_wreck", Email: "james@example.com",PasswordHash:"hashed_password_789"}
	// if(err != nil){
	// 	panic(err)
	// }	
	// AddUser(user);

	// UpdateUser()
	// user := Users{Username: "john_smith", Email: "john_smith@example.com", PasswordHash:"hashed_password_666", Id: 2}
	// if(err != nil){
	// 	panic(err)
	// }
	// UpdateUser(user);

	// DeleteUserById()
	// err = DeleteUserById(3)
	// if(err != nil){
	// 	panic(err)
	// }

}

func GetUsers() ([]Users, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, username, email, password_hash, created_at  from users" // Go is strictly positional when scanning rows, If someone later changes the table or column order change it will break.
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []Users{}
	for rows.Next() {
		user := Users{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUsersX() ([]Users, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, username, email, password_hash, created_at  from users" 
	users := []Users{}
	err = db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id int) (*Users, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	user := Users{}
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &user, nil

}

func GetUserXById(id int) (*Users, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	user := Users{}
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE id = ?"
	err = db.Get(&user, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &user, nil

}

func AddUser(user Users) error {
	tx, err := db.Begin()
	if err != nil {
    	return err
	}

	query := `
			INSERT INTO users (username, email, password_hash)
			VALUES (?, ?, ?)
			`
	result, err := tx.Exec(query,
		user.Username,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
    	return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback() // we can rollback
		return err
	}
	if affected <= 0 {
		return errors.New("cannot insert")
	}

	err = tx.Commit()
	if err != nil {
    	return err
	}

	return nil
}

func UpdateUser(user Users) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE users
		SET username = ?, email = ?, password_hash = ?
		WHERE id = ?
	`

	result, err := db.Exec(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Id,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteUserById(id int) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	query := `
		delete from users
		where id = ?
	`

	result, err := db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

