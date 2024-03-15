package entities

import "time"

type Ci_Class struct {
	Id                uint      `json:"id" gorm:"primarykey"`
	Id_Superclass     int       `json:"id_superclass"`
	Namespace         string    `json:"name_space" gorm:"type:varchar(255)"`
	Class_Name        string    `json:"class_name" gorm:"type:varchar(255)"`
	Status            string    `json:"status" gorm:"type:varchar(255)"`
	Description       string    `json:"description" gorm:"type:text"`
	Custom_Properties string    `json:"custom_properties" gorm:"type:text"`
	Final             bool      `json:"final" gorm:"type:tinyint(1)"`
	Singleton         bool      `json:"singleton" gorm:"type:tinyint(1)"`
	CreatedAt         time.Time `json:"created_at" gorm:"type:datetime"`
	CreatedBy         int       `json:"created_by" gorm:"type:int"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"type:datetime"`
	UpdatedBy         int       `json:"updated_by" gorm:"type:int"`
	DeletedAt         time.Time `json:"deleted_at" gorm:"type:datetime"`
	DeletedBy         int       `json:"deleted_by" gorm:"type:int"`
	Is_Default        bool      `json:"is_default" gorm:"type:tinyint(1)"`
	Class_Type        string    `json:"class_type" gorm:"type:varchar(255)"`
}

func (Ci_Class) TableName() string {
	return "ci_class"
}
