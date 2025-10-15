package main

import (
	"fmt"

	"github.com/KittiArm/golang/getuser"
	"github.com/KittiArm/golang/getuuid"
)

func main () {
	fmt.Println("Hello,", getuser.GetUser())
	fmt.Println("UUID :", uuid.GetUUID())

	// 1: set variable with type >> var <var> <type>
	var fullName string 
	// 2: set variable with type + initial value >> var <var> <type> = <value>
	var age int = 25
	// 3: set variable without type + initial value >> var <var> = <value>
	var userRole = "Software Engineer"
	// 4: set variable with initial value (shorthand) >> <var> := <value>
	userAddress := "Bangkok, Thailand"
	// var >> able to re-assign value
	fullName = "KittiArm"
	// 5: set variable with constant value (anable to re-assign value)
	const userNationality = "Thai"

	fmt.Printf("Hello %s, age = %d. %s from %s (%s).", fullName, age, userRole, userAddress, userNationality)
}