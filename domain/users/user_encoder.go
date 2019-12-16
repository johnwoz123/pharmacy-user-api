package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type InternalUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Role        string `json:"role"`
}

func (users Users) Encode(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, u := range users {
		result[index] = u.Encode(isPublic)
	}
	return result
}

func (user *User) Encode(isPublic bool) interface{} {
	if isPublic {
		pJson, _ := json.Marshal(user)
		var publicUserJson PublicUser
		if err := json.Unmarshal(pJson, &publicUserJson); err != nil {
			return nil
		}
		return publicUserJson
	}
	uJson, _ := json.Marshal(user)
	var internalUserJson InternalUser
	if err := json.Unmarshal(uJson, &internalUserJson); err != nil {
		return nil
	}
	return internalUserJson
}
