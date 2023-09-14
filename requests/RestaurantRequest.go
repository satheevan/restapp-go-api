package requests

type RestaurantCreateRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name,omitempty" validate:"required"`
	Contact string `json:"contact,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
}

// func (ucr UserCreateRequest) EncryptPassword() string {
// 	password, err := bcrypt.GenerateFromPassword([]byte(ucr.Password), 14)
// 	if err != nil {
// 		log.Fatal("Error in encrypting the password")
// 	}
// 	return string(password)
// }
