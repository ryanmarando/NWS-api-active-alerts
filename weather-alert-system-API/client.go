package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var eventTypeLibrary = map[string]int{
	"Tornado Warning":             1,
	"Severe Thunderstorm Warning": 2,
	"Flash Flood Warning":         3,
	"Flood Warning":               4,
	"Tornado Watch":               5,
	"Severe Thunderstorm Watch":   6,
	"Flood Watch":                 7,
	"Wind Advisory":               8,
	"Coastal Flood Advisory":      9,
	"Frost Advisory":              10,
}

type Response struct {
	// Response struct to access all the features
	Features []Features `json:"features"` // features contain an array of properties in struct defined below

}
type Features struct {
	Properties Alert `json:"properties"` // properties contain alert strings defined in the alert struct
}
type Alert struct {
	AreaDesc 	string `json:"areaDesc"`
	Event 		string `json:"event"`
	Effective	string `json:"effective"`
	Expires		string `json:"expires"`
	Headline	string `json:"headline"`
	Priority	int `json:"priority"`
}

var alertList []Alert
var countyList = map[string]int{}

func addStateIdToCountyList(stateId string) {
	countyListLength := len(countyList)
	count := 0
	for county := range countyList {
		countyState := county + " " + stateId
		countyList[countyState] = 1
		count++
		if count == countyListLength {
			return
		}
	}
}

func addCounties(countyListArr []string) {
	for _,county := range countyListArr {
		countyList[county] = 1
	}
}

func inCountyListCheck(singleAlert *Alert) bool {
	areaDescList := strings.Split(singleAlert.AreaDesc, "; ")
	filteredAreaDescList := []string{}
	if len(filteredAreaDescList) == 0 { return false }
	for _,location := range areaDescList {
		_,ok := countyList[location]
		if ok {
			filteredAreaDescList = append(filteredAreaDescList, location)
		}
	}
	singleAlert.AreaDesc = strings.Join(filteredAreaDescList, "; ")
	return true
}

func changeTimeOutputAndHeadline(singleAlert *Alert) {
	timeStringEffective := singleAlert.Effective
	timeStringExpires := singleAlert.Expires

	timeEffective, _ := time.Parse(time.RFC3339, timeStringEffective)
	timeExpires, _ := time.Parse(time.RFC3339, timeStringExpires)
	localTimeEffective, localTimeExpires := timeEffective.Local(), timeExpires.Local()
	timeEffectiveTimeOutput := localTimeEffective.Format("3:04PM")
	timeExpiresTimeOutput := localTimeExpires.Format("3:04PM")
	

	singleAlert.Effective = timeEffectiveTimeOutput
	singleAlert.Expires = timeExpiresTimeOutput
	singleAlert.Headline = singleAlert.Event + " until " + singleAlert.Expires
}

func removeCommas(singleAlert *Alert) {
	locations := singleAlert.AreaDesc
	locations = strings.ReplaceAll(locations, ",", "")
	singleAlert.AreaDesc = locations
}

func readInDataFromNWS(responseData []byte) Response {
	var alertData Response
	json.Unmarshal(responseData, &alertData)
	return alertData
}

func sortAlertsByPriority(i, j int) bool {
	return alertList[i].Priority < alertList[j].Priority
	
}

func getWarningPriority(warningEvent string) int {
	if eventTypeLibrary[warningEvent] == 0 { return 1000 }
	return eventTypeLibrary[warningEvent]
}

func appendAndSortAlerts(alertListResponse Response, stateId string) {
	addCounties([]string{})
	addStateIdToCountyList(stateId)
	for _, alertFeatures := range alertListResponse.Features {
		singleAlert := alertFeatures.Properties
		singleAlert.Priority = getWarningPriority(singleAlert.Event)
		removeCommas(&singleAlert)
		changeTimeOutputAndHeadline(&singleAlert)
		if len(countyList) > 0 && inCountyListCheck(&singleAlert) {
			alertList = append(alertList, []Alert{singleAlert}...)
		} else if len(countyList) == 0 {
			alertList = append(alertList, []Alert{singleAlert}...)
		}
	}
	sort.Slice(alertList, sortAlertsByPriority)
}

func getActiveAlertsFromNWS(stateId string) {
	const BASE_URL = "https://api.weather.gov"
	response, err := http.Get(BASE_URL + "/alerts/active?area=" + stateId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rawAlertData := readInDataFromNWS(responseData)
	appendAndSortAlerts(rawAlertData, stateId)
}

func getAlerts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, alertList)
}

func getState(c *gin.Context) {
	alertList = []Alert{}
	arrayStates := c.Param("arrayStates")

	states := strings.Split(arrayStates, ",")
	for _, state := range states {
		getActiveAlertsFromNWS(state)
	}
	//c.JSON(200, gin.H{"states": states})
	//state := c.Param("state")
	//fmt.Println(state)
	//getActiveAlertsFromNWS(state)
	c.IndentedJSON(http.StatusOK, alertList)
}

func main() {
	//getActiveAlertsFromNWS("GA")
	router := gin.Default()
	router.GET("/alerts", getAlerts)
	//router.GET("/alerts/:state", getState)
	router.GET("/alerts/:arrayStates", getState)
	router.Run("localhost:8080")

}