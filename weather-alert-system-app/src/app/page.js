"use client";
import { useState, useEffect } from "react";
import React from "react";
import Select from "react-select";
import Image from "next/image";
import Clock from "@/components/Clock";
import LoadingAnimation from "@/components/LoadingAnimation";
import CountdownTimer from "@/components/Counter";
import { Navbar } from "@/components/Navbar";
import BlueArrow from "../../public/assets/blue-button.svg";
import { Footer } from "@/components/Footer";
import { Hero } from "@/components/Hero";

export default function Home() {
  return (
    <div>
      <Navbar />
      <Hero />
      <Footer />
    </div>
  );
}
