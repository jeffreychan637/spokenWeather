package main

import (
    "bufio"
    // "encoding/json"
    "fmt"
    "net/url"
    "net/http"
    "io/ioutil"
    "os"
    "strings"
)

var baseForecastURL string = "https://api.darksky.net/forecast/";

func checkError(err error) {
    if (err != nil) {
        panic(err)
    }
}

func getDarkSkyAPIkey() string {
    data, err := ioutil.ReadFile("secret")
    checkError(err)

    // Get first line without newline
    return strings.Split(string(data), "\n")[0]
}

func getLocations() map[string]string {
    file, err := os.Open("locations")
    checkError(err)
    defer file.Close()

    locations := make(map[string]string)

    reader := bufio.NewReader(file)
    // todo: use a while loop and loop till end
    // of file instead of just getting first value
    firstLine, err := reader.ReadString('\n')
    checkError(err)

    var locationLine []string = strings.Split(firstLine, " : ")
    locations[locationLine[0]] = strings.Split(locationLine[1], "\n")[0] // remove newline

    fmt.Println(locations)
    return locations
}

func buildURL(apiKey string, location string) string {
    queryParams := url.Values{}
    queryParams.Set("exclude", "flags")
    qPs := queryParams.Encode()
    forecastURL := baseForecastURL + apiKey + "/" + location + "?" + qPs

    fmt.Println(forecastURL)
    return forecastURL
}

func makeForecastRequest(forecastURL string) []byte {
    response, err := http.Get(forecastURL)
    checkError(err)
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    checkError(err)

    fmt.Println(string(body))
    return body
}

func formatForecastResponse() {

}

func processForecastData() {

}

func outputForecast() {

}

func main() {
    apiKey := getDarkSkyAPIkey()
    locations := getLocations()
    location := locations["San Francisco"]
    // todo: update with concurrency with one
    // thread per location in locations
    forecastURL := buildURL(apiKey, location)
    makeForecastRequest(forecastURL)
}
