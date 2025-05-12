package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	"html/template"

	// "io"
	"log"
	"net/http"
	"net/http/cgi"
	"net/url"

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

type formData struct {
	Name   string
	Email  string
	Tel    string
	Date   string
	Sex    string
	Langs  []string
	Bio    string
	Errors map[string]string
}

var (
	nameRegex  = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ ]{3,150}$`)
	telRegex   = regexp.MustCompile(`^\+[0-9]{11}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	cookieStoragePeriod = 365 * 24 * 60 * 60
)

func in(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

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

func clearErrorCookies(w http.ResponseWriter) {
	cookiesToClear := []string{
		"name_error", "email_error", "tel_error",
		"name_value", "email_value", "tel_value",
	}

	for _, name := range cookiesToClear {
		setCookie(w, name, "", -1)
	}
}

func setCookie(w http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     "/",
		HttpOnly: true,
	})
}

func loadFormData(r *http.Request, data *formData) {
	for _, cookie := range r.Cookies() {
		switch cookie.Name {
		case "name_error":
			data.Errors["Name"], _ = url.QueryUnescape(cookie.Value)
		case "email_error":
			data.Errors["Email"], _ = url.QueryUnescape(cookie.Value)
		case "tel_error":
			data.Errors["Tel"], _ = url.QueryUnescape(cookie.Value)
		case "name_value":
			data.Name, _ = url.QueryUnescape(cookie.Value)
		case "email_value":
			data.Email, _ = url.QueryUnescape(cookie.Value)
		case "tel_value":
			data.Tel, _ = url.QueryUnescape(cookie.Value)
		case "sex_value":
			data.Sex = cookie.Value
		case "date_value":
			data.Date = cookie.Value
		case "langs_value":
			langs, _ := url.QueryUnescape(cookie.Value)
			data.Langs = strings.Split(langs, ",")
		case "bio_value":
			data.Bio, _ = url.QueryUnescape(cookie.Value)
		}
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(".", "static", r.URL.Path[1:])

	if !strings.HasPrefix(filepath.Clean(path), filepath.Clean(filepath.Join(".", "static"))) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	http.ServeFile(w, r, path)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	savedData := formData{
		Errors: make(map[string]string),
	}
	funcMap := template.FuncMap{
		"in": in,
	}

	loadFormData(r, &savedData)

	tmpl := template.Must(template.New("form.html").Funcs(funcMap).ParseFiles("static/src/template/form.html"))
	tmpl.Execute(w, savedData)
}

func appFormParse(r *http.Request) (appForm, error) {
	err := r.ParseForm()
	if err != nil {
		return appForm{}, err
	}

	var appF appForm
	appF.Name = r.FormValue("fullName")
	appF.Email = r.FormValue("email")
	appF.Tel = r.FormValue("phone")
	appF.Date = r.FormValue("date")
	appF.Sex = r.FormValue("sex")
	appF.Langs = r.Form["lang"]
	appF.Bio = r.FormValue("bio")

	return appF, nil
}

func userAddHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	clearErrorCookies(w)
	// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// r.Header.Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "POST" {
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
		return
	}

	// if r.Header.Get("Content-Type") != "application/json" {
	// 	http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
	// 	return
	// }

	// var appF appForm
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&appF); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	appF, err := appFormParse(r)
	if err != nil {
		http.Error(w, "Ошибка декодирования", http.StatusBadRequest)
		return
	}

	hasError := false

	if !nameRegex.MatchString(appF.Name) {
		setCookie(w, "name_error", "Имя может содержать только символы латиницы, кириллицы и пробелы, и не может превышать 150 символов", cookieStoragePeriod)
		hasError = true
	}
	setCookie(w, "name_value", appF.Name, cookieStoragePeriod)

	if !telRegex.MatchString(appF.Tel) {
		setCookie(w, "tel_error", "Телефон должен быть формата +01234567890", cookieStoragePeriod)
		hasError = true
	}
	setCookie(w, "tel_value", appF.Tel, cookieStoragePeriod)

	if !emailRegex.MatchString(appF.Email) {
		setCookie(w, "email_error", "Email должен быть формата adress@domen.tld", cookieStoragePeriod)
		hasError = true
	}
	setCookie(w, "email_value", appF.Email, cookieStoragePeriod)

	setCookie(w, "sex_value", appF.Sex, cookieStoragePeriod)

	setCookie(w, "date_value", appF.Date, cookieStoragePeriod)

	setCookie(w, "langs_value", strings.Join(appF.Langs, ","), cookieStoragePeriod)

	setCookie(w, "bio_value", appF.Bio, cookieStoragePeriod)
	appF.Bio = convertToMnemonic(appF.Bio)

	if hasError {
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
		return
	}

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
	http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		formHandler(w, r)
	case "/user/add/":
		userAddHandler(w, r)
	default:
		staticHandler(w, r)
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("PANIC: %v", err)
			fmt.Println("Status: 500 Internal Server Error")
		}
	}()
}
