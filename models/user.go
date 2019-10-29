package models

const(
	tableName = "users"
)

type Users struct {
	Id 			int
	Name		string
	email 		string
	Username	string
	Password 	string
}