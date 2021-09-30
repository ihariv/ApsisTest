package users

import "sync"

type (
	Counter struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		SumForUsers int `json:"sum_for_users"`
		*Users
	}

	Counters struct {
		*sync.RWMutex
		Counters []Counter
		Users *Users
		SumForUsers int
	}
)

var counters = Counters{
	new(sync.RWMutex),
	[]Counter{
		{ID: "1", Name: "Yellow", Users: new(Users)},
		{ID: "2", Name: "Orange", Users: new(Users)},
		{ID: "3", Name: "Green", Users: new(Users)},
	},
	nil,
	0,
}

func GetCounters(userList *Users) *Counters {
	counters.Users = userList
	return &counters
}
