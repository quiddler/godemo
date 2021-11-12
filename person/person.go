package person

import "time"

type Person struct {
	First  string    `json:"first"`
	Last   string    `json:"last"`
	Middle string    `json:"middle"`
	Phone  string    `json:"phone"`
	Dob    time.Time `json:"dob"`
}

func New() *Person {
	return &Person{
		First:  "Eliot",
		Last:   "Easterling",
		Middle: "D",
		Phone:  "234-703-9147",
		Dob:    time.Date(1982, 10, 4, 11, 30, 0, 0, time.FixedZone("EST", (4*60*60))),
	}
}
