package customer

var Name = "Customer Package"

// unexported struct
type Secret struct {
    name string
    age  int
}

// we can use getter or setter to access unexported fields

func (p Secret) GetName() string {
	return p.name
}

func (p *Secret) SetName(name string) {
	p.name = name
}	

func (p Secret) GetAge() int {
	return p.age
}

func (p *Secret) SetAge(age int) {
	p.age = age
}