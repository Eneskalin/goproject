package student

type Student struct {
	Name    string
	Surname string
	ID      int
}

func StudentsList() []Student {
	student1 := Student{
		Name:    "Enes",
		Surname: "Kalin",
		ID:      1,
	}
	student2 := Student{
		Name:    "Yigithan",
		Surname: "Sari",
		ID:      2,
	}
	student3 := Student{
		Name:    "Ahmet",
		Surname: "Yilmaz",
		ID:      3,
	}
	student4 := Student{
		Name:    "Mehmet",
		Surname: "Akinci",
		ID:      4,
	}
	student5 := Student{
		Name:    "Mustafa",
		Surname: "Uysal",
		ID:      5,
	}
	student6 := Student{
		Name:    "Turan",
		Surname: "Ozturk",
		ID:      6,
	}

	studentsList := []Student{student1, student2, student3, student4, student5, student6}

	return studentsList
}
