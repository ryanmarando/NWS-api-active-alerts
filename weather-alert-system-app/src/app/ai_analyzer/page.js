"use client";
import React, { useState } from "react";
import axios from "axios";

const REACT_APP_OPEN_AI_KEY =
    "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";

const OPEN_AI_KEY = process.env.OPEN_AI_KEY;

export default function AiAnalyzer() {
    const [userInput, setUserInput] = useState("");
    const [response, setResponse] = useState("");
    const [loading, setLoading] = useState(false);
    const [videoFile, setVideoFile] = useState(null);
    const url = "http://localhost:8080";

    // Handle user input change
    const handleInputChange = (e) => {
        setUserInput(e.target.value);
    };

    // Handle video file change
    const handleFileChange = (e) => {
        setVideoFile(e.target.files[0]); // Capture the file selected by the user
    };

    // Handle form submission to upload video and get a response
    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!videoFile) {
            alert("Please upload a video file.");
            return;
        }

        setLoading(true);

        const formData = new FormData();
        formData.append("video", videoFile); // Append video file to FormData
        formData.append("userInput", userInput); // You can include any other input data if necessary

        try {
            // Send video file and user input to backend for processing
            const response = await axios.post(url + "/uploadVideo", formData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            });

            // Handle backend response (e.g., transcript)
            setResponse(response.data.transcript);
        } catch (error) {
            console.error("Error uploading video:", error);
            setResponse("Sorry, something went wrong.");
        }

        setLoading(false);
    };

    return (
        <div className="p-4 bg-white rounded shadow-md w-96 mx-auto">
            <h1 className="text-xl font-bold mb-4">Quick AI Response</h1>
            <form onSubmit={handleSubmit}>
                <textarea
                    className="w-full p-2 mb-4 border border-gray-300 rounded-md"
                    rows="3"
                    value={userInput}
                    onChange={handleInputChange}
                    placeholder="Ask something..."
                />

                {/* Video file input */}
                <input
                    type="file"
                    onChange={handleFileChange}
                    accept="video/*"
                    className="w-full p-2 mb-4 border border-gray-300 rounded-md"
                />

                <button
                    type="submit"
                    className="w-full p-2 bg-blue-500 text-white rounded-md"
                    disabled={loading}
                >
                    {loading ? "Loading..." : "Upload Video and Analyze"}
                </button>
            </form>

            {response && (
                <div className="mt-4 p-2 bg-gray-100 rounded-md">
                    <h2 className="text-lg font-semibold">AI Response:</h2>
                    <p>{response}</p>
                </div>
            )}
        </div>
    );
}
