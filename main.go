package da_fish

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/BHunter2889/da-fish/alexa"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"os"
	"encoding/base64"
)



func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "TodaysFishRatingIntent":
		response = HandleTodaysFishRatingIntent(request)
	case "FrontpageDealIntent":
		response = HandleFrontpageDealIntent(request)
	case "PopularDealIntent":
		response = HandlePopularDealIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	default:
		response = HandleAboutIntent(request)
	}
	return response
}

//TODO
func HandleTodaysFishRatingIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Today's Fishing Forecast", "Fish Rating is provided here.")
}

//TODO Remove
func HandleFrontpageDealIntent(request alexa.Request) alexa.Response {
	feedResponse, _ := RequestFeed("frontpage")
	var builder alexa.SSMLBuilder
	builder.Say("Here are the current frontpage deals:")
	builder.Pause("1000")
	for _, item := range feedResponse.Channel.Item {
		builder.Say(item.Title)
		builder.Pause("1000")
	}
	return alexa.NewSSMLResponse("Frontpage Deals", builder.Build())
}

//TODO Remove
func HandlePopularDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	// TODO
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("Give me the frontpage deals.")
	builder.Pause("1000")
	builder.Say("Give me the popular deals.")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Da Fish was created by HuntX in Saint Louis, Missouri so that he couldn't talk himself out of going fishing by using the excuse that conditions may not be optimal and figuring it out takes too much time to look up.")
}

// TODO - Delete this or rework for JSON
type FeedResponse struct {
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func RequestFeed(mode string) (FeedResponse, error) {
	endpoint, _ := url.Parse("https://slickdeals.net/newsearch.php")
	queryParams := endpoint.Query()
	queryParams.Set("mode", mode)
	queryParams.Set("searcharea", "deals")
	queryParams.Set("searchin", "first")
	queryParams.Set("rss", "1")
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		return FeedResponse{}, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var feedResponse FeedResponse
		xml.Unmarshal(data, &feedResponse)
		return feedResponse, nil
	}
}





func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}