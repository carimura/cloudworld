package main

import (
	"encoding/json"
	"os"
	"time"
)

type reservation struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departing"`
	ReturnDate    string `json:"returning"`
}

func main() {
	time.Sleep(3 * time.Second)

	r := new(reservation)
	err := json.NewDecoder(os.Stdin).Decode(r)
	if err != nil {
		panic(err)
	}
	str := "Processing Reservation " + r.ID + " for " + r.Email
	os.Stderr.WriteString(str)
}
