package main

import "fmt"

type person struct {
	name string
	age uint8
}

type student struct {
	person
	class string
}

func main() {
	// Print String
	fmt.Println("Hello World")

	//Numerical
	number := 2.65
	fmt.Printf("Decimal Number: %f\n", number)
	fmt.Printf("Decimal Number: %.3f\n", number)

	//Boolean
	isExist := true
	fmt.Println("isExist value: ",isExist)

	//Variable Declaration
	const fullname = "Prasetya"
	fmt.Println("Hello", fullname)

	//Function
	firstName, secondName := swap("Prasetya", "Priyadi")
	fmt.Println(firstName, secondName)

	//Array and Loop
	var barisAngka = [5]string{"satu","dua","tiga","empat","lima"}
	for i, angka := range barisAngka {
		fmt.Println("Baris",angka,",index ke",i)
	}
	//Struct
	// List of student
	studentLists := [3]student{}
	studentLists[0].name, studentLists[0].age, studentLists[0].class = "Prasetya Priyadi", 25, "Golang Web"
	studentLists[1].name, studentLists[1].age, studentLists[1].class = "Andhika Ramdhan", 29, "Javascript Mastery"
	studentLists[2].name, studentLists[2].age, studentLists[2].class = "Prama Ramadhan", 22, "Data Science with Python"

	//Looping
	for i, student := range studentLists {
		fmt.Printf("Siswa ke %d, %s, berumur %d tahun, mengambil kelas %s \n", i+1, student.name, student.age, student.class)
	}
}

func swap(x,y string) (string,string) {
	return y, x
}