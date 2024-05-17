package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var eventTypeLibrary = map[string]int{
"Tsunami Warning":	1,
"Tornado Warning":	2,
"Extreme Wind Warning":	3,
"Severe Thunderstorm Warning":	4,
"Flash Flood Warning":	5,
"Flash Flood Statement":	6,
"Severe Weather Statement":	7,
"Shelter In Place Warning":	8,
"Evacuation Immediate":	9,
"Civil Danger Warning":	10,
"Nuclear Power Plant Warning":	11,
"Radiological Hazard Warning":	12,
"Hazardous Materials Warning":	13,
"Fire Warning":	14,
"Civil Emergency Message":	15,
"Law Enforcement Warning":	16,
"Storm Surge Warning":	17,
"Hurricane Force Wind Warning":	18,
"Hurricane Warning":	19,
"Typhoon Warning":	20,
"Special Marine Warning":	21,
"Blizzard Warning":	22,
"Snow Squall Warning":	23,
"Ice Storm Warning":	24,
"Winter Storm Warning":	25,
"High Wind Warning":	26,
"Tropical Storm Warning":	27,
"Storm Warning":	28,
"Tsunami Advisory":	29,
"Tsunami Watch":	30,
"Avalanche Warning":	31,
"Earthquake Warning":	32,
"Volcano Warning":	33,
"Ashfall Warning":	34,
"Coastal Flood Warning":	35,
"Lakeshore Flood Warning":	36,
"Flood Warning":	37,
"High Surf Warning":	38,
"Dust Storm Warning":	39,
"Blowing Dust Warning":	40,
"Lake Effect Snow Warning":	41,
"Excessive Heat Warning":	42,
"Tornado Watch":	43,
"Severe Thunderstorm Watch":	44,
"Flash Flood Watch":	45,
"Gale Warning":	46,
"Flood Statement":	47,
"Wind Chill Warning":	48,
"Extreme Cold Warning":	49,
"Hard Freeze Warning":	50,
"Freeze Warning":	51,
"Red Flag Warning":	52,
"Storm Surge Watch":	53,
"Hurricane Watch":	54,
"Hurricane Force Wind Watch":	55,
"Typhoon Watch":	56,
"Tropical Storm Watch":	57,
"Storm Watch":	58,
"Hurricane Local Statement":	59,
"Typhoon Local Statement":	60,
"Tropical Storm Local Statement":	61,
"Tropical Depression Local Statement":	62,
"Avalanche Advisory":	63,
"Winter Weather Advisory":	64,
"Wind Chill Advisory":	65,
"Heat Advisory":	66,
"Urban and Small Stream Flood Advisory":	67,
"Small Stream Flood Advisory":	68,
"Arroyo and Small Stream Flood Advisory":	69,
"Flood Advisory":	70,
"Hydrologic Advisory":	71,
"Lakeshore Flood Advisory":	72,
"Coastal Flood Advisory":	73,
"High Surf Advisory":	74,
"Heavy Freezing Spray Warning":	75,
"Dense Fog Advisory":	76,
"Dense Smoke Advisory":	77,
"Small Craft Advisory":	78,
"Brisk Wind Advisory":	79,
"Hazardous Seas Warning":	80,
"Dust Advisory":	81,
"Blowing Dust Advisory":	82,
"Lake Wind Advisory":	83,
"Wind Advisory":	84,
"Frost Advisory":	85,
"Ashfall Advisory":	86,
"Freezing Fog Advisory":	87,
"Freezing Spray Advisory":	88,
"Low Water Advisory":	89,
"Local Area Emergency":	90,
"Avalanche Watch":	91,
"Blizzard Watch":	92,
"Rip Current Statement":	93,
"Beach Hazards Statement":	94,
"Gale Watch":	95,
"Winter Storm Watch":	96,
"Hazardous Seas Watch":	97,
"Heavy Freezing Spray Watch":	98,
"Coastal Flood Watch":	99,
"Lakeshore Flood Watch":	100,
"Flood Watch":	101,
"High Wind Watch":	102,
"Excessive Heat Watch":	103,
"Extreme Cold Watch":	104,
"Wind Chill Watch":	105,
"Lake Effect Snow Watch":	106,
"Hard Freeze Watch":	107,
"Freeze Watch":	108,
"Fire Weather Watch":	109,
"Extreme Fire Danger":	110,
"911 Telephone Outage":	111,
"Coastal Flood Statement":	112,
"Lakeshore Flood Statement":	113,
"Special Weather Statement":	114,
"Marine Weather Statement":	115,
"Air Quality Alert":	116,
"Air Stagnation Advisory":	117,
"Hazardous Weather Outlook":	118,
"Hydrologic Outlook":	119,
"Short Term Forecast":	120,
"Administrative Message":	121,
"Test":	122,
"Child Abduction Emergency":	123,
"Blue Alert": 124,

}

var eventTypeColorLibrary = map[string]string{
	"Tsunami Warning":	"#FD6347",
	"Tornado Warning":	"#FF0000",
	"Extreme Wind Warning":	"#FF8C00",
	"Severe Thunderstorm Warning":	"#FFA500",
	"Flash Flood Warning":	"#8B0000",
	"Flash Flood Statement":	"#8B0000",
	"Severe Weather Statement":	"#00FFFF",
	"Shelter In Place Warning":	"#FA8072",
	"Evacuation Immediate":	"#7FFF00",
	"Civil Danger Warning":	"#FFB6C1",
	"Nuclear Power Plant Warning":	"#4B0082",
	"Radiological Hazard Warning":	"#4B0082",
	"Hazardous Materials Warning":	"#4B0082",
	"Fire Warning":	"#A0522D",
	"Civil Emergency Message":	"#FFB6C1",
	"Law Enforcement Warning":	"#C0C0C0",
	"Storm Surge Warning":	"#B524F7",
	"Hurricane Force Wind Warning":	"#CD5C5C",
	"Hurricane Warning":	"#DC143C",
	"Typhoon Warning":	"#DC143C",
	"Special Marine Warning":	"#FFA500",
	"Blizzard Warning":	"#FF4500",
	"Snow Squall Warning":	"#C71585",
	"Ice Storm Warning":	"#8B008B",
	"Winter Storm Warning":	"#FF69B4",
	"High Wind Warning":	"#DAA520",
	"Tropical Storm Warning":	"#B22222",
	"Storm Warning":	"#9400D3",
	"Tsunami Advisory":	"#D2691E",
	"Tsunami Watch":	"#FF00FF",
	"Avalanche Warning":	"#1E90FF",
	"Earthquake Warning":	"#8B4513",
	"Volcano Warning":	"#2F4F4F",
	"Ashfall Warning":	"#A9A9A9",
	"Coastal Flood Warning":	"#228B22",
	"Lakeshore Flood Warning":	"#228B22",
	"Flood Warning":	"#00FF00",
	"High Surf Warning":	"#228B22",
	"Dust Storm Warning":	"#FFE4C4",
	"Blowing Dust Warning":	"#FFE4C4",
	"Lake Effect Snow Warning":	"#008B8B",
	"Excessive Heat Warning":	"#C71585",
	"Tornado Watch":	"#FFFF00",
	"Severe Thunderstorm Watch":	"#DB7093",
	"Flash Flood Watch":	"#2E8B57",
	"Gale Warning":	"#DDA0DD",
	"Flood Statement":	"#00FF00",
	"Wind Chill Warning":	"#B0C4DE",
	"Extreme Cold Warning":	"#0000FF",
	"Hard Freeze Warning":	"#9400D3",
	"Freeze Warning":	"#483D8B",
	"Red Flag Warning":	"#FF1493",
	"Storm Surge Watch":	"#DB7FF7",
	"Hurricane Watch":	"#FF00FF",
	"Hurricane Force Wind Watch":	"#9932CC",
	"Typhoon Watch":	"#FF00FF",
	"Tropical Storm Watch":	"#F08080",
	"Storm Watch":	"#FFE4B5",
	"Hurricane Local Statement":	"#FFE4B5",
	"Typhoon Local Statement":	"#FFE4B5",
	"Tropical Storm Local Statement":	"#FFE4B5",
	"Tropical Depression Local Statement":	"#FFE4B5",
	"Avalanche Advisory":	"#CD853F",
	"Winter Weather Advisory":	"#7B68EE",
	"Wind Chill Advisory":	"#AFEEEE",
	"Heat Advisory":	"#FF7F50",
	"Urban and Small Stream Flood Advisory":	"#00FF7F",
	"Small Stream Flood Advisory":	"#00FF7F",
	"Arroyo and Small Stream Flood Advisory":	"#00FF7F",
	"Flood Advisory":	"#00FF7F",
	"Hydrologic Advisory":	"#00FF7F",
	"Lakeshore Flood Advisory":	"#7CFC00",
	"Coastal Flood Advisory":	"#7CFC00",
	"High Surf Advisory":	"#BA55D3",
	"Heavy Freezing Spray Warning":	"#00BFFF",
	"Dense Fog Advisory":	"#708090",
	"Dense Smoke Advisory":	"#F0E68C",
	"Small Craft Advisory":	"#D8BFD8",
	"Brisk Wind Advisory":	"#D8BFD8",
	"Hazardous Seas Warning":	"#D8BFD8",
	"Dust Advisory":	"#BDB76B",
	"Blowing Dust Advisory":	"#BDB76B",
	"Lake Wind Advisory":	"#D2B48C",
	"Wind Advisory":	"#D2B48C",
	"Frost Advisory":	"#6495ED",
	"Ashfall Advisory":	"#696969",
	"Freezing Fog Advisory":	"#8080",
	"Freezing Spray Advisory":	"#00BFFF",
	"Low Water Advisory":	"#A52A2A",
	"Local Area Emergency":	"#C0C0C0",
	"Avalanche Watch":	"#F4A460",
	"Blizzard Watch":	"#ADFF2F",
	"Rip Current Statement":	"#40E0D0",
	"Beach Hazards Statement":	"#40E0D0",
	"Gale Watch":	"#FFC0CB",
	"Winter Storm Watch":	"#4682B4",
	"Hazardous Seas Watch":	"#483D8B",
	"Heavy Freezing Spray Watch":	"#BC8F8F",
	"Coastal Flood Watch":	"#66CDAA",
	"Lakeshore Flood Watch":	"#66CDAA",
	"Flood Watch":	"#2E8B57",
	"High Wind Watch":	"#B8860B",
	"Excessive Heat Watch":	"#800000",
	"Extreme Cold Watch":	"#0000FF",
	"Wind Chill Watch":	"#5F9EA0",
	"Lake Effect Snow Watch":	"#87CEFA",
	"Hard Freeze Watch":	"#4169E1",
	"Freeze Watch":	"#00FFFF",
	"Fire Weather Watch":	"#FFDEAD",
	"Extreme Fire Danger":	"#E9967A",
	"911 Telephone Outage":	"#C0C0C0",
	"Coastal Flood Statement":	"#6B8E23",
	"Lakeshore Flood Statement":	"#6B8E23",
	"Special Weather Statement":	"#FFE4B5",
	"Marine Weather Statement":	"#FFDAB9",
	"Air Quality Alert":	"#808080",
	"Air Stagnation Advisory":	"#808080",
	"Hazardous Weather Outlook":	"#EEE8AA",
	"Hydrologic Outlook":	"#90EE90",
	"Short Term Forecast":	"#98FB98",
	"Administrative Message":	"#C0C0C0",
	"Test":	"#F0FFFF",
	"Child Abduction Emergency":	"#FFFFFF",
	"Blue Alert": "#FFFFFF",
	
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
	Color		string `json:"color"`
}

type Path struct {
	Path string `json:"path"`
}

type NoAlert struct {
	Output string `json:"output"`
}

type userAlertType struct {
	Data []string `json:"data"`
}

var alertList []Alert
var countyList = map[string]int{}
var userEnteredPath Path
var NoAlertStatement []NoAlert
var userAlertTypes userAlertType

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
	timeEffectiveTimeOutput := timeEffective.Format("3:04PM")
	timeExpiresTimeOutput := timeExpires.Format("3:04PM")

	

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

func getWarningColor(warningEvent string) string {
	if eventTypeColorLibrary[warningEvent] == "" { return "white"}
	return eventTypeColorLibrary[warningEvent]
}

func appendAndSortAlerts(alertListResponse Response, stateId string) {
	addStateIdToCountyList(stateId)
	for _, alertFeatures := range alertListResponse.Features {
		singleAlert := alertFeatures.Properties
		singleAlert.Priority = getWarningPriority(singleAlert.Event)
		singleAlert.Color = getWarningColor(singleAlert.Event)
		removeCommas(&singleAlert)
		changeTimeOutputAndHeadline(&singleAlert)
		if len(userAlertTypes.Data) > 0 && len(userAlertTypes.Data) < 124 && !checkIfAlertInUserAlert(singleAlert.Event) {
			continue
		}
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
	countyList = map[string]int{}
	arrayStates := c.Param("arrayStates")

	states := strings.Split(arrayStates, ",")
	for _, state := range states {
		getActiveAlertsFromNWS(state)
	}
	if (len(alertList) == 0) {
		emptyHeadline := "All clear! Currently there are no active alerts for " + arrayStates
		emptyAlerts := Alert{Headline: emptyHeadline}
		alertList = append(alertList, emptyAlerts)
	}
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
	if (len(alertList) == 0) {
		emptyHeadline := "All clear! Currently there are no active alerts for " + arrayStates
		emptyAlerts := Alert{Headline: emptyHeadline, AreaDesc: arrayCounties}
		alertList = append(alertList, emptyAlerts)
	}
	c.IndentedJSON(http.StatusOK, alertList)
}


func getUserAlertTypes(c *gin.Context) {
	if err := c.BindJSON(&userAlertTypes); err != nil {
		return
	}
	fmt.Println(userAlertTypes.Data, len(userAlertTypes.Data))
	c.IndentedJSON(http.StatusCreated, userAlertTypes)
}

func checkIfAlertInUserAlert(event string) bool {
	return (slices.Contains(userAlertTypes.Data, event))
}


func main() {
	//getActiveAlertsFromNWS("GA")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // http://localhost:3000 https://nws-api-active-alerts.vercel.app
		AllowMethods: []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.GET("/alerts", getAlerts)
	//router.GET("/alerts/:state", getState)
	router.GET("/alerts/:arrayStates", getState)
	router.GET("/alerts/:arrayStates/:arrayCounties", getStateWithCounties)
	router.POST("/userAlertTypes", getUserAlertTypes)
	router.Run("localhost:8080") //localhost:8080 :10000

}