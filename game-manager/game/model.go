package game

type Game struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
