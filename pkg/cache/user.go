package v1

import "sync"

type User struct {
	Id       int64
	Db       interface{}
	TotalMem int64
	UsedMem  int64
}

type UserStore interface {
	GetUser(id int64) (*User, error)
	GetUsers() (map[int64]User, error)
	CreateUser(user *User) error
}

type userStore struct {
	userMap map[int64]User
}

func (u userStore) GetUser(id int64) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userStore) GetUsers() (map[int64]User, error) {
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
				userMap: make(map[int64]User),
			}
		}
	})
	return cache
}
