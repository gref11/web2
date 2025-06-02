package storage

import (
	"database/sql"
	"fmt"

	"web3/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type DataBaseStorage struct {
	db                 *sql.DB
	stmtUserInsert     *sql.Stmt
	stmtUserUpdate     *sql.Stmt
	stmtUserAuthInsert *sql.Stmt
	stmtUserLangInsert *sql.Stmt
}

func NewDataBaseStorage() (DataBaseStorage, error) {
	User := "u69196"
	pass := "8946883"
	dbName := "u69196"
	host := "localhost"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", User, pass, host, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserInsert, err := db.Prepare(`INSERT INTO user (name, tel, email, sex, bio) 
		VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserUpdate, err := db.Prepare(`UPDATE user
		SET name = ?,
    	tel = ?,
    	email = ?, 
    	sex = ?,
    	bio = ?
		WHERE user_id = ?;`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserAuthInsert, err := db.Prepare(`insert into user_auth(user_id, login, password_hash) values (?, ?, ?);`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserLangInsert, err := db.Prepare(`insert into user_lang (user_id, lang_id) select ?, lang_id from lang where lang_name = ? limit 1;`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	st := DataBaseStorage{
		db:                 db,
		stmtUserInsert:     stmtUserInsert,
		stmtUserUpdate:     stmtUserUpdate,
		stmtUserAuthInsert: stmtUserAuthInsert,
		stmtUserLangInsert: stmtUserLangInsert,
	}

	return st, nil
}

func (st DataBaseStorage) InsertUserData(userData models.UserData) error {
	res, err := st.stmtUserInsert.Exec(userData.Name, userData.Tel, userData.Email, userData.Sex, userData.Bio)
	if err != nil {
		return err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for _, lang := range userData.Langs {
		_, err = st.stmtUserLangInsert.Exec(userID, lang)
		if err != nil {
			return err
		}
	}

	return nil
}

func (st DataBaseStorage) NewUserAuth(userID int, login, passwordHash string) error {
	if _, err := st.stmtUserAuthInsert.Exec(userID, login, passwordHash); err != nil {
		return err
	}

	return nil
}

func (st DataBaseStorage) UpdateUserData(userID int, userData models.UserData) error {
	if _, err := st.stmtUserUpdate.Exec(userData.Name,
		userData.Tel, userData.Email, userData.Sex, userData.Bio, userID); err != nil {
		return err
	}

	return nil
}

func (st DataBaseStorage) Close() {
	st.stmtUserAuthInsert.Close()
	st.stmtUserInsert.Close()
	st.stmtUserLangInsert.Close()
	st.stmtUserUpdate.Close()
}
