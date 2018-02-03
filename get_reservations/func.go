/*
This is a stub that mocks the return of a couple cars from a fake getcars API for testing
the getcars func.
*/
package main

import (
	"bytes"
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
type reservations struct {
	Reservations []reservation `json:"reservations"`
}

func main() {
	time.Sleep(3 * time.Second)
	r1 := newReservation("1005", "joe@yahoo.com", "Hawaii", "5/5/2018", "5/10/2018")
	r2 := newReservation("1006", "jane@hotmail.com", "Vietnam", "6/5/2018", "6/10/2018")
	r3 := newReservation("1007", "sarah@gmail.com", "Connecticut", "8/5/2018", "8/10/2018")
	r4 := newReservation("1008", "gary@dogpile.com", "Bahamas", "12/8/2018", "12/20/2018")
	r5 := newReservation("1009", "susan@mailanator.com", "Japan", "8/1/2018", "8/10/2018")
	r6 := newReservation("1010", "fred@gmail.com", "Portland", "9/20/2018", "9/22/2018")
	r7 := newReservation("1011", "phil@hotmail.com", "Bakersfield", "9/5/2018", "9/6/2018")
	r8 := newReservation("1012", "dave@yahoo.com", "Fiji", "11/11/2018", "11/18/2018")
	reservations := &reservations{
		Reservations: []reservation{*r1, *r2, *r3, *r4, *r5, *r6, *r7, *r8},
	}
	b, _ := json.Marshal(reservations)

	var buffer bytes.Buffer

	buffer.WriteString("Found " + string(len(reservations.Reservations)) + "Reservations:\n")

	for _, r := range reservations.Reservations {
		buffer.WriteString("Reservation " + r.ID + ": " + r.Email + " is going to " + r.Destination + " from " + r.DepartureDate + " to " + r.ReturnDate + "\n")
	}

	os.Stderr.WriteString(buffer.String())
	os.Stdout.Write(b)
}

func newReservation(id string, email string, destination string, departuredate string, returndate string) *reservation {
	r := new(reservation)
	r.ID = id
	r.Email = email
	r.Destination = destination
	r.DepartureDate = departuredate
	r.ReturnDate = returndate
	return r
}
