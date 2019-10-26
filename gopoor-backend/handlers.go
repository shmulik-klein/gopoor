package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/unrolled/render"
)

type purchase struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func PurchaseHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		db, err := sql.Open("mysql", "root:D4exufru@/gopoor?charset=utf8")
		checkErr(err)
		stmt, err := db.Prepare("INSERT purchases SET name=?, price=?")
		checkErr(err)
		if req.Body == nil {
			http.Error(w, err.Error(), 400)
		}
		var p purchase
		err = json.NewDecoder(req.Body).Decode(&p)
		_, err = stmt.Exec(p.Name, p.Price)
		checkErr(err)
		w.Header().Set("Location", "/purchases/")
		formatter.JSON(w, http.StatusCreated, p)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
