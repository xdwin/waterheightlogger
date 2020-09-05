package controller

import (
	"fmt"
	"net/http"

	"github.com/xdwin/waterheightlogger/db"

	db "github.com/xdwin/waterheightcontroller/db"
	"gorm.io/gorm"
)

type height struct {
	gorm.Model
	Height    int
	CreatedAt string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() error %v", err)
		return
	}
	height := r.FormValue("height")
}

func save(height int) {
	instance := db.Instance
}

// Test is
func Test(w http.ResponseWriter, r *http.Request) {
	db.Save()
}
