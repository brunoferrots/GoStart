package main

import forl "learngo/forL"

func main() {
	//var personP *map[string]string
	person := map[string]string{
		"name":  "Bruno Ferro",
		"age":   "26",
		"email": "bruno@example.com",
	}
	//personP = &person

	forl.ForPrint(&person)

}
