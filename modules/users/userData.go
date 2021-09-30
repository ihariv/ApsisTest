package users

import "sync"

type (
	User struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		PasswordSecure string `json:"-"`
		*CounterUniq
	}

	CounterUniq struct {
		*sync.RWMutex
		Count int
	}
	Users []*User
)

var users = Users{
	{ID: "1", Name: "Lena Smith ", Email: "LSmith@gmail.com", PasswordSecure: "test1pass", CounterUniq: new(CounterUniq) },
	{ID: "2", Name: "Nicol Green ", Email: "NGreen@gmail.com", PasswordSecure: "test2pass", CounterUniq: new(CounterUniq)},
	{ID: "3", Name: "Tom Taylor ", Email: "TTaylor@gmail.com", PasswordSecure: "test3pass", CounterUniq: new(CounterUniq)},
}

func GetUsers() *Users{
	for i, _ := range users {
		users[i].RWMutex = new(sync.RWMutex)
	}
	return &users
}

func (us *Users) LoginUser(email string, password string) bool {
	for _, u := range *us{
		if u.Email == email {
			if u.PasswordSecure == password {
				return true
			}
			return false
		}
	}
	return false
}

func (us *Users) Append(user *User) bool {
	for _, u := range *us{
		if u.Email == user.Email {
			return false
		}
	}
	*us = append(*us, user)
	return false
}

func (us *Users) Remove(userId string) bool {
	for i, u := range *us{
		if u.ID == userId {
			(*us)[i] = (*us)[len(*us)-1]
			(*us) = (*us)[:len(*us)-1]
			return true
		}
	}
	return false
}

func (cu *CounterUniq) Increate(){
	if cu.RWMutex == nil {
		cu.RWMutex = new(sync.RWMutex)
	}
	cu.Lock()
	defer cu.Unlock()
	cu.Count++
}
