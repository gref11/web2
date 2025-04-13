package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "io"
	"log"
	"net/http"
	"net/http/cgi"

	// "os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type appForm struct {
	Name  string   `json:"fullName"`
	Email string   `json:"email"`
	Tel   string   `json:"phone"`
	Date  string   `json:"date"`
	Sex   string   `json:"sex"`
	Langs []string `json:"lang"`
	Bio   string   `json:"bio"`
}

var (
	allowedGenders = map[string]bool{"male": true, "female": true, "other": true}

	allowedLanguages = map[string]bool{
		"GoLang":     true,
		"C":          true,
		"C++":        true,
		"JavaScript": true,
		"PHP":        true,
		"Python":     true,
		"Java":       true,
		"Haskel":     true,
		"Clojure":    true,
		"Prolog":     true,
		"Scala":      true,
	}

	nameRegex  = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ ]{3,150}$`)
	telRegex   = regexp.MustCompile(`^\+[0-9]{11}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func convertToMnemonic(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp")
	s = strings.ReplaceAll(s, "<", "&lt")
	s = strings.ReplaceAll(s, ">", "&gt")
	s = strings.ReplaceAll(s, "'", "&#39")
	s = strings.ReplaceAll(s, `"`, "&quot")
	return s
}

func insertAppF(db *sql.DB, appF appForm) {
	stmtApp, err := db.Prepare(`INSERT INTO application (name, tel, email, sex, bio) 
		VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmtApp.Close()

	res, err := stmtApp.Exec(appF.Name, appF.Tel, appF.Email, appF.Sex, appF.Bio)
	if err != nil {
		log.Fatal(err)
	}

	appId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	stmtAppAb, err := db.Prepare(`insert into app_ability (app_id, lang_id) select ?, lang_id from lang where lang_name = ? limit 1;`)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmtAppAb.Close()

	for _, lang := range appF.Langs {
		_, err = stmtAppAb.Exec(appId, lang)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(".", "static", r.URL.Path[1:])

	if !strings.HasPrefix(filepath.Clean(path), filepath.Clean(filepath.Join(".", "static"))) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	http.ServeFile(w, r, path)
}

func userAddHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// r.Header.Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var appF appForm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&appF); err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !nameRegex.MatchString(appF.Name) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Имя может содержать только символы латиницы, кириллицы и пробелы, и не может превышать 150 символов")
		return
	}

	if !telRegex.MatchString(appF.Tel) {
		http.Error(w, "Телефон должен быть формата +01234567890", http.StatusBadRequest)
		return
	}
	if !emailRegex.MatchString(appF.Email) {
		http.Error(w, "Email должен быть формата adress@domen.tld", http.StatusBadRequest)
		return
	}
	if !allowedGenders[appF.Sex] {
		http.Error(w, "Выбран неверный пол", http.StatusBadRequest)
		return
	}
	for _, lang := range appF.Langs {
		if !allowedLanguages[lang] {
			http.Error(w, "Выбран неверный язык", http.StatusBadRequest)
			return
		}
	}
	appF.Bio = convertToMnemonic(appF.Bio)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Данные успешно сохранены")

	user := "u69196"
	pass := "8946883"
	dbName := "u69196"
	host := "localhost"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", user, pass, host, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	insertAppF(db, appF)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		staticHandler(w, r)
	case "/user/add/":
		userAddHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	// mux.Handle("/", http.FileServer(http.Dir("./static")))
	// mux.HandleFunc("/", staticHandler)
	mux.HandleFunc("/", handler)
	// err := http.ListenAndServe(":8080", mux)
	err := cgi.Serve(mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
