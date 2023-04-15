package handlers

import (
	"html/template"
	"net/http"
	"time"
)

type clock struct {
	ID        string
	Location  string
	Offset    int
	Formatted string
}

func updateClocks(clocks []*clock) {
	for _, c := range clocks {
		now := time.Now()
		utc := now.UTC().Add(time.Duration(now.UTC().Minute()) * time.Minute).Unix()
		loc := time.FixedZone(c.Location, c.Offset*60*60)
		time := time.Unix(utc, 0).In(loc)
		c.Formatted = time.Format("15:04:05")
	}
}

func main() {
	clocks := []*clock{
		{"new-york", "America/New_York", -5, ""},
		{"berlin", "Europe/Berlin", 1, ""},
		{"tokyo", "Asia/Tokyo", 9, ""},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		updateClocks(clocks)
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, clocks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
