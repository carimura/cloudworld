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
	time.Sleep(5 * time.Second)

	r := new(reservation)
	err := json.NewDecoder(os.Stdin).Decode(r)
	if err != nil {
		panic(err)
	}
	str := "Booking hotel for reservation " + r.ID + " checking in " + r.DepartureDate + " and checking out " + r.ReturnDate
	os.Stderr.WriteString(str)
}
