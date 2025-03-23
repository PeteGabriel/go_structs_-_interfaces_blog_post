package structs

type Person struct {
	Name     string
	LastName string
	Age      int64
}


func InitializationByPosition() Person{
	return Person{"Jonh", "Wick", 41}
}

func InitializationByFieldName() Person{
	return Person {
		Name: "John",
		LastName: "Wick",
		Age: 41,
	}
}

func InitializationByNew() *Person {
	p := new(Person)
	p.Name = "John"
	p.LastName = "Wick"
	p.Age = 41
	return p
}

