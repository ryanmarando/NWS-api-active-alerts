"use client";
import { Navbar } from "../../components/Navbar";
import { Footer } from "../../components/Footer";
import YouTubeEmbed from "@/components/Youtube";

export default function Reel() {
    const videos = [
        {
          title: 'Wet Bulb Globe Temperature Explained',
          url: 'https://www.youtube.com/watch?v=g8IAGWvhe-s',
        },
        {
          title: 'Understanding Weather Patterns',
          url: 'https://www.youtube.com/watch?v=dQw4w9WgXcQ',
        },
        {
          title: 'Climate Change and Its Effects',
          url: 'https://www.youtube.com/watch?v=3JZ_D3ELwOQ',
        },
        {
          title: 'Meteorology Basics',
          url: 'https://www.youtube.com/watch?v=9bZkp7q19f0',
        },
        // Add more video objects as needed
      ];

    return (
        <div>
            <Navbar />
            <div className="container mx-auto p-4">
                <h1 className="text-2xl font-bold text-center mb-6">YouTube Videos</h1>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {videos.map((video, index) => (
                <div key={index} className="flex justify-center">
                     <YouTubeEmbed key={index} title={video.title} url={video.url} />
                </div>
                ))}
            </div>
            </div>
            <Footer />
        </div>
    )
}