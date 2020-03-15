package models

import "time"

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Avatar   string `json:"avatar"`
	Role     int    `json:"role"`
	Source   int    `json:"source"`
	IsMuted  int    `json:"is_muted"`
}

type UserResponse struct {
	ID         int        `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Username   string     `json:"username"`
	URL        string     `json:"url"`
	Avatar     string     `json:"avatar"`
	Role       int        `json:"role"`
	Source     int        `json:"source"`
	IsMuted    int        `json:"is_muted"`
	Token      string     `json:"token"`
}

type UsersSerializer struct {
	Users []*User
}

func (u *User) Response() UserResponse {
	user := UserResponse{
		ID:         u.ID,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Username:   u.Username,
		URL:        u.URL,
		Avatar:     u.Avatar,
		Role:       u.Role,
		Source:     u.Source,
	}
	return user
}

func (u *User) ResponseWithToken(token string) UserResponse {
	user := UserResponse{
		ID:         u.Model.ID,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Username:   u.Username,
		URL:        u.URL,
		Avatar:     u.Avatar,
		Role:       u.Role,
		Source:     u.Source,
		Token:      token,
	}
	return user
}

func (s *UsersSerializer) Response() []UserResponse {
	var users []UserResponse
	for _, user := range s.Users {
		users = append(users, user.Response())
	}
	return users
}
