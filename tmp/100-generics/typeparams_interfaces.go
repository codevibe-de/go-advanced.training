package _100_generics

import "fmt"

type Repository[T any] interface {
	Save(item T)
	FindAll() []T
}

type UserRepository struct {
	users []string
}

func (r *UserRepository) Save(user string) {
	r.users = append(r.users, user)
}

func (r *UserRepository) FindAll() []string {
	return r.users
}

func main() {
	var repo Repository[string] = &UserRepository{}
	repo.Save("Alice")
	repo.Save("Bob")
	fmt.Println(repo.FindAll()) // ["Alice", "Bob"]
}
