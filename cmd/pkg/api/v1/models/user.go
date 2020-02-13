package models

type UserStatus int

const (
	ActiveUserStatus   UserStatus = 0
	InactiveUserStatus UserStatus = 1
)

type User struct {
	Id         int64      `json:"id"`
	Username   string     `json:"username"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Phone      string     `json:"phone"`
	UserStatus UserStatus `json:"userStatus"`
}
