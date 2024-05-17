import { useState } from "react";

export function AlertForm() {
  const [checkedItems, setCheckedItems] = useState([]);
  const [submittedItems, setSubmittedItems] = useState([]);
  const [selectAll, setSelectAll] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");
  const [buttonText, setButtonText] = useState("Submit Alert Types");

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
    <div className="p-4">
      <p className="w-full flex items-center justify-center">
        Optional: skip for all alerts
      </p>
      <input
        type="text"
        placeholder="Search..."
        value={searchTerm}
        onChange={handleSearchChange}
        className="path-bar mb-6 p-1 w-full"
      />
      <form onSubmit={handleSubmit}>
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
        <div className=" grid grid-cols-3 gap-2 gap-x-6">
          {filteredItems.map((item, index) => (
            <div key={index} className="flex items-center">
              <input
                type="checkbox"
                id={`checkbox-${index}`}
                value={item}
                checked={checkedItems.includes(item)}
                onChange={handleCheckboxChange}
                className="mr-2 text-center"
              />
              <label htmlFor={`checkbox-${index}`}>{item}</label>
            </div>
          ))}
        </div>
      </form>
      <div className="w-full flex items-center justify-center">
        <button
          type="submit"
          onClick={handleSubmit}
          className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-white my-[5px] mx-[15px] mt-[14px]"
        >
          {buttonText}
        </button>
      </div>
    </div>
  );
}
