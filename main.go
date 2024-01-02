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

func max(num1, num2 int) int {

	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	max(1, 2)
	fmt.Println("a")

}
