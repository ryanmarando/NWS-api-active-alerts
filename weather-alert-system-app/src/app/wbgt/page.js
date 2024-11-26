"use client";
import { Navbar } from "@/components/Navbar";
import { Footer } from "@/components/Footer";
import { MapContainer, TileLayer, useMapEvents } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import { useState, useEffect } from "react";
import WBGTLevels from "@/components/WBGTLevels";
import LoadingAnimation from "@/components/LoadingAnimation";
import dynamic from "next/dynamic";

export default function WbgtPage() {
  const [clientReady, setClientReady] = useState(false);
  const ClickableMap = dynamic(() => import("@/components/ClickableMap"), {
    ssr: false,
  });

  useEffect(() => {
    setClientReady(true); // This ensures code only runs in the client
  }, []);

  return (
    <>
      <Navbar />
      <div>
        <h1 className="font-bold flex items-center justify-center">
          Wet Bulb Globe Temperature Forecast
        </h1>
        <h3 className="flex items-center justify-center">
          Click a location to get the forecast:
        </h3>
        <div className="relative">
          {!clientReady ? (
            <div className="inset-0 bg-white bg-opacity-75 flex items-center justify-center z-10">
              <LoadingAnimation />
            </div>
          ) : (
            <ClickableMap />
          )}
        </div>
        <WBGTLevels />
      </div>
      <Footer />
    </>
  );
}
