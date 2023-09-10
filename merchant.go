
type Merchant struct {
	Id    int64  `json:"id"` // google id
	Email string `json:"email"`

	Paypal  string `json:"paypal"`
	Toss    string `json:"toss"`
	Osmosis string `json:"osmosis"`
}

var merchants = make(map[int64]Merchant) // id -> merchant

func register(c *gin.Context) {
	
}