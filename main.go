package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/fwiedmann/evaluate-casbin/pkg/user"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policies.csv")
	if err != nil {
		panic(err)
	}

	userDB := &user.RepositoryMariaDB{}
	userService := &user.Service{Repo: userDB}
	if err := userService.Create(user.Model{
		ID:      "1",
		Name:    "Han",
		Surname: "Solo",
	}); err != nil {
		fmt.Println(err)
	}

	user, err := userService.FindById("1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
	allowed, err := e.Enforce("1", "user:delete")
	if err != nil {
		panic(err)
	}
	fmt.Println(allowed)
}
