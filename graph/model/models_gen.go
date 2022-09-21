// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Diary struct {
	Diaryid     string   `json:"Diaryid"`
	Word        string   `json:"Word"`
	Imageurl    string   `json:"Imageurl"`
	User        *User    `json:"User"`
	Emotion     *Emotion `json:"Emotion"`
	Englishword string   `json:"Englishword"`
	CreatedAt   string   `json:"CreatedAt"`
	UpdatedAt   string   `json:"UpdatedAt"`
}

type Emotion struct {
	Diaryid   string `json:"Diaryid"`
	Happy     string `json:"Happy"`
	Angry     string `json:"Angry"`
	Surprise  string `json:"Surprise"`
	Sad       string `json:"Sad"`
	Fear      string `json:"Fear"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type Me struct {
	User     *User        `json:"User"`
	Diary    []*Diary     `json:"Diary"`
	Followee []*UserDiary `json:"Followee"`
	Follower []*UserDiary `json:"Follower"`
}

type NewDiary struct {
	Word        string `json:"Word"`
	Englishword string `json:"Englishword"`
	Userid      string `json:"Userid"`
	Imageurl    string `json:"Imageurl"`
}

type NewEmotion struct {
	Diaryid  string `json:"Diaryid"`
	Happy    string `json:"Happy"`
	Angry    string `json:"Angry"`
	Surprise string `json:"Surprise"`
	Sad      string `json:"Sad"`
	Fear     string `json:"Fear"`
}

type NewFollow struct {
	Followee string `json:"Followee"`
	Follower string `json:"Follower"`
}

type NewUser struct {
	Userid string `json:"Userid"`
	Name   string `json:"Name"`
}

type User struct {
	Userid string `json:"Userid"`
	Name   string `json:"Name"`
}

type UserDiary struct {
	User  *User    `json:"User"`
	Diary []*Diary `json:"Diary"`
}
