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
	"Tsunami Warning":                        1,
	"Tornado Warning":                        2,
	"Extreme Wind Warning":                   3,
	"Severe Thunderstorm Warning":            4,
	"Flash Flood Warning":                    5,
	"Flash Flood Statement":                  6,
	"Severe Weather Statement":               7,
	"Shelter In Place Warning":               8,
	"Evacuation Immediate":                   9,
	"Civil Danger Warning":                   10,
	"Nuclear Power Plant Warning":            11,
	"Radiological Hazard Warning":            12,
	"Hazardous Materials Warning":            13,
	"Fire Warning":                           14,
	"Civil Emergency Message":                15,
	"Law Enforcement Warning":                16,
	"Storm Surge Warning":                    17,
	"Hurricane Force Wind Warning":           18,
	"Hurricane Warning":                      19,
	"Typhoon Warning":                        20,
	"Special Marine Warning":                 21,
	"Blizzard Warning":                       22,
	"Snow Squall Warning":                    23,
	"Ice Storm Warning":                      24,
	"Winter Storm Warning":                   25,
	"High Wind Warning":                      26,
	"Tropical Storm Warning":                 27,
	"Storm Warning":                          28,
	"Tsunami Advisory":                       29,
	"Tsunami Watch":                          30,
	"Avalanche Warning":                      31,
	"Earthquake Warning":                     32,
	"Volcano Warning":                        33,
	"Ashfall Warning":                        34,
	"Coastal Flood Warning":                  35,
	"Lakeshore Flood Warning":                36,
	"Flood Warning":                          37,
	"High Surf Warning":                      38,
	"Dust Storm Warning":                     39,
	"Blowing Dust Warning":                   40,
	"Lake Effect Snow Warning":               41,
	"Excessive Heat Warning":                 42,
	"Tornado Watch":                          43,
	"Severe Thunderstorm Watch":              44,
	"Flash Flood Watch":                      45,
	"Gale Warning":                           46,
	"Flood Statement":                        47,
	"Wind Chill Warning":                     48,
	"Extreme Cold Warning":                   49,
	"Hard Freeze Warning":                    50,
	"Freeze Warning":                         51,
	"Red Flag Warning":                       52,
	"Storm Surge Watch":                      53,
	"Hurricane Watch":                        54,
	"Hurricane Force Wind Watch":             55,
	"Typhoon Watch":                          56,
	"Tropical Storm Watch":                   57,
	"Storm Watch":                            58,
	"Hurricane Local Statement":              59,
	"Typhoon Local Statement":                60,
	"Tropical Storm Local Statement":         61,
	"Tropical Depression Local Statement":    62,
	"Avalanche Advisory":                     63,
	"Winter Weather Advisory":                64,
	"Wind Chill Advisory":                    65,
	"Heat Advisory":                          66,
	"Urban and Small Stream Flood Advisory":  67,
	"Small Stream Flood Advisory":            68,
	"Arroyo and Small Stream Flood Advisory": 69,
	"Flood Advisory":                         70,
	"Hydrologic Advisory":                    71,
	"Lakeshore Flood Advisory":               72,
	"Coastal Flood Advisory":                 73,
	"High Surf Advisory":                     74,
	"Heavy Freezing Spray Warning":           75,
	"Dense Fog Advisory":                     76,
	"Dense Smoke Advisory":                   77,
	"Small Craft Advisory":                   78,
	"Brisk Wind Advisory":                    79,
	"Hazardous Seas Warning":                 80,
	"Dust Advisory":                          81,
	"Blowing Dust Advisory":                  82,
	"Lake Wind Advisory":                     83,
	"Wind Advisory":                          84,
	"Frost Advisory":                         85,
	"Ashfall Advisory":                       86,
	"Freezing Fog Advisory":                  87,
	"Freezing Spray Advisory":                88,
	"Low Water Advisory":                     89,
	"Local Area Emergency":                   90,
	"Avalanche Watch":                        91,
	"Blizzard Watch":                         92,
	"Rip Current Statement":                  93,
	"Beach Hazards Statement":                94,
	"Gale Watch":                             95,
	"Winter Storm Watch":                     96,
	"Hazardous Seas Watch":                   97,
	"Heavy Freezing Spray Watch":             98,
	"Coastal Flood Watch":                    99,
	"Lakeshore Flood Watch":                  100,
	"Flood Watch":                            101,
	"High Wind Watch":                        102,
	"Excessive Heat Watch":                   103,
	"Extreme Cold Watch":                     104,
	"Wind Chill Watch":                       105,
	"Lake Effect Snow Watch":                 106,
	"Hard Freeze Watch":                      107,
	"Freeze Watch":                           108,
	"Fire Weather Watch":                     109,
	"Extreme Fire Danger":                    110,
	"911 Telephone Outage":                   111,
	"Coastal Flood Statement":                112,
	"Lakeshore Flood Statement":              113,
	"Special Weather Statement":              114,
	"Marine Weather Statement":               115,
	"Air Quality Alert":                      116,
	"Air Stagnation Advisory":                117,
	"Hazardous Weather Outlook":              118,
	"Hydrologic Outlook":                     119,
	"Short Term Forecast":                    120,
	"Administrative Message":                 121,
	"Test":                                   122,
	"Child Abduction Emergency":              123,
	"Blue Alert":                             124,
}

var eventTypeColorLibrary = map[string]string{
	"Tsunami Warning":                        "#FD6347",
	"Tornado Warning":                        "#FF0000",
	"Extreme Wind Warning":                   "#FF8C00",
	"Severe Thunderstorm Warning":            "#FFA500",
	"Flash Flood Warning":                    "#8B0000",
	"Flash Flood Statement":                  "#8B0000",
	"Severe Weather Statement":               "#00FFFF",
	"Shelter In Place Warning":               "#FA8072",
	"Evacuation Immediate":                   "#7FFF00",
	"Civil Danger Warning":                   "#FFB6C1",
	"Nuclear Power Plant Warning":            "#4B0082",
	"Radiological Hazard Warning":            "#4B0082",
	"Hazardous Materials Warning":            "#4B0082",
	"Fire Warning":                           "#A0522D",
	"Civil Emergency Message":                "#FFB6C1",
	"Law Enforcement Warning":                "#C0C0C0",
	"Storm Surge Warning":                    "#B524F7",
	"Hurricane Force Wind Warning":           "#CD5C5C",
	"Hurricane Warning":                      "#DC143C",
	"Typhoon Warning":                        "#DC143C",
	"Special Marine Warning":                 "#FFA500",
	"Blizzard Warning":                       "#FF4500",
	"Snow Squall Warning":                    "#C71585",
	"Ice Storm Warning":                      "#8B008B",
	"Winter Storm Warning":                   "#FF69B4",
	"High Wind Warning":                      "#DAA520",
	"Tropical Storm Warning":                 "#B22222",
	"Storm Warning":                          "#9400D3",
	"Tsunami Advisory":                       "#D2691E",
	"Tsunami Watch":                          "#FF00FF",
	"Avalanche Warning":                      "#1E90FF",
	"Earthquake Warning":                     "#8B4513",
	"Volcano Warning":                        "#2F4F4F",
	"Ashfall Warning":                        "#A9A9A9",
	"Coastal Flood Warning":                  "#228B22",
	"Lakeshore Flood Warning":                "#228B22",
	"Flood Warning":                          "#00FF00",
	"High Surf Warning":                      "#228B22",
	"Dust Storm Warning":                     "#FFE4C4",
	"Blowing Dust Warning":                   "#FFE4C4",
	"Lake Effect Snow Warning":               "#008B8B",
	"Excessive Heat Warning":                 "#C71585",
	"Tornado Watch":                          "#FFFF00",
	"Severe Thunderstorm Watch":              "#DB7093",
	"Flash Flood Watch":                      "#2E8B57",
	"Gale Warning":                           "#DDA0DD",
	"Flood Statement":                        "#00FF00",
	"Wind Chill Warning":                     "#B0C4DE",
	"Extreme Cold Warning":                   "#0000FF",
	"Hard Freeze Warning":                    "#9400D3",
	"Freeze Warning":                         "#483D8B",
	"Red Flag Warning":                       "#FF1493",
	"Storm Surge Watch":                      "#DB7FF7",
	"Hurricane Watch":                        "#FF00FF",
	"Hurricane Force Wind Watch":             "#9932CC",
	"Typhoon Watch":                          "#FF00FF",
	"Tropical Storm Watch":                   "#F08080",
	"Storm Watch":                            "#FFE4B5",
	"Hurricane Local Statement":              "#FFE4B5",
	"Typhoon Local Statement":                "#FFE4B5",
	"Tropical Storm Local Statement":         "#FFE4B5",
	"Tropical Depression Local Statement":    "#FFE4B5",
	"Avalanche Advisory":                     "#CD853F",
	"Winter Weather Advisory":                "#7B68EE",
	"Wind Chill Advisory":                    "#AFEEEE",
	"Heat Advisory":                          "#FF7F50",
	"Urban and Small Stream Flood Advisory":  "#00FF7F",
	"Small Stream Flood Advisory":            "#00FF7F",
	"Arroyo and Small Stream Flood Advisory": "#00FF7F",
	"Flood Advisory":                         "#00FF7F",
	"Hydrologic Advisory":                    "#00FF7F",
	"Lakeshore Flood Advisory":               "#7CFC00",
	"Coastal Flood Advisory":                 "#7CFC00",
	"High Surf Advisory":                     "#BA55D3",
	"Heavy Freezing Spray Warning":           "#00BFFF",
	"Dense Fog Advisory":                     "#708090",
	"Dense Smoke Advisory":                   "#F0E68C",
	"Small Craft Advisory":                   "#D8BFD8",
	"Brisk Wind Advisory":                    "#D8BFD8",
	"Hazardous Seas Warning":                 "#D8BFD8",
	"Dust Advisory":                          "#BDB76B",
	"Blowing Dust Advisory":                  "#BDB76B",
	"Lake Wind Advisory":                     "#D2B48C",
	"Wind Advisory":                          "#D2B48C",
	"Frost Advisory":                         "#6495ED",
	"Ashfall Advisory":                       "#696969",
	"Freezing Fog Advisory":                  "#8080",
	"Freezing Spray Advisory":                "#00BFFF",
	"Low Water Advisory":                     "#A52A2A",
	"Local Area Emergency":                   "#C0C0C0",
	"Avalanche Watch":                        "#F4A460",
	"Blizzard Watch":                         "#ADFF2F",
	"Rip Current Statement":                  "#40E0D0",
	"Beach Hazards Statement":                "#40E0D0",
	"Gale Watch":                             "#FFC0CB",
	"Winter Storm Watch":                     "#4682B4",
	"Hazardous Seas Watch":                   "#483D8B",
	"Heavy Freezing Spray Watch":             "#BC8F8F",
	"Coastal Flood Watch":                    "#66CDAA",
	"Lakeshore Flood Watch":                  "#66CDAA",
	"Flood Watch":                            "#2E8B57",
	"High Wind Watch":                        "#B8860B",
	"Excessive Heat Watch":                   "#800000",
	"Extreme Cold Watch":                     "#0000FF",
	"Wind Chill Watch":                       "#5F9EA0",
	"Lake Effect Snow Watch":                 "#87CEFA",
	"Hard Freeze Watch":                      "#4169E1",
	"Freeze Watch":                           "#00FFFF",
	"Fire Weather Watch":                     "#FFDEAD",
	"Extreme Fire Danger":                    "#E9967A",
	"911 Telephone Outage":                   "#C0C0C0",
	"Coastal Flood Statement":                "#6B8E23",
	"Lakeshore Flood Statement":              "#6B8E23",
	"Special Weather Statement":              "#FFE4B5",
	"Marine Weather Statement":               "#FFDAB9",
	"Air Quality Alert":                      "#808080",
	"Air Stagnation Advisory":                "#808080",
	"Hazardous Weather Outlook":              "#EEE8AA",
	"Hydrologic Outlook":                     "#90EE90",
	"Short Term Forecast":                    "#98FB98",
	"Administrative Message":                 "#C0C0C0",
	"Test":                                   "#F0FFFF",
	"Child Abduction Emergency":              "#FFFFFF",
	"Blue Alert":                             "#FFFFFF",
}

var countiesInStateLibrary = map[string][]string{
	"AK": {"Admiralty Island", "Alaska Peninsula", "Aleutians East", "Aleutians West", "Anchorage", "Annette Island", "Baldwin Peninsula",
		"Bering Strait Coast", "Bethel", "Bristol Bay", "Cape Fairweather to Lisianski Straight", "Central Aleutians",
		"Central Arctic Plains", "Central Beaufort Sea Coast", "Central Books Range", "Central Interior",
		"Chatanika River Valley", "Chugach", "City and Borough of Juneau", "City and Borough of Sitka",
		"City and Borough of Yakutat", "City of Hyder", "Copper River", "Copper River Basin", "Dalton Highway Summits",
		"Delta Junction", "Denali", "Dillingham", "East Turnagain Arm", "Eastern Alaska Range North of Trims Camp",
		"Eastern Alaska Range South of Trims Camp", "Eastern Aleutians", "Eastern Beaufort Sea Coast",
		"Eastern Chichagof Island", "Eastern Norton Sound and Nulato Hills", "Eielson AFB and Salcha",
		"Fairbanks Metro Area", "Fairbanks North Star", "Fortymile Country", "Glacier Bay", "Goldstream Valley and Nenana Hills",
		"Haines", "Haines Borough and Klukwan", "Hoonah-Angoon", "Howard Pass and the Delong Mountains", "Interior Seward Peninsula",
		"Juneau", "Kenai Peninsula", "Ketchikan Gateway", "Ketchikan Gateway Borough", "Kivalina and Red Dog Dock", "Kodiak Island", "Kuskokwim Delta",
		"Kusilvak", "Lake and Peninsula", "Lower Kobuk Valley", "Lower Koyukuk Valley", "Lower Kuskokwim Valley", "Lower Yukon and Innoko Valleys",
		"Lower Yukon River", "Matanuska Valley", "Matanuska-Susitna", "Middle Yukon Valley", "Municipality of Skagway", "Nenana",
		"Noatak Valley", "Nome", "North Slope", "North Slopes of the Western Alaska Range", "Northeast Prince William Sound",
		"Northern Arctic Coast", "Northern Denali Borough", "Northern Seward Peninsula", "Northwest Arctic", "Northwest Arctic Coast",
		"Petersburg Borough", "Pribilof Islands", "Prince of Wales Island", "Prince of Wales-Hyder", "Romanzof Mountains", "Shishmaref",
		"Sitka", "Skagway", "South Slopes Of The Central Brooks Range", "South Slopes Of The Eastern Brooks Range",
		"South Slopes of the Western Brooks Range", "Southeast Fairbanks", "Southeast Prince William Sound",
		"Southern Denali Borough", "Southern Seward Peninsula Coast", "St Lawrence Island", "Susitna Valley",
		"Tanana Flats", "Two Rivers", "Upper Chena River Valley", "Upper Kobuk Valleys", "Upper Koyukuk Valley",
		"Upper Kuskokwim Valley", "Upper Tanana Valley", "Western Aleutians", "Western Arctic Coast",
		"Western Arctic Plains", "Western Kenai Peninsula", "Western Kupreanof and Kuiu Island", "Western Prince William Sound",
		"White Mountains and High Terrain South of the Yukon River", "Wrangell", "Yakutat", "Yukon-Koyukuk", "Yukon Delta Coast", "Yukon Flats",
	},
	"AL": {"Autauga", "Baldwin", "Baldwin Inland", "Baldwin Central", "Baldwin Coastal", "Barbour", "Bibb", "Blount", "Bullock", "Butler", "Calhoun",
		"Chambers", "Cherokee", "Chilton", "Choctaw", "Clarke", "Clay", "Clebrune", "Coffee", "Colbert",
		"Conecuh", "Coosa", "Convington", "Crenshaw", "Cullman", "Dale", "Dallas", "DeKalb", "Elmore",
		"Escambia", "Etowah", "Fayette", "Franklin", "Geneva", "Greene", "Hale", "Henry", "Houston",
		"Jackson", "Jefferson", "Lamar", "Lauderdale", "Lawrence", "Lee", "Limestone", "Lowndes", "Macon",
		"Madison", "Marengo", "Marion", "Marshall", "Mobile", "Mobile Inland", "Mobile Central", "Mobile Coastal", "Monroe", "Montgomery", "Morgan", "Perry",
		"Pickens", "Pike", "Randolph", "Russell", "St. Clair", "Shelby", "Sumter", "Talladega", "Tallapoosa",
		"Tuscaloosa", "Walker", "Washington", "Wilcox", "Winston",
	},
	"AZ": {"Aguila Valley", "Apache", "Apache Junction/Gold Canyon", "Baboquivari Mountains including Kitt Peak", "Black Mesa Area",
		"Buckeye/Avondale", "Cave Creek/New River", "Central La Paz", "Central Phoenix", "Chinle Valley", "Chiricahua Mountains including Chiricahua National Monument",
		"Chuska Mountains and Defiance Plateau", "Cochise", "Coconino", "Coconino Plateau", "Deer Valley", "Dragoon/Mule/Huachuca and Santa Rita Mountains including Bisbee/Canelo Hills/Madera Canyon",
		"Dripping Springs", "East Valley", "Eastern Cochise County Below 5000 Feet including Douglas/Willcox",
		"Eastern Mogollon Rim", "Fountain Hills/East Mesa", "Galiuro and Pinaleno Mountains including Mount Graham",
		"Gila", "Gila Bend", "Gila River Valley", "Globe/Miami", "Graham", "Grand Canyon Country", "Greenlee",
		"Kaibab Plateau", "Kofa", "La Paz", "Lake Havasu and Fort Mohave", "Lake Mead National Recreation Area",
		"Little Colorado River Valley in Apache County", "Little Colorado River Valley in Coconino County",
		"Little Colorado River Valley in Navajo County", "Marble and Glen Canyons", "Maricopa", "Mazatzal Mountains",
		"Mohave", "Navajo", "New River Mesa", "North Phoenix/Glendale", "Northeast Plateaus and Mesas Hwy 264 Northward",
		"Northeast Plateaus and Mesas South of Hwy 264", "Northern Gila County", "Northwest Deserts", "Northwest Pinal County",
		"Northwest Plateau", "Northwest Valley", "Oak Creek and Sycamore Canyons", "Parker Valley",
		"Pima", "Pinal", "Pinal/Superstition Mountains", "Rio Verde/Salt River", "San Carlos", "Santa Cruz",
		"Santa Catalina and Rincon Mountains including Mount Lemmon/Summerhaven", "Scottsdale/Paradise Valley",
		"Sonoran Desert Natl Monument", "South Central Pinal County including Eloy/Picacho Peak State Park",
		"South Mountain/Ahwatukee", "Southeast Gila County", "Southeast Pinal County including Kearny/Mammoth/Oracle",
		"Southeast Valley/Queen Creek", "Southeast Yuma County", "Superior", "Tohono O'odham Nation including Sells",
		"Tonopah Desert", "Tonto Basin", "Tucson Metro Area including Tucson/Green Valley/Marana/Vail",
		"Upper Gila River and Aravaipa Valleys including Clifton/Safford", "Upper San Pedro River Valley including Sierra Vista/Benson",
		"Upper Santa Cruz River and Altar Valleys including Nogales", "West Pinal County", "Western Mogollon Rim",
		"Western Pima County Including Ajo/Organ Pipe Cactus National Monument", "White Mountains of Graham and Greenlee Counties including Hannagan Meadow",
		"White Mountains", "Yavapai", "Yavapai County Mountains", "Yavapai County Valleys and Basins", "Yuma",
	},
	"AR": {"Arkansas", "Ashley", "Baxter", "Benton", "Boone", "Boone County Except Southwest", "Boone County Higher Elevations", "Bradley", "Calhoun", "Carroll",
		"Central and Eastern Montgomery County", "Central and Southern Scott County", "Chicot", "Clark",
		"Clay", "Cleburne", "Cleveland", "Columbia", "Conway", "Craighead", "Crawford", "Crittenden", "Cross",
		"Dallas", "Desha", "Drew", "Eastern, Central, and Southern Searcy County Higher Elevations", "Faulkner", "Franklin", "Fulton", "Garland", "Grant", "Greene", "Hempstead",
		"Hot Spring", "Howard", "Independence", "Izard", "Jackson", "Jefferson", "Johnson", "Johnson County Higher Elevations", "Lafayette", "Lawrence",
		"Lee", "Lincoln", "Little River", "Logan", "Lonoke", "Madison", "Marion", "Miller", "Mississippi", "Monroe",
		"Montgomery", "Nevada", "Newton", "Newton County Higher Elevations", "Newton County Lower Elevations",
		"Northern Montgomery County Higher Elevations", "Northern Polk County Higher Elevations", "Northern Scott County",
		"Northwest Searcy County Higher Elevations", "Northwest Yell County", "Ouachita", "Perry", "Phillips", "Pike", "Poinsett", "Polk",
		"Polk County Lower Elevations", "Pope County Higher Elevations", "Pope",
		"Prairie", "Pulaski", "Randolph", "Saline", "Scott", "Searcy", "Searcy County Lower Elevations", "Sebastian", "Sevier", "Sharp",
		"Southeast Polk County Higher Elevations", "Southeast Van Buren County", "Southern and Eastern Logan County", "Southern Johnson County",
		"Southern Pope County", "Southwest Montgomery County Higher Elevations", "St. Francis",
		"Stone", "Union", "Van Buren", "Van Buren County Higher Elevations", "Washington", "Western and Northern Logan County", "White", "Woodruff", "Yell", "Yell Excluding Northwest",
	},
	"DE": {"New Castle", "Kent", "Sussex", "Inland Sussex", "Delaware Beaches"},
	"IN": {"Adams", "Allen", "Bartholomew", "Benton", "Blackford", "Boone", "Brown", "Carroll", "Cass",
		"Clark", "Clay", "Clinton", "Crawford", "Daviess", "Dearborn", "Decatur", "DeKalb", "Delaware",
		"Dubois", "Eastern St. Joseph", "Elkhart", "Fayette", "Floyd", "Fountain", "Franklin", "Fulton", "Gibson", "Grant", "Greene",
		"Hamilton", "Hancock", "Harrison", "Hendricks", "Henry", "Howard", "Huntington", "Jackson",
		"Jasper", "Jay", "Jefferson", "Jennings", "Johnson", "Knox", "Kosciusko", "La Porte", "Legrange", "Lake",
		"Lawrence", "Madison", "Marion", "Marshall", "Martin", "Miami", "Monroe", "Montgomery",
		"Morgan", "Newton", "Noble", "Northern Kosciusko", "Northern La Porte", "Ohio", "Orange", "Owen", "Parke", "Perry", "Pike", "Porter", "Posey",
		"Pulaski", "Putnam", "Randolph", "Ripley", "Rush", "Scott", "Shelby", "Southern Kosciusko", "Southern La Porte", "Spencer", "St. Joseph",
		"Starke", "Steuben", "Sullivan", "Switzerland", "Tippecanoe", "Tipton", "Union", "Vanderburgh",
		"Vermillion", "Vigo", "Wabash", "Warren", "Warrick", "Washington", "Wayne", "Wells", "Western St. Joseph", "White", "Whitley",
	},
	"OH": {"Adams", "Allen", "Ashland", "Ashtabula", "Ashtabula Lakeshore", "Athens", "Auglaize", "Belmont", "Brown",
		"Butler", "Carroll", "Champaign", "Clark", "Clermont", "Clinton", "Columbiana",
		"Coshocton", "Crawford", "Cuyahoga", "Darke", "Defiance", "Delaware", "Erie",
		"Fairfield", "Fayette", "Franklin", "Fulton", "Gallia", "Geagua", "Greene",
		"Guernsey", "Hamilton", "Hancock", "Hardin", "Harrison", "Henry", "Highland",
		"Hocking", "Holmes", "Huron", "Jackson", "Jefferson", "Knox", "Lake", "Lawrence",
		"Licking", "Logan", "Lorain", "Lucas", "Madison", "Mahoning", "Marion", "Medina",
		"Meigs", "Mercer", "Miami", "Monroe", "Montgomery", "Morgan", "Morrow", "Muskingum",
		"Noble", "Ottawa", "Paulding", "Perry", "Pickaway", "Pike", "Porage", "Preble",
		"Putnam", "Richland", "Ross", "Sandusky", "Scioto", "Seneca", "Shelby", "Stark",
		"Summit", "Trumbull", "Tuscarawas", "Union", "Van Wert", "Vinton", "Warren",
		"Washington", "Wayne", "Williams", "Wood", "Wyandot",
	},
}

type Response struct {
	// Response struct to access all the features
	Features []Features `json:"features"` // features contain an array of properties in struct defined below

}
type Features struct {
	Properties Alert `json:"properties"` // properties contain alert strings defined in the alert struct
}
type Alert struct {
	AreaDesc    string `json:"areaDesc"`
	Event       string `json:"event"`
	Effective   string `json:"effective"`
	Expires     string `json:"expires"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Color       string `json:"color"`
	SenderName  string `json:"senderName"`
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

type userNWSOffice struct {
	OfficeList []string `json:"officeList"`
}

var alertList []Alert
var countyList = map[string]int{}
var NoAlertStatement []NoAlert
var userAlertTypes userAlertType
var userNWSOffices userNWSOffice

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
	for _, county := range countyListArr {
		countyList[county] = 1
	}
}

func inCountyListCheck(singleAlert *Alert) bool {
	areaDescList := strings.Split(singleAlert.AreaDesc, "; ")
	filteredAreaDescList := []string{}
	for _, location := range areaDescList {
		_, ok := countyList[location]
		if ok {
			filteredAreaDescList = append(filteredAreaDescList, location)
		}
	}
	if len(filteredAreaDescList) == 0 {
		return false
	}
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

func removeCommas(AreaDesc string) string {
	return strings.ReplaceAll(AreaDesc, ",", "")

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
	if eventTypeLibrary[warningEvent] == 0 {
		return 1000
	}
	return eventTypeLibrary[warningEvent]
}

func getWarningColor(warningEvent string) string {
	if eventTypeColorLibrary[warningEvent] == "" {
		return "white"
	}
	return eventTypeColorLibrary[warningEvent]
}

func appendAndSortAlerts(alertListResponse Response, stateId string) {
	addStateIdToCountyList(stateId)
	for _, alertFeatures := range alertListResponse.Features {
		singleAlert := alertFeatures.Properties
		fmt.Println(singleAlert.SenderName)
		if len(userNWSOffices.OfficeList) > 0 && !checkIfAlertIssedByOffice(singleAlert.SenderName) {
			continue
		}
		singleAlert.Priority = getWarningPriority(singleAlert.Event)
		singleAlert.Color = getWarningColor(singleAlert.Event)
		singleAlert.AreaDesc = removeCommas(singleAlert.AreaDesc)
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
	if len(alertList) == 0 {
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
	//if len(alertList) == 0 {
	//	emptyHeadline := "All clear! Currently there are no active alerts for " + arrayStates
	//	arrayCounties = strings.ReplaceAll(arrayCounties, ",", "; ")
	//	emptyAlerts := Alert{Headline: emptyHeadline, AreaDesc: arrayCounties}
	//	alertList = append(alertList, emptyAlerts)
	//}
	c.IndentedJSON(http.StatusOK, alertList)
}

func getUserAlertTypes(c *gin.Context) {
	if err := c.BindJSON(&userAlertTypes); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, userAlertTypes)
}

func checkIfAlertInUserAlert(event string) bool {
	return (slices.Contains(userAlertTypes.Data, event))
}

func getNWSOffices(c *gin.Context) {
	if err := c.BindJSON(&userNWSOffices); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, userNWSOffices)
}

func checkIfAlertIssedByOffice(eventSenderName string) bool {
	return (slices.Contains(userNWSOffices.OfficeList, eventSenderName))
}

func getCountiesByState(c *gin.Context) {
	state := c.Param("state")
	countyListLibrary := countiesInStateLibrary[state]
	c.IndentedJSON(http.StatusCreated, countyListLibrary)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // http://localhost:3000 https://nws-api-active-alerts.vercel.app", "https://www.ryanmarando.com
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	router.GET("/alerts", getAlerts)
	//router.GET("/alerts/:state", getState)
	router.GET("/alerts/:arrayStates", getState)
	router.GET("/alerts/:arrayStates/:arrayCounties", getStateWithCounties)
	router.POST("/userAlertTypes", getUserAlertTypes)
	router.POST("/getNWSOffices", getNWSOffices)
	router.GET("/countiesByState/:state", getCountiesByState)
	router.Run("localhost:8080") //localhost:8080 :10000

}
