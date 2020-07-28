package main

import (

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// insert record into the table called Music3


type Music3 struct {
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

	music := Music3{
		Artist: "Sonu Nigam",
		SongTitle: "Hello Again",
		AlbumTitle: "My Album",
		Awards: 2,
	}

	musicMap, err := dynamodbattribute.MarshalMap(music)

	if err != nil {
		panic("Cannot map the values given in music struct... ")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Music3"),
		Item:      musicMap,
	}



	resp, err := db.PutItem(params)

	if err != nil {
		panic("some problem while inserting")
	}

	fmt.Println("Record Inserted... ")
	fmt.Println(resp)
}
