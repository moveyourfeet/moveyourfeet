package game

import "time"

type Game struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	CreateTime time.Time `json:"create_time"`
	JoinSecret string    `json:"join_secret"`
}
