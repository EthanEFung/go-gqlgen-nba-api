package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"gqlgen-nba/graph/generated"
	"gqlgen-nba/graph/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func (r *queryResolver) Scoreboard(ctx context.Context) (*model.ScoreboardResponse, error) {
	var t string
	t = time.Now().Format("20060102")
	res, err := http.Get("http://data.nba.net/10s/prod/v1/" + t + "/scoreboard.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed *model.ScoreboardResponse

	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}
	return parsed, err
}

func (r *queryResolver) Teams(ctx context.Context) (*model.TeamsResponse, error) {
	t := time.Now().Year()
	res, err := http.Get("http://data.nba.net/10s/prod/v2/" + strconv.Itoa(t) + "/teams.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed *model.TeamsResponse
	err = json.Unmarshal(body, &parsed)
	return parsed, err
}

func (r *queryResolver) Players(ctx context.Context) (*model.PlayersResponse, error) {
	res, err := http.Get("http://data.nba.net/prod/v1/2020/players.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed *model.PlayersResponse
	err = json.Unmarshal(body, &parsed)
	return parsed, err
}

func (r *queryResolver) Coaches(ctx context.Context) (*model.CoachesResponse, error) {
	t := time.Now().Year()
	res, err := http.Get("http://data.nba.net/prod/v1/" + strconv.Itoa(t) + "/coaches.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed *model.CoachesResponse
	err = json.Unmarshal(body, &parsed)
	return parsed, err
}

func (r *queryResolver) Schedule(ctx context.Context) (*model.ScheduleResponse, error) {
	t := time.Now().Year()
	res, err := http.Get("http://data.nba.net/prod/v1/" + strconv.Itoa(t) + "/schedule.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var parsed *model.ScheduleResponse
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		fmt.Println("unmarshal issue")
	}
	return parsed, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
