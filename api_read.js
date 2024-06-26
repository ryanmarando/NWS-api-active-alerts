const { convertArrayToCSV } = require("convert-array-to-csv");
const fs = require("fs");

const baseURL = "https://api.weather.gov";
let alertList = new Array();
const eventTypePriorityLibrary = new Map([
  ["Tornado Warning", 1],
  ["Severe Thunderstorm Warning", 2],
  ["Flash Flood Warning", 3],
  ["Flood Warning", 4],
  ["Tornado Watch", 5],
  ["Severe Thunderstorm Watch", 6],
  ["Flood Watch", 7],
  ["Wind Advisory", 8],
  ["Coastal Flood Advisory", 9],
  ["Frost Advisory", 10],
]);
var stateLibrary = new Set([
  "AL",
  "AK",
  "AZ",
  "AR",
  "AS",
  "CA",
  "CO",
  "CT",
  "DE",
  "DC",
  "FL",
  "GA",
  "HI",
  "ID",
  "IL",
  "IN",
  "IA",
  "KS",
  "LA",
  "ME",
  "MD",
  "MA",
  "MI",
  "MN",
  "MS",
  "MO",
  "MT",
  "NE",
  "NV",
  "NH",
  "NJ",
  "NM",
  "NY",
  "NC",
  "ND",
  "OH",
  "OK",
  "OR",
  "PA",
  "RI",
  "SC",
  "SD",
  "TN",
  "TX",
  "UT",
  "VT",
  "VA",
  "WA",
  "WV",
  "WI",
  "WY",
  "Q",
]);
let countyList = new Set([]);

function addStateIdToCountyList(stateId) {
  const originalCountyListLength = countyList.size;
  const countySetToArray = [...countyList];
  for (i = 0; i < countyList.size - 1; i++) {
    countyList.add(countySetToArray[i] + " " + stateId);
  }
}

function addCounties(countyListArr) {
  for (county of countyListArr) {
    countyList.add(county);
  }
}

function changeTimeOutput(singleAlertTime) {
  const dateTime = new Date(singleAlertTime);
  dayPart = "AM";
  newDateHour = dateTime.getHours();
  if (newDateHour > 12) {
    newDateHour = newDateHour - 12;
    dayPart = "PM";
  } else if (newDateHour == 0) {
    newDateHour = 8;
    dayPart = "PM";
    //return "further notice";
  }
  newDateMinutes = dateTime.getMinutes().toString();
  if (newDateMinutes === 0 || newDateMinutes.length === 1)
    newDateMinutes = newDateMinutes + "0";
  newDateTime = newDateHour + ":" + newDateMinutes + dayPart;
  return newDateTime;
}

function changeHeadlineOutput(singleAlert) {
  singleAlert.headline = singleAlert.event + " until " + singleAlert.expires;
}

function createAlert(areaDesc, event, effective, expires, headline, priority) {
  const Alert = {
    areaDesc,
    event,
    effective,
    expires,
    headline,
    priority,
  };
  return Alert;
}

function setWarningPriority(warningEvent) {
  if (eventTypePriorityLibrary.get(warningEvent) === undefined) return 1000;
  return eventTypePriorityLibrary.get(warningEvent);
}

function sortAlertsByPriority(alertList) {
  alertList.sort((a, b) => a.priority - b.priority);
}

function printAllAlerts() {
  console.log("\n-------------------------------------------------------");
  for (const [idx, alert] of alertList.entries()) {
    console.log(
      "Warning",
      idx + 1,
      ": ",
      alert.headline,
      "| priority:",
      alert.priority,
      "\nIssued for:",
      alert.areaDesc
    );
  }
  console.log("-------------------------------------------------------");
}

function removeCommas(singleAlert) {
  const locations = singleAlert.areaDesc;
  let locationsNoCommaAddSemicolon = locations.replaceAll(",", "");
  singleAlert.areaDesc = locationsNoCommaAddSemicolon;
}

function readInData(stateId) {
  url = "http://api.weather.gov/alerts/active?area=" + stateId;
  const results = fetch(url)
    .then((data) => data.json())
    .then((data) => data);
  return results;
}

function inCountyListCheck(singleAlert) {
  const areaDescList = singleAlert.areaDesc.split("; ");
  let filteredAreaDescList = [];
  for (location of areaDescList) {
    if (countyList.has(location)) {
      filteredAreaDescList.push(location);
    }
  }
  if (filteredAreaDescList.length == 0) {
    return false;
  }
  singleAlert.areaDesc = filteredAreaDescList.join("; ");
  return true;
}

async function getActiveAlerts(stateId) {
  results = await readInData(stateId);
  const countiesAdded = countyList.size > 0;
  for (const [idx, alert] of results["features"].entries()) {
    warningProperties = alert["properties"];
    singleAlert = createAlert(
      warningProperties["areaDesc"],
      warningProperties["event"],
      warningProperties["effective"],
      warningProperties["expires"],
      warningProperties["headline"],
      setWarningPriority(warningProperties["event"])
    );
    singleAlert.expires = changeTimeOutput(singleAlert.expires);
    singleAlert.effective = changeTimeOutput(singleAlert.effective);
    changeHeadlineOutput(singleAlert);
    removeCommas(singleAlert);
    if (countiesAdded && inCountyListCheck(singleAlert)) {
      alertList.push(singleAlert);
    } else if (!countiesAdded) {
      alertList.push(singleAlert);
    }
  }
  sortAlertsByPriority(alertList);
}

async function getAlertsList(stateList, countyList) {
  addCounties(countyList);
  for (state of stateList) {
    addStateIdToCountyList(state);
    await getActiveAlerts(state);
  }
  printAllAlerts();
  newOutputToCSV();
  //outputToCSV();
}

function newOutputToCSV() {
  const { convertArrayToCSV } = require("convert-array-to-csv");
  const fs = require("fs");
  const headers = Object.keys(alertList[0]);
  const csvFromArrayOfArrays = convertArrayToCSV(alertList, {
    headers,
    separator: ",",
  });
  fs.writeFile(
    "C:/Users/Ryan Marando/program_files/course_careers/final-project/outputTEST.csv",
    csvFromArrayOfArrays,
    (err) => {
      if (err) {
        console.log(18, err);
      }
      console.log("CSV file saved successfully");
    }
  );
}

function outputToCSV() {
  const titleKeys = Object.keys(alertList[0]);
  const refinedData = [];
  refinedData.push(titleKeys);

  alertList.forEach((item) => {
    refinedData.push(Object.values(item));
  });

  let csvContent = "";
  refinedData.forEach((row) => {
    csvContent += row.join(",") + "\n";
  });

  const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8," });
  const objUrl = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.setAttribute("href", objUrl);
  link.setAttribute("download", "alert-list.csv");
  link.click();

  document.querySelector("body").append(link);
}

getAlertsList(["FL"], []);
/*
let stateList = [];
while (true) {
  let stateAbbr = prompt(
    "Enter the state abbreviation which you wish to receive active warnings or q to get results: "
  ).toUpperCase();

  const notCorrectInput = stateAbbr.length > 2 || stateAbbr === null;
  const notAState = !stateLibrary.has(stateAbbr);
  const noStateChosen = stateAbbr === "Q" && stateList.length === 0;
  const stateEntered = stateAbbr !== "Q";
  if (notCorrectInput || notAState) {
    alert("Please enter two letter abbreviation");
    continue;
  }
  if (noStateChosen) {
    alert("Please choose at least one state");
    continue;
  }
  if (stateEntered) {
    alert(stateAbbr + " added!");
    stateList.push(stateAbbr);
    console.log(stateList);
    continue;
  }
  getAlertsList(stateList);
  break;
}
*/
