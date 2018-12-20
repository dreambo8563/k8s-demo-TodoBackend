package model

// User struct
type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	ID       string `json:"id,omitempty"`
}

//GetID - get uid
func (u *User) GetID() string {
	return u.ID
}

//GetName -
func (u *User) GetName() string {
	return u.Name
}
