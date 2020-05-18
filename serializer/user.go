package serializer

import "giligili/model"

// User 用户序列化器
type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"user_name"`
	Nickname   string `json:"nickname"`
	Status     string `json:"status"`
	Avatar     string `json:"avatar"`
	CreatedAt  int64  `json:"created_at"`
	Permission uint   `json:"permission"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:         user.ID,
		UserName:   user.UserName,
		Nickname:   user.Nickname,
		Status:     user.Status,
		Avatar:     user.AvatarURL(),
		CreatedAt:  user.CreatedAt.Unix(),
		Permission: user.Permission,
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
