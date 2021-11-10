package group

type Member struct {
	Name string
	Age  int
}

func NewMember(name string, age int) *Member {
	return &Member{Name: name, Age: age}
}
