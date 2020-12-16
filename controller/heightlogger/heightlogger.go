package heightlogger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	dbInstance "github.com/xdwin/waterheightlogger/db"
	"gorm.io/gorm"
)

// WaterHeight is model of waterheight that saves to DB and shown to user
type WaterHeight struct {
	gorm.Model
	Height int
}

// WaterHeightSimplified is Simplified model of WaterHeight
type WaterHeightSimplified struct {
	Height    int
	CreatedAt time.Time
}

// WaterHeightArr is model of WaterHeightSimplified in arr
type WaterHeightArr struct {
	Data []WaterHeightSimplified
}

// Handler is the main function for this file
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/log/read":
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		if from == "" && to == "" {
			data := read()
			write(w, data)
		} else if from != "" && to == "" {
			fromConcatenated := from + " 00:00:00"
			data := readFrom(fromConcatenated)
			writeArr(w, data)
		} else {
			fromConcatenated := from + " 00:00:00"
			toConcatenated := to + " 00:00:00"
			data := readFromTo(fromConcatenated, toConcatenated)
			writeArr(w, data)
		}
	case "/log/save":
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "ParseForm() error %v", err)
			return
		}
		hhgt := r.Header.Get("Header")
		fmt.Println(hhgt)
		height, err := strconv.Atoi(hhgt)
		fmt.Println(err)
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

func readFrom(from string) WaterHeightArr {
	db := dbInstance.Instance
	result := WaterHeightArr{}
	var waterheights []WaterHeight

	db.Where("created_at >= ?", from).Find(&waterheights)

	for _, data := range waterheights {
		waterHeightSimplified := WaterHeightSimplified{Height: data.Height, CreatedAt: data.CreatedAt}
		result.Data = append(result.Data, waterHeightSimplified)
	}
	return result
}

func readFromTo(from string, to string) WaterHeightArr {
	db := dbInstance.Instance
	result := WaterHeightArr{}
	var waterheights []WaterHeight

	db.Where("created_at BETWEEN ? AND ?", from, to).Find(&waterheights)

	for _, data := range waterheights {
		waterHeightSimplified := WaterHeightSimplified{Height: data.Height, CreatedAt: data.CreatedAt}
		result.Data = append(result.Data, waterHeightSimplified)
	}
	return result
}

func write(w http.ResponseWriter, data WaterHeight) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&data)
}

func writeArr(w http.ResponseWriter, data WaterHeightArr) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&data)
}
