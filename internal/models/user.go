package models

type User struct {
	ID       int        `json:"user_id"`
	Name     string     `json:"name"`
	Comments *[]Comment `json:"comments"`
	Likes    *[]Like    `json:"likes"`
}

type Comment struct {
	ID      int    `json:"comment_id"`
	VideoID int    `json:"video_id"`
	Text    string `json:"text"`
}

type Like struct {
	ID      int `json:"like_id"`
	VideoID int
}
