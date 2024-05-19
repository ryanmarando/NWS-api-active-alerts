"use client";
import Link from "next/link";
import { Navbar } from "@/components/Navbar";
import { Footer } from "@/components/Footer";

export default function PricingPage() {
  return (
    <div>
      <Navbar />
      <div className="flex items-center justify-center">
        <div className="max-w-md w-full bg-white shadow-lg rounded-lg overflow-hidden">
          <div className="px-6 py-8">
            <div className="flex justify-center">
              <h2 className="text-xl font-bold text-gray-800">
                Weather Alert System Pro
              </h2>
            </div>
            <div className="mt-4 flex justify-center items-center">
              <span className="text-3xl font-bold text-gray-800">$50</span>
              <span className="text-gray-500 ml-1">/ month</span>
            </div>
            <div className="mt-8">
              <ul>
                <li className="flex items-center mt-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6 text-green-500 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                  <span className="text-gray-700">Automated Alerts</span>
                </li>
                <li className="flex items-center mt-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6 text-green-500 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                  <span className="text-gray-700">Customizable warnings</span>
                </li>
                <li className="flex items-center mt-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6 text-green-500 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                  <span className="text-gray-700">Savable Settings</span>
                </li>
                <li className="flex items-center mt-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6 text-green-500 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                  <span className="text-gray-700">Tech support</span>
                </li>
              </ul>
            </div>
            <div className="mt-6 flex justify-center">
              <Link href="/payment">
                <button className="bg-[#4328EB] hover:text-gray-500 w-33 py-2 px-4 rounded-[8px] text-white my-[5px] mx-[15px] mt-[14px]">
                  Subscribe
                </button>
              </Link>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}
