package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Person struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"desc"`
	Address   Address `json:"address"`
}

var Users []Person

func init() {
	Users = []Person{
		Person{FirstName: "bob", LastName: "mcgee", Address: Address{Street: "1234 W. Here", City: "San Francisco", State: "CA", Zip: "94105"}},
		Person{FirstName: "steve", LastName: "mcgerkin", Address: Address{Street: "555 W. Here", City: "Tucson", State: "AZ", Zip: "85742"}},
		{"John", "Doe", Address{"123 Main St", "New York", "NY", "12345"}},
		{"Mary", "Johnson", Address{"456 Oak St", "Los Angeles", "CA", "67890"}},
		{"Mike", "Smith", Address{"789 Pine St", "San Francisco", "CA", "24680"}},
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllUsers")
	json.NewEncoder(w).Encode(Users)
}

func CountUsersByState(w http.ResponseWriter, r *http.Request) {
	//txn := newrelic.FromContext(r.Context())

	stateCounts := make(map[string]int)

	timey := time.Now().UnixMilli()

	if timey%2 == 0 {
		http.Error(w, "Session timeout", 503)
		//txn.NoticeError(errors.New("Session timeout"))
		return
	}

	for _, person := range Users {
		state := strings.ToUpper(person.Address.State)

		if _, exists := stateCounts[state]; exists {
			stateCounts[state]++
		} else {
			stateCounts[state] = 1
		}
	}

	json.NewEncoder(w).Encode(stateCounts)
}
