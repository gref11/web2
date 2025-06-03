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
	stmtUserLangDelete *sql.Stmt
	stmtUserSelect     *sql.Stmt
	stmtUserLangSelect *sql.Stmt
	stmtPasswordSelect *sql.Stmt
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

	stmtUserLangDelete, err := db.Prepare(`delete from user_lang where user_id = ?;`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserSelect, err := db.Prepare(`select name, tel, email, sex, bio from user where user_id = ? limit 1;`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtUserLangSelect, err := db.Prepare(`select lang_name from lang join user_lang on lang.lang_id = user_lang.lang_id where user_id = ?;`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	stmtPasswordSelect, err := db.Prepare(`select user_id, password_hash from user_auth where login = ? limit 1`)
	if err != nil {
		return DataBaseStorage{}, err
	}

	st := DataBaseStorage{
		db:                 db,
		stmtUserInsert:     stmtUserInsert,
		stmtUserUpdate:     stmtUserUpdate,
		stmtUserAuthInsert: stmtUserAuthInsert,
		stmtUserLangInsert: stmtUserLangInsert,
		stmtUserLangDelete: stmtUserLangDelete,
		stmtUserSelect:     stmtUserSelect,
		stmtUserLangSelect: stmtUserLangSelect,
		stmtPasswordSelect: stmtPasswordSelect,
	}

	return st, nil
}

func (st DataBaseStorage) InsertUserData(userData models.UserData) (int, error) {
	res, err := st.stmtUserInsert.Exec(userData.Name, userData.Tel, userData.Email, userData.Sex, userData.Bio)
	if err != nil {
		return 0, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, lang := range userData.Langs {
		_, err = st.stmtUserLangInsert.Exec(userID, lang)
		if err != nil {
			return 0, err
		}
	}

	return int(userID), nil
}

func (st DataBaseStorage) NewUserAuth(userID int, login, passwordHash string) error {
	if _, err := st.stmtUserAuthInsert.Exec(userID, login, passwordHash); err != nil {
		return err
	}

	return nil
}

func (st DataBaseStorage) UpdateUserData(userID string, userData models.UserData) error {
	// Начинаем транзакцию
	tx, err := st.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Обновляем основные данные пользователя
	if _, err := tx.Stmt(st.stmtUserUpdate).Exec(
		userData.Name,
		userData.Tel,
		userData.Email,
		userData.Sex,
		userData.Bio,
		userID,
	); err != nil {
		return fmt.Errorf("user update failed: %v", err)
	}

	// 2. Удаляем старые языки
	if _, err := tx.Stmt(st.stmtUserLangDelete).Exec(userID); err != nil {
		return fmt.Errorf("failed to delete user languages: %v", err)
	}

	// 3. Добавляем новые языки
	for _, lang := range userData.Langs {
		res, err := tx.Stmt(st.stmtUserLangInsert).Exec(userID, lang)
		if err != nil {
			return fmt.Errorf("failed to insert language '%s': %v", lang, err)
		}

		// Проверяем, что язык был действительно добавлен
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to check affected rows for language '%s': %v", lang, err)
		}
		if rowsAffected == 0 {
			return fmt.Errorf("language '%s' not found in lang table", lang)
		}
	}

	// Коммитим транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (st DataBaseStorage) GetUserByID(userID string) (models.UserData, error) {
	var userData models.UserData

	if err := st.stmtUserSelect.QueryRow(userID).Scan(&userData.Name, &userData.Tel, &userData.Email, &userData.Sex, &userData.Bio); err != nil {
		return models.UserData{}, err
	}

	rows, err := st.stmtUserLangSelect.Query(userID)
	if err != nil {
		return models.UserData{}, err
	}

	userData.Langs = make([]string, 0)
	for rows.Next() {
		var langName string

		if err := rows.Scan(&langName); err != nil {
			return models.UserData{}, err
		}

		userData.Langs = append(userData.Langs, langName)
	}

	if rows.Err() != nil {
		return models.UserData{}, err
	}

	return userData, nil
}

func (st DataBaseStorage) GetPasswordHash(login string) (string, string, error) {
	var user_id, password_hash string

	if err := st.stmtPasswordSelect.QueryRow(login).Scan(&user_id, &password_hash); err != nil {
		return "", "", err
	}

	return user_id, password_hash, nil
}

func (st DataBaseStorage) Close() {
	st.stmtUserAuthInsert.Close()
	st.stmtUserInsert.Close()
	st.stmtUserLangInsert.Close()
	st.stmtUserLangDelete.Close()
	st.stmtUserUpdate.Close()
	st.stmtUserSelect.Close()
	st.stmtUserLangSelect.Close()
	st.stmtPasswordSelect.Close()
}
