package v1

import (
	"errors"
	"sync"
)

type User struct {
	Id       string
	Db       []interface{}
	TotalMem int64
	UsedMem  int64
}

type UserStore interface {
	GetUser(id string) User
	GetUsers() map[string]User
	CreateUser(user User) error
}

type userStore struct {
	userMap map[string]User
}

func (u *userStore) GetUser(id string) User {
	return u.userMap[id]
}

func (u *userStore) GetUsers() map[string]User {
	return u.userMap
}

func (u *userStore) CreateUser(user User) error {
	if _, ok := u.userMap[user.Id]; ok {
		return errors.New("user already exists")
	}
	u.userMap[user.Id] = user
	return nil
}

var cache userStore

func GetUserStore() UserStore {
	once := sync.Once{}
	once.Do(func() {
		if cache.userMap == nil {
			cache = userStore{
				userMap: make(map[string]User),
			}
		}
	})
	return &cache
}
