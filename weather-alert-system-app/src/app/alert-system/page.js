"use client";
import { useState, useEffect } from "react";
import React from "react";
import Select from "react-select";
import Image from "next/image";
import Clock from "@/components/Clock";
import LoadingAnimation from "@/components/LoadingAnimation";
import CountdownTimer from "@/components/Counter";
import { Navbar } from "@/components/Navbar";
import BlueArrow from "../../../public/assets/blue-button.svg";
import { Footer } from "@/components/Footer";
import { useUser } from "@clerk/clerk-react";
import { clerkClient } from "@clerk/nextjs/server";

export default function AlertSystem() {
  const options = [
    { value: "AL", label: "AL" },
    { value: "AK", label: "AK" },
    { value: "AZ", label: "AZ" },
    { value: "AR", label: "AR" },
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
  const [isLoading, setIsLoading] = useState(false);
  const [showSettings, setShowSettings] = useState(true);
  const [countyListSaved, setCountyListSaved] = useState("");
  const [inputValue, setInputValue] = useState("");
  const { isLoaded, user } = useUser();

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
    setStateList([...stateList]);
  }

  async function getDataFromOwnAPIWithCounties() {
    const stateListString = stateList.join(",");
    const results = await fetch(
      "https://nws-api-active-alerts.onrender.com/alerts/" + // http://localhost:8080 https://nws-api-active-alerts.onrender.com
        stateListString +
        "/" +
        countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
        setIsLoading(false);
      });

    console.log("successful county export");
  }

  async function getDataFromOwnAPI() {
    if (stateList.length === 0) return alert("Please choose a state.");
    setIsLoading(true);
    if (countyList.length > 0) return getDataFromOwnAPIWithCounties();
    const stateListString = stateList.join(",");
    const results = await fetch(
      "https://nws-api-active-alerts.onrender.com/alerts/" +
        stateListString +
        countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
        setIsLoading(false);
      });
    console.log("successful state wide export");
  }

  function ChangeSettings() {
    return setShowSettings(!showSettings);
  }

  const AutomaticOutput = () => {
    const myFunction = () => {
      console.log("Function called!");
      getDataFromOwnAPI();
    };

    const toggleFunction = () => {
      setIsRunning((prev) => !prev);
    };

    useEffect(() => {
      let intervalId;
      if (isRunning) {
        intervalId = setInterval(myFunction, 15000); // Call every 30 seconds
      } else {
        clearInterval(intervalId);
      }
      return () => clearInterval(intervalId);
    }, [isRunning]);
    return (
      <div>
        <button
          className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[25px]"
          onClick={toggleFunction}
        >
          {isRunning ? (
            <CountdownTimer initialTime={15} />
          ) : (
            "Start Automatic Output"
          )}
        </button>
      </div>
    );
  };

  const StateListComponent = () => {
    let frontendStateList = "";
    if (stateList.length === 1) {
      frontendStateList = stateList[0];
      return (
        <div className="flex items-center justify-center gap-x-2">
          <span>
            <Image src={BlueArrow} alt="blue arrow" />
          </span>
          <p className="text-xl my-[10px]">States Added: {frontendStateList}</p>
        </div>
      );
    }
    for (let i = 0; i < stateList.length; i++) {
      //const state of stateList
      frontendStateList += ", " + stateList[i];
    }
    frontendStateList = frontendStateList.slice(2);
    return (
      <div className="flex items-center justify-center gap-x-2">
        <span>
          <Image src={BlueArrow} alt="blue arrow" />
        </span>
        <p className="text-xl my-[10px]">States Added: {frontendStateList}</p>
      </div>
    );
  };
  /*
  const saveCountyListInput = () => {
    setCountyListSaved(countyList);
    alert("Saved counties to your profile!");
  };
  */

  const saveCountyListInput = async () => {
    console.log(user?.publicMetadata.countyList);
    console.log(user?.id);

    setCountyListSaved(countyList);
    console.log(countyListSaved);
    try {
      await clerkClient.users.updateUserMetadata(user?.id, {
        publicMetadata: {
          countyList: countyListSaved,
        },
      });
      alert("Input value saved successfully!");
    } catch (error) {
      console.error("Error saving input value:", error);
      alert("Failed to save input value. Please try again.");
    }
  };

  const populateCountyListInput = () => {
    setCountyList(user?.publicMetadata.countyList);
  };

  return (
    <div>
      <Navbar />
      {showSettings && (
        <div className="alert-system">
          <label className="label">
            Please select state(s) to receive alerts:
          </label>
          <Select
            defaultValue={selectedOption}
            onChange={setSelectedOption}
            options={stateOptions}
          />
          <div className="flex w-full items-center justify-center">
            <button
              onClick={addState}
              className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-white my-[5px] mx-[15px] mt-[14px]"
            >
              Add State
            </button>
            <button
              onClick={removeState}
              className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[15px] mt-[14px]"
            >
              Remove Last State
            </button>
          </div>
          <StateListComponent />
          <label className="label">Want to enter counties?</label>
          <div className="flex w-full items-center justify-center ml-6 ">
            <input
              className="path-bar"
              type="text"
              value={countyList}
              placeholder="County,County,County"
              onChange={(e) => setCountyList(e.target.value)}
            />
            <div className="flex items-center justify-center flex-col ml-6 mr-6">
              <button
                onClick={saveCountyListInput}
                className="bg-[#4328EB] hover:text-gray-500 py-1 px-2 rounded-[8px] text-white mb-2"
              >
                Save Counties
              </button>
              <button
                onClick={populateCountyListInput}
                className="bg-[#4328EB] hover:text-gray-500 py-1 px-2 rounded-[8px] text-white mb-2 "
              >
                Populate Counties
              </button>
            </div>
          </div>
          <br></br>
          <button
            className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[15px]"
            type="alert-button"
            onClick={getDataFromOwnAPI}
            disabled={isLoading}
          >
            {isLoading ? <LoadingAnimation /> : "Get Alerts"}
          </button>
        </div>
      )}
      <div className="flex w-full items-center justify-between px-[20px] py-[8px]">
        <button
          className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[25px]"
          onClick={ChangeSettings}
        >
          {showSettings ? "Hide Settings" : "Show Settings"}
        </button>
        <AutomaticOutput />
      </div>
      <div className="clock">
        <Clock />
      </div>
      <div className="alert-output">
        <ul>
          {alertList.map((obj, idx) => (
            <li
              className="alert-list"
              style={{ backgroundColor: obj.color }}
              key={obj.id}
            >
              {/* Access object properties and render them */}
              <span className="effective">{obj.effective}</span>{" "}
              <span className="headline">{obj.headline}</span>
              <br></br>
              <div className="areaDesc">{obj.areaDesc}</div>
            </li>
          ))}
        </ul>
      </div>
      <Footer />
    </div>
  );
}
