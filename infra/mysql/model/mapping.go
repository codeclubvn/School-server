package model

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "school__user"
}
