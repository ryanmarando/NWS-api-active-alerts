package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
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

type Path struct {
	Path string `json:"path"`
}

var alertList []Alert
var countyList = map[string]int{}
var userEnteredPath Path

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
	for _,location := range areaDescList {
		_,ok := countyList[location]
		if ok {
			filteredAreaDescList = append(filteredAreaDescList, location)
		}
	}
	if len(filteredAreaDescList) == 0 { return false }
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

func exportToCSV(path string) {
	//path := "C:\Users\Ryan Marando\program_files\course_careers\final-project\data.csv"
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"AreaDesc", "Event", "Effective", "Expires", "Headline", "Priority"}
	if err := writer.Write(header); err != nil {
		return
	}

	// Write each record to the CSV file
	for _, alert := range alertList {
		record := []string{alert.AreaDesc, alert.Event, alert.Effective, alert.Expires, alert.Headline, strconv.FormatInt(int64(alert.Priority), 10)}
		if err := writer.Write(record); err != nil {
			return
		}
	}

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
	countyList = map[string]int{}
	arrayStates := c.Param("arrayStates")
	fmt.Println("User's path:", userEnteredPath.Path)

	states := strings.Split(arrayStates, ",")
	for _, state := range states {
		getActiveAlertsFromNWS(state)
	}
	exportToCSV(userEnteredPath.Path)
	c.IndentedJSON(http.StatusOK, alertList)
}

func getStateWithCounties(c *gin.Context) {
	alertList = []Alert{}
	countyList = map[string]int{}
	arrayStates := c.Param("arrayStates")
	arrayCounties := c.Param("arrayCounties")

	counties := strings.Split(arrayCounties, ",")
	addCounties(counties)

	states := strings.Split(arrayStates, ",")
	for _, state := range states {
		getActiveAlertsFromNWS(state)
	}
	exportToCSV(userEnteredPath.Path)
	c.IndentedJSON(http.StatusOK, alertList)
}



func getExportPath(c *gin.Context) {
	if err := c.BindJSON(&userEnteredPath); err != nil {
		return 
	}
	userEnteredPath.Path = userEnteredPath.Path + "/currentwarnings.csv"
	c.IndentedJSON(http.StatusCreated, userEnteredPath)
}

func main() {
	//getActiveAlertsFromNWS("GA")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://nws-api-active-alerts.vercel.app"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.GET("/alerts", getAlerts)
	//router.GET("/alerts/:state", getState)
	router.GET("/alerts/:arrayStates", getState)
	router.GET("/alerts/:arrayStates/:arrayCounties", getStateWithCounties)
	router.POST("/path", getExportPath)
	router.Run("localhost:8080") //localhost:8080


}