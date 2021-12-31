package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u User) Name() string {
	return u.email
}

func (u *User) Email(e string) {
	u.email = e
}

func main() {
	u1 := User{
		name: "andrew",
		email: "andrew@gmail.com",
	}
	fmt.Printf("u1: %v\n", u1)
	u1.Email("andrew@hotmail.com")
	fmt.Printf("u1: %v\n", u1)

	u2 := &User{
		name: "tony",
		email: "tony@gmail.com",
	}
	fmt.Printf("u2: %v\n", u2)
	u2.Email("tony@hotmail.com")
	fmt.Printf("u2: %v\n", u2)

	fmt.Printf("u2: %v\n", *u2)
}
