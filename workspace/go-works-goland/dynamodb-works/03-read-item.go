package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


type Music1 struct {
	Artist string
	SongTitle string
	AlbumTitle string
	Awards int
}
func main()  {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))

	db:= dynamodb.New(sess);

	// query parameters
	artist :="Sonu Nigam"
	songTitle :="Hello Again"

	params := &dynamodb.GetItemInput{
		TableName: aws.String("Music3"),
		Key: map[string]*dynamodb.AttributeValue{
			"Artist" : {
				S:aws.String(artist),
			},
			"SongTitle" : {
				S:aws.String(songTitle),
			},
		},
	}
	resp, err := db.GetItem(params)
	if err != nil {
		panic("Sorry item not found... ")
	}
	fmt.Println(resp.Item)

	 music  := Music1{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &music)


	if err != nil {
		panic("Sorry item not found... ")
	}
	fmt.Println("Unmarshelled Item %+v :", music )


}