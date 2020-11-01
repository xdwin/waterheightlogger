package heightlogger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dbInstance "github.com/xdwin/waterheightlogger/db"
	"gorm.io/gorm"
)

// WaterHeight is
type WaterHeight struct {
	gorm.Model
	Height int
}

// Handler is the main function for this file
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/log/read":
		data := read()
		write(w, data)
	case "/log/save":
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "ParseForm() error %v", err)
			return
		}
		height, _ := strconv.Atoi(r.FormValue("height"))
		data := save(height)
		write(w, data)
	}
}

func save(height int) WaterHeight {
	db := dbInstance.Instance
	data := WaterHeight{Height: height}

	db.Create(&data)
	return data
}

func read() WaterHeight {
	result := &WaterHeight{}

	db := dbInstance.Instance
	db.Last(result)
	return *result
}

func write(w http.ResponseWriter, data WaterHeight) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
