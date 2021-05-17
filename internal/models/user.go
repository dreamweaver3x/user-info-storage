package models

type User struct {
	ID       int        `json:"user_id" db:"id"`
	Name     string     `json:"name" db:"name"`
	Comments *[]Comment `json:"comments" db:"comments"`
	Likes    *[]Like    `json:"likes" db:"likes"`
}

type Comment struct {
	ID      int    `json:"comment_id" db:"id"`
	VideoID int    `json:"video_id" db:"video_id"`
	Text    string `json:"text" db:"text"`
}

type Like struct {
	ID      int `json:"like_id" db:"id"`
	VideoID int `json:"video_id" db:"video_id"`
}

//db:"video_id"
