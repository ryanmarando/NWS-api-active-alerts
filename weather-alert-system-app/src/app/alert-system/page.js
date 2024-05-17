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
import { IoIosArrowBack } from "react-icons/io";
import { IoIosArrowDown } from "react-icons/io";
import { Popup } from "@/components/Popup";

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
  const items = [
    "Tsunami Warning",
    "Tornado Warning",
    "Extreme Wind Warning",
    "Severe Thunderstorm Warning",
    "Flash Flood Warning",
    "Flash Flood Statement",
    "Severe Weather Statement",
    "Shelter In Place Warning",
    "Evacuation Immediate",
    "Civil Danger Warning",
    "Nuclear Power Plant Warning",
    "Radiological Hazard Warning",
    "Hazardous Materials Warning",
    "Fire Warning",
    "Civil Emergency Message",
    "Law Enforcement Warning",
    "Storm Surge Warning",
    "Hurricane Force Wind Warning",
    "Hurricane Warning",
    "Typhoon Warning",
    "Special Marine Warning",
    "Blizzard Warning",
    "Snow Squall Warning",
    "Ice Storm Warning",
    "Winter Storm Warning",
    "High Wind Warning",
    "Tropical Storm Warning",
    "Storm Warning",
    "Tsunami Advisory",
    "Tsunami Watch",
    "Avalanche Warning",
    "Earthquake Warning",
    "Volcano Warning",
    "Ashfall Warning",
    "Coastal Flood Warning",
    "Lakeshore Flood Warning",
    "Flood Warning",
    "High Surf Warning",
    "Dust Storm Warning",
    "Blowing Dust Warning",
    "Lake Effect Snow Warning",
    "Excessive Heat Warning",
    "Tornado Watch",
    "Severe Thunderstorm Watch",
    "Flash Flood Watch",
    "Gale Warning",
    "Flood Statement",
    "Wind Chill Warning",
    "Extreme Cold Warning",
    "Hard Freeze Warning",
    "Freeze Warning",
    "Red Flag Warning",
    "Storm Surge Watch",
    "Hurricane Watch",
    "Hurricane Force Wind Watch",
    "Typhoon Watch",
    "Tropical Storm Watch",
    "Storm Watch",
    "Hurricane Local Statement",
    "Typhoon Local Statement",
    "Tropical Storm Local Statement",
    "Tropical Depression Local Statement",
    "Avalanche Advisory",
    "Winter Weather Advisory",
    "Wind Chill Advisory",
    "Heat Advisory",
    "Urban and Small Stream Flood Advisory",
    "Small Stream Flood Advisory",
    "Arroyo and Small Stream Flood Advisory",
    "Flood Advisory",
    "Hydrologic Advisory",
    "Lakeshore Flood Advisory",
    "Coastal Flood Advisory",
    "High Surf Advisory",
    "Heavy Freezing Spray Warning",
    "Dense Fog Advisory",
    "Dense Smoke Advisory",
    "Small Craft Advisory",
    "Brisk Wind Advisory",
    "Hazardous Seas Warning",
    "Dust Advisory",
    "Blowing Dust Advisory",
    "Lake Wind Advisory",
    "Wind Advisory",
    "Frost Advisory",
    "Ashfall Advisory",
    "Freezing Fog Advisory",
    "Freezing Spray Advisory",
    "Low Water Advisory",
    "Local Area Emergency",
    "Avalanche Watch",
    "Blizzard Watch",
    "Rip Current Statement",
    "Beach Hazards Statement",
    "Gale Watch",
    "Winter Storm Watch",
    "Hazardous Seas Watch",
    "Heavy Freezing Spray Watch",
    "Coastal Flood Watch",
    "Lakeshore Flood Watch",
    "Flood Watch",
    "High Wind Watch",
    "Excessive Heat Watch",
    "Extreme Cold Watch",
    "Wind Chill Watch",
    "Lake Effect Snow Watch",
    "Hard Freeze Watch",
    "Freeze Watch",
    "Fire Weather Watch",
    "Extreme Fire Danger",
    " Telephone Outage",
    "Coastal Flood Statement",
    "Lakeshore Flood Statement",
    "Special Weather Statement",
    "Marine Weather Statement",
    "Air Quality Alert",
    "Air Stagnation Advisory",
    "Hazardous Weather Outlook",
    "Hydrologic Outlook",
    "Short Term Forecast",
    "Administrative Message",
    "Test",
    "Child Abduction Emergency",
    "Blue Alert",
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
  const [stateListSaved, setStateListSaved] = useState("");
  const [inputValue, setInputValue] = useState("");
  const { isLoaded, isSignedIn, user } = useUser();
  const [showWarningSettings, setShowWarningSettings] = useState(false);
  const [checkedItems, setCheckedItems] = useState([]);
  const [selectAll, setSelectAll] = useState(false);
  const [showPopup, setShowPopup] = useState(false);

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
    const data = { data: checkedItems };
    const response = await fetch("http://localhost:8080/userAlertTypes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (response.ok) {
      console.log("Warnings submitted!");
    }
    const results = await fetch(
      "http://localhost:8080/alerts/" + // http://localhost:8080 https://nws-api-active-alerts.onrender.com
        stateListString +
        "/" +
        countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
        setIsLoading(false);
      });

    console.log("Successful county export");
  }

  async function getDataFromOwnAPI() {
    if (stateList.length === 0) return alert("Please choose a state.");
    setIsLoading(true);
    if (countyList.length > 0) return getDataFromOwnAPIWithCounties();
    const stateListString = stateList.join(",");
    const data = { data: checkedItems };
    const response = await fetch("http://localhost:8080/userAlertTypes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (response.ok) {
      console.log("Warnings submitted!");
    }
    const results = await fetch(
      "http://localhost:8080/alerts/" + stateListString + countyList
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
        setIsLoading(false);
      });
    console.log("Successful state wide export");
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
      if (!user) {
        return setShowPopup(true);
      }
      setIsRunning((prev) => !prev);
    };

    useEffect(() => {
      let intervalId;
      if (isRunning) {
        intervalId = setInterval(myFunction, 15000);
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

  const AlertForm = () => {
    const [searchTerm, setSearchTerm] = useState("");
    const handleCheckboxChange = (event) => {
      const { value, checked } = event.target;
      setCheckedItems((prevCheckedItems) =>
        checked
          ? [...prevCheckedItems, value]
          : prevCheckedItems.filter((item) => item !== value)
      );
    };

    const handleSelectAllChange = (event) => {
      const { checked } = event.target;
      setSelectAll(checked);
      setCheckedItems(checked ? filteredItems : []);
    };

    const handleSearchChange = (event) => {
      setSearchTerm(event.target.value);
    };

    const filteredItems = items.filter((item) =>
      item.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
      <div className="p-4 w-full">
        <p className="w-full flex items-center justify-center">
          Optional: skip for all alerts
        </p>
        <div className="w-full">
          <input
            type="text"
            placeholder="Search..."
            value={searchTerm}
            onChange={handleSearchChange}
            className="path-bar mb-6 p-1 w-full"
          />
        </div>
        <form className="p-2">
          <div className="w-full flex items-center justify-center mb-2 row-span-2">
            <input
              type="checkbox"
              id="select-all"
              checked={selectAll}
              onChange={handleSelectAllChange}
              className="mr-2"
            />
            <label htmlFor="select-all">Select All</label>
          </div>
          <div className=" grid grid-cols-2 lg:grid-cols-3 gap-2 gap-x-6 -ml-4 lg:ml-2">
            {filteredItems.map((item, index) => (
              <div key={index} className="flex items-center pl-4">
                <input
                  type="checkbox"
                  id={`checkbox-${index}`}
                  value={item}
                  checked={checkedItems.includes(item)}
                  onChange={handleCheckboxChange}
                  className="p-2 mr-2 text-center"
                />
                <label htmlFor={`checkbox-${index}`}>{item}</label>
              </div>
            ))}
          </div>
        </form>
      </div>
    );
  };

  const saveDataListInput = async () => {
    if (!isSignedIn) return alert("Please login to save your settings.");
    if (stateList.length === 0)
      return alert("You must add at least once state to save data.");
    try {
      const response = await fetch("/api/updateUserMetadata", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          userId: user?.id,
          metadata: {
            countyList: countyList,
            stateList: stateList.join(","),
            warningTypes: checkedItems,
            subscription: "false",
          },
        }),
      });

      if (response.ok) {
        alert("Settings updated successfully! Refresh to update");
      } else {
        const errorData = await response.json();
        alert(`Failed to update metadata: ${errorData.error}`);
      }
    } catch (error) {
      console.error("Error updating metadata:", error);
      alert("Failed to update metadata. Please try again.");
    }
  };

  async function populateDataInput() {
    if (!isSignedIn) return alert("Please login to load your saved data.");
    const savedStateListArr = user?.publicMetadata.stateList.split(",");
    setStateList(savedStateListArr);
    setCountyList(user?.publicMetadata.countyList);
    const savedWarningTypeArr = user?.publicMetadata.warningTypes;
    setCheckedItems(savedWarningTypeArr);
  }

  const closePopup = () => {
    setShowPopup(false);
  };

  return (
    <div>
      <Navbar />
      <Popup show={showPopup} onClose={closePopup} />
      {showSettings && (
        <div className="alert-system">
          <div className="flex relative w-full items-center justify-center">
            <button
              onClick={populateDataInput}
              className="bg-[#4328EB] lg:absolute lg:top-0 lg:right-0 hover:text-gray-500 py-1 px-2 w-50 mt-2 mr-1 lg:mr-0 lg:mt-0 lg:w-50 rounded-[8px] text-white "
            >
              Populate Saved Settings
            </button>
          </div>
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
          <input
            className="path-bar"
            type="text"
            value={countyList}
            placeholder="County,County,County"
            onChange={(e) => setCountyList(e.target.value)}
          />
          <div className="flex gap-1">
            <label className="label">Choose your specific warnings:</label>
            <button
              onClick={() => {
                if (!user) return setShowPopup(true);
                setShowWarningSettings(!showWarningSettings);
              }}
            >
              {showWarningSettings ? (
                <IoIosArrowDown className="mt-[13px] w-8 hover:text-gray-500" />
              ) : (
                <IoIosArrowBack className="mt-[13px] w-8 hover:text-gray-500" />
              )}
            </button>
          </div>
          {showWarningSettings ? <AlertForm /> : ""}
          <button
            className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[15px]"
            type="alert-button"
            onClick={getDataFromOwnAPI}
            disabled={isLoading}
          >
            {isLoading ? <LoadingAnimation /> : "Get Alerts"}
          </button>
          <div className="flex relative w-full items-center justify-center">
            <button
              onClick={saveDataListInput}
              className="bg-[#4328EB] lg:absolute lg:bottom-0 lg:right-0 hover:text-gray-500 py-1 px-2 w-50 lg:mr-0 lg:mt-0 lg:w-50 rounded-[8px] text-white "
            >
              Save Settings
            </button>
          </div>
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
        <ul className="p-2">
          {alertList.map((obj, idx) => (
            <li
              className="alert-list shadow-md border-spacing-1"
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
