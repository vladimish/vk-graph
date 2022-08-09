package app

import (
	"fmt"
	"github.com/vladimish/vk-graph/internal/vk"
)

func NewApp(token string) *App {
	return &App{
		vk: vk.NewVK(token),
	}
}

type App struct {
	vk VK
}

func (a *App) Run(id int) error {
	user, err := a.vk.GetPages([]int{id})
	if err != nil {
		return err
	}
	fmt.Println(user[0].FirstName, user[0].LastName)

	f, err := a.vk.GetFriends(id)
	if err != nil {
		return err
	}

	for i := range f {
		fmt.Println(f[i].FirstName, f[i].LastName)
	}

	return nil
}
