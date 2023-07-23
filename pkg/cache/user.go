package v1

import "sync"

type User struct {
	Id       string
	Db       interface{}
	TotalMem int64
	UsedMem  int64
}

type UserStore interface {
	GetUser(id string) (*User, error)
	GetUsers() (map[string]User, error)
	CreateUser(user *User) error
}

type userStore struct {
	userMap map[string]User
}

func (u userStore) GetUser(id string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userStore) GetUsers() (map[string]User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userStore) CreateUser(user *User) error {
	//TODO implement me
	panic("implement me")
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
	return cache
}
