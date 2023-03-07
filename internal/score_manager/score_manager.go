package scoremanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/streadway/amqp"
)

type Details struct {
	MatchId string `json:"match_id"`
	OwnerId int    `json:"owner_id"`
}

type Summary struct {
	MatchId  string  `json:"match_id"`
	DataType string  `json:"data_type"`
	Score    Details `json:"Score"`
}

type ScoreManager struct {
	dynamodb *dynamodb.DynamoDB
	ch       *amqp.Channel
}

func New(dynamodb *dynamodb.DynamoDB, ch *amqp.Channel) *ScoreManager {
	return &ScoreManager{
		dynamodb,
		ch,
	}
}

func (sm *ScoreManager) Run() {

	msgs, err := sm.ch.Consume("myQueue", "", true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			go sm.process(d)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

func (sm *ScoreManager) Close() {

}

func (sm *ScoreManager) process(d amqp.Delivery) {
	var s Details
	json.Unmarshal(d.Body, &s)
	r := Summary{
		MatchId:  "match_2",
		DataType: "match_2_raw",
		Score:    s,
	}
	fmt.Println(s)
	av, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	_, err = sm.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("IMatches"),
		Item:      av,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}
}
