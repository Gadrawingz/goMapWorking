package main

import "fmt"

func main() {

	// 1.a How to declare a Map?
	// i)Using var: var mapX = map[TypeOfKey] TypeOfValue { key1:value1, key2:value2,...}
	// ii)Shorthand: mapX := map[TypeOfKey] TypeOfValue { key1:value1, key2:value2,...}
	var englishTeams = map[int]string{1: "Arsenal", 2: "Everton", 3: "Leeds"}
	italianTeams := map[int]string{1: "Napoli", 2: "Juventis", 3: "AC Milan"}
	studentGender := map[string]string{"Gad": "Male", "Sindy": "Female", "Lee": "Male"}
	fmt.Printf("\n\a")
	fmt.Println(englishTeams)
	fmt.Println(italianTeams)
	fmt.Println(studentGender)

	// 1.b How to create a Map using map?
	// a)Syntax: make(map[type of key] type of value)

	studentMarks := make(map[string]int) // string key, int value
	fmt.Println(studentMarks)            // map[]

	// 2. Adding items to a map
	studentMarks["Gad"] = 90
	studentMarks["Dan"] = 78
	studentMarks["Wayne"] = 69
	studentMarks["Sarah"] = 88
	// Print out the map contents!
	fmt.Println(studentMarks) // map[Dan:78 Gad:90 Sarah:88 Wayne:69]

	// 3. Initialize a map during declaration itself.
	traineeMarks := map[string]int{
		"Bryant": 100,
		"Clara":  95,
		"Iggy":   75,
	}
	fmt.Println(traineeMarks) // map[Bryant:100 Clara:95 Iggy:75]

	// Later, A new element with key Jonas is added!
	traineeMarks["Jonas"] = 82
	fmt.Println(traineeMarks) // map[Bryant:100 Clara:95 Iggy:75 Jonas:82]

	// 4. It's not necessary that only string types should be keys
	var studPos = map[int]string{1: "Abdul", 2: "Benny", 3: "Chris", 4: "Mario"}
	fmt.Println(studPos) // map[1:Abdul 2:Benny 3:Chris 4:Mario]

	// 5.a. Accessing Maps:
	employeeList := map[int]string{1: "Lee", 2: "Pete", 3: "Rob"}
	fmt.Println(employeeList[2]) // Pete (As he's located at 2nd index)!

	// 5. Retrieving value for a key from a map
	employeeHeight := map[string]float64{
		"Rebecca":    1.80,
		"Kardashian": 1.91,
		"Stephen":    1.7,
		"Keshia":     2.0,
	}
	// Let's use one of employees
	selectedEmployee := "Stephen"
	height := employeeHeight[selectedEmployee]
	fmt.Println("The height of", selectedEmployee, "is ", height)

	// 5.X. Add or Update elements
	rappersRating := map[float64]string{1.1: "Drake", 2.4: "NF", 4.5: "Nas"}
	fmt.Println(rappersRating)    // map[1.1:Drake 2.4:NF 4.5:Nas]
	rappersRating[1.1] = "Nathan" // Update1
	rappersRating[4.5] = "J.Cole" // Update2
	rappersRating[5.0] = "Eminem" // Newly Added
	fmt.Println(rappersRating)    // map[1.1:Nathan 2.4:NF 4.5:J.Cole 5:Emisnem]

	// 6. How to Check if a key exists in a map?
	// Ex: I wanna know whether a key is present in the employeeHeight map
	// Syntax: value, ok := map[key]
	newEmployee := "Jeanette"
	value, okay := employeeHeight[newEmployee]
	if okay == true {
		fmt.Println("The Height of ", newEmployee, "is", value)
		return
	}
	fmt.Println(newEmployee, "not found")

	// Code number 2!
	playerMarks := map[string]float64{
		"Angela":  34.5,
		"Bruno":   49.5,
		"Gilbert": 27.0,
		"Liliane": 50.5,
	}

	// Search for a player called Bruno in our map
	playerToFind := "Bruno"
	value1, found := playerMarks[playerToFind]
	if found == true {
		fmt.Println("Loading..................")
		fmt.Println("Marks of player: ", playerToFind, "is", value1)
		return
	}
	fmt.Println(playerToFind, "not found")

	// Check for none-existing key:
	value3, okay3 := playerMarks["James"]
	fmt.Println(value3, okay3)

	value4, okay4 := playerMarks["Blaise"]
	fmt.Println(value4, okay4)

	// 7. Iterate over all elements in a map...
	// The 'range' form the 'for loop' is used 2iterate over all elements of a map
	fmt.Printf("wwwwwwwwwwwwwwwwwwwwwwwwwwwwww\n")
	for key2, value2 := range playerMarks {
		fmt.Printf("playerMarks[%s] has %.2f\n", key2, value2)
	}

	// Deleting items from a map
	// Syntax: delete(map, key)
	fmt.Println("Before deleting: ", playerMarks)
	// Before deleting:  map[Angela:34.5 Bruno:49.5 Gilbert:27 Liliane:50.5]
	delete(playerMarks, "Angela")
	fmt.Println("Afterr deleting: ", playerMarks)
	// Before deleting:  map[Bruno:49.5 Gilbert:27 Liliane:50.5]

	// 8. How to determine the length of the map?
	// (How 2 Get Count of Map Elements?)
	// We use len function
	schoolFees := map[string]int32{
		"GSP":        200000,
		"ENDPK":      400000,
		"Gabiro":     240000,
		"Kigali TSS": 155000,
	}
	fmt.Println("Original Fees: ", schoolFees)
	updatedFees := schoolFees
	updatedFees["ENDPK"] = 340000
	updatedFees["Kigali TSS"] = 167500
	fmt.Println("Updated. Fees: ", schoolFees)

	// 9. NB: Map Equality (Maps Maps can't be compared using the ==
	// operator. The == can be only used to check if a map is nil.
	numbersMap1 := map[string]int{
		"one":   1,
		"two":   2,
		"zero":  0,
		"three": 3,
	}
	numbersMap2 := numbersMap1
	fmt.Println(numbersMap1) // map[one:1 three:3 two:2 zero:0]
	fmt.Println(numbersMap2) // map[one:1 three:3 two:2 zero:0]
	// For Comparing maps https://golangbot.com/reflection/ come in handy
	// if numbersMap1 == numbersMap2 { } this will bring errors

}
