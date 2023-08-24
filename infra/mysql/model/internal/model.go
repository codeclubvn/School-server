package internal

import "time"

type User struct {
	Id                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email                string    `gorm:"column:userid__c;type:varchar(80)"`
	Type                 string    `gorm:"column:type__c;type:varchar(18)"`
	Name                 string    `gorm:"column:name;type:varchar(80)"`
	IsMiddleManager      bool      `gorm:"column:ismiddlemanager__c"`
	IsProxy              bool      `gorm:"column:isproxy__c"`
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
