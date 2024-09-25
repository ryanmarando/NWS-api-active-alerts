"use client";
import { Navbar } from "@/components/Navbar";
import { Footer } from "@/components/Footer";
import { MapContainer, TileLayer, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { useState } from 'react';
import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale } from 'chart.js';


ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale);


export default function WbgtPage() {
    const URL = "http://localhost:8080" //https://nws-api-active-alerts.onrender.com
    const [position, setPosition] = useState(null);
    const [wbgtForecast, setWbgtForecast] = useState([])
    const [wbgtLocation, setWbgtLocation] = useState("")
    const [loading, setLoading] = useState(false)

    function extractDateTime(timestamp) {
        const match = timestamp.match(/^([^/]+)\s*/);
        return match ? match[1] : timestamp;
      }
      
      function formatDateToAMPM(dateString) {
        const date = new Date(dateString);
        return new Intl.DateTimeFormat('en-US', {
          dateStyle: 'short',
          timeStyle: 'short',
          hour12: true
        }).format(date);
      }

const ClickableMap = () => {
  async function getWBGTForecast(pos) {
    let results = await fetch(
        URL + "/getWBGTForecastData/" +
          pos.lat + "," + pos.lng
      )
        .then((data) => data.json())
        .then((data) => {
          setWbgtForecast(data);
        });

    results = await fetch(
        URL + "/getWBGTForecastCityState"
      )
        .then((data) => data.json())
        .then((data) => {
          setWbgtLocation(data);
          setLoading(false)
        });
    }



  // This component listens to click events on the map
  const LocationMarker = () => {
    useMapEvents({
      click(e) {
        const clickedPosition = e.latlng;
        setPosition(clickedPosition); // Set the position based on the click
        getWBGTForecast(clickedPosition); // Call the API with the clicked position
        setLoading(true)
      },
    });

    return null;
  };

  return (
    <div>
      <MapContainer
        center={[37.8, -96]} // Center of the US
        zoom={4}
        style={{ height: '500px', width: '100%' }}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        <LocationMarker />
      </MapContainer>

      {/* Display the clicked position */}
      {position &&  (
        <div>
          <WBGTChart />
          <p>Clicked Position:</p>
          <p>Latitude: {position.lat}, Longitude: {position.lng}</p>
        </div>
      )}
    </div>
  );
};

const WBGTChart = () => {
    const chartData = {
      labels: wbgtForecast.map(item => formatDateToAMPM(extractDateTime(item.validTime))),
      datasets: [
        {
          label: 'WBGT Value',
          data: wbgtForecast.map(item => item.value),
          borderColor: 'rgba(75, 192, 192, 1)',
          backgroundColor: 'rgba(75, 192, 192, 0.2)',
          fill: true,
        },
      ],
    };
  
    const options = {
      responsive: true,
      scales: {
        x: {
          title: {
            display: true,
            text: 'Time',
          },
        },
        y: {
          title: {
            display: true,
            text: 'Value (°F)',
          },
        },
      },
    };
  
    return (
      <div>
        {loading ? <h1 className="w-full items-center justify-center flex">Loading...</h1>:
        <h1 className="w-full items-center justify-center flex">WBGT Forecast Chart for {wbgtLocation.properties.relativeLocation.properties.city}, {wbgtLocation.properties.relativeLocation.properties.state}</h1>
        }
        <Line data={chartData} options={options} />
      </div>
    );
  }


    return (
        <>
        <Navbar/>
        <div>
            <h1 className="font-bold flex items-center justify-center">Wet Bulb Globe Temperature Forecast</h1>
            <h3 className="flex items-center justify-center">Click a location to get the forecast:</h3>
            <ClickableMap />
        </div>
        <Footer/>
        </>
    )
}