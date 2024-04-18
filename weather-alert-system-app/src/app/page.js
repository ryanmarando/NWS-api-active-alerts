"use client";
import { useState, useEffect } from "react";
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
  const [countyList, setCountyList] = useState([]);

  function addState() {
    const state = selectedOption.value;
    if (stateList.includes(state)) {
      const alertString =
        "You've already entered " + state + ". Choose another state.";
      alert(alertString);
      return;
    }
    setStateList([...stateList, state]);
    setCountyList(["Brooks", "Brantley"]);
    setSelectedOption(null);
  }

  async function getAlerts() {
    let alertList = await getAlertsList(stateList, countyList);
    console.log(alertList);
    setStateList([]);
  }

  function alerting() {
    alert("getting long");
  }

  useEffect(() => {
    //if (stateList.length > 2) alerting();
  }, [stateList]);

  return (
    <div>
      <h1>Welcome To The Weather Alert System</h1>
      <h4>Please select state(s) to receive alerts:</h4>
      <Select
        defaultValue={selectedOption}
        onChange={setSelectedOption}
        options={stateOptions}
      />
      <button type="button" onClick={addState}>
        Add State
      </button>
      <div>
        <p>{stateList}</p>
      </div>
      <button onClick={getAlerts}>Get Alerts!</button>
    </div>
  );
}
