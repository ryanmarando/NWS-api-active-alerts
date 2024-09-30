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
	"unicode"

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
	"AL": {"Autauga", "Baldwin", "Baldwin Inland", "Baldwin Central", "Baldwin Coastal", "Barbour", "Bibb", "Blount", "Bullock", "Butler", "Calhoun",
		"Chambers", "Cherokee", "Chilton", "Choctaw", "Clarke", "Clay", "Clebrune", "Coffee", "Colbert",
		"Conecuh", "Coosa", "Convington", "Crenshaw", "Cullman", "Dale", "Dallas", "DeKalb", "Elmore",
		"Escambia", "Etowah", "Fayette", "Franklin", "Geneva", "Greene", "Hale", "Henry", "Houston",
		"Jackson", "Jefferson", "Lamar", "Lauderdale", "Lawrence", "Lee", "Limestone", "Lowndes", "Macon",
		"Madison", "Marengo", "Marion", "Marshall", "Mobile", "Mobile Inland", "Mobile Central", "Mobile Coastal", "Monroe", "Montgomery", "Morgan", "Perry",
		"Pickens", "Pike", "Randolph", "Russell", "St. Clair", "Shelby", "Sumter", "Talladega", "Tallapoosa",
		"Tuscaloosa", "Walker", "Washington", "Wilcox", "Winston",
	},
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
	"CA": {"Alameda", "Alpine", "Amador", "Antelope Valley", "Apple and Lucerne Valleys", "Bakersfield", "Buena Vista",
		"Burney Basin / Eastern Shasta County", "Butte", "Buttonwillow - Lost Hills - I5", "Cadiz Basin", "Calabasas and Agoura Hills",
		"Calaveras", "Carquinez Strait and Delta", "Caruthers - San Joaquin - Selma", "Catalina and Santa Barbara Islands",
		"Central Sacramento Valley", "Central Siskiyou County", "Central Ventura County Valleys", "Chiriaco Summit",
		"Chuckwalla Mountains", "Chuckwalla Valley", "Coachella Valley", "Coalinga - Avenal", "Coastal Del Norte",
		"Coastal North Bay Including Point Reyes National Seashore, CA", "Colusa", "Contra Costa", "Cuyama Valley",
		"Death Valley National Park", "Del Norte", "Del Norte Interior", "Delano-Wasco-Shafter", "East Bay Hills",
		"East Bay Interior Valleys", "Eastern Antelope Valley Foothills", "Eastern Mojave Desert, Including the Mojave National Preserve",
		"Eastern San Fernando Valley", "Eastern San Gabriel Mountains", "Eastern Santa Clara Hills",
		"Eastern Santa Monica Mountains Recreational Area", "Eastern Sierra Slopes of Inyo County", "El Dorado",
		"Frazier Mountain Communities", "Fresno", "Fresno-Clovis", "Fresno-Tulare Foothills", "Fresno-Tulare Lower Sierra",
		"Glenn", "Grant Grove Area", "Grapevine", "Greater Lake Tahoe Area", "Humboldt", "Imperial", "Imperial County Southeast",
		"Imperial County Southwest", "Imperial County West", "Imperial Valley", "Indian Wells Valley", "Interstate 5 Corridor",
		"Inyo", "Joshua Tree NP East", "Joshua Tree NP West", "Kaiser to Rodgers Ridge", "Kern", "Kern River Valley",
		"Kings", "Kings Canyon NP", "Lake", "Lake Casitas", "Lassen", "Lassen-Eastern Plumas-Eastern Sierra Counties",
		"Los Angeles", "Los Angeles County Beaches", "Los Angeles County Inland Coast including Downtown Los Angeles",
		"Los Angeles County San Gabriel Valley", "Los Banos - Dos Palos", "Madera", "Malibu Coast", "Marin", "Marin Coastal Range",
		"Mariposa", "Mariposa Madera Foothills", "Mariposa-Madera Lower Sierra", "Mendocino", "Mendocino Coast", "Merced", "Merced - Madera - Mendota",
		"Modoc", "Mojave Desert Slopes", "Mojave Desert", "Mono", "Morongo Basin", "Monterey", "Motherlode",
		"Mountains Of San Benito County And Interior Monterey County Including Pinnacles National Park",
		"Mountains Southwestern Shasta County to Western Colusa County", "Napa", "Nevada", "North Bay Interior Mountains",
		"North Bay Interior Valleys", "North Central and Southeast Siskiyou County", "Northeast Foothills/Sacramento Valley",
		"Northeast Siskiyou and Northwest Modoc Counties", "Northeastern Mendocino Interior", "Northern Humboldt Coast",
		"Northern Humboldt Interior", "Northern Lake County", "Northern Monterey Bay", "Northern Sacramento Valley",
		"Northern Salinas Valley/Hollister Valley and Carmel Valley", "Northern San Joaquin Valley", "Northern Trinity",
		"Northern Ventura County Mountains", "Northwestern Mendocino Interior", "Ojai Valley", "Orange", "Orange County Coastal",
		"Orange County Inland", "Owens Valley", "Palo Verde Valley", "Palos Verdes Hills", "Piute Walker Basin", "Placer",
		"Planada - Le Grand - Snelling", "Plumas", "Riverside", "Riverside County Mountains", "Sacramento",
		"Salton Sea", "San Benito", "San Bernardino", "San Bernardino and Riverside County Valleys-The Inland Empire",
		"San Bernardino County Mountains", "San Bernardino County-Upper Colorado River Valley", "San Diego", "San Diego County Coastal Areas",
		"San Diego County Deserts", "San Diego County Inland Valleys", "San Diego County Mountains", "San Francisco",
		"San Francisco Bay Shoreline", "San Francisco Peninsula Coast", "San Gorgonio Pass Near Banning", "San Joaquin",
		"San Joaquin River Canyon", "San Luis Obispo County Beaches", "San Luis Obispo County Inland Central Coast",
		"San Luis Obispo County Interior Valleys", "San Luis Obispo County Mountains",
		"San Luis Obispo", "San Mateo", "San Miguel and Santa Rosa Islands", "Santa Ana Mountains and Foothills",
		"Santa Barbara", "Santa Barbara County Central Coast Beaches", "Santa Barbara County Inland Central Coast",
		"Santa Barbara County Interior Mountains", "Santa Barbara County Southeastern Coast", "Santa Barbara County Southwestern Coast",
		"Santa Clara", "Santa Clara Valley Including San Jose", "Santa Clarita Valley", "Santa Cruz",
		"Santa Cruz and Anacapa Islands", "Santa Cruz Mountains", "Santa Lucia Mountains and Los Padres National Forest",
		"Santa Lucia Mountains", "Santa Susana Mountains", "Santa Ynez Mountains Eastern Range", "Santa Ynez Mountains Western Range",
		"Santa Ynez Valley", "Sequoia NP", "Shasta", "Shasta Lake Area / Northern Shasta County", "Sierra", "Siskiyou",
		"Solano", "Sonoma", "Sonoma Coastal Range", "South Central Siskiyou County", "South End of the Lower Sierra",
		"South End of the Upper Sierra", "South End San Joaquin Valley", "South End Sierra Foothills",
		"Southeast San Joaquin Valley", "Southeastern Mendocino Interior", "Southeastern Ventura County Valleys",
		"Southern Humboldt Interior", "Southern Lake County", "Southern Monterey Bay and Big Sur Coast",
		"Southern Sacramento Valley", "Southern Salinas Valley", "Southern Salinas Valley/Arroyo Seco and Lake San Antonio",
		"Southern Trinity", "Southern Ventura County Mountains", "Southwestern Humboldt", "Southwestern Mendocino Interior",
		"Stanislaus", "Surprise Valley California", "Sutter", "Tehachapi", "Tehama", "Trinity", "Tulare", "Tuolumne",
		"Upper San Joaquin River", "Ventura", "Ventura County Beaches", "Ventura County Inland Coast", "Visalia - Porterville - Reedley",
		"West Side Mountains north of 198", "West Side Mountains South of 198", "West Side of Fresno and Kings Counties",
		"West Slope Northern Sierra Nevada", "Western Antelope Valley Foothills", "Western Mojave Desert",
		"Western Plumas County/Lassen Park", "Western San Fernando Valley", "Western San Gabriel Mountains and Highway 14 Corridor",
		"Western Santa Monica Mountains Recreational Area", "Western Siskiyou County", "White Mountains of Inyo County",
		"Yolo", "Yosemite NP outside of the valley", "Yosemite Valley", "Yuba",
	},
	"CO": {"Adams", "Alamosa", "Alamosa Vicinity/Central San Luis Valley Below 8500 Ft", "Animas River Basin", "Arapahoe",
		"Archuleta", "Baca", "Bent", "Boulder", "Boulder And Jefferson Counties Below 6000 Feet/West Broomfield County",
		"Broomfield", "Canon City Vicinity/Eastern Fremont County", "Central and East Adams and Arapahoe Counties",
		"Central and South Weld County", "Central and Southeast Park County", "Central Chaffee County Below 9000 Ft",
		"Central Colorado River Basin", "Central Gunnison and Uncompahgre River Basin", "Central Yampa River Basin",
		"Chaffee", "Cheyenne", "Clear Creek", "Colorado Springs Vicinity/Southern El Paso County/Rampart Range Below 7400 Ft",
		"Conejos", "Costilla", "Crowley", "Custer", "Debeque to Silt Corridor", "Del Norte Vicinity/Northern San Luis Valley Below 8500 Ft",
		"Delta", "Denver", "Dolores", "Douglas", "Eagle", "El Paso", "Elbert", "Eastern Kiowa County", "Eastern Las Animas County",
		"Eastern San Juan Mountains Above 10000 Ft", "Eastern Sawatch Mountains above 11000 Ft", "Elbert/Central and East Douglas Counties Above 6000 Feet",
		"Elkhead and Park Mountains", "Flat Tops", "Four Corners/Upper Dolores River", "Fremont",
		"Garfield", "Gilpin", "Gore and Elk Mountains/Central Mountain Valleys", "Grand", "Grand and Battlement Mesas",
		"Grand and Summit Counties Below 9000 Feet", "Grand Valley", "Gunnison", "Hinsdale", "Huerfano", "Jackson", "Jackson County Below 9000 Feet",
		"Jefferson", "Jefferson and West Douglas Counties Above 6000 Feet/Gilpin/Clear Creek/Northeast Park Counties Below 9000 Feet",
		"Kiowa", "Kit Carson", "La Garita Mountains Above 10000 Ft", "La Junta Vicinity/Otero County", "La Plata",
		"Lake", "Lamar Vicinity/Prowers County", "Larimer", "Larimer and Boulder Counties Between 6000 and 9000 Feet",
		"Larimer County Below 6000 Feet/Northwest Weld County", "Las Animas", "Las Animas Vicinity/Bent County",
		"Leadville Vicinity/Lake County Below 11000 Ft", "Lincoln", "Logan", "Lower Yampa River Basin", "Mesa", "Mineral", "Moffat", "Montezuma", "Montrose", "Morgan",
		"North and Northeast Elbert County Below 6000 Feet/North Lincoln County", "North Douglas County Below 6000 Feet/Denver/West Adams and Arapahoe Counties/East Broomfield County",
		"Northeast Weld County", "Northern El Paso County/Monument Ridge/Rampart Range Below 7500 Ft", "Northern Sangre de Cristo Mountains above 11000 Ft",
		"Northern Sangre de Cristo Mountains Between 8500 And 11000 Ft", "Northwestern Fremont County Above 8500Ft", "Northwestern San Juan Mountains", "Otero",
		"Paradox Valley/Lower Dolores River", "Park", "Philips", "Pikes Peak above 11000 Ft", "Pitkin", "Prowers", "Pueblo",
		"Pueblo Vicinity/Pueblo County Below 6300 Feet", "Rio Blanco", "Rio Grande", "Roan and Tavaputs Plateaus", "Routt", "Saguache",
		"Saguache County East of Continental Divide below 10000 Ft", "Saguache County West of Continental Divide Below 10000 Ft", "San Juan",
		"San Juan River Basin", "San Miguel", "Sedgwick",
		"South and East Jackson/Larimer/North and Northeast Grand/Northwest Boulder Counties Above 9000 Feet",
		"South and Southeast Grand/West Central and Southwest Boulder/Gilpin/Clear Creek/Summit/North and West Park Counties Above 9000 Feet",
		"Southeast Elbert County Below 6000 Feet/South Lincoln County", "Southern San Luis Valley", "Southern Sangre De Cristo Mountains Above 11000 Ft",
		"Southern Sangre De Cristo Mountains Between 7500 and 11000 Ft", "Southwest San Juan Mountains", "Springfield Vicinity/Baca County",
		"Summit", "Teller", "Teller County/Rampart Range above 7500fT/Pike's Peak Between 7500 And 11000 Ft",
		"Trinidad Vicinity/Western Las Animas County Below 7500 Ft", "Uncompahgre Plateau/Dallas Divide",
		"Upper Gunnison River Valley", "Upper Rio Grande Valley/Eastern San Juan Mountains Below 10000 Ft",
		"Upper Yampa River Basin", "Walsenburg Vicinity/Upper Huerfano River Basin Below 7500 Ft", "Washington", "Weld",
		"West Elk and Sawatch Mountains", "West Jackson and West Grand Counties Above 9000 Feet", "Western Chaffee County Between 9000 and 11000 Ft",
		"Western Kiowa County", "Western Mosquito Range/East Chaffee County above 9000Ft", "Western Mosquito Range/East Lake County Above 11000 Ft",
		"Western/Central Fremont County Below 8500 Ft", "Wet Mountain Valley Below 8500 Ft", "Wet Mountains above 10000 Ft",
		"Wet Mountains between 6300 and 10000Ft", "Yuma",
	},
	"CT": {"Fairfield", "Hartford", "Litchfield", "Middlesex", "New Haven", "New London", "Northern Fairfield",
		"Northern Litchfield", "Northern Middlesex", "Northern New Haven", "Northern New London", "Southern Fairfield",
		"Southern Litchfield", "Southern Middlesex", "Southern New Haven", "Southern New London", "Tolland", "Windham"},
	"DE": {"New Castle", "Kent", "Sussex", "Inland Sussex", "Delaware Beaches"},
	"DC": {"District of Columbia"},
	"FL": {"Alachua", "Baker", "Bay", "Bradford", "Brevard", "Broward", "Calhoun", "Central Marion", "Central Walton",
		"Charlotte", "Citrus", "Clay", "Coastal Bay", "Coastal Broward County", "Coastal Charlotte", "Coastal Citrus",
		"Coastal Collier County", "Coastal Dixie", "Coastal Duval", "Coastal Flagler", "Coastal Franklin", "Coastal Gulf",
		"Coastal Hernando", "Coastal Hillsborough", "Coastal Indian River", "Coastal Jefferson", "Coastal Lee", "Coastal Levy",
		"Coastal Manatee", "Coastal Martin", "Coastal Miami-Dade County", "Coastal Nassau", "Coastal Palm Beach County",
		"Coastal Pasco", "Coastal Sarasota", "Coastal St. Johns", "Coastal St. Lucie", "Coastal Taylor", "Coastal Volusia",
		"Coastal Wakulla", "Collier", "Columbia", "DeSoto", "Dixie", "Duval", "Eastern Alachua", "Eastern Clay",
		"Eastern Hamilton", "Eastern Marion", "Eastern Putnam", "Escambia", "Escambia Coastal", "Escambia Inland",
		"Far South Miami-Dade County", "Flagler", "Franklin", "Gadsden", "Gilchrist", "Glades", "Gulf", "Hamilton",
		"Hardee", "Hendry", "Hernando", "Highlands", "Hillsborough", "Holmes", "Indian River", "Inland Bay",
		"Inland Broward County", "Inland Charlotte", "Inland Citrus", "Inland Collier County", "Inland Dixie", "Inland Flagler",
		"Inland Franklin", "Inland Gulf", "Inland Hernando", "Inland Hillsborough", "Inland Indian River", "Inland Jefferson",
		"Inland Lee", "Inland Levy", "Inland Manatee", "Inland Martin", "Inland Miami-Dade County", "Inland Nassau",
		"Inland Northern Brevard", "Inland Palm Beach County", "Inland Pasco", "Inland Sarasota", "Inland Southern Brevard",
		"Inland St. Johns", "Inland St. Lucie", "Inland Taylor", "Inland Volusia", "Inland Wakulla", "Jackson", "Jefferson", "Lafayette",
		"Lake", "Lee", "Leon", "Levy", "Liberty", "Madison", "Mainland Monroe", "Mainland Northern Brevard", "Mainland Southern Brevard",
		"Manatee", "Marion", "Martin", "Metro Broward County", "Metro Palm Beach County", "Miami-Dade", "Metropolitan Miami-Dade",
		"Monroe Lower Keys", "Monroe Middle Keys", "Monroe Upper Keys", "Nassau", "North Walton", "Northern Brevard Barrier Islands",
		"Northern Columbia", "Northern Lake County",
		"Okaloosa", "Okaloosa Coastal", "Okaloosa Inland", "Okeechobee", "Orange", "Osceola", "Palm Beach", "Pasco", "Pinellas", "Polk",
		"Putnam", "Santa Rosa", "Santa Rosa Coastal", "Santa Rosa Inland", "Sarasota", "Seminole", "South Central Duval",
		"South Walton", "Southeastern Columbia", "Southern Brevard Barrier Islands", "Southern Lake County", "Southwestern Columbia", "St. Johns", "St. Lucie",
		"Sumter", "Suwannee", "Taylor", "Trout River", "Union", "Volusia", "Wakulla", "Waslton",
		"Washington", "Western Alachua", "Western Clay", "Western Duval", "Western Hamilton", "Western Marion", "Western Putnam",
	},
	"GA": {"Appling", "Atkinson", "Bacon", "Baker", "Baldwin", "Banks", "Barrow", "Bartow", "Ben Hill", "Berrien", "Bibb", "Bleckley",
		"Brantley", "Brooks", "Bryan", "Bulloch", "Burke", "Butts", "Calhoun", "Camden", "Candler", "Carroll", "Catoosa", "Charlton",
		"Chatham", "Chattahoochee", "Chattooga", "Cherokee", "Clarke", "Clay", "Clayton", "Clinch", "Coastal Bryan",
		"Coastal Camden", "Coastal Chatham", "Coastal Glynn", "Coastal Liberty", "Coastal McIntosh", "Cobb", "Coffee", "Colquitt",
		"Columbia", "Cook", "Coweta", "Crisp", "Dade", "Dawson", "Decatur", "DeKalb", "Dodge", "Dooly", "Dougherty", "Douglas",
		"Early", "Echols", "Effingham", "Elbert", "Emanuel", "Evans", "Fannin", "Fayette", "Floyd", "Forsyth", "Franklin", "Fulton",
		"Gilmer", "Glascock", "Glynn", "Gordon", "Grady", "Greene", "Gwinnett", "Habersham", "Hall", "Hancock", "Haralson", "Harris",
		"Hart", "Heard", "Henry", "Houston", "Inland Bryan", "Inland Camden", "Inland Chatham", "Inland Glynn", "Inland Liberty",
		"Inland McIntosh", "Irwin", "Jackson", "Jasper", "Jeff Davis", "Jefferson", "Jenkins", "Johnson", "Jones",
		"Lamar", "Lanier", "Laurens", "Lee", "Liberty", "Lincoln", "Long", "Lowndes", "Lumpkin", "Macon", "Madison", "Marion", "McDuffie",
		"McIntosh", "Meriwether", "Miller", "Mitchell", "Monroe", "Montgomery", "Morgan", "Murray", "Muscogee", "Newton",
		"North Fulton", "Northeastern Charlton", "Northern Ware", "Oconee",
		"Oglethorpe", "Paulding", "Peach", "Pickens", "Pierce", "Pike", "Polk", "Pulaski", "Putnam", "Quitman", "Rabun", "Randolph",
		"Richmond", "Rockdale", "Schley", "Screven", "Seminole", "South Fulton", "Southern Ware", "Spalding", "Stephens", "Stewart", "Sumter", "Talbot", "Taliaferro",
		"Tattnall", "Taylor", "Telfair", "Terrell", "Thomas", "Tift", "Toombs", "Towns", "Treutlen", "Troup", "Turner", "Twiggs", "Union",
		"Upson", "Walker", "Walton", "Ware", "Warren", "Washington", "Wayne", "Webster", "Western Charlton", "Wheeler", "White", "Whitfield", "Wilcox",
		"Wilkes", "Wilkinson", "Worth",
	},
	"HI": {"Big Island East", "Big Island Interior", "Big Island North", "Big Island South", "Big Island Southeast", "Big Island Summit",
		"Central Oahu", "East Honolulu", "Ewq Plain", "Haleakala Summit", "Honolulu Metro", "Kahoolawe", "Kauai East", "Kauai Mountains",
		"Kauai North", "Kauai South", "Kauai Southwest", "Kipahulu", "Kohala", "Kona", "Koolau Leeward", "Koolau Windward",
		"Lanai Leeward", "Lanai Mauka", "Lanai South", "Lanai Windward", "Maui Central Valley North", "Maui Central Valley South",
		"Maui Leeward West", "Maui Windward West", "Molokai Leeward South", "Molokai North", "Molokai Southeast", "Molokai West",
		"Molokai Westward", "Niihau", "Oahu North Shore", "Olomana", "South Haleakala", "South Maui/Upcountry", "Waianaie",
		"Waianae Mountains", "Windward Haleakala"},
	"ID": {"Ada", "Adams", "Arco/Mud Lake Desert", "Bannock", "Bear Lake", "Bear Lake Valley", "Bear River Range",
		"Beaverhead/Lemhi Highlands", "Benewah", "Big Hole Mountains", "Big Lost Highlands/Copper Basin", "Bingham",
		"Blackfoot Mountains", "Blaine", "Boise", "Boise Mountains", "Bonner", "Bonneville", "Boundary", "Butte",
		"Camas", "Camas Prairie", "Canyon", "Caribou", "Caribou Range", "Cassia", "Centennial Mountains/Island Park",
		"Central Panhandle Mountains", "Challis/Pahsimeroi Valleys", "Clark", "Clearwater", "Coeur d'Alene Area", "Custer",
		"Eastern Lemhi County", "Eastern Magic Valley", "Elmore", "Frank Church Wilderness", "Franklin", "Franklin/Eastern Oneida Region", "Fremont", "Gem", "Gooding",
		"Idaho", "Idaho Palouse", "Jefferson", "Jerome", "Kootenai", "Latah", "Lemhi", "Lewis",
		"Lewis and Southern Nez Perce Counties", "Lewiston Area", "Lincoln", "Lost River Range", "Lost River Valleys",
		"Lower Hells Canyon/Salmon River Region", "Lower Snake River Plain", "Lower Treasure Valley", "Madison",
		"Marsh and Arbon Highlands", "Minidoka", "Nez Perce", "Northern Clearwater Mountains", "Northern Panhandle", "Oneida",
		"Orofino/Grangeville Region", "Owyhee", "Owyhee Mountains", "Payette", "Power", "Raft River Region",
		"Sawtooth/Stanley Basin", "Shoshone", "Shoshone/Lava Beds", "Southern Clearwater Mountains", "Southern Hills/Albion Mountains",
		"Southern Twin Falls County", "Southwest Highlands", "Sun Valley Region", "Teton", "Teton Valley", "Twin Falls",
		"Upper Snake River Plain", "Upper Treasure Valley", "Upper Weiser River", "Valley", "Washington",
		"West Central Mountains", "Western Lemhi County", "Western Magic Valley", "Wood River Foothills",
	},
	"IL": {"Adams", "Alexander", "Bond", "Boone", "Brown", "Bureau", "Calhoun", "Carroll", "Cass", "Champaign", "Central Cook", "Christian", "Clark", "Clay", "Clinton",
		"Coles", "Cook", "Crawford", "Cumberland", "De Kalb", "De Witt", "Douglas", "DuPage", "Eastern Will", "Edgar", "Edwards", "Effingham", "Fayette", "Ford", "Franklin",
		"Fulton", "Gallatin", "Greene", "Grundy", "Hamilton", "Hancock", "Hardin", "Henderson", "Henry", "Iroquois", "Jackson", "Jasper", "Jefferson",
		"Jersey", "Jo Daviess", "Johnson", "Kane", "Kankakee", "Kendall", "Knox", "La Salle", "Lake", "Lawrence", "Lee", "Livingston", "Logan", "Macon",
		"Macoupin", "Madison", "Marion", "Marshall", "Mason", "Massac", "McDonough", "McHenry", "McLean", "Menard", "Mercer", "Monroe", "Montgomery",
		"Morgan", "Moultrie", "Northern Cook", "Northern Will", "Ogle", "Peoria", "Perry", "Piatt", "Pike", "Pope", "Pulaski", "Putnam", "Randolph", "Richland", "Rock Island", "Saline",
		"Sangamon", "Schuyler", "Scott", "Shelby", "Southern Cook", "Southern Will", "St. Clair", "Stark", "Stephenson", "Tazewell", "Union", "Vermilion", "Wabash", "Warren", "Washington",
		"Wayne", "White", "Whiteside", "Will", "Williamson", "Winnebago", "Woodford",
	},
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
	"IA": {"Adair", "Adams", "Allamakee", "Appanoose", "Audubon", "Benton", "Black Hawk", "Boone", "Bremer", "Buchanan", "Buena Vista", "Butler",
		"Calhoun", "Carroll", "Cass", "Cedar", "Cerro Gordo", "Cherokee", "Chickasaw", "Clarke", "Clay", "Clayton", "Clinton", "Crawford", "Dallas",
		"Davis", "Decatur", "Delaware", "Des Moines", "Dickinson", "Dubuque", "Emmet", "Fayette", "Floyd", "Franklin", "Fremont", "Greene", "Grundy",
		"Guthrie", "Hamilton", "Hancock", "Hardin", "Harrison", "Henry", "Howard", "Humboldt", "Ida", "Iowa", "Jackson", "Jasper", "Jefferson",
		"Johnson", "Jones", "Keokuk", "Kossuth", "Lee", "Linn", "Louisa", "Lucas", "Lyon", "Madison", "Mahaska", "Marion", "Marshall", "Mills",
		"Mitchell", "Monona", "Monroe", "Montgomery", "Muscatine", "O'Brien", "Osceola", "Page", "Palo Alto", "Plymouth", "Pocahontas", "Polk",
		"Pottawattamie", "Poweshiek", "Ringgold", "Sac", "Scott", "Shelby", "Sioux", "Story", "Tama", "Taylor", "Union", "Van Buren", "Wapello",
		"Warren", "Washington", "Wayne", "Webster", "Winnebago", "Winneshiek", "Woodbury", "Worth", "Wright",
	},
	"KS": {"Allen", "Anderson", "Atchison", "Barber", "Barton", "Bourbon", "Brown", "Butler", "Chase", "Chautauqua", "Cherokee", "Cheyenne", "Clark", "Clay",
		"Cloud", "Coffey", "Comanche", "Cowley", "Crawford", "Decatur", "Dickinson", "Doniphan", "Douglas", "Edwards", "Elk", "Ellis", "Ellsworth",
		"Finney", "Ford", "Franklin", "Geary", "Gove", "Graham", "Grant", "Gray", "Greeley", "Greenwood", "Hamilton", "Harper", "Harvey", "Haskell",
		"Hodgeman", "Jackson", "Jefferson", "Jewell", "Johnson", "Kearny", "Kingman", "Kiowa", "Labette", "Lane", "Leavenworth", "Lincoln", "Linn",
		"Logan", "Lyon", "Marion", "Marshall", "McPherson", "Meade", "Miami", "Mitchell", "Montgomery", "Morris", "Morton", "Nemaha", "Neosho", "Ness",
		"Norton", "Osage", "Osborne", "Ottawa", "Pawnee", "Phillips", "Pottawatomie", "Pratt", "Rawlins", "Reno", "Republic", "Rice", "Riley", "Rooks",
		"Rush", "Russell", "Saline", "Scott", "Sedgwick", "Seward", "Shawnee", "Sheridan", "Sherman", "Smith", "Stafford", "Stanton", "Stevens", "Sumner",
		"Thomas", "Trego", "Wabaunsee", "Wallace", "Washington", "Wichita", "Wilson", "Woodson", "Wyandotte",
	},
	"KY": {"Adair", "Allen", "Anderson", "Ballard", "Barren", "Bath", "Bell", "Boone", "Bourbon", "Boyd", "Boyle", "Bracken", "Breathitt", "Breckinridge",
		"Bullitt", "Butler", "Caldwell", "Calloway", "Campbell", "Carlisle", "Carroll", "Carter", "Casey", "Christian", "Clark", "Clay", "Clinton",
		"Crittenden", "Cumberland", "Daviess", "Edmonson", "Elliott", "Estill", "Fayette", "Fleming", "Floyd", "Franklin", "Fulton", "Gallatin", "Garrard",
		"Grant", "Graves", "Grayson", "Green", "Greenup", "Hancock", "Hardin", "Harlan", "Harrison", "Hart", "Henderson", "Henry", "Hickman", "Hopkins", "Jackson",
		"Jefferson", "Jessamine", "Johnson", "Kenton", "Knott", "Knox", "Larue", "Laurel", "Lawrence", "Lee", "Leslie", "Letcher", "Lewis", "Lincoln", "Livingston",
		"Logan", "Lyon", "Madison", "Magoffin", "Marion", "Marshall", "Martin", "Mason", "McCracken", "McCreary", "McLean", "Meade", "Menifee", "Mercer", "Metcalfe",
		"Monroe", "Montgomery", "Morgan", "Muhlenberg", "Nelson", "Nicholas", "Ohio", "Oldham", "Owen", "Owsley", "Pendleton", "Perry", "Pike", "Powell", "Pulaski",
		"Robertson", "Rockcastle", "Rowan", "Russell", "Scott", "Shelby", "Simpson", "Spencer", "Taylor", "Todd", "Trigg", "Trimble", "Union", "Warren",
		"Washington", "Wayne", "Webster", "Whitley", "Wolfe", "Woodford",
	},
	"LA": {"Acadia Parish", "Allen", "Allen Parish", "Ascension Parish", "Assumption", "Assumption Parish", "Avoyelles",
		"Avoyelles Parish", "Beauregard", "Beauregard Parish", "Bienville", "Bienville Parish", "Bossier",
		"Bossier Parish", "Caddo", "Caddo Parish", "Calcasieu Parish", "Caldwell", "Caldwell Parish", "Cameron Parish",
		"Catahoula", "Catahoula Parish", "Central Plaquemines", "Central Tangipahoa", "Claiborne", "Claiborne Parish",
		"Coastal Jefferson", "Concordia", "Concordia Parish", "De Soto", "De Soto Parish,", "East Baton Rouge", "East Baton Rouge Parish",
		"East Cameron", "East Caroll", "East Carroll Parish", "East Feliciana", "East Feliciana Parish", "Eastern Ascension", "Eastern Orleans",
		"Evangeline", "Evangeline Parish", "Franklin", "Franklin Parish", "Grant", "Grant Parish", "Iberia", "Iberia Parish",
		"Iberville", "Iberville Parish", "Jackson", "Jackson Parish", "Jefferson Davis Parish", "Jefferson Parish", "La Salle", "La Salle Parish",
		"Lafayette", "Lafayette Parish", "Lafourche Parish", "Lincoln", "Lincoln Parish", "Livingston Parish",
		"Lower Iberia", "Lower Jefferson", "Lower Lafourche", "Lower Plaquemines", "Lower St. Bernard", "Lower St. Martin", "Lower St. Mary",
		"Lower Tangipahoa", "Lower Terrebonne", "Lower Vermilion", "Madison", "Madison Parish", "Morehouse", "Morehouse Parish", "Natchitoches", "Natchitoches Parish",
		"Northern Acadia", "Northern Calcasieu", "Northern Jefferson Davis", "Northern Livingston", "Northern St. Tammany", "Northern Tangipahoa",
		"Orleans Parish", "Ouachita", "Ouachita Parish", "Plaquemines Parish", "Pointe Coupee", "Pointe Coupee Parish", "Rapides", "Rapides Parish",
		"Red River", "Red River Parish", "Richland", "Richland Parish", "Sabine", "Sabine Parish", "Southeast St. Tammany",
		"Southern Acadia", "Southern Calcasieu", "Southern Jefferson Davis", "Southern Livingston", "Southwestern St. Tammany",
		"St. Bernard Parish", "St. Charles", "St. Charles Parish", "St. Helena", "St. Helena Parish",
		"St. James", "St. James Parish", "St. John The Baptist", "St. John The Baptist Parish", "St. Landry",
		"St. Landry Parish", "St. Martin Parish", "St. Mary Parish", "St. Tammany Parish", "Tangiphahoa Parish", "Tensas", "Tensas Parish", "Terrebonne Parish",
		"Union Parish", "Upper Iberia", "Upper Jefferson", "Upper Lafourche", "Upper Plaquemines", "Upper St.Bernard", "Upper St. Martin",
		"Upper St. Mary", "Upper Terrebonne", "Upper Vermilion", "Vermilion Parish", "Vernon", "Vernon Parish", "Washington", "Washington Parish",
		"Webster", "Webster Parish", "West Baton Rouge", "West Baton Rouge Parish", "West Cameron", "West Carroll", "West Carroll Parish",
		"West Feliciana", "West Feliciana Parish", "Western Ascension", "Western Orleans", "Winn", "Winn Parish",
	},
	"ME": {"Androscoggin", "Aroostook", "Central Interior Cumberland", "Central Penobscot", "Central Piscataquis", "Central Somerset", "Central Washington",
		"Coastal Cumberland", "Coastal Hancock", "Coastal Waldo", "Coastal Washington", "Coastal York", "Cumberland", "Franklin", "Hancock",
		"Interior Cumberland Highlands", "Interior Hancock", "Interior Waldo", "Interior York", "Kennebec", "Knox", "Lincoln",
		"Northeast Aroostook", "Northern Franklin", "Northern Oxford", "Northern Penobscot", "Northern Piscataquis", "Northern Somerset",
		"Northern Washington", "Northwest Aroostook", "Oxford", "Penobscot", "Piscataquis",
		"Sagadahoc", "Somerset", "Southeast Aroostook", "Southern Franklin", "Southern Oxford", "Southern Penobscot",
		"Southern Piscataquis", "Southern Somerset", "Waldo", "Washington", "York",
	},
	"MD": {"Allegany", "Anne Arundel", "Baltimore City", "Baltimore", "Calvert", "Caroline", "Carroll", "Cecil",
		"Central and Eastern Allegany", "Central and Southeast Howard", "Central and Southeast Montgomery", "Charles", "Dorchester", "Extreme Western Allegany", "Frederick",
		"Garrett", "Harford", "Howard", "Inland Worcester", "Kent", "Maryland Beaches", "Montgomery", "Northern Baltimore", "Northwest Harford",
		"Northwest  Howard", "Northwest Montgomery", "Prince Georges", "Queen Anne's", "Somerset", "Southeast Harford",
		"Southern Baltimore", "St. Marys", "Talbot", "Washington", "Wicomico", "Worcester",
	},
	"MA": {"Barnstable", "Berkshire", "Bristol", "Central Middlesex County", "Dukes", "Essex", "Eastern Essex", "Eastern Franklin", "Eastern Hampden",
		"Eastern Hampshire", "Eastern Norfolk", "Eastern Plymouth", "Franklin", "Hampden", "Hampshire", "Middlesex", "Nantucket", "Norfolk",
		"Northern Berkshire", "Northern Bristol", "Northern Worcester", "Northwest Middlesex", "Plymouth", "Southern Berkshire", "Southern Bristol",
		"Southern Plymouth", "Southern Worcester", "Suffolk", "Western Essex", "Western Franklin", "Western Hampden", "Western Hampshire", "Western Norfolk",
		"Worcester"},
	"MI": {"Alcona", "Alger", "Allegan", "Alpena", "Antrim", "Arenac", "Baraga", "Barry", "Bay",
		"Beaver Island and surrounding islands", "Benzie", "Berrien", "Branch", "Calhoun",
		"Cass", "Central Chippewa", "Charlevoix", "Cheboygan", "Chippewa", "Clare", "Clinton", "Crawford", "Delta", "Dickinson",
		"Eastern Mackinac", "Eaton", "Emmet",
		"Genesee", "Gladwin", "Gogebic", "Grand Traverse", "Gratiot", "Hillsdale", "Houghton", "Huron", "Ingham", "Ionia",
		"Iosco", "Iron", "Isabella", "Jackson", "Kalamazoo", "Kalkaska", "Kent", "Keweenaw", "Lake", "Lapeer", "Leelanau",
		"Lenawee", "Livingston", "Luce", "Mackinac", "Mackinac Island/Bois Blanc Island", "Macomb", "Manistee", "Marquette", "Mason", "Mecosta", "Menominee",
		"Midland", "Missaukee", "Monroe", "Montcalm", "Montmorency", "Muskegon", "Newaygo",
		"Northern Berrien", "Northern Schoolcraft", "Oakland", "Oceana", "Ogemaw",
		"Ontonagon", "Osceola", "Oscoda", "Otsego", "Ottawa", "Presque Isle", "Roscommon", "Saginaw", "Sanilac", "Schoolcraft",
		"Shiawassee", "Southeast Chippewa", "Southern Berrien", "Southern Houghton", "Southern Schoolcraft", "St. Clair",
		"St. Joseph", "Tuscola", "Van Buren", "Washtenaw", "Wayne", "Western Chippewa", "Western Mackinac", "Wexford",
	},
	"MN": {"Aitkin", "Anoka", "Becker", "Beltrami", "Benton", "Big Stone", "Blue Earth", "Brown", "Carlton",
		"Carlton/South St. Louis", "Carver", "Cass", "Central St. Louis",
		"Chippewa", "Chisago", "Clay", "Clearwater", "Cook", "Cottonwood", "Crow Wing", "Dakota", "Dodge", "Douglas",
		"East Becker", "East Marshall", "East Otter Tail", "East Polk", "Faribault",
		"Fillmore", "Freeborn", "Goodhue", "Grant", "Hennepin", "Houston", "Hubbard", "Isanti", "Jackson", "Kanabec", "Kandiyohi",
		"Kittson", "Koochiching", "Lac qui Parle", "Lake", "Lake Of The Woods", "Le Sueur", "Lincoln", "Lyon", "Mahnomen", "Marshall", "Martin",
		"McLeod", "Meeker", "Mille Lacs", "Morrison", "Mower", "Murray", "Nicollet", "Nobles", "Norman",
		"North Beltrami", "North Cass", "North Clearwater", "North Itasca", "North St. Louis", "Northern Aitkin", "Northern Cook/Northern Lake", "Olmsted", "Otter Tail",
		"Pennington", "Pine", "Pipestone", "Polk", "Pope", "Ramsey", "Red Lake", "Redwood", "Renville", "Rice", "Rock", "Roseau",
		"Scott", "Sherburne", "Sibley", "South Aitkin", "South Beltrami", "South Cass", "South Clearwater", "South Itasca",
		"Southern Cook/North Shore", "Southern Lake/North Shore", "St. Louis", "Stearns", "Steele", "Stevens", "Swift", "Todd", "Traverse", "Wabasha",
		"Wadena", "Waseca", "Washington", "Wantonwan", "West Becker", "West Marshall", "West Otter Tail", "West Polk", "Wilkin", "Winona", "Wright", "Yellow Medicine",
	},
	"MS": {"Adams", "Alcorn", "Amite", "Attala", "Benton", "Bolivar", "Calhoun", "Carroll", "Chickasaw", "Choctaw", "Claiborne",
		"Clarke", "Clay", "Coahoma", "Copiah", "Covington", "DeSoto", "Forrest", "Franklin", "George", "Greene", "Grenada",
		"Hancock", "Harrison", "Hinds", "Holmes", "Humphreys", "Issaquena", "Itawamba", "Jackson", "Jasper", "Jefferson",
		"Jefferson Davis", "Jones", "Kemper", "Lafayette", "Lamar", "Lauderdale", "Lawrence", "Leake", "Lee", "Leflore",
		"Lincoln", "Lowndes", "Madison", "Marion", "Marshall", "Monroe", "Montgomery", "Neshoba", "Newton",
		"Northern Hancock", "Northern Harrison", "Northern Jackson", "Noxubee",
		"Oktibbeha", "Panola", "Pearl River", "Perry", "Pike", "Pontotoc", "Prentiss", "Quitman", "Rankin", "Scott", "Sharkey",
		"Simpson", "Smith", "Southern Hancock", "Southern Harrison", "Southern Jackson", "Stone", "Sunflower",
		"Tallahatchie", "Tate", "Tippah", "Tishomingo", "Tunica", "Union", "Walthall",
		"Warren", "Washington", "Wayne", "Webster", "Wilkinson", "Winston", "Yalobusha", "Yazoo",
	},
	"MO": {"Adair", "Andrew", "Atchison", "Audrain", "Barry", "Barton", "Bates", "Benton", "Bollinger", "Boone", "Buchanan",
		"Butler", "Caldwell", "Callaway", "Camden", "Cape Girardeau", "Carroll", "Carter", "Cass", "Cedar", "Chariton",
		"Christian", "Clark", "Clay", "Clinton", "Cole", "Cooper", "Crawford", "Dade", "Dallas", "Daviess", "DeKalb", "Demt",
		"Douglas", "Dunklin", "Franklin", "Gasconade", "Gentry", "Greene", "Grundy", "Harrison", "Henry", "Hickory", "Holt",
		"Howard", "Howell", "Iron", "Jackson", "Jasper", "Jefferson", "Johnson", "Knox", "Laclede", "Lafayette", "Lawrence",
		"Lewis", "Lincoln", "Linn", "Livingston", "Macon", "Madison", "Maries", "Marion", "McDonald", "Mercer", "Miller",
		"Mississippi", "Moniteau", "Monroe", "Montgomery", "Morgan", "New Madrid", "Newton", "Nodaway", "Oregon", "Osage",
		"Ozark", "Pemiscot", "Perry", "Pettis", "Phelps", "Pike", "Platte", "Polk", "Pulaski", "Putnam", "Ralls", "Randolph",
		"Ray", "Reynolds", "Ripley", "Saline", "Schuyler", "Scotland", "Scott", "Shannon", "Shelby", "St. Charles",
		"St. Clair", "St. Francois", "St. Louis City", "St. Louis", "Ste. Genevieve", "Stoddard", "Stone", "Sullivan",
		"Taney", "Texas", "Vernon", "Warren", "Washington", "Wayne", "Webster", "Worth", "Wright",
	},
	"MT": {"Absaroka/Beartooth Mountains", "Bears Paw Mountains and Southern Blaine", "Beartooth Foothills", "Beaverhead",
		"Beaverhead and Western Madison below 6000ft", "Big Belt, Bridger and Castle Mountains", "Big Horn", "Bighorn Canyon",
		"Bitterroot/Sapphire Mountains", "Blaine", "Broadwater", "Butte/Blackfoot Region", "Canyon Ferry Area", "Carbon",
		"Carter", "Cascade", "Cascade County below 5000ft", "Central and Southeast Phillips", "Central and Southern Valley",
		"Chouteau", "Crazy Mountains", "Custer", "Daniels", "Dawson", "Deer Lodge",
		"East Glacier Park Region", "Eastern Glacier, Western Toole, and Central Pondera", "Eastern Pondera and Eastern Teton",
		"Eastern Roosevelt", "Eastern Toole and Liberty", "Elkhorn and Boulder Mountains", "Fallon", "Fergus",
		"Fergus County below 4500ft", "Flathead", "Flathead/Mission Valleys", "Gallatin",
		"Gallatin and Madison County Mountains and Centennial Mountains", "Gallatin Valley", "Garfield",
		"Gates of the Mountains", "Glacier", "Golden Valley", "Granite", "Helena Valley",
		"Hill", "Jefferson", "Judith Basin", "Judith Basin County and Judith Gap", "Judith Gap",
		"Kootenai/Cabinet Region", "Lake", "Lewis and Clark", "Liberty", "Lincoln", "Little Belt and Highwood Mountains",
		"Livingston Area", "Lower Clark Fork Region", "Madison", "Madison River Valley", "McCone", "Meagher", "Meagher County Valleys",
		"Melville Foothills", "Mineral", "Missoula", "Missoula/Bitterroot Valleys", "Missouri Headwaters", "Musselshell",
		"Northeastern Yellowstone", "Northern Big Horn", "Northern Blaine County", "Northern Carbon", "Northern High Plains",
		"Northern Park", "Northern Phillips", "Northern Rosebud", "Northern Stillwater", "Northern Sweet Grass", "Northern Valley",
		"Northwest Beaverhead County", "Paradise Valley", "Park", "Petroleum", "Phillips", "Pondera",
		"Potomac/Seeley Lake Region", "Powder River", "Powell", "Prairie", "Pryor/Northern Bighorn Mountains",
		"Red Lodge Foothills", "Ravalli", "Richland", "Roosevelt", "Rosebud", "Ruby Mountains and Southern Beaverhead Mountains",
		"Sanders", "Sheridan", "Silver Bow", "Snowy and Judith Mountains", "Southeastern Carbon", "Southern Big Horn",
		"Southern High Plans", "Southern Rocky Mountain Front", "Southern Rosebud", "Southern Wheatland",
		"Southwest Phillips", "Southwestern Yellowstone", "Stillwater",
		"Sweet Grass", "Teton", "Toole", "Treasure", "Upper Blackfoot and MacDonald Pass", "Valley",
		"West Glacier Region", "Western and Central Chouteau County", "Western Roosevelt", "Wheatland", "Wibaux", "Yellowstone",
	},
	"NE": {"Adams", "Antelope", "Arthur", "Banner", "Blaine", "Boone", "Box Butte", "Boyd", "Brown", "Buffalo", "Burt", "Butler",
		"Cass", "Cedar", "Chase", "Cherry", "Cheyenne", "Clay", "Colfax", "Cuming", "Custer", "Dakota", "Dawes", "Dawson",
		"Deuel", "Dixon", "Dodge", "Douglas", "Dundy", "Eastern Cherry", "Fillmore", "Franklin", "Frontier", "Furnas", "Gage", "Garden", "Garfield",
		"Gosper", "Grant", "Greeley", "Hall", "Hamilton", "Harlan", "Hayes", "Hitchcock", "Holt", "Hooker", "Howard", "Jefferson",
		"Johnson", "Kearney", "Keith", "Keya Paha", "Kimball", "Knox", "Lancaster", "Lincoln", "Logan", "Loup", "Madison",
		"McPherson", "Merrick", "Morrill", "Nance", "Nemaha", "Nuckolls", "Otoe", "Pawnee", "Perkins", "Phelps", "Pierce", "Platte",
		"Polk", "Red Willow", "Richardson", "Rock", "Saline", "Sarpy", "Saunders", "Scotts Bluff", "Seward", "Sheridan", "Sherman",
		"Sioux", "Stanton", "Thayer", "Thomas", "Thurston", "Valley", "Washington", "Wayne", "Webster", "Wheeler", "Western Cherry", "York",
	},
	"NV": {"Carson City", "Churchill", "Clark", "Douglas", "Elko", "Esmeralda", "Esmeralda and Central Nye County", "Eureka",
		"Greater Lake Tahoe Area", "Greater Reno-Carson City-Minden Area", "Humboldt", "Lake Mead National Recreation Area", "Lander",
		"Las Vegas Valley", "Lincoln", "Lyon", "Mineral", "Mineral and Southern Lyon Counties",
		"Northeast Clark County", "Northeastern Nye", "Northern Elko", "Northern Lander County and Northern Eureka County",
		"Northern Washoe County", "Northern Nye County", "Nye", "Pershing", "Rush Mountains and East Humboldt Range",
		"Sheep Range", "South Central Elko County", "Southeastern Elko County", "Southern Clark County",
		"Southern Lander County and Southern Eureka County", "Southwest Elko County", "Spring Mountains-Red Rock Canyon", "Storey",
		"Washoe", "Western Clark an Southern Nye County", "Western Nevada Basin and Range including Pyramid Lake", "White Pine",
	},
	"NH": {"Belknap", "Carroll", "Cheshire", "Coastal Rockingham", "Coos", "Eastern Hillsborough", "Grafton", "Hillsborough",
		"Interiror Rockingham", "Merrimack", "Northern Carroll", "Northern Coos", "Northern Grafton", "Rockingham",
		"Southern Carroll", "Southern Coos", "Southern Grafton", "Strafford", "Sullivan", "Western And Central Hillsborough",
	},
	"NJ": {"Atlantic", "Atlantic Coastal Cape May", "Coastal Atlantic", "Coastal Ocean", "Bergen", "Burlington", "Camden", "Cape May", "Cumberland",
		"Easter Bergen", "Eastern Essex", "Eastern Monmouth", "Eastern Passaic", "Eastern Union", "Essex", "Gloucester", "Hudson", "Hunterdon",
		"Mercer", "Middlesex", "Monmouth", "Morris", "Northwestern Burlington", "Ocean", "Passaic", "Salem", "Somerset",
		"Southeastern Burlington", "Sussex", "Union", "Warren", "Western Bergen", "Western Essex", "Western Monmouth", "Western Passaic", "Western Union",
	},
	"NM": {"Bernalillo", "Catron", "Central Grant County/Silver City Area", "Central Highlands", "Central Lea County",
		"Chaves", "Chaves County Plains", "Chuska Mountains", "Cibola", "Curry", "De Baca", "Dona Ana",
		"East Central Tularosa Basin/Alamogordo", "East Slopes Sacramento Mountains Below 7500 Feet", "East Slopes Sangre de Cristo Mountains",
		"Eastern Black Range Foothills", "Eastern Lincoln County", "Eastern San Miguel County", "Eddy", "Eddy County Plains",
		"Espanola Valley", "Estancia Valley", "Far Northeast Highlands", "Far Northwest Highlands", "Glorieta Mesa Including Glorieta Pass",
		"Grant", "Guadalupe", "Guadalupe Mountains of Eddy County", "Harding", "Hidalgo", "Jemez Mountains",
		"Johnson and Bartlett Mesas Including Raton Pass", "Lea", "Lincoln", "Los Alamos",
		"Lower Rio Grande Valley", "Lowlands of the Bootheel", "Luna", "Middle Rio Grande Valley/Albuquerque Metro Area",
		"McKinley", "Mora", "Northeast Highlands", "Northern Dona Ana County", "Northern Lea County", "Northern Sangre de Cristo Mountains",
		"Northwest Highlands", "Northwest Plateau", "Otero", "Otero Mesa", "Quay", "Rio Arriba", "Roosevelt",
		"Sacramento Mountains Above 7500 Feet", "San Agustin Plains and Adjacent Lowlands", "San Francisco River Valley", "San Juan",
		"San Miguel", "Sandia/Manzano Mountains Including Edgewood", "Sandoval", "Santa Fe", "Santa Fe Metro Area", "Sierra",
		"Sierra County Lakes", "Socorro", "South Central Highlands", "South Central Mountains", "Southeast Tularosa Basin",
		"Southern Dona Ana County/Mesilla Valley", "Southern Gila Foothills/Mimbres Valley", "Southern Gila Region Highlands/Black Range",
		"Southern Lea County", "Southern Sangre de Cristo Mountains", "Southwest Chaves County", "Southwest Desert/Lower Gila River Valley",
		"Southwest Desert/Mimbres Basin", "Southwest Mountains", "Taos", "Torrance", "Tusas Mountains Including Chama", "Union",
		"Uplands of the Bootheel", "Upper Gila River Valley,", "Upper Rio Grande Valley", "Upper Tularosa Valley", "Valencia",
		"West Central Highlands", "West Central Mountains", "West Central Plateau", "West Central Tularosa Basin/White Sands",
		"West Slopes Sacramento Mountains Below 7500 Feet",
	},
	"NY": {"Albany", "Allegany", "Bronx", "Broome", "Cattaraugus", "Cayuga", "Chautauqua", "Chemung", "Chenango", "Clinton", "Columbia",
		"Cortland", "Delaware", "Dutchess", "Eastern Albany", "Eastern Clinton", "Eastern Columbia", "Eastern Dutchess",
		"Eastern Essex", "Eastern Greene", "Eastern Rensselaer", "Eastern Schenectady", "Eastern Ulster", "Erie",
		"Essex", "Franklin", "Fulton", "Genesee", "Greene", "Hamilton", "Herkimer", "Jefferson", "Kings", "Kings (Brooklyn)", "Lewis",
		"Livingston", "Madison", "Monroe", "Montgomery", "Nassau", "New York (Manhattan)", "Niagra", "Northeast Suffolk", "Northern Cayuga",
		"Northern Erie", "Northern Franklin", "Northern Fulton", "Northern Herkimer", "Northern Nassau", "Northern Oneida",
		"Northern Queens", "Northern Saratoga", "Northern St. Lawrence", "Northern Warren", "Northern Wasington", "Northern Westchester",
		"Northwest Suffolk", "Oneida", "Onondaga", "Ontario", "Orange", "Orleans", "Oswego", "Otsego", "Putnam", "Queens", "Rensselaer", "Richmond",
		"Richmond (Staten Is.)", "Rockland", "Saratoga", "Schenectady", "Schoharie", "Schuyler", "Seneca", "Southeast Suffolk",
		"Southeast Warren", "Southeastern St. Lawrence", "Southern Cayuga", "Southern Erie", "Southern Franklin", "Southern Fulton",
		"Southern Herkimer", "Southern Nassau", "Southern Oneida", "Southern Queens", "Southern Saratoga", "Southern Washington",
		"Southern Westchester", "Southwest Suffolk", "Southwestern St. Lawrence", "St. Lawrence", "Steuben", "Suffolk", "Sullivan",
		"Tioga", "Tompkins", "Ulster", "Warren", "Washington", "Wayne", "Westchester", "Western Albany", "Western Clinton",
		"Western Columbia", "Western Dutchess", "Western Essex", "Western Greene", "Western Rensselaer",
		"Western Schenectady", "Western Ulster", "Wyoming", "Yates",
	},
	"NC": {"Alamance", "Alexander", "Alleghany", "Anson", "Ashe", "Avery", "Beaufort", "Bertie", "Bladen", "Brunswick", "Buncombe",
		"Burke", "Burke Mountains", "Cabarrus", "Caldwell", "Caldwell Mountains", "Camden", "Carteret", "Caswell", "Catawba", "Chatham",
		"Cherokee", "Chowan", "Clay", "Cleveland", "Coastal Brunswick", "Coastal New Hanover", "Coastal Onslow",
		"Coastal Pender", "Columbus", "Craven", "Cumberland", "Currituck", "Dare", "Davidson", "Davie", "Duplin", "Durham",
		"East Carteret", "Eastern Currituck", "Eastern McDowell", "Eastern Polk", "Edgecombe", "Forsyth", "Franklin", "Gaston",
		"Gates", "Graham", "Granville", "Greater Burke", "Greater Caldwell", "Greater Rutherford", "Greene", "Guilford", "Halifax",
		"Harnett", "Hatteras Island", "Haywood", "Henderson", "Hertford", "Hoke", "Hyde",
		"Inland Brunswick", "Inland New Hanover", "Inland Onslow", "Inland Pender", "Iredell", "Jackson",
		"Johnston", "Jones", "Lee", "Lenoir", "Lincoln", "Macon", "Madison", "Mainland Dare", "Mainland Hyde", "Martin", "McDowell",
		"McDowell Mountains", "Mecklenburg", "Mitchell", "Montgomery", "Moore", "Nash",
		"New Hanover", "Northampton", "Northern Craven", "Northern Jackson", "Northern Outer Banks", "Ocracoke Island",
		"Onslow", "Orange", "Pamlico", "Pasquotank", "Pender", "Perquimans", "Person", "Pitt", "Polk", "Polk Mountains",
		"Randolph", "Richmond", "Robeson", "Rockingham", "Rowan", "Rutherford", "Rutherford Mountains", "Sampson", "Scotland",
		"Southern Craven", "Southern Jackson", "Stanly", "Stokes", "Surry", "Swain", "Transylvania", "Tyrrell", "Union", "Vance",
		"Wake", "Warren", "Washington", "Watauga", "Wayne", "West Carteret", "Western Currituck",
		"Wilkes", "Wilson", "Yadkin", "Yancey",
	},
	"ND": {"Adams", "Barnes", "Benson", "Billings", "Bottineau", "Bowman", "Burke", "Burleigh", "Cass", "Cavalier", "Dickey",
		"Divide", "Dunn", "Eastern Walsh County", "Eddy", "Emmons", "Foster", "Golden Valley", "Grand Forks", "Grant", "Griggs", "Hettinger", "Kidder",
		"LaMoure", "Logan", "McHenry", "McIntosh", "McKenzie", "McLean", "Mercer", "Morton", "Mountrail", "Nelson", "Oliver",
		"Pembina", "Pierce", "Ramsey", "Ransom", "Renville", "Richland", "Rolette", "Sargent", "Sheridan", "Sioux", "Slope",
		"Stark", "Steele", "Stutsman", "Towner", "Traill", "Walsh", "Ward", "Wells", "Western Walsh County", "Williams",
	},
	"OH": {"Adams", "Allen", "Ashland", "Ashtabula", "Ashtabula Inland", "Ashtabula Lakeshore", "Athens", "Auglaize", "Belmont", "Brown",
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
	"OK": {"Adair", "Alfalfa", "Atoka", "Beaver", "Beckham", "Blaine", "Bryan", "Caddo", "Canadian", "Carter", "Cherokee",
		"Choctaw", "Cimarron", "Cleveland", "Coal", "Comanche", "Cotton", "Craig", "Creek", "Custer", "Delaware", "Dewey",
		"Ellis", "Garfield", "Garvin", "Grady", "Grant", "Greer", "Harmon", "Harper", "Haskell", "Hughes", "Jackson",
		"Jefferson", "Johnston", "Kay", "Kingfisher", "Kiowa", "Latimer", "Le Flore", "Lincoln", "Logan", "Love", "Major",
		"Marshall", "Mayes", "McClain", "McCurtain", "McIntosh", "Murray", "Muskogee", "Noble", "Nowata", "Okfuskee",
		"Oklahoma", "Okmulgee", "Osage", "Ottawa", "Pawnee", "Payne", "Pittsburg", "Pontotoc", "Pottawatomie",
		"Pushmataha", "Roger Mills", "Rogers", "Seminole", "Sequoyah", "Stephens", "Texas", "Tillman", "Tulsa", "Wagoner",
		"Washington", "Washita", "Woods", "Woodward",
	},
	"OR": {"Baker", "Benton", "Benton County Lowlands", "Cascade Foothills of Marion and Linn Counties",
		"Cascades of Lane County", "Cascades of Marion and Linn Counties", "Central and Eastern Lake County", "Central Coast of Oregon",
		"Central Columbia River Gorge I-84 Corridor", "Central Douglas County", "Central Oregon Coast Range Lowlands",
		"Central Oregon Coast Range", "Central Oregon", "Clackamas", "Clackamas County Cascade Foothills", "Clatsop",
		"Clatsop County Coast", "Columbia", "Coos", "Crook", "Curry", "Curry County Coast", "Deschutes", "Douglas",
		"East Central Willamette Valley", "East Portland Metro", "East Slopes of the Oregon Cascades",
		"Eastern Columbia River Gorge of Oregon", "Eastern Curry County and Josephine County", "Eastern Douglas County Foothills",
		"Foothills of the Northern Blue Mountains of Oregon", "Foothills of the Southern Blue Mountains of Oregon", "Gilliam",
		"Grande Ronde Valley", "Grant", "Harney", "Hood River", "Inner Portland Metro", "Jackson", "Jefferson",
		"John Day Basin", "Josephine", "Klamath", "Klamath Basin", "Lake", "Lane", "Lane County Cascade Foothills",
		"Lane County Lowlands", "Lincoln", "Linn", "Linn County Lowlands", "Lower Columbia Basin of Oregon", "Lower Columbia River",
		"Malheur", "Marion", "Marrow", "Multnomah", "North Central Oregon", "North Oregon Cascades",
		"North Oregon Coast Range Lowlands", "North Oregon Coast Range", "Northern and Eastern Klamath County and Western Lake County",
		"Northern Blue Mountains of Oregon", "Ochoco-John Day Highlands", "Oregon Lower Treasure Valley", "Outer Southeast Portland Metro",
		"Polk", "Sherman", "Siskiyou Mountains and Southern Oregon Cascades", "South Central Oregon Cascades",
		"South Central Oregon Coast", "Southern Blue Mountains of Oregon", "Tillamook", "Tillamook County Coast", "Tualatin Valley",
		"Umatilla", "Union", "Upper Hood River Valley", "Wallowa", "Wasco", "Washington",
		"West Central Willamette Valley", "West Columbia River Gorge I-84 Corridor",
		"West Columbia River Gorge of Oregon above 500 ft", "West Hills and Chehalem Mountains", "Wheeler", "Yamhill",
	},
	"PA": {"Adams", "Allegheny", "Armstrong", "Beaver", "Bedford", "Berks", "Blair", "Bradford", "Bucks", "Butler", "Cambria",
		"Cameron", "Carbon", "Centre", "Chester", "Clarion", "Clearfield", "Clinton", "Columbia", "Crawford", "Cumberland",
		"Dauphin", "Delaware", "Eastern Chester", "Eastern Montgomery", "Elk", "Erie", "Fayette", "Fayette Ridges",
		"Forest", "Franklin", "Fulton", "Greene", "Higher Elevations of Indiana", "Huntingdon", "Indiana",
		"Jefferson", "Juniata", "Lackawanna", "Lancaster", "Lawrence", "Lebanon", "Lehigh", "Lower Bucks", "Luzerne", "Lycoming", "McKean",
		"Mercer", "Mifflin", "Monroe", "Montgomery", "Montour", "Northampton", "Northumberland", "Northern Centre",
		"Northern Clinton", "Northern Erie", "Northern Lycoming", "Northern Wayne", "Perry", "Philadelphia",
		"Pike", "Potter", "Schuylkill", "Snyder", "Somerset", "Southern Centre", "Southern Clinton", "Southern Erie",
		"Southern Lycoming", "Southern Wayne", "Sullivan", "Susquehanna", "Tioga", "Union", "Upper Bucks", "Venango",
		"Warren", "Washington", "Wayne", "Western Chester", "Western Montgomery", "Westmoreland", "Westmoreland Ridges", "Wyoming", "York",
	},
	"RI": {"Block Island", "Bristol", "Eastern Kent", "Kent", "Newport", "Northwest Providence", "Providence",
		"Southeast Providence", "Washington", "Western Kent",
	},
	"SC": {"Abbeville", "Aiken", "Allendale", "Anderson", "Bamberg", "Barnwell", "Beaufort", "Berkeley", "Calhoun",
		"Central Greenville", "Central Horry", "Central Orangeburg", "Charleston",
		"Cherokee", "Chester", "Chesterfield", "Clarendon", "Coastal Colleton", "Coastal Georgetown", "Coastal Horry",
		"Coastal Jasper", "Colleton", "Darlington", "Dillon", "Dorchester", "Edgefield",
		"Fairfield", "Florence", "Greater Oconee", "Greater Pickens", "Greater Mountains", "Georgetown", "Greenville",
		"Greenwood", "Hampton", "Horry", "Inland Berkeley", "Inland Colleton", "Inland Jasper", "Jasper", "Kershaw", "Lancaster",
		"Laurens", "Lee", "Lexington", "Marion", "Marlboro", "McCormick", "Newberry", "Northern Horry",
		"Northern Lancaster", "Northern Spartanburg", "Northern Orangeburg", "Oconee", "Oconee Mountains", "Orangeburg", "Pickens",
		"Pickens Mountains", "Richland", "Saluda", "Southeastern Orangeburg", "Southeastern Greenville", "Southeastern Lancaster",
		"Southeastern Spartanburg", "Spartanburg", "Sumter", "Tidal Berkeley", "Union", "Williamsburg", "York",
	},
	"SD": {"Aurora", "Beadle", "Bennett", "Bon Homme", "Brookings", "Brown", "Brule", "Buffalo", "Butte", "Campbell", "Central Black Hills", "Charles Mix",
		"Clark", "Clay", "Codington", "Corson", "Custer", "Custer Co Plains", "Davison", "Day", "Deuel", "Dewey", "Douglas", "Eastern Fall River", "Edmunds", "Fall River",
		"Faulk", "Grant", "Gregory", "Haakon", "Hamlin", "Hand", "Hanson", "Harding", "Hermosa Foothills", "Hughes", "Hutchinson", "Hyde", "Jackson",
		"Jerauld", "Jones", "Kingsbury", "Lake", "Lawrence", "Lincoln", "Lyman", "Marshall", "McCook", "McPherson", "Meade",
		"Mellette", "Miner", "Minnehaha", "Moody", "Northern Black Hills", "Northern Foothills", "Northern Jackson",
		"Northern Meade Co Plains", "Northern Oglala Lakota", "Northern Perkins", "Pennington Co Plains", "Oglala Lakota", "Pennington", "Perkins", "Perkins", "Potter", "Roberts",
		"Sanborn", "Southern Black Hills", "Southern Foothills", "Southern Jackson", "Southern Meade Co Plains",
		"Southern Oglala Lakota", "Southern Perkins", "Spink", "Stanley", "Sturgis/Piedmont Foothills", "Sully", "Todd",
		"Tripp", "Turner", "Union", "Walworth", "Western Fall River", "Yankton", "Zieback",
	},
	"TN": {"Anderson", "Bedford", "Benton", "Bledsoe", "Blount", "Blount Smoky Mountains", "Bradley", "Campbell", "Cannon", "Carroll", "Carter", "Cheatham",
		"Chester", "Claiborne", "Clay", "Cocke Smoky Mountains", "Cocke", "Coffee", "Crockett", "Cumberland", "De Kalb", "Decatur", "Dickson", "Dyer", "East Polk",
		"Fayette", "Fentress", "Franklin", "Gibson", "Giles", "Grainger", "Greene", "Grundy", "Hamblen", "Hamilton", "Hancock",
		"Hardeman", "Hardin", "Hawkins", "Haywood", "Henderson", "Henry", "Hickman", "Houston", "Humphreys", "Jackson",
		"Jefferson", "Johnson", "Knox", "Lake", "Lauderdale", "Lawrence", "Lewis", "Lincoln", "Loudon", "Macon", "Madison",
		"Marion", "Marshall", "Maury", "McMinn", "McNairy", "Meigs", "Monroe", "Montgomery", "Moore", "Morgan",
		"North Sevier", "Northwwest Carter", "Northwest Cocke", "Northwest Greene", "Northwest Monroe", "NW Blount", "Obion",
		"Overton", "Perry", "Pickett", "Polk", "Putnam", "Rhea", "Roane", "Robertson", "Rutherford", "Scott", "Sequatchie",
		"Sevier Smoky Mountains", "Sevier", "Shelby", "Smith", "Southeast Carter", "Southeast Greene", "Southeast Monroe",
		"Stewart", "Sullivan", "Sumner", "Tipton", "Trousdale", "Unicoi", "Union", "Van Buren",
		"Warren", "Washington", "Wayne", "Weakley", "West Polk", "White", "Williamson", "Wilson",
	},
	"TX": {"Andreson", "Andrews", "Angelina", "Aransas", "Aransas Islands", "Archer", "Armstrong", "Atascosa", "Austin",
		"Bailey", "Bandera", "Bastrop", "Baylor", "Bee", "Bell", "Bexar", "Blanco", "Bolivar Peninsula", "Borden", "Bosque", "Bowie", "Brazoria",
		"Brazos", "Brewster", "Briscoe", "Brooks", "Brown", "Burleson", "Burnet", "Caldwell", "Calhoun", "Calhoun Islands,", "Callahan", "Cameron", "Cameron Island",
		"Brazoria Islands", "Camp", "Carson", "Cass", "Castro", "Central Brewster County", "Chambers", "Cherokee",
		"Childress", "Chinati Mountains", "Chisos Basin", "Clay", "Coastal Aransas", "Coastal Brazoria", "Coastal Calhoun",
		"Coastal Cameron", "Coastal Galveston", "Coastal Harris", "Coastal Jackson", "Coastal Kenedy", "Coastal Kleberg", "Coastal Matagorda",
		"Coastal Nueces", "Coastal Refugio", "Coastal San Patricio", "Coastal Willacy", "Cochran", "Coke", "Coleman", "Collin",
		"Collingsworth", "Colorado", "Comal", "Comanche", "Concho", "Cooke", "Coryell", "Cottle", "Crane", "Crockett", "Crosby",
		"Culberson", "Dallam", "Dallas", "Davis Mountains Foothills", "Davis Mountains", "Dawson", "Deaf Smith", "Delta", "Denton", "DeWitt", "Dickens", "Dimmit", "Donley",
		"Duval", "Eastern Culberson County", "Eastern/Central El Paso County", "Eastland", "Ector", "Edwards", "El Paso", "Ellis", "Erath", "Falls", "Fannin", "Fayette", "Fisher", "Floyd",
		"Foard", "Fort Bend", "Franklin", "Freestone", "Frio", "Gaines", "Galveston", "Galveston Island", "Garza", "Gillespie", "Glasscock",
		"Goliad", "Gonzales", "Gray", "Grayson", "Gregg", "Grimes", "Guadalupe", "Guadalupe and Delaware Mountains",
		"Guadalupe Mountains Above 7000 Feet", "Hale", "Hall", "Hamilton", "Hansford",
		"Hardeman", "Hardin", "Harris", "Harrison", "Hartley", "Haskell", "Hays", "Hemphill", "Henderson", "Hidalgo", "Hill",
		"Hockley", "Hood", "Hopkins", "Houston", "Howard", "Hudspeth", "Hunt", "Hutchinson", "Inland Brazoria",
		"Inland Calhoun", "Inland Cameron", "Inland Galveston", "Inland Harris", "Inland Jackson", "Inland Kenedy",
		"Inland Kleberg", "Inland Matagorda", "Inland Nueces", "Inland Refugio", "Inland San Patricio", "Inland Willacy", "Irion", "Jack", "Jackson",
		"Jasper", "Jeff Davis", "Jefferson", "Jim Hogg", "Jim Wells", "Johnson", "Jones", "Karnes", "Kaufman", "Kendall",
		"Kenedy", "Kenedy Island", "Kent", "Kerr", "Kimble", "King", "Kinney", "Kleberg", "Kleberg Islands", "Knox", "La Salle", "Lamar", "Lamb", "Lampasas",
		"Lavaca", "Lee", "Leon", "Liberty", "Limestone", "Lipscomb", "Live Oak", "Llano", "Loving", "Lower Brewster County",
		"Lower Jefferson", "Lubbock", "Lynn", "Madison", "Marfa Plateau",
		"Marion", "Martin", "Mason", "Matagorda", "Matagorda Islands", "Maverick", "McCulloch", "McLennan", "McMullen", "Medina", "Menard",
		"Midland", "Milam", "Mills", "Mitchell", "Montague", "Montgomery", "Moore", "Morris", "Motley", "Nacogdoches", "Navarro",
		"Newton", "Nolan", "Northern Hidalgo", "Northern Hudspeth Highlands/Hueco Mountains", "Northern Jasper",
		"Northern Liberty", "Northern Newton", "Northern Orange", "Nueces", "Nueces Island", "Ochiltree", "Oldham", "Orange",
		"Palo Duro Canyon", "Palo Pinto", "Panola", "Parker", "Parmer", "Pecos",
		"Polk", "Potter", "Presidio", "Presidio Valley", "Rains", "Randall", "Reagan", "Real", "Red River", "Reeves", "Reeves County Plains", "Refugio",
		"Rio Grande Valley of Eastern El Paso/Western Hudspeth Counties", "Rio Grande Valley of Eastern Hudspeth County", "Roberts",
		"Robertson", "Rockwall", "Runnels", "Rusk", "Sabine", "San Augustine", "San Jacinto", "San Patricio", "San Saba",
		"Schleicher", "Scurry", "Shackelford", "Shelby", "Sherman", "Smith", "Somervell", "Southern Hidalgo",
		"Southern Hudspeth Highlands", "Southern Jasper", "Southern Liberty", "Southern Newton", "Southern Orange", "Starr", "Stephens", "Sterling",
		"Stonewall", "Sutton", "Swisher", "Tarrant", "Taylor", "Terrel", "Terry", "Throckmorton", "Titus", "Tom Green",
		"Travis", "Trinity", "Tyler", "Upper Jefferson", "Upshur", "Upton", "Uvalde", "Val Verde",
		"Van Horn and Highway 54 Corridor", "Van Zandt", "Victoria", "Walker", "Waller",
		"Ward", "Washington", "Webb", "Western El Paso County", "Wharton", "Wheeler", "Wichita", "Wilbarger", "Willacy",
		"Willacy Island", "Williamson", "Wilson", "Winkler",
		"Wise", "Wood", "Yoakum", "Young", "Zapata", "Zavala",
	},
	"UT": {"Arches/Grand Flat", "Bear Lake and Bear River Valley", "Beaver", "Box Elder", "Bryce Canyon Country", "Cache",
		"Cache Valley/Utah Portion", "Canyonlands/Natural Bridges", "Capitol Reef National Park and Vicinity", "Carbon",
		"Castle Country", "Central Mountains", "Daggett", "Davis", "Duchesne", "Eastern Box Elder County",
		"Eastern Juab/Millard Counties", "Eastern Uinta Basin", "Eastern Uinta Mountains", "Emery", "Garfield", "Grand",
		"Glen Canyon Recreation Area/Lake Powell", "Great Salt Lake Desert and Mountains", "Iron",
		"Juab", "Kane", "La Sal and Abajo Mountains", "Lower Washington County", "Millard", "Morgan",
		"Northern Wasatch Front", "Piute", "Rich", "Salt Lake", "Salt Lake Valley", "San Juan",
		"San Rafael Swell", "Sanpete", "Sanpete Valley", "Sevier", "Sevier Valley", "South Central Utah",
		"Southeast Utah", "Southern Mountains", "Southwest Utah", "Summit", "Tavaputs Plateau", "Tooele", "Tooele and Rush Valleys",
		"Uintah", "Upper Sevier River Valleys", "Utah", "Utah Valley", "Wasatch", "Wasatch Back",
		"Wasatch Mountains I-80 North", "Wasatch Mountains South of I-80", "Wasatch Mountains South of I-80", "Washington", "Wayne", "Weber",
		"Western Canyonlands", "Western Millard and Juab Counties", "Western Uinta Basin", "Western Uinta Mountains", "Zion National Park",
	},
	"VT": {"Addison", "Bennington", "Caledonia", "Chittenden", "Eastern Addison", "Eastern Chittenden", "Eastern Franklin",
		"Eastern Rutland", "Eastern Windham", "Eastern Windsor", "Essex", "Franklin", "Grand Isle", "Lamoille", "Orange", "Orleans",
		"Rutland", "Washington", "Windham", "Windsor", "Western Addison", "Western Chittenden", "Western Franklin",
		"Western Rutland", "Western Windham", "Western Windsor",
	},
	"VA": {"Accomack", "Albemarle", "Alleghany", "Amelia", "Amherst", "Appomattox", "Arlington", "Arlington/Falls Church/Alexandria", "Augusta", "Bath", "Bedford",
		"Bland", "Botetourt", "Brunswick", "Buchanan", "Buckingham", "Campbell", "Caroline", "Carroll",
		"Central and Southeast Prince William/Manassas/Manassas Park", "Central Virginia Blue Ridge", "Charles City",
		"Charlotte", "Chesterfield", "City of Alexandria", "City of Bristol", "City of Buena Vista", "City of Charlottesville",
		"City of Chesapeake", "City of Colonial Heights", "City of Convington", "City of Danville", "City of Emporia",
		"City of Fairfax", "City of Falls Church", "City of Franklin", "City of Fredericksburg", "City of Galax",
		"City of Hampton", "City of Harrisonburg", "City of Hopewell", "City of Lexington", "City of Lynchburg",
		"City of Manassas Park", "City of Martinsville", "City of Newport News", "City of Norfolk", "City of Norton",
		"City of Petersburg", "City of Poquoson", "City of Portsmouth", "City of Radford", "City of Richmond", "City of Roanoke",
		"City of Salem", "City of Staunton", "City of Suffolk", "City of Virginia Beach", "City of Waynesboro", "City of Williamsburg",
		"City of Winchester", "Clarke", "Craig", "Culpeper", "Cumberland", "Dickenson", "Dinwiddie", "Eastern Chesterfield (Including Col. Heights)",
		"Eastern Essex", "Eastern Hanover", "Eastern Henrico", "Eastern Highland", "Eastern King and Queen", "Eastern King William",
		"Eastern Loudoun", "Eastern Louisa", "Essex", "Fairfax", "Fauquier",
		"Floyd", "Fluvanna", "Franklin", "Frederick", "Giles", "Gloucester", "Goodchland", "Grayson", "Greene", "Greensville",
		"Halifax", "Hampton/Poquoson", "Hanover", "Henrico", "Henry", "Highland", "Isle of Wight", "James City", "King and Queen", "King George",
		"King William", "Lancaster", "Lee", "Loudoun", "Louisa", "Lunenburg", "Madison", "Mathews", "Mecklenburg", "Middlesex",
		"Montgomery", "Nelson", "New Kent", "Newport News", "Norfolk/Portsmouth", "Northampton", "Northern Fauquier",
		"Northern Virginia Blue Ridge", "Northumberland", "Northwest Prince William", "Nottoway", "Orange", "Page", "Patrick", "Pittsylvania",
		"Powhatan", "Prince Edward", "Prince George", "Prince William", "Pulaski", "Rappahannock", "Richmond", "Roanoke",
		"Rockbridge", "Rockingham", "Russell", "Scott", "Shenandoah", "Smyth", "Southampton", "Southern Fauquier", "Spotsylvania", "Stafford",
		"Surry", "Sussex", "Tazewell", "Virginia Beach", "Warren", "Washington", "Western Chesterfield", "Western Essex",
		"Western Hanover", "Western Henrico (Including the City of Richmond)", "Western Highland", "Western King and Queen",
		"Western King William", "Western Loudoun", "Western Louisa", "Westmoreland", "Wise", "Wythe", "York",
	},
	"WA": {"Adams", "Admiralty Inlet Area", "Asotin", "Benton", "Bellevue and Vicinity", "Bremerton and Vicinity", "Central Chelan County",
		"Central Coast", "Central Columbia River Gorge - SR 14", "Chelan", "Clallam", "Clark", "Columbia", "Cowlitz", "Cowlitz County Lowlands",
		"Douglas", "East Clark County Lowlands", "East Puget Sound Lowlands", "Eastern Columbia River Gorge of Washington",
		"Eastern Strait of Juan de Fuca", "Everett and Vicinity", "Ferry", "Foothills of the Blue Mountains of Washington", "Franklin",
		"Garfield", "Grant", "Grays Harbor", "Hood Canal Area", "Inner Vancouver Metro", "Island", "Jefferson", "King", "Kitsap", "Kittitas",
		"Kittitas Valley", "Klickitat", "Lewis", "Lincoln", "Lower Chehalis Valley Area", "Lower Columbia Basin of Washington",
		"Lower Garfield and Asotin Counties", "Lower Slopes of the Eastern Washington Cascades Crest", "Mason", "Moses Lake Area",
		"North Clark County Lowlands", "North Coast", "Northeast Blue Mountains", "Northeast Mountains", "Northwest Blue Mountains",
		"Okanogan", "Okanogan Highlands", "Okanogan Valley", "Olympics", "Pacific", "Pend Oreille", "Pierce", "San Juan",
		"Seattle and Vicinity", "Simcoe Highlands", "Skagit", "Skamania", "Snohomish", "South Washington Cascade Foothills",
		"South Washington Cascades", "South Washington Coast", "Southwest Interior",
		"Spokane", "Spokane Area", "Stevens", "Tacoma Area", "Thurston", "Upper Columbia Basin",
		"Upper Slopes of the Eastern Washington Cascades Crest", "Wahkiakum", "Walla Walla", "Washington Palouse",
		"Waterville Plateau", "Wenatchee Area", "West Columbia River Gorge - SR 14", "West Slopes North Cascades and Passes",
		"West Slopes North Central Cascades and Passes", "West Slopes South Central Cascades and Passes", "Western Chelan County",
		"Western Okanogan County", "Western Skagit County", "Western Strait of Juan De Fuca", "Western Whatcom County",
		"Whatcom", "Whitman", "Willapa and Wahkiakum Lowlands", "Willapa Hills", "Yakima", "Yakima Valley",
	},
	"WV": {"Barbour", "Berkeley", "Boone", "Braxton", "Brooke", "Cabell", "Calhoun", "Clay", "Doddridge", "Eastern Grant",
		"Eastern Greenbrier", "Eastern Mineral", "Eastern Pendleton", "Eastern Preston", "Eastern Tucker", "Fayette", "Gilmer",
		"Grant", "Greenbrier", "Hampshire", "Hancock", "Hardy", "Harrison", "Jackson", "Jefferson", "Kanawha", "Lewis",
		"Lincoln", "Logan", "McDowell", "Marion", "Marshall", "Mason", "Mercer", "Mineral", "Mingo", "Monongalia", "Monroe",
		"Morgan", "Nicholas", "Northwest Fayette", "Northwest Nicholas", "Northwest Pocahontas", "Northwest Raleight",
		"Northwest Randolph", "Northwest Webster", "Ohio", "Pendleton", "Pleasants", "Pocahontas", "Preston", "Putnam",
		"Ridges of Eastern Monongalia and Northwestern Preston", "Raleigh", "Randolph",
		"Ritchie", "Roane", "Southeast Fayette", "Southeast Nicholas", "Southeast Pocahontas", "Southeast Raleigh",
		"Southeast Randolph", "Southeast Webster", "Summers", "Taylor", "Tucker", "Tyler", "Upshur", "Wayne",
		"Western Grant", "Western Greenbrier", "Western Mineral", "Western Pendleton", "Western Tucker", "Webster", "Wetzel", "Wirt", "Wood",
		"Wyoming",
	},
	"WI": {"Adams", "Ashland", "Barron", "Bayfield", "Brown", "Buffalo", "Burnett", "Calumet", "Chippewa", "Clark", "Columbia",
		"Crawford", "Dane", "Dodge", "Door", "Douglas", "Dunn", "Eau Claire", "Florence", "Fond du Lac", "Forest", "Grant",
		"Green", "Green Lake", "Iowa", "Iron", "Jackson", "Jefferson", "Juneau", "Kenosha", "Kewaunee", "La Crosse", "Lafayette",
		"Langlade", "Lincoln", "Manitowoc", "Marathon", "Marinette", "Marquette", "Menominee", "Milwaukee", "Monroe",
		"Northern Marinette County", "Northern Oconto County", "Oconto",
		"Oneida", "Outagamie", "Ozaukee", "Pepin", "Pierce", "Polk", "Portage", "Price", "Racine", "Richland", "Rock", "Rusk", "Sauk",
		"Sawyer", "Shawano", "Sheboygan", "Southern Marinette County", "Southern Oconto County", "St. Croix", "Taylor",
		"Trempealeau", "Vernon", "Vilas", "Walworth", "Washburn",
		"Washington", "Waukesha", "Waupaca", "Waushara", "Winnebago", "Wood",
	},
	"WY": {"Absaroka Mountains", "Albany", "Big Horn", "Bighorn Mountains Southeast", "Bighorn Mountains West",
		"Campbell", "Carbon", "Casper Mountain", "Central Carbon County", "Central Laramie County", "Central Laramie Range and Southwest Platte County",
		"Cody Foothills", "Converse", "Converse County Lower Elevations", "Crook", "East Laramie County", "East Platte County", "East Sweetwater County",
		"Ferris/Seminoe/Shirley Mountains", "Flaming Gorge", "Fremont", "Goshen", "Green Mountains and Rattlesnake Range", "Hot Springs",
		"Jackson Hole", "Johnson", "Lander Foothills", "Laramie", "Laramie Valley", "Lincoln", "Natrona",
		"Natrona County Lower Elevations", "Newcastle", "Niobrara", "North Bighorn Basin", "North Laramie Range",
		"North Snowy Range Foothills", "Northeast Bighorn Mountains", "Northeast Johnson County", "Northeastern Crook",
		"Northern Campbell", "Owl Creek and Bridger Mountains", "Park", "Platte", "Rock Springs and Green River",
		"Salt River and Wyoming Ranges", "Sheridan", "Sheridan Foothills", "Shirley Basin", "Sierra Madre Range",
		"Snowy Range", "South Laramie Range Foothills", "South Laramie Range", "South Lincoln County",
		"Southeast Bighorn Basin", "Southeast Johnson County", "Southern Campbell", "Southwest Bighorn Basin",
		"Southwest Carbon County", "Southwest Wyoming", "Star Valley", "Sublette", "Sweetwater", "Teton", "Teton and Gros Ventre Mountains",
		"Uinta", "Upper Green River Basin Foothills", "Upper Green River Basin", "Upper North Platte River Basin",
		"Upper Wind River Basin", "Washakie", "Western Crook", "Weston", "Weston County Plains", "Wind River Basin",
		"Wind River Mountains East", "Wind River Mountains West", "Wyoming Black Hills", "Yellowstone National Park",
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
	ID          string `json:"id"`
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

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type wbgtURL struct {
	Properties struct {
		Url string `json:"forecastGridData"`
	} `json:"properties"`
}

type TemperatureValue struct {
	ValidTime string  `json:"validTime"`
	Value     float64 `json:"value"`
}

type WetBulbGlobeTemperature struct {
	UOM    string             `json:"uom"`
	Values []TemperatureValue `json:"values"`
}

type WBGTForecast struct {
	Properties struct {
		WetBulbGlobeTemperature WetBulbGlobeTemperature `json:"wetBulbGlobeTemperature"`
	} `json:"properties"`
}

type WBGTForecastLocation struct {
	Properties struct {
		RelativeLocation struct {
			Properties struct {
				City  string `json:"city"`
				State string `json:"state"`
			} `json:"properties"`
		} `json:"relativeLocation"`
	} `json:"properties"`
}

var alertList []Alert
var countyList = map[string]int{}
var NoAlertStatement []NoAlert
var userAlertTypes userAlertType
var userNWSOffices userNWSOffice
var wbgtDataLocation WBGTForecastLocation
var wbgtRawData WBGTForecast

func addCounties(countyListArr []string) {
	for _, county := range countyListArr {
		countyList[county] = 1
	}
}

var transformedCountyList = map[string]int{}

// Step 1: Refactor addStateIdToCountyList to avoid modifying the original countyList
func addStateIdToCountyList(stateId string) {
	transformedCountyList = make(map[string]int)
	for county := range countyList {
		countyState := county + " " + stateId
		transformedCountyList[countyState] = 1
		transformedCountyList[county] = 1
	}
}

// Step 2: Use transformedCountyList in inCountyListCheck
func inCountyListCheck(singleAlert *Alert) bool {
	areaDescList := strings.Split(strings.TrimSpace(singleAlert.AreaDesc), "; ")

	// Sort areaDescList to ensure consistent order
	sort.Strings(areaDescList)

	// Create filtered list
	filteredAreaDescList := []string{}

	// Normalize and sanitize both locations and countyList entries
	for _, location := range areaDescList {
		location = strings.TrimSpace(location)
		normalizedLocation := strings.ToLower(removeNonASCII(location))

		// Normalize the county names and compare
		for county := range transformedCountyList {
			normalizedCounty := strings.ToLower(removeNonASCII(county))

			// If they match, append the location
			if normalizedLocation == normalizedCounty {
				filteredAreaDescList = append(filteredAreaDescList, location)
				break
			}
		}
	}

	// Sort filteredAreaDescList to ensure consistent order of results
	sort.Strings(filteredAreaDescList)

	// If no matching locations found, return false
	if len(filteredAreaDescList) == 0 {
		return false
	}

	// Join filtered list and update AreaDesc
	singleAlert.AreaDesc = strings.Join(filteredAreaDescList, "; ")

	return true
}

// Utility function to remove non-ASCII characters for simpler comparison
func removeNonASCII(str string) string {
	var result []rune
	for _, r := range str {
		if r <= unicode.MaxASCII {
			result = append(result, r)
		}
	}
	return string(result)
}

func changeTimeOutputAndHeadline(singleAlert *Alert) {
	timeStringEffective := singleAlert.Effective
	timeStringExpires := singleAlert.Expires

	timeEffective, _ := time.Parse(time.RFC3339, timeStringEffective)
	timeExpires, _ := time.Parse(time.RFC3339, timeStringExpires)
	timeEffectiveTimeOutput := timeEffective.Format("3:04PM")
	timeExpiresTimeOutput := timeExpires.Format("3:04PM Jan 2, 2006")

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
		//fmt.Println(singleAlert.SenderName)
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

func getWBGTForecastLocation(lat, long string) {
	const BASE_URL = "https://api.weather.gov"
	response, err := http.Get(BASE_URL + "/points/" + lat + "," + long)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var responseURL wbgtURL
	json.Unmarshal(responseData, &responseURL)
	// Save Location
	json.Unmarshal(responseData, &wbgtDataLocation)
	// Go in to get the WBGT
	response, err = http.Get(responseURL.Properties.Url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	responseData, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Unmarshal(responseData, &wbgtRawData)
	wbgtRawData.Properties.WetBulbGlobeTemperature.Values = convertTemperatures(wbgtRawData.Properties.WetBulbGlobeTemperature.Values)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertTemperatures(temperatures []TemperatureValue) []TemperatureValue {
	var converted []TemperatureValue

	for _, temp := range temperatures {
		// Convert the temperature value to Fahrenheit
		fahrenheit := celsiusToFahrenheit(temp.Value)
		converted = append(converted, TemperatureValue{
			ValidTime: temp.ValidTime,
			Value:     fahrenheit,
		})
	}

	return converted
}

func getWBGTForecastData(c *gin.Context) {
	loc := strings.Split(c.Param("loc"), ",")
	lat := loc[0]
	long := loc[1]
	getWBGTForecastLocation(lat, long)
	c.IndentedJSON(http.StatusCreated, wbgtRawData.Properties.WetBulbGlobeTemperature.Values)
}

func getWBGTForecastCityState(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, wbgtDataLocation)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://nws-api-active-alerts.vercel.app", "https://www.ryanmarando.com"}, // http://localhost:3000 https://nws-api-active-alerts.vercel.app", "https://www.ryanmarando.com
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
	router.GET("/getWBGTForecastCityState", getWBGTForecastCityState)
	router.GET("/getWBGTForecastData/:loc", getWBGTForecastData)
	router.Run(":10000") //localhost:8080 :10000

}
