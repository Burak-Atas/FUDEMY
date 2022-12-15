package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"id"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	Token     string             `json:"token"`
	Type      string             `json:"type"`

	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`

	UserID string `json:"userid"`
}

//education
type Product struct {
	ID        primitive.ObjectID `json:"id" bson:"id"`
	ProductId string             `json:"productid"`
	Category  string             `json:"category"`
	Title     string             `json:"name"`
	Details   string
	CreatedAt time.Time
	Price     string `json:"price"`
	TeacherID string `json:"teacherid"`
}

//edu --> Video

type Video struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	VideoId      string             `json:"Videoid"`
	VideoTitle   string             `json:"videoname"`
	VideoUrl     string             `json:"videourl"`
	VideoDetails string
}

type Blog struct {
	ID           primitive.ObjectID
	Blog_ID      int
	Blog_Phost   string
	Blog_Title   string
	Blog_Content string
	TeacherId    string

	CreateDate time.Time
}

type Card struct {
	UserId     string
	Cardnumber string
	Csv        string
}

type Sepet struct {
	ID     primitive.ObjectID
	UserId string
	EduId  string
	Title  string
}

type Yorumlar struct {
	VideoId        string
	UserId         string
	S_UserId       string
	EduId          string
	CommentDetails string
}
