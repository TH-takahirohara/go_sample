package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {
	var id int
		var name string
		arr := []struct {
			Id   int
			Name string
		}{}
		rows, err := app.db.Query("select id, name from users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, struct {
				Id   int
				Name string
			}{id, name})
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		jsonData, err := json.Marshal(arr)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(jsonData)
}

func (app *application) postUser(w http.ResponseWriter, r *http.Request) {
	// JSONリクエストの読み込み
		// 参考：https://www.twihike.dev/docs/golang-web/json-request
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "application/json") {
			http.Error(w, "send in JSON format", http.StatusUnsupportedMediaType)
			return
		}

		var ns struct{ Name string `json:"name"` }
		err := json.NewDecoder(r.Body).Decode(&ns)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid json: %v", err), http.StatusBadRequest)
			return
		}

		// POSTされたNameのuserを登録する
		// 参考：http://go-database-sql.org/modifying.html
		stmt, err := app.db.Prepare("INSERT INTO users(name) VALUES(?)")
		if err != nil {
			http.Error(w, fmt.Sprintf("server error: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(ns.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("server error: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
}