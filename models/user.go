package models

const(
	tableName = "users"
)

type Users struct {
	Id 			int 	`json:"id,omitempty"`
	Name		string 	`json:"name,omitempty"`
	Email 		string 	`json:"email,omitempty"`
	Username	string 	`json:"username,omitempty"`
	Password 	string 	`json:"password,omitempty"`
}