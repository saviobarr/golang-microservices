package domain

type User struct {
	Id        uint64 `json:"id"`
	FirtsName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
