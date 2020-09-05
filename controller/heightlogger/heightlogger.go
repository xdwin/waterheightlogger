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
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() error %v", err)
		return
	}
	height, _ := strconv.Atoi(r.FormValue("height"))
	data := save(height)
	write(w, data)
}

func save(value int) WaterHeight {
	db := dbInstance.Instance
	data := WaterHeight{Height: value}

	db.Create(&data)
	return data
}

func write(w http.ResponseWriter, data WaterHeight) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
