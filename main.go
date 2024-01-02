package main

import (
	"fmt"
)

type UserAdmin struct {
	name    string `bson:"first_name,omitempty"`
	phone   string `bson:"phone,omitempty"`
	address string `bson:"address,omitempty"`
	cvv     string `bson:"cvv,omitempty"`
	id      string `bson:"id,omitempty"`
}

func getUser() UserAdmin {
	u := UserAdmin{
		name:    "jo",
		phone:   "6343",
		address: "aaaa",
		cvv:     "666",
	}

	return u
}

//func homeLink(w http.ResponseWriter, r *http.Request) {
//    getUser()
//
//    fmt.Fprintf(w, "Welcome home!")
//}

func main() {
	fmt.Println("a")
	go_lib.func_in_lib()
}
