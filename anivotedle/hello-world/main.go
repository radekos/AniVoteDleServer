package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type AnimeObject struct {
	Picture       string
	Name          string
	Rating        float32
	URL           string
	Summary       string
	Genres        []string
	YearOfRelease int32
	EpisodeCount  int32
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//dateFromPath := request.PathParameters["date"]

	animeObjects := []AnimeObject{
		{
			Picture:       "FAKEPICTURE",
			Name:          "FAKENAME",
			Rating:        5.54,
			URL:           "asdasdasdassd.com",
			Summary:       "Fake anime",
			Genres:        []string{"Fake", "Isekai", "Food"},
			YearOfRelease: 2004,
			EpisodeCount:  102,
		}, {
			Picture:       "FAKEPICTURE",
			Name:          "FAKENAME",
			Rating:        7.54,
			URL:           "asdasdasdassd.com",
			Summary:       "Fake anime",
			Genres:        []string{"Fake", "Isekai", "Food"},
			YearOfRelease: 2004,
			EpisodeCount:  102,
		}, {
			Picture:       "FAKEPICTURE",
			Name:          "FAKENAME",
			Rating:        2.54,
			URL:           "asdasdasdassd.com",
			Summary:       "Fake anime",
			Genres:        []string{"Fake", "Isekai", "Food"},
			YearOfRelease: 2004,
			EpisodeCount:  102,
		}, {
			Picture:       "FAKEPICTURE",
			Name:          "FAKENAME",
			Rating:        1.54,
			URL:           "asdasdasdassd.com",
			Summary:       "Fake anime",
			Genres:        []string{"Fake", "Isekai", "Food"},
			YearOfRelease: 2004,
			EpisodeCount:  102,
		}, {
			Picture:       "FAKEPICTURE",
			Name:          "FAKENAME",
			Rating:        9.54,
			URL:           "asdasdasdassd.com",
			Summary:       "Fake anime",
			Genres:        []string{"Fake", "Isekai", "Food"},
			YearOfRelease: 2004,
			EpisodeCount:  102,
		},
	}

	jsonResponseBody, err := json.Marshal(animeObjects)

	if err != nil {
		//This is where we log it and exit???
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponseBody),
		StatusCode: 200,
	}, nil
}

func main() {
	//This is where we call the database

	lambda.Start(handler)
}
