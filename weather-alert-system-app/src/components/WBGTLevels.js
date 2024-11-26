import React from "react";

const WBGTLevels = () => {
  // Data representing WBGT levels and corresponding safety measures
  const wbgtLevels = [
    {
      level: "Low",
      range: "< 80°F",
      description: "Safe for most outdoor activities.",
      color: "#00e676",
    }, // Green
    {
      level: "Moderate",
      range: "80°F – 84°F",
      description:
        "Caution: risk for high-intensity activities. Take breaks and hydrate.",
      color: "#ffff00",
    }, // Yellow
    {
      level: "High",
      range: "84°F – 88°F",
      description:
        "Increased risk of heat illness. Frequent rest breaks recommended.",
      color: "#ffa726",
    }, // Orange
    {
      level: "Very High",
      range: "88°F – 90°F",
      description:
        "Heat illness likely with prolonged activity. Limit outdoor activities.",
      color: "#ff7043",
    }, // Dark Orange
    {
      level: "Extreme",
      range: "> 90°F",
      description:
        "Extreme caution! Heat stroke is possible. Avoid strenuous outdoor activities.",
      color: "#f44336",
    }, // Red
  ];

  return (
    <div className="p-4">
      <div className="alert-system">
        <h1 className="flex items-center justify-center text-lg font-bold mb-4">
          Wet Bulb Globe Temperature (WBGT) Levels
        </h1>
        <p className="mb-4">
          The Wet Bulb Globe Temperature (WBGT), according to the National
          Weather Service, is a measure of heat stress in direct sunlight, which
          takes into account: temperature, humidity, wind speed, sun angle, and
          cloud cover (solar radiation). This differs from the heat index, which
          takes into consideration temperature and humidity and is calculated
          for shady areas.
        </p>
        <div className="overflow-x-auto w-[80%]">
          <table className="w-full border-collapse table-auto mt-4">
            <thead>
              <tr className="bg-gray-200 text-left">
                <th className="p-3 border-b-2 border-gray-300">Level</th>
                <th className="p-3 border-b-2 border-gray-300">
                  Temperature Range
                </th>
                <th className="p-3 border-b-2 border-gray-300">Description</th>
              </tr>
            </thead>
            <tbody>
              {wbgtLevels.map((level, index) => (
                <tr
                  key={index}
                  className="text-left"
                  style={{ backgroundColor: level.color }}
                >
                  <td className="p-3 border-b border-gray-300">
                    {level.level}
                  </td>
                  <td className="p-3 border-b border-gray-300">
                    {level.range}
                  </td>
                  <td className="p-3 border-b border-gray-300">
                    {level.description}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default WBGTLevels;
