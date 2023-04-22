package model

type User struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	InterestsID []int64 `json:"interests_id"`
	CityID      *int64  `json:"city_id"`
}
