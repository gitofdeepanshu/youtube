package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response models the JSON structure
// that we get back from the YouTube API
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

// Items stores the ID + Statistics for
// a given channel
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

// Stats stores the information we care about
// so how many views the channel has, how many subscribers
// how many video etc.
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

//GetSubscribers ...
func GetSubscribers() (Items, error) {

	var response Response
	resp, err := http.Get("https://youtube.googleapis.com/youtube/v3/channels?id=UCUMZ7gohGI9HcU9VNsr2FJQ&key=AIzaSyAddtgzQpt0SMLnSvq7dhzkBVMvgWGpNsw&part=statistics")
	if err != nil {
		return Items{}, err
	}
	defer resp.Body.Close()
	// req, err := http.NewRequest("Get", "https://youtube.googleapis.com/youtube/v3/channels?id=UCUMZ7gohGI9HcU9VNsr2FJQ&key=AIzaSyCrq1S6WAtDc5UxCi-2SaGYZCOXhiu02yw&part=statistics", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return Item{}, err
	// }

	// q := req.URL.Query()
	// q.Add("part", "statistics")
	// q.Add("key","AIzaSyCrq1S6WAtDc5UxCi-2SaGYZCOXhiu02yw")
	// q.Add("id","UCUMZ7gohGI9HcU9VNsr2FJQ")

	// // https://youtube.googleapis.com/youtube/v3/channels?part=statistics&id=UCUMZ7gohGI9HcU9VNsr2FJQ&key=AIzaSyCrq1S6WAtDc5UxCi-2SaGYZCOXhiu02yw

	// req.URL.RawQuery = q.Encode()
	// fmt.Println(req)

	// client := &http.Client{}
	// resp, err := client.Do(req)

	fmt.Println("Response Status: ", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)

	if err != nil {
		return Items{}, err
	}

	return response.Items[0], nil
}
