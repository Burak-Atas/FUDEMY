package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
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

// education
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	ProductId string             `json:"productid"`
	Category  string             `json:"category"`
	Title     string             `json:"name"`
	Details   string
	Url       string
	CreatedAt time.Time
	Price     string `json:"price"`
	TeacherID string `json:"teacherid"`
}

type Control struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	ProductId string             `json:"productid"`
	UserID    string             `json:"userid"`
}

//edu --> Video

type Video struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	VideoId    string             `json:"videoid"`
	VideoTitle string             `json:"videoname"`
	VideoUrl   string             `json:"videourl"`
	EduId      string             `json:"eduid"`
}

type Card struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	UserId     string
	Name       string
	Cardnumber string
	Csv        string
}

type Sepet struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserId string
	EduId  string `json:"eduid"`
}

type Yorumlar struct {
	ID             primitive.ObjectID `json:"_id" bson:"id"`
	UserId         string
	EduId          string
	CommentDetails string
}
