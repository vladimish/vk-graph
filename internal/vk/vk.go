package vk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/vladimish/vk-graph/internal/models"
	"log"
	"strconv"
)

type VK struct {
	vk *api.VK
}

func NewVK(token string) *VK {
	vk := api.NewVK(token)
	return &VK{vk}
}

func (v VK) GetFriends(userId int) ([]models.Friend, error) {
	fp := params.NewFriendsGetBuilder().UserID(userId)
	friends, err := v.vk.FriendsGet(fp.Params)
	if err != nil {
		log.Fatal(err)
	}

	f, err := v.GetPages(friends.Items)
	if err != nil {
		// Maybe it could be a good idea to not throw out already scalped data
		return nil, err
	}

	return f, nil
}

func (v VK) GetPages(userIds []int) ([]models.Friend, error) {
	// Convert userIds to string array
	ids := make([]string, len(userIds))
	for i := range userIds {
		ids[i] = strconv.Itoa(userIds[i])
	}

	p := params.NewUsersGetBuilder().UserIDs(ids)
	pages, err := v.vk.UsersGet(p.Params)
	if err != nil {
		log.Fatal(err)
	}

	return toFriends(pages), err
}

func toFriends(f []object.UsersUser) []models.Friend {
	res := make([]models.Friend, len(f))
	for i := range f {
		res[i] = models.Friend{
			ID:        f[i].ID,
			FirstName: f[i].FirstName,
			LastName:  f[i].LastName,
			Sex:       f[i].Sex,
			PhotoURL:  f[i].PhotoID,
		}
	}

	return res
}

func (v VK) GetFriendsIds(userId int) ([]int, error) {
	panic("not implemented")
}
