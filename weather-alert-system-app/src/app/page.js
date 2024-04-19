"use client";
import { useState, useEffect, useRef } from "react";
import React from "react";
import Select from "react-select";
import { saveAs } from "file-saver";
import {
  getAlertsList,
  printAllAlerts,
  getActiveAlerts,
} from "@/components/api_read";

export default function Home() {
  const options = [
    { value: "AL", label: "AL" },
    { value: "AK", label: "AK" },
    { value: "AZ", label: "AZ" },
    { value: "AR", label: "AR" },
    { value: "AS", label: "AS" },
    { value: "CA", label: "CA" },
    { value: "CO", label: "CO" },
    { value: "CT", label: "CT" },
    { value: "DE", label: "DE" },
    { value: "DC", label: "DC" },
    { value: "FL", label: "FL" },
    { value: "GA", label: "GA" },
    { value: "HI", label: "HI" },
    { value: "ID", label: "ID" },
    { value: "IL", label: "IL" },
    { value: "IN", label: "IN" },
    { value: "IA", label: "IA" },
    { value: "KS", label: "KS" },
    { value: "LA", label: "LA" },
    { value: "ME", label: "ME" },
    { value: "MD", label: "MD" },
    { value: "MA", label: "MA" },
    { value: "MI", label: "MI" },
    { value: "MN", label: "MN" },
    { value: "MS", label: "MS" },
    { value: "MO", label: "MO" },
    { value: "MT", label: "MT" },
    { value: "NE", label: "NE" },
    { value: "NV", label: "NV" },
    { value: "NH", label: "NH" },
    { value: "NJ", label: "NJ" },
    { value: "NM", label: "NM" },
    { value: "NY", label: "NY" },
    { value: "NC", label: "NC" },
    { value: "ND", label: "ND" },
    { value: "OH", label: "OH" },
    { value: "OK", label: "OK" },
    { value: "OR", label: "OR" },
    { value: "PA", label: "PA" },
    { value: "RI", label: "RI" },
    { value: "SC", label: "SC" },
    { value: "SD", label: "SD" },
    { value: "TN", label: "TN" },
    { value: "TX", label: "TX" },
    { value: "UT", label: "UT" },
    { value: "VT", label: "VT" },
    { value: "VA", label: "VA" },
    { value: "WA", label: "WA" },
    { value: "WV", label: "WV" },
    { value: "WI", label: "WI" },
    { value: "WY", label: "WY" },
  ];
  const [stateOptions, setStateOptions] = useState(options);
  const [selectedOption, setSelectedOption] = useState(null);
  const [stateList, setStateList] = useState([]);
  const [countyList, setCountyList] = useState("");
  const [alertList, setAlertList] = useState([]);
  const [isRunning, setIsRunning] = useState(false);

  function addState() {
    if (!selectedOption) return alert("Please enter a state.");
    const state = selectedOption.value;
    if (stateList.includes(state)) {
      const alertString =
        "You've already entered " + state + ". Choose another state.";
      alert(alertString);
      return;
    }
    setStateList([...stateList, state]);
  }

  function removeState() {
    if (stateList.length === 0)
      return alert("You must add at least once state to remove.");
    if (stateList.length <= 1) {
      setStateList([]);
      return;
    }
    stateList.pop();
    setStateList([stateList]);
  }
  /*
  useEffect(() => {
    fetch("http://localhost:8080/alerts/GA")
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
      });
  }, [stateList]);
*/
  async function getDataFromOwnAPIWithCounties() {
    //if (stateList.length === 0) return alert("Please choose a state.");
    const stateListString = stateList.join(",");
    const results = await fetch(
      "http://localhost:8080/alerts/" + stateListString + "/" + countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
      });
    console.log("successful county export");
  }

  async function getDataFromOwnAPI() {
    if (stateList.length === 0) return alert("Please choose a state.");
    if (countyList.length > 0) return getDataFromOwnAPIWithCounties();
    const stateListString = stateList.join(",");
    const results = await fetch(
      "http://localhost:8080/alerts/" + stateListString + countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
      });
    console.log("successful state wide export");
  }

  const MyComponent = () => {
    // Function to be called
    const myFunction = () => {
      console.log("Function called!");
      getDataFromOwnAPI();
    };

    // Toggle function execution
    const toggleFunction = () => {
      setIsRunning((prev) => !prev);
    };

    // Run the function when isRunning is true
    useEffect(() => {
      let intervalId;
      if (isRunning) {
        intervalId = setInterval(myFunction, 30000); // Example: call every 30 seconds
      } else {
        clearInterval(intervalId);
      }
      return () => clearInterval(intervalId);
    }, [isRunning]);

    return (
      <div>
        <button onClick={toggleFunction}>
          {isRunning ? "Stop Automatic Output" : "Start Automatic Output"}
        </button>
      </div>
    );
  };

  return (
    <div>
      <h1>Welcome To The Weather Alert System</h1>
      <h4>Please select state(s) to receive alerts:</h4>
      <Select
        defaultValue={selectedOption}
        onChange={setSelectedOption}
        options={stateOptions}
      />
      <span>
        <button type="button" onClick={addState}>
          Add State
        </button>
      </span>
      <span>
        <button type="button" onClick={removeState}>
          /Remove Last State
        </button>
      </span>
      <div>
        <p>States Added: {stateList}</p>
      </div>
      <label>Want to enter counties?</label>
      <input
        type="text"
        value={countyList}
        placeholder="County,County,County"
        onChange={(e) => setCountyList(e.target.value)}
      />
      <br></br>
      <button type="alert-button" onClick={getDataFromOwnAPI}>
        Get Alerts!
      </button>
      <MyComponent />
    </div>
  );
}
