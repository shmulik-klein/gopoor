package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func createBuyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusCreated,
			struct{ Test string }{"This is a test"})
	}
}
