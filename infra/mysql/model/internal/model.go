package internal

import "time"

type User struct {
	Id                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email                string    `gorm:"column:mail__e;type:varchar(80)"`
	Type                 string    `gorm:"column:type__c;type:varchar(18)"`
	Name                 string    `gorm:"column:name;type:varchar(80)"`
	Password             string    `gorm:"column:password__c;type:varchar(255)"`
	Status               string    `gorm:"column:status__c"`
	DepartmentName       string    `gorm:"column:departmentname__c;type:varchar(255)"`
	LastLoginAt          time.Time `gorm:"column:lastlogin_at__c;type:timestamp"`
	LastChangePasswordAt time.Time `gorm:"column:lastchangepassword_at__c;type:timestamp"`
	CreatedAt            time.Time `gorm:"column:recordcreatetime__c;autoCreateTime;type:timestamp"`
	UpdatedAt            time.Time `gorm:"column:recordupdatetime__c;autoUpdateTime;type:timestamp"`
	ExternalId           string    `gorm:"column:externalid__c;varchar(50)"`
	RecordCreateUserId   string    `gorm:"column:recordcreateuserid__c;varchar(18)"`
	RecordUpdateUserId   string    `gorm:"column:recordupdateuserid__c;varchar(18)"`
}

type UserToken struct {
	Id                 int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserId             int       `gorm:"column:user_id;index:school_usertoken_i_userid;not null"`
	Token              string    `gorm:"column:token;type:varchar(255);not null"`
	ExpireAt           time.Time `gorm:"column:expire_at;type:timestamp;not null"`
	IsDeleted          bool      `gorm:"column:is_deleted"`
	ResendTime         time.Time `gorm:"column:resend_time;type:timestamp;not null"`
	RecordCreateUserId string    `gorm:"column:record_create_user_id;type:varchar(18)"`
	RecordCreateTime   time.Time `gorm:"column:record_create_time;autoCreateTime;type:timestamp"`
	RecordUpdateUserId string    `gorm:"column:record_update_user_id;type:varchar(18)"`
	RecordUpdateTime   time.Time `gorm:"column:record_update_time;autoUpdateTime;type:timestamp"`
}

type Course struct {
	Id			int			`gorm:"column:id;primaryKey;autoIncrement"`
	Name		string		`gorm:"column:name;type:varchar(80);notnull"`
	Description	string		`gorm:"column:description"`
	CreatedAt	time.Time	`gorm:"column:created_at;autoCreateTime;type:timestamp"`
	CreatedBy	int			`gorm:"column:created_by;notnull"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;autoUpdateTime;type:timestamp"`
	UpdatedBy	int			`gorm:"column:updated_by;notnull"`
	Visible		bool		`gorm:"column:visible;type:bool;default:false"`
}

type Unit struct {
	Id			int			`gorm:"column:id;primaryKey;autoIncrement"`
	Name		string		`gorm:"column:name;notnull"`
	Description	string		`gorm:"column:description"`
	StartDate	time.Time	`gorm:"column:start_date;type:timestamp;notnull"`
	EndDate		time.Time	`gorm:"column:end_date;type:timestamp;notnull"`
	CourseId	int			`gorm:"column:course_id"`
	CreatedBy	int			`gorm:"column:created_by;notnull"`
	CreatedAt	time.Time	`gorm:"column:created_at;autoCreateTime;type:timestamp"`
	Visible		bool		`gorm:"column:visible;type:bool;default:false"`
}

type Assignment struct {
	Id				int			`gorm:"column:id;primaryKey;autoIncrement"`
	Title			string		`gorm:"column:title;type:varchar(80);notnull"`
	Description		string		`gorm:"column:description"`
	DueTime			time.Time	`gorm:"column:due_at;type:timestamp"`
	OpenAt			time.Time	`gorm:"column:open_at;type:timestamp"`
	ClosedAt		time.Time	`gorm:"column:closed_at;type:timestamp"`
	IsGraded		bool		`gorm:"column:is_graded;type:bool;default:false"`
	MaxGrade		uint		`gorm:"column:max_grade;type:uint"`
	SubmissionLimit	int			`gorm:"column:submission_limit"`
	UnitId			int			`gorm:"column:unit_id"`
	CreatedBy		int			`gorm:"column:created_by;notnull"`
	UpdatedBy		int			`gorm:"column:updated_by;notnull"`
	CreatedAt		time.Time	`gorm:"column:created_at;autoCreateTime;type:timestamp"`
	UpdatedAt		time.Time	`gorm:"column:updated_at;autoUpdateTime;type:timestamp"`
	Visibility		int			`gorm:"column:visibility"`
}

type Rubric struct {
	Id				int			`gorm:"column:id;primaryKey;autoIncrement"`
	Title			string		`gorm:"column:title;type:varchar(80);notnull"`
	TotalPoint		float64		`gorm:"column:total_point"`
	HideTotalPoint	bool		`gorm:"column:hide_total_point;type:bool;default:false"`
}

type RubricCriteria struct {
	Id				int			`gorm:"column:id;primaryKey;autoIncrement"`
	Description		string		`gorm:"column:description"`
	LongDescription	string		`gorm:"column:long_description"`
	Point			float64		`gorm:"column:point"`
	RubricId		int			`gorm:"column:rubric_id"`
	UseRangeRating	bool		`gorm:"column:use_range_rating;type:bool;default:false"`
}

type RubricCriteriaRating struct {
	Id					int			`gorm:"column:id;primaryKey;autoIncrement"`
	Title				string		`gorm:"column:title;type:varchar(80);notnull"`
	Description			string		`gorm:"column:description"`
	MinPoint			float64		`gorm:"column:min_point"`
	MaxPoint			float64		`gorm:"column:max_point"`
	RubricCriteriaId	int			`gorm:"column:rubric_criteria_id"`
	Point				float64		`gorm:"column:point;"`
}





