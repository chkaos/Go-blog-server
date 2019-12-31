package models

type UserModel struct {
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
	ID         int `json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
	Username string  `json:"username"`
	URL      string `json:"url"`
	Avatar   string `json:"avatar"`
	Role     int    `json:"role"`
	Source   int    `json:"source"`
	IsMuted  int    `json:"is_muted,omitempty"`
	Token    string  `json:"token,omitempty"`
}

func (u *UserModel) UserResponse() UserResponse {
	user := UserResponse{
		ID: u.ID,
		CreatedAt: u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Username: u.Username,
		URL:    u.URL,
		Avatar:    u.Avatar,
		Role:    u.Role,
		Source:      u.Source,
	}
	return user
}

func (u *UserModel) UserResponseWithToken(token string) UserResponse {
	user := UserResponse{
		ID: u.ID,
		CreatedAt: u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Username: u.Username,
		URL:    u.URL,
		Avatar:    u.Avatar,
		Role:    u.Role,
		Source:      u.Source,
		Token:    token,
	}
	return user
}
