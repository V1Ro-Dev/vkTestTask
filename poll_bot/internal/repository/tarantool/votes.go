package tarantool

import (
	"encoding/json"
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"poll_bot/internal/models"
)

type TarantoolRepository struct {
	tarantoolConn *tarantool.Connection
}

func NewTarantoolRepository() *TarantoolRepository {
	tarantoolConn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{
		User: "guest",
	})
	if err != nil {
		log.Fatal("tarantool connection error:", err)
	}

	return &TarantoolRepository{tarantoolConn: tarantoolConn}
}

func (t *TarantoolRepository) CreatePoll(poll models.Poll) error {
	jsonPoll, err := json.Marshal(poll)

	if err != nil {
		return fmt.Errorf("unable to parse struct: %v", err)
	}

	_, err = t.tarantoolConn.Insert("polls", []interface{}{poll.ID, poll.ChannelID, jsonPoll})
	if err != nil {
		fmt.Println("Insertion error:", err)
	}

	return nil
}

func (t *TarantoolRepository) GetPollResults(poll models.Poll) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) EndPoll(poll models.Poll) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) DelPoll(poll models.Poll) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) Vote(poll models.Poll) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) GetPoll(channelId string, pollId string) (models.Poll, error) {
	resp, err := t.tarantoolConn.Select("polls", "primary", 0, 1, tarantool.IterEq, []interface{}{pollId, channelId})
	if err != nil {
		return models.Poll{}, fmt.Errorf("get query error: %v", err)
	}

	if len(resp.Data) == 0 {
		return models.Poll{}, fmt.Errorf("poll not found")
	}

	return models.Poll{}, err

}

func (t *TarantoolRepository) Close() {
	t.tarantoolConn.Close()
}
