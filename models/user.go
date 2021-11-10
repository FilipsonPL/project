package models

type User struct {
	ID        	uint64 `json:"id" gorm:"primary_key"`
	Name 		string `json:"name"`
	Surname  	string `json:"surname"`
	Age			int	   `json:"age"`
	Email		string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
}

type CreateUserInput struct {
	Name 		string `json:"name" binding:"required"`
	Surname  	string `json:"surname" binding:"required"`
	Age			int	   `json:"age" binding:"required"`
	Email		string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UpdateUser struct {
	ID			uint64 `json:"id" binding:"required"`
	Name 		string `json:"name"`
	Surname  	string `json:"surname"`
	Age			int	   `json:"age"`
	Email		string `json:"email"`
	Password    string `json:"password"`
}

