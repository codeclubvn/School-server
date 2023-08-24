package model

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "school__user"
}

func (UserToken) TableName() string {
	return "school__user_token"
}
