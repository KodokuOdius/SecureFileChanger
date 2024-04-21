package v1

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/KodokuOdius/SecureFileChanger/db"
)

type Content struct {
	Status int
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("assets/html/home.page.html")
	if err != nil {
		log.Println(err.Error())
		fmt.Fprintln(w, "500 - internal server error")
		return
	}

	if r.URL.Path != "/" {
		ts.Execute(w, Content{Status: http.StatusNotFound})
		return
	}

	err = ts.Execute(w, Content{Status: http.StatusOK})

	if err != nil {
		log.Println(err.Error())
		return
	}
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error parsing form")
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("specify a filen anem please")
		return
	}

	f, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("something went wrong")
		return
	}
	defer f.Close()
	fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + name + fileExtension

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("file went wrong")
		return
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("copy went wrong")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("file upload")
}

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	name := r.URL.Query().Get("name")

	directory := filepath.Join("files", name)

	_, err := os.Open(directory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to open file")
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(directory))
	http.ServeFile(w, r, directory)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error parsing form")
		return
	}

	email := r.Form.Get("email")
	if email == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("specify a user email please")
		return
	}

	pass := r.Form.Get("pass")
	if pass == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("specify a user pass please")
		return
	}

	user := db.NewUser(email, pass, false)
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User not create")
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error occured during marshaling")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(string(j))
	fmt.Println(user)
}
