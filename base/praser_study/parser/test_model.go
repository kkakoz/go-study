package parser

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}


func (item *User) ToUserDTO() *UserDTO {
	return &UserDTO{
		ID:   item.ID,
		Name: item.Name,
		Age:  item.Age,
	}
}

