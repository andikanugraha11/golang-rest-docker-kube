package models

const(
	tableName = "users"
)

type Users struct {
	Id 			int 	`gorm:"AUTO_INCREMENT",json:"id,omitempty"`
	Name		string 	`gorm:"type:varchar(100)",json:"name,omitempty"`
	Email 		string 	`gorm:"type:varchar(100);unique",json:"email,omitempty"`
	Username	string 	`gorm:"type:varchar(100);unique",json:"username,omitempty"`
	Password 	string 	`gorm:"type:varchar(100)"json:"password,omitempty"`
}