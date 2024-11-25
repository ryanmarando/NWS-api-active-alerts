"use client";
import { Navbar } from "../../components/Navbar";
import { Footer } from "../../components/Footer";
import YouTubeEmbed from "@/components/Youtube";
import Image from 'next/image';

export default function Broadcasting() {
    const videos = [
        {
          title: 'Every Day',
          url: 'https://www.youtube.com/watch?v=g8IAGWvhe-s',
          autoplay: true,
        },
        {
          title: 'Helene Tracker',
          url: 'https://youtu.be/o0SnyxxKjB4',
          autoplay: false,
        },
        {
          title: 'Wet Bulb Globe Temperature Explainer',
          url: 'https://youtu.be/LAhao2jE1qk',
          autoplay: false,
        },
        {
          title: '3D Set',
          url: 'https://youtu.be/Cq5njUmZkJs',
          autoplay: false,
        },
        // Add more video objects as needed
      ];

      const Hero = () => {
        return (
          <section className="relative h-[40vh] w-full flex justify-center items-center overflow-hidden">
            <div className="absolute bottom-0 left-0 p-4 z-9">
            <p className="text-white p-2 text-sm w-[70%] mb-8 lg:text-base lg:w-[40%] lg:mb-28">
  Colleagues call me an engaging and compelling meteorologist who blends storytelling and science with ease. 
  As a Meteorologist for WHIO in Dayton, Ohio, <strong>I shape my forecasts to help the community understand the impacts of everyday to dangerous, severe weather.</strong>
</p>
              <span className="p-2 border-2 rounded-md border-white text-2xl md:text-4xl font-bold text-white">Ryan Marando</span>
            </div>
            <Image 
              src="/assets/demo_page_heropic.png" 
              alt="Ryan Marando"
              layout="fill"
              objectPosition="60% 35%"
              objectFit="cover"
              quality={100}
              className="z-[-1] brightness-75"
            />
          </section>
        );
      };

      const Graphics = () => {
        return (
          <div className="p-4">
            <div className="flex items-center sm:max-h-[300px] ">
            <h1 className="sm: text-md ml-2 lg:text-2xl font-bold mb-6 text-center">Graphics Powered With Programming</h1>
            <p className="text-sm text-center w-[70%] mb-8 lg:text-base">
              These graphics use more than the every day graphics system of the Max system or Baron Lynx.
              Using computer programming from python to automate graphics for social media and on air use.
            </p>
            </div>
            {/* Top row: 3 images with titles */}
            <div className="lg:grid lg:grid-cols-3 lg:gap-4">
              <div className="text-center">
                <h2 className="font-medium text-lg">Alert System</h2>
                <img
                  src="/assets/alert_system_coded.png"
                  alt="Graphic 1"
                  className="w-full object-cover"
                />
              </div>
              <div className="text-center">
                <h2 className="font-medium text-lg">Hurricane Forecast Ingredients</h2>
                <img
                  src="/assets/hurricane_coded.png"
                  alt="Graphic 2"
                  className="w-full object-cover"
                />
              </div>
              <div className="text-center">
                <h2 className="font-medium text-lg">Rainfall Totals</h2>
                <img
                  src="/assets/rainfall_coded.png"
                  alt="Graphic 3"
                  className="w-full object-cover"
                />
              </div>
            </div>
      
            {/* Second row: 2 images centered horizontally */}
            <div className="lg:flex lg:justify-center lg:gap-4 lg:mt-4">
              <div className="text-center">
                <h2 className="font-medium text-lg">Temperature Readings</h2>
                <img
                  src="/assets/temps_coded.png"
                  alt="Graphic 4"
                  className="w-full lg:h-60 object-cover"
                />
              </div>
              <div className="text-center">
                <h2 className="font-medium text-lg">Wet Bulb Globe Temperature Forecasts</h2>
                <img
                  src="/assets/wbgt_coded.png"
                  alt="Graphic 5"
                  className="w-full lg:h-60 object-cover"
                />
              </div>
            </div>
          </div>
        );
      };

      
      
      

    return (
        <div>
            <Navbar />
            <Hero />
            <div className="container mx-auto p-4 mb-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {videos.map((video, index) => (
              <div key={index} className="flex justify-center">
                <YouTubeEmbed 
                key={index} 
                title={video.title} 
                url={video.url} 
                autoplay={video.autoplay}
                className="w-full sm:max-w-[300px] md:max-w-[600px] lg:max-w-[500px] aspect-video"
              />
              </div>
              ))}
            </div>
            </div>
            <Graphics />
            <Footer />
        </div>
    )
}