package main

import "fmt"

func main() {
	// Declare map so that it can be dynamically allocated
	Student := make(map[string]int)

	// rest similar to python
	Student["Will"] = 90
	fmt.Printf("%v\n", Student["Will"])
	fmt.Printf("%v\n", len(Student))
	fmt.Printf("%v\n", Student)

	Student["Greg"] = 100
	fmt.Printf("%v\n", Student["Greg"])
	fmt.Printf("%v\n", len(Student))
	fmt.Printf("%v\n", Student)

	// Remove entry
	delete(Student, "Greg")
	fmt.Printf("%v\n", Student)

	// map inside of a map
	// map of a string to map of string with value string
	// add comma at end always
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": map[string]string{
			"RealName": "Bruce WAyne",
			"City":     "Gotham City",
		},
	}

	// Reference the hero type within the "superhero" map within map
	// hero is the boolean value of whether or not superhero["Superman"] is being in the map
	// if in the map, temp gets assigned the value of superhero["Superman"]
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp)
		fmt.Println(hero)
		fmt.Println(temp["RealName"], temp["City"])
	}
}
