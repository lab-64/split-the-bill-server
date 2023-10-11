package wire

import "split-the-bill-server/types"

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r RegisterUser) ToUser() types.User {
	return types.CreateUser(r.Username, r.Email)
}
