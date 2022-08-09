package app

import "github.com/vladimish/vk-graph/internal/models"

type VK interface {
	GetFriends(userId int) ([]models.Friend, error)
	GetPages(userIds []int) ([]models.Friend, error)
}
