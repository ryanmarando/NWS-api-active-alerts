"use client";

import { useEffect, useState } from "react";
import { useUser } from "@clerk/nextjs";
import { Navbar } from "@/components/Navbar";
import { Footer } from "@/components/Footer";
import { API_URL } from "@/lib/constants";

export default function AdminPage() {
  const { user, isLoaded } = useUser();
  const [isAdmin, setIsAdmin] = useState(false);
  const [titles, setTitles] = useState(["", "", "", "", ""]);
  const [urls, setUrls] = useState(["", "", "", "", ""]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    const checkAdmin = async () => {
      if (!user) return;
      try {
        const response = await fetch(
          `/api/updatePrivateMetadata?userId=${user?.id}`
        );
        const data = await response.json();
        setIsAdmin(data.privateMetadata.admin === true);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    if (isLoaded) checkAdmin();
  }, [isLoaded, user]);

  useEffect(() => {
    const fetchYoutubeData = async () => {
      try {
        const res = await fetch(`${API_URL}/youtubeURLs`);
        const data = await res.json();
        setTitles(data.titles || ["", "", "", "", ""]);
        setUrls(data.urls || ["", "", "", "", ""]);
      } catch (error) {
        console.error("Error fetching YouTube data:", error);
        setMessage("Failed to load YouTube links.");
      }
    };

    if (isAdmin) fetchYoutubeData();
  }, [isAdmin]);

  const handleUpdate = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${API_URL}/youtubeURLs`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ titles, urls }),
      });

      if (!response.ok) throw new Error("Update failed");

      const result = await response.json();
      setMessage("YouTube URLs updated successfully!");
      console.log("Updated:", result);
    } catch (error) {
      console.error("PATCH error:", error);
      setMessage("Failed to update YouTube URLs.");
    }
  };

  if (!isAdmin) {
    return (
      <div>
        <Navbar />
        <h1 className="flex justify-center">Not Authorized</h1>
        <Footer />
      </div>
    );
  }

  return (
    <div>
      <Navbar />
      <h1 className="text-center my-4">Admin Page</h1>

      <form onSubmit={handleUpdate} className="max-w-2xl mx-auto space-y-6">
        <h2 className="text-xl font-semibold">Edit YouTube Titles and URLs</h2>

        {[...Array(5)].map((_, idx) => (
          <div key={idx} className="border p-4 rounded shadow">
            <label className="block mb-2 font-medium">
              Title {idx + 1}
              <input
                type="text"
                value={titles[idx] || ""}
                onChange={(e) =>
                  setTitles((prev) => {
                    const updated = [...prev];
                    updated[idx] = e.target.value;
                    return updated;
                  })
                }
                className="w-full border p-2 rounded mt-1"
              />
            </label>
            <label className="block mt-4 font-medium">
              URL {idx + 1}
              <input
                type="url"
                value={urls[idx] || ""}
                onChange={(e) =>
                  setUrls((prev) => {
                    const updated = [...prev];
                    updated[idx] = e.target.value;
                    return updated;
                  })
                }
                className="w-full border p-2 rounded mt-1"
              />
            </label>
          </div>
        ))}

        <button
          type="submit"
          className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700"
        >
          Update YouTube Links
        </button>

        {message && (
          <p className="text-center mt-4 text-green-600">{message}</p>
        )}
      </form>

      <Footer />
    </div>
  );
}
