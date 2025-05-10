"use client";

import React, { useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
  CartesianGrid,
} from "recharts";
import axios from "axios";
import dayjs from "dayjs";
import { API_URL } from "@/lib/constants";
import utc from "dayjs/plugin/utc";
import { Navbar } from "@/components/Navbar";
import { Footer } from "@/components/Footer";

export default function ModelTrendVisualizer() {
  const [data, setData] = useState([]);
  const [allData, setAllData] = useState([]); // raw data for filtering
  const [runTime, setRunTime] = useState(null);
  const [location, setLocation] = useState("KDAY");
  const [selectedRunHour, setSelectedRunHour] = useState(null); // selected UTC hour
  const [selectedParameters, setSelectedParameters] = useState({
    TMP: true,
    DPT: true,
    WSP: true,
    WDR: true,
    SKY: true,
  });
  const [fetching, setFetching] = useState(false);

  const fetchModelData = async () => {
    if (!location) return;

    setFetching(true);

    try {
      // Step 1: Get the latest runTime for this location
      const latestRunRes = await axios.get(
        `${API_URL}/modelTrender/getLatestRunTime?location=${location}`
      );
      const latestRunTime = latestRunRes.data.latestRunTime;

      if (!latestRunTime) {
        console.warn("No latest runTime found");
        setData([]);
        setRunTime(null);
        setFetching(false);
        return;
      }

      // Extract the hour from latestRunTime for filtering
      dayjs.extend(utc);
      const latestHour = dayjs(latestRunTime).utc().hour();
      setRunTime(latestRunTime);
      setSelectedRunHour(latestHour);

      // Step 2: Get the full model data for this location
      const modelRes = await axios.get(
        `${API_URL}/modelTrender/getModelComparison?location=${location}`
      );
      const raw = modelRes.data;

      if (raw.length === 0) {
        setData([]);
        setFetching(false);
        return;
      }

      setAllData(raw); // store full raw data for filtering later
      processAndSetData(raw, latestHour);
    } catch (err) {
      console.error("Failed to fetch model data:", err);
    } finally {
      setFetching(false);
    }
  };

  const processAndSetData = (raw, filterHour) => {
    const filtered = raw.filter(
      (item) => dayjs(item.runTime).hour() === filterHour
    );

    if (filtered.length === 0) {
      setData([]);
      return;
    }

    const groupedByParameter = filtered.reduce((acc, item) => {
      const key = item.parameter;
      if (!acc[key]) acc[key] = [];
      acc[key].push(item);
      return acc;
    }, {});

    const byValidTime = {};

    Object.keys(groupedByParameter).forEach((param) => {
      groupedByParameter[param].forEach((point) => {
        const vt = point.validTime;
        if (!byValidTime[vt]) {
          byValidTime[vt] = { validTime: vt };
        }
        byValidTime[vt][param] = point.value;
        byValidTime[vt][`${param}_Trend`] = point.trend;
        byValidTime[vt][`${param}_Delta`] = point.delta;
        byValidTime[vt][`${param}_Original`] = point.original;
      });
    });

    const final = Object.values(byValidTime).sort(
      (a, b) => new Date(a.validTime) - new Date(b.validTime)
    );

    setData(final);
  };

  const handleRunHourClick = (hour) => {
    setSelectedRunHour(hour);
    processAndSetData(allData, hour);
  };

  const formatTime = (iso) => dayjs(iso).format("dd ha");

  const renderCustomTooltip = ({ active, payload, label }) => {
    if (active && payload?.length) {
      const pt = payload[0].payload;
      return (
        <div className="bg-white p-2 rounded shadow text-sm">
          <div>
            <strong>{formatTime(label)}</strong>
          </div>
          {Object.keys(pt)
            .filter(
              (k) =>
                !k.includes("validTime") &&
                !k.includes("Trend") &&
                !k.includes("Delta") &&
                !k.includes("Original")
            )
            .map((key) => {
              let unit = "°F";
              if (key === "WSP") unit = "mph";
              if (key === "SKY") unit = "%";
              return (
                <div key={key}>
                  {key}: {pt[key]}
                  {unit} ({pt[`${key}_Trend`]}, Δ{pt[`${key}_Delta`]})
                </div>
              );
            })}
        </div>
      );
    }
    return null;
  };

  const handleToggleChange = (param) => {
    setSelectedParameters((prev) => ({
      ...prev,
      [param]: !prev[param],
    }));
  };

  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      <header className="bg-blue-600 text-white text-xl font-semibold py-4 px-6">
        NBM Model Trend Visualizer
      </header>

      <main className="flex-1 p-4">
        <div className="max-w-md mx-auto mb-4">
          <label
            htmlFor="stationCode"
            className="block text-sm font-medium text-gray-700 mb-1"
          >
            Enter Station Code:
          </label>
          <input
            id="stationCode"
            type="text"
            value={location}
            onChange={(e) => setLocation(e.target.value.toUpperCase())}
            placeholder="e.g., KDAY"
            className="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring focus:border-blue-300"
          />
        </div>

        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Select Parameters to Display:
          </label>
          <div className="space-y-2">
            {["TMP", "DPT", "WSP", "WDR", "SKY"].map((param) => (
              <div key={param} className="flex items-center">
                <input
                  type="checkbox"
                  checked={selectedParameters[param]}
                  onChange={() => handleToggleChange(param)}
                  id={param}
                  className="mr-2"
                />
                <label htmlFor={param} className="text-sm">
                  {param}
                </label>
              </div>
            ))}
          </div>
        </div>

        <div className="text-center mb-4">
          <button
            onClick={fetchModelData}
            className="bg-blue-600 text-white px-4 py-2 rounded"
            disabled={fetching}
          >
            {fetching ? "Loading..." : "Fetch Data"}
          </button>
        </div>

        {runTime && (
          <div className="text-center text-gray-500 text-sm mb-4">
            Latest Model Run Time: {dayjs(runTime).format("YYYY-MM-DD HH:mm")}{" "}
            UTC
          </div>
        )}

        {/* UTC Run Hour Selector */}
        <div className="flex flex-wrap justify-center gap-2 mb-6">
          {[...Array(24)].map((_, i) => (
            <button
              key={i}
              onClick={() => handleRunHourClick(i)}
              className={`px-3 py-1 rounded border ${
                selectedRunHour === i
                  ? "bg-blue-600 text-white"
                  : "bg-white text-gray-700 hover:bg-gray-100"
              }`}
            >
              {i}Z
            </button>
          ))}
        </div>

        {data.length === 0 ? (
          <div className="text-center text-gray-600">
            No data found for {selectedRunHour}Z
          </div>
        ) : (
          <div className="w-full h-[500px]">
            <ResponsiveContainer width="100%" height="100%">
              <LineChart
                data={data}
                margin={{ top: 20, right: 30, left: 10, bottom: 5 }}
              >
                <CartesianGrid strokeDasharray="3 3" stroke="#000000" />
                <XAxis dataKey="validTime" tickFormatter={formatTime} />
                <YAxis />
                <Tooltip content={renderCustomTooltip} />
                <Legend
                  content={({ payload }) => (
                    <div className="flex flex-wrap justify-center gap-6 mt-4">
                      {payload.map((entry, index) => {
                        const isOriginal = entry.value.includes("(original)");
                        return (
                          <div
                            key={`item-${index}`}
                            className="flex items-center space-x-2 text-sm"
                          >
                            <div
                              style={{
                                width: 24,
                                height: 0,
                                borderBottom: isOriginal
                                  ? `2px dashed ${entry.color}`
                                  : `2px solid ${entry.color}`,
                              }}
                            />
                            <span>{entry.value}</span>
                          </div>
                        );
                      })}
                    </div>
                  )}
                />

                {["TMP", "DPT", "WSP", "WDR", "SKY"].map((param, idx) => {
                  if (!selectedParameters[param]) return null;
                  return (
                    <React.Fragment key={param}>
                      <Line
                        type="monotone"
                        dataKey={param}
                        name={`${param} (latest)`}
                        stroke={
                          [
                            "#ef4444",
                            "#3b82f6",
                            "#10b981",
                            "#6366f1",
                            "#f59e0b",
                            "#8b5cf6",
                            "#34d399",
                          ][idx]
                        }
                        strokeWidth={2}
                        dot={false}
                        isAnimationActive={false}
                      />
                      <Line
                        type="monotone"
                        dataKey={`${param}_Original`}
                        name={`${param} (original)`}
                        stroke={
                          [
                            "#ef4444",
                            "#3b82f6",
                            "#10b981",
                            "#6366f1",
                            "#f59e0b",
                            "#8b5cf6",
                            "#34d399",
                          ][idx]
                        }
                        strokeDasharray="4 2"
                        strokeWidth={1.5}
                        dot={false}
                        isAnimationActive={false}
                      />
                    </React.Fragment>
                  );
                })}
              </LineChart>
            </ResponsiveContainer>
          </div>
        )}
      </main>

      <Footer />
    </div>
  );
}
