package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func testCopier() {
	user := User{
		Name:   "John",
		Age:    30,
		Role:   "Admin",
		Salary: 5000,
		contacts: []Contact{
			{
				contactType: "phone",
				contactInfo: "123456789",
			},
			{
				contactType: "email",
				contactInfo: "john@example.com",
			},
		},
	}

	userOut := UserOut{Salary: 150000}

	// copier.Copy(&employee, &user) // employee 裡面有15000 , 不會被蓋掉

	copier.CopyWithOption(&userOut, &user, copier.Option{DeepCopy: true})

	// copier.Copy(&userOut, &user
	// copier.Copy(&userOut, &user)
	// copier.Copy(&userOut.contactOuts, &user.contacts)

	fmt.Printf("%#v \n", userOut)

	// fmt.Printf("%#v \n", user)
	// fmt.Printf("%#v \n", employee)
}
