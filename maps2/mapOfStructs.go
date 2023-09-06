package main

import "fmt"

// This student struct contains: weight, empCity fields
type student struct {
	weight  float64
	empCity string
}

func main() {

	/*
		Given that we have the following map
		8. What if we want to store the student city too?
		We use the map of structs
		-----------------------------------------------
	*/

	// Let's create three students: student1, student2, student3
	student1 := student{
		weight:  74.4,
		empCity: "Huye",
	}

	student2 := student{
		weight:  55.0,
		empCity: "Nyanza",
	}

	student3 := student{
		weight:  66.4,
		empCity: "Gatsibo",
	}
	student4 := student{
		weight:  100.3,
		empCity: "Nyagatare",
	}

	// EMPLOYEES
	studentInfo := map[string]student{
		"Mark":     student1,
		"Gomez":    student2,
		"Daniella": student3,
		"Rosette":  student4,
	}

	fmt.Println(studentInfo)
	// map[Daniella:{66.4 Gatsibo} Gomez:{55 Nyanza} etc...]

	for name, info := range studentInfo {
		fmt.Printf("STD: %s , WEIGH: %f, CITY: %s\n", name, info.weight, info.empCity)
	}

}
