package model

import (
	"elearning/infra/mysql/model/internal"
	"gorm.io/gorm"
)

type BaseModel struct {
	DeletedAt gorm.DeletedAt `gorm:"column:record_delete_time;type:timestamp"`
}

type User struct {
	internal.User
}

type UserToken struct {
	internal.UserToken
}

type Course struct {
	internal.Course
}

type Unit struct {
	internal.Unit
}

type Assignment struct {
	internal.Assignment
}

type Rubric struct {
	internal.Rubric
}

type RubricCriteria struct {
	internal.RubricCriteria
}

type RubricCriteriaRating struct {
	internal.RubricCriteriaRating
}


