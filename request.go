package hive

import (
	"encoding/json"
	"io"
	"net/http"
)

func buildRequest(args ...string) string {
	var str = "https://api.playhive.com/v0/game"
	for _, arg := range args {
		str += "/" + arg
	}
	return str
}

/*type GlobalRequest struct {
	time Time
	game Game
}

func (r *GlobalRequest) Do() {
	time := r.time.String()
	game := r.game.String()

	_ = buildRequest(time, game)
}*/

type PlayerRequest struct {
	time   Time
	game   Game
	player string
}

func NewPlayerRequest(time Time, game Game, player string) *PlayerRequest {
	return &PlayerRequest{
		time:   time,
		game:   game,
		player: player,
	}
}

func (r *PlayerRequest) Do() (PlayerResponse, error) {
	var playerResponse PlayerResponse

	time := r.time.String()
	game := r.game.String()

	req := buildRequest(time, game, r.player)

	response, err := http.Get(req)
	if err != nil {
		return playerResponse, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return playerResponse, err
	}

	err = json.Unmarshal(data, &playerResponse)
	if err != nil {
		return PlayerResponse{}, err
	}

	return playerResponse, nil
}
