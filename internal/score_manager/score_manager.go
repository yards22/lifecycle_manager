package scoremanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/streadway/amqp"
	"github.com/yards22/lcmanager/internal/entities"
)

type Summary struct {
	MatchId  string             `json:"match_id"`
	DataType string             `json:"data_type"`
	Score    entities.ScoreItem `json:"Score"`
}

type GetEntry struct {
	MatchId  string `json:"match_id"`
	DataType string `json:"data_type"`
}

type ScoreManager struct {
	dynamodb *dynamodb.DynamoDB
	ch       *amqp.Channel
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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
	var s entities.ScoreItem
	json.Unmarshal(d.Body, &s)

	if s.Undo {

		return
	}

	r := Summary{
		MatchId:  s.MatchId,
		DataType: s.MatchId + "_raw",
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

	sm.generateCommentry_1(s)

	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}
}

func (sm *ScoreManager) generateCommentry_1(raw entities.ScoreItem) {
	// get the prev match status (network call)

	k := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"match_id": {
				S: aws.String("match_2"),
			},
			"data_type": {
				S: aws.String("match_2_commentry"),
			},
		},
		TableName: aws.String("IMatches"),
	}

	data, err := sm.dynamodb.GetItem(k)

	if err != nil {
		panic(err)
	}
	S := data.Item["score"].L
	//TODO: load S into commentry array as the schema presently is not proper in the table doing this dummy thing.
	var commentryArray []entities.CommentryF
	fmt.Println(S)

	var SS entities.ScoreSummary

	// have cur_batsmen score details .....
	SS.StrikerBatsman = raw.PlayersInAction.StrikerBatsman
	SS.NonStrikerBatsman = raw.PlayersInAction.NonStrikerBatsman
	if raw.InningsDetails.Balls == 1 && raw.InningsDetails.Overs > 0 {
		SS.PrevBowler = SS.CurBowler
	}
	SS.CurBowler = raw.PlayersInAction.Bowler
	// have extra details ...

	// have prev balls and commentry of 12 balls ...

	if len(commentryArray) > 0 {

		for i := len(commentryArray) - 1; i >= Max(0, len(commentryArray)-13); i-- {

		}
	}

}

func (sm *ScoreManager) generateCommentry_2(raw entities.ScoreItem) entities.CommentryF {
	// create commentry and send it back to process function ..
	if raw.WicketDetails.IsWicket {
		r := entities.Wicket{
			Overs:             raw.InningsDetails.Overs,
			Balls:             raw.InningsDetails.Balls,
			WicketType:        raw.WicketDetails.WicketType,
			StrikerBatsman:    raw.PlayersInAction.StrikerBatsman,
			NonStrikerBatsman: raw.PlayersInAction.NonStrikerBatsman,
			Bowler:            raw.PlayersInAction.Bowler,
			RunsScored:        raw.RunsDetails.RunsScored,
			WagonDirection:    raw.RunsDetails.WagonDirection,
			IsFielder:         raw.WicketDetails.IsFielder,
			FieldedBy:         raw.WicketDetails.FieldedBy,
			IsExtra:           raw.ExtraDetails.IsExtra,
			ExtraType:         raw.ExtraDetails.ExtraType,
		}
		return r.WicketC()
	}
	if raw.RunsDetails.RunsScored != 0 {
		r := entities.Runs{
			Overs:             raw.InningsDetails.Overs,
			Balls:             raw.InningsDetails.Balls,
			RunsScored:        raw.RunsDetails.RunsScored,
			StrikerBatsman:    raw.PlayersInAction.StrikerBatsman,
			NonStrikerBatsman: raw.PlayersInAction.NonStrikerBatsman,
			Bowler:            raw.PlayersInAction.Bowler,
			IsBoundary:        raw.RunsDetails.IsBoundary,
			BoundaryType:      raw.RunsDetails.BoundaryType,
			WagonDirection:    raw.RunsDetails.WagonDirection,
			IsExtra:           raw.ExtraDetails.IsExtra,
			ExtraType:         raw.ExtraDetails.ExtraType,
		}
		return r.RunsC()
	}
	if raw.ExtraDetails.IsExtra {
		r := entities.Extra{
			Overs:             raw.InningsDetails.Overs,
			Balls:             raw.InningsDetails.Balls,
			ExtraType:         raw.ExtraDetails.ExtraType,
			StrikerBatsman:    raw.PlayersInAction.StrikerBatsman,
			NonStrikerBatsman: raw.PlayersInAction.NonStrikerBatsman,
			Bowler:            raw.PlayersInAction.Bowler,
		}
		return r.ExtraC()
	}
	return entities.CommentryF{}
}
