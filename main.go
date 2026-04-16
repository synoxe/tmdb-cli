package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Movie struct {
	Title       string  `json:"title"`
	VoteAverage float64 `json:"vote_average"`
	ReleaseDate string  `json:"release_date"`
}

type TMDBResponse struct {
	Results []Movie `json:"results"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env yuklemede sikinti")
	}

	apiKey := os.Getenv("TMDB_API_KEY")

	if apiKey == "" {
		fmt.Println("api key bos")
		return
	}

	movieType := flag.String("type", "popular", "Hangi filmleri görmek istiyorsun? (playing, popular, top, upcoming)")

	flag.Parse()

	mainUrl := "https://api.themoviedb.org/3/movie/"

	var endpoint string

	switch *movieType {
	case "popular":
		endpoint = "popular"

	case "playing":
		endpoint = "now_playing"
	case "top":
		endpoint = "top_rated"
	case "upcoming":
		endpoint = "upcoming"
	default:
		fmt.Println("Gecerli bir tur secin (playing, popular, top, upcoming)")
		return
	}

	getMovie(endpoint, mainUrl, apiKey)

}

func getMovie(movieType, mainUrl, apiKey string) {

	url := fmt.Sprintf("%s%s?api_key=%s", mainUrl, movieType, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("get isteginde sikinti", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("body okumada sikiniti", err)
		return
	}

	var results TMDBResponse

	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println("unmarshall da sikinti", err)
		return
	}
	fmt.Printf("%s Filmler: \n", movieType)
	for _, v := range results.Results {

		vote := strconv.FormatFloat(v.VoteAverage, 'f', 2, 64)

		fmt.Printf("Baslik:%s \nPuani:%s \nCikis Tarihi:%s \n---------\n", v.Title, vote, v.ReleaseDate)

	}
}
