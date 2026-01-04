package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	 name = "mohammad mujamil"
	 surnmae := "mohammad"
	 fmt.Println(name)
	fmt.Println(surnmae)

	//writing constants
	const (
		port = 5050
		host = "local host"
	)
	fmt.Println(port, host)

	//for loop
	i := 1
	for i <= 5 {
		fmt.Println("muz")
		i = i + 1
	}
	for j := 1; j <= 4; j++ {
		fmt.Println("this is for loop")

	}
	//writing range code

	for k := range 4 {
		fmt.Println(k)
	}
	//if-else code
	age := 20

	if age >= 20 {
		fmt.Println("person is adult")

	} else if age < 18 {
		fmt.Println("person is teenager")

	}
	//Boolean

	var role = "admin"
	var haspermission = true

	if role == "admin" || haspermission {
		fmt.Println("yes")
	}

	//we can declare a variable inside a if condition
	if agep := 20; agep > 18 {
		fmt.Println("person is younger")

	} else {
		fmt.Println("person is a child")
	}

	switch statements

		i := 5

		switch i {
		case 1:
			fmt.Println("one bro")
		case 2:
			fmt.Println("one bro")
		case 3:
			fmt.Println("one bro")
		default:
			fmt.Println("other ")
		}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("stop")
	}

	//type switch
	who := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Print("its an  integer", v)
		case string:
			fmt.Print("its an string", v)
		case float32:
			fmt.Print("ita float", v)
		default:
			fmt.Print("other")
		}

	 }
	 
	 who("mujju")

}
