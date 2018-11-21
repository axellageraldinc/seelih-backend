package request

type RegisterRequest struct {
	Email string
	Password string
	Fullname string
	Phone string
	FullAddress string
	CityCode uint
}