package resume

type Resume struct {
	Id         string `json:"id"`
	About      string `json:"about"`
	Experience string `json:"experience"`
	UserID     int
}
