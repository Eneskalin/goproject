package main

import (
	"fmt"
	"kamp/student"
)

func main() {
	var id int
	fmt.Println("Total User Count")
	studentsList := student.StudentsList()
	fmt.Printf("%d\n", len(studentsList))

	fmt.Print("Enter id of user you want to appear: ")
	fmt.Scan(&id)

	foundStudent := findStudentByID(studentsList, id)

	if foundStudent != nil {
		fmt.Printf("Student Found:\nName: %s\nSurname: %s\nID: %d\n", foundStudent.Name, foundStudent.Surname, foundStudent.ID)
	} else {
		fmt.Println("Student not found.")
	}
}

func findStudentByID(studentsList []student.Student, id int) *student.Student {
	for _, s := range studentsList {
		if s.ID == id {
			return &s
		}
	}
	return nil
}
