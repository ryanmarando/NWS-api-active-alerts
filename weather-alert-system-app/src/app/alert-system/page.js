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
    { value: "KY", label: "KY" },
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
  const NWSofficeFullList = [
    "NWS Birmingham AL",
    "NWS Mobile AL",
    "NWS Huntsville AL",
    "NWS Anchorange AK",
    "NWS Fairbanks AK",
    "NWS Juneau AK",
    "NWS Flagstaff AZ",
    "NWS Phoenix AZ",
    "NWS Tucson AZ",
    "NWS Little Rock AR",
    "NWS Key West FL",
    "NWS Jacksonville FL",
    "NWS Melbourne FL",
    "NWS Miami FL",
    "NWS Talahassee FL",
    "NWS Tampa Bay Ruskin FL",
    "NWS Peachtree City GA",
    "NWS Chicago IL",
    "NWS Northern Indiana",
    "NWS Indianapolis IN",
    "NWS Caribou ME",
    "NWS Burlington VT",
    "NWS Wilmington NC",
    "NWS Mount Holly NJ",
    "NWS Albany NY",
    "NWS Upton NY",
    "NWS Duluth MN",
    "NWS Wilmington OH",
    "NWS Medford OR",
    "NWS Charleston SC",
  ]
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
  const [NWSselectedOfficeList, setNWSselectedOfficeList] = useState([])
  const [showNWSOfficeFullList, setShowNWSOfficeFullList] = useState(false);
  const [checkedNWSOffices, setCheckedNWSOffices] = useState([]);
  const [countiesByState, setCountiesByState] = useState({});
  const [selectedCounties, setSelectedCounties] = useState({});
  const [showCountiesForSelectedStates, setShowCountiesForSelectedStates] = useState(false);
  const [hasSearchedForAlerts, setHasSearchedForAlerts] = useState(false);
  const URL = "http://localhost:8080" //https://nws-api-active-alerts.onrender.com

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

  const CountyListLibrary = () => {

  // Fetch counties for a state when it's added
  const getCountiesFromState = async (state) => {
    const results = await fetch(URL + `/countiesByState/` + state)
      .then((data) => data.json())
      .then((data) => {
        setCountiesByState((prevState) => ({
          ...prevState,
          [state]: data, // Store counties for the selected state
        }));
      });
  };

 // Handle county checkbox change
 const handleCheckboxChange = (event, state, county) => {
  const isChecked = event.target.checked;

  setSelectedCounties((prevState) => {
    const updatedCounties = isChecked
      ? [...(prevState[state] || []), county]
      : prevState[state].filter((item) => item !== county);

    return {
      ...prevState,
      [state]: updatedCounties,
    };
  });
};

  // Fetch counties when a new state is added to the list
  useEffect(() => {
    stateList.forEach((state) => {
      if (!countiesByState[state]) {
        getCountiesFromState(state); // Fetch counties only if not already fetched
      }
    });
  }, [stateList]);

  return (
    <div className="p-4">
      <p className="w-full flex items-center justify-center">
          Optional: skip for all statewide alerts
        </p>
      {/* Display selected states and allow selecting counties for each */}
      {/* Conditional rendering based on stateList */}
      {stateList.length === 0 ? (
        <p>Please enter a state.</p>
      ) : (
        stateList.map((state, index) => (
          <div key={index} className="mb-4">
            <label>Select Counties for {state}:</label>
            <div className="grid grid-cols-6 md:grid-cols-6 lg:grid-cols-6 gap-1">
              {countiesByState[state] ? (
                countiesByState[state].map((county, countyIndex) => (
                  <div key={countyIndex} className="flex items-center">
                    <input
                      type="checkbox"
                      id={`checkbox-${state}-${countyIndex}`}
                      value={county}
                      checked={
                        selectedCounties[state]?.includes(county) || false
                      }
                      onChange={(event) =>
                        handleCheckboxChange(event, state, county)
                      }
                    />
                    <label
                      htmlFor={`checkbox-${state}-${countyIndex}`}
                      className="ml-2"
                    >
                      {county}
                    </label>
                  </div>
                ))
              ) : (
                <p>Loading counties...</p>
              )}
            </div>
          </div>
        ))
      )}
    </div>
  );
  };

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
    const tempAlertList = [];
    const data = { data: checkedItems };
    const response = await fetch(
      URL + "/userAlertTypes",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      }
    );
    if (response.ok) {
      console.log("Warnings submitted county!");
    }
    const NWSOfficeList = { officeList: checkedNWSOffices}
    const NWSOfficeResponse = await fetch(
      URL + "/getNWSOffices",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(NWSOfficeList),
      }
    );
    if (NWSOfficeResponse.ok) {
      console.log("Offices submitted county!");
    }
    for (let i = 0; i < Object.entries(selectedCounties).length; i++) {
      if (Object.values(selectedCounties)[i].length > 0) {
      const results = await fetch(
        URL + "/alerts/" +
         Object.keys(selectedCounties)[i] +
          "/" +
          Object.values(selectedCounties)[i]
      )
        .then((data) => data.json())
        .then((data) => {
          tempAlertList.push(...data)
        });
    } else {
      const results = await fetch(
        URL + "/alerts/" +
         Object.keys(selectedCounties)[i]
      )
        .then((data) => data.json())
        .then((data) => {
          tempAlertList.push(...data)
        });
  }
  }
    setAlertList(tempAlertList);
    setIsLoading(false);
    setHasSearchedForAlerts(true)
    console.log("Successful county export");
  }

  async function getDataFromOwnAPI() {
    if (stateList.length === 0) return alert("Please choose a state.");
    setIsLoading(true);
    if (Object.entries(selectedCounties).length > 0) {
      if (Object.values(selectedCounties)[0].length > 0) return getDataFromOwnAPIWithCounties();}
    const stateListString = stateList.join(",");
    const data = { data: checkedItems };
    const response = await fetch(
      URL + "/userAlertTypes",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      }
    );
    if (response.ok) {
      console.log("Warnings submitted!");
    }
    const NWSOfficeList = { officeList: checkedNWSOffices}
    const NWSOfficeResponse = await fetch(
      URL + "/getNWSOffices",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(NWSOfficeList),
      }
    );
    if (NWSOfficeResponse.ok) {
      console.log("Offices submitted!");
    }
    const results = await fetch(
      URL + "/alerts/" +
        stateListString
    )
      .then((data) => data.json())
      .then((data) => {
        setAlertList(data);
        setIsLoading(false);
      });
    setHasSearchedForAlerts(true)
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

    const handleSubmit = async (event) => {
      event.preventDefault();
      setSubmittedItems(checkedItems);

      const data = { data: checkedItems };
      console.log(data);

      const response = await fetch(
        URL + "/userAlertTypes",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );

      if (response.ok) {
        console.log("Warnings submitted!");
      }
      setButtonText("Submitted!");
      setTimeout(() => {
        setButtonText("Submit Alert Types");
      }, 3000);
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

  const NWSOfficeListForm = () => {
    const [searchOfficeTerm, setSearchOfficeTerm] = useState("");

  // Handle checkbox change
  const handleCheckboxChange = (event) => {
    const { value, checked } = event.target;

    setCheckedNWSOffices((prevCheckedItems) =>
      checked
        ? [...prevCheckedItems, value] // Add to checked list
        : prevCheckedItems.filter((item) => item !== value) // Remove from checked list
    );
  };

  // Handle search term change
  const handleSearchChange = (event) => {
    setSearchOfficeTerm(event.target.value);
  };

  // Filter the office list based on search term
  const filteredOfficeItems = NWSofficeFullList.filter((item) =>
    item.toLowerCase().includes(searchOfficeTerm.toLowerCase())
  );

  return (
    <div className="p-4 w-full">
      <p className="w-full flex items-center justify-center">
        Optional: skip for any office that applies
      </p>
      <div className="w-full">
        <input
          type="text"
          placeholder="Search..."
          value={searchOfficeTerm}
          onChange={handleSearchChange}
          className="path-bar mb-6 p-1 w-full"
        />
      </div>
      <form className="p-2">
        <div className="grid grid-cols-2 lg:grid-cols-3 gap-2 gap-x-6 -ml-4 lg:ml-2">
          {filteredOfficeItems.map((item, index) => (
            <div key={item} className="flex items-center pl-4">
              <input
                type="checkbox"
                id={`checkbox-${index}`}
                value={item}
                checked={checkedNWSOffices.includes(item)} // Check if it's selected
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
            selectedCounties: selectedCounties,
            stateList: stateList.join(","),
            warningTypes: checkedItems,
            checkedNWSOffices: checkedNWSOffices
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
    const savedStateListArr = user?.publicMetadata.stateList.split(",");
    setStateList(savedStateListArr);
    setSelectedCounties(user?.publicMetadata.selectedCounties);
    const savedWarningTypeArr = user?.publicMetadata.warningTypes;
    setCheckedItems(savedWarningTypeArr);
    setCheckedNWSOffices(user?.publicMetadata.checkedNWSOffices)
  }

  const handleSelectOfficeChange = (e) => {
    setNWSselectedOfficeList(e.target.value);
  };



  const AlertList = ({ alertList }) => {
    const [expandedItemIndex, setExpandedItemIndex] = useState(null);

    const toggleExpansion = (index) => {
      setExpandedItemIndex((prevIndex) => (prevIndex === index ? null : index));
    };

    if (alertList.length === 0 && hasSearchedForAlerts) {
      return (
        <ul className="p-2">
          <li
            className="alert-list shadow-md border-spacing-1"
            style={{ backgroundColor: "white" }}>
                <span className="headline items-center">All clear!</span>
            </li>
        </ul>
      )
    }

    return (
      <ul className="p-2">
        {alertList.map((obj, idx) => (
          <li
            key={obj.id}
            className="alert-list shadow-md border-spacing-1"
            style={{ backgroundColor: obj.color }}
          >
            <button
              className="w-full text-left"
              onClick={() => toggleExpansion(idx)}
            >
              {/* Access object properties and render them */}
              <span className="effective">{obj.effective}</span>
              <span className="headline">{obj.headline}</span>
              <br></br>
              <div className="areaDesc">{obj.areaDesc}</div>
            </button>
            <div
              className={`expanded-content ${
                expandedItemIndex === idx
                  ? "expanded"
                  : expandedItemIndex !== null
                  ? "collapsing"
                  : ""
              }`}
            >
              {expandedItemIndex === idx && (
                <pre>
                  <div className="p-4 text-wrap">{obj.description}</div>
                </pre>
              )}
            </div>
          </li>
        ))}
      </ul>
    );
  };

  return (
    <div>
      <Navbar />
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
          <div className="flex gap-1">
            <label className="label">Choose your specific counties:</label>
            <button
              onClick={() => {
                setShowCountiesForSelectedStates(!showCountiesForSelectedStates);
              }}
            >
              {showCountiesForSelectedStates ? (
                <IoIosArrowDown className="mt-[13px] w-8 hover:text-gray-500" />
              ) : (
                <IoIosArrowBack className="mt-[13px] w-8 hover:text-gray-500" />
              )}
            </button>
          </div>
          {showCountiesForSelectedStates ? <CountyListLibrary statesList={stateList}/> : ""}
          <div className="flex gap-1">
            <label className="label">Choose your specific warnings:</label>
            <button
              onClick={() => {
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
          <div className="flex gap-1">
            <label className="label">Choose your issuing office:</label>
            <button
              onClick={() => {
                setShowNWSOfficeFullList(!showNWSOfficeFullList);
              }}
            >
              {showNWSOfficeFullList ? (
                <IoIosArrowDown className="mt-[13px] w-8 hover:text-gray-500" />
              ) : (
                <IoIosArrowBack className="mt-[13px] w-8 hover:text-gray-500" />
              )}
            </button>
          </div>
          {showNWSOfficeFullList ? <NWSOfficeListForm /> : ""}
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
        <AlertList alertList={alertList} />
      </div>
      <Footer />
    </div>
  );
}
