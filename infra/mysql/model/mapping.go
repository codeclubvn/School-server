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

func (Course) TableName() string {
	return "school__course"
}

func (Unit) TableName() string {
	return "school__unit"
}

func (Assignment) TableName() string {
	return "school__assignment"
}

func (Rubric) TableName() string {
	return "school__rubric"
}

func (RubricCriteria) TableName() string {
	return "school__rubric_criteria"
}

func (RubricCriteriaRating) TableName() string {
	return "school__rubric_criteria_rating"
}


