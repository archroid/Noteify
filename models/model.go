package models


type User struct{
	ID int
	Username string
	Password string
	CREATED_AT int
	Token string
}

type Post struct{
	PostID int
	PostTitle string
	PostDate string
	Deleted int
	OwnerID int
}