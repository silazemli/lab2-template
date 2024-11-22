package reservation

type Hotel struct {
	ID       int    `json:"id"`
	HotelUID string `json:"uid"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Stars    int    `json:"stars"`
	Price    int    `json:"price"`
}
