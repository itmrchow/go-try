package main

type User struct {
	Name         string
	Role         string
	Age          int32
	EmployeeCode int64 `copier:"EmployeeNum"`
	Salary       int

	contacts []Contact `copier:"contact"`
}

type Contact struct {
	contactType string
	contactInfo string
}

func (user *User) DoubleAge() int32 {
	return user.Age * 2
}

type UserOut struct {
	Name string `copier:"must"`
	Age  int32  `copier:"must,nopanic"` // 如果這裡沒有複製到跳error
	// Salary     int    `copier:"-"`            //忽略
	Salary     int
	DoubleAge  int32
	EmployeeId int64 `copier:"EmployeeNum"`
	SuperRole  string

	contactOuts []ContactOuts `copier:"contact"`
}

type ContactOuts struct {
	contactType string
	contactInfo string
}

// func main() {
// 	testCopier()

// }
