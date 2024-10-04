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
        },
        {
          title: 'Helene Tracker',
          url: 'https://youtu.be/o0SnyxxKjB4',
        },
        {
          title: 'Wet Bulb Globe Temperature Explainer',
          url: 'https://youtu.be/LAhao2jE1qk',
        },
        {
          title: '3D Set',
          url: 'https://youtu.be/Cq5njUmZkJs',
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
          <div>
            <h1 className="ml-2">Broadcasting Graphics</h1>
          </div>
        )
      }

    return (
        <div>
            <Navbar />
            <Hero />
            <div className="container mx-auto p-4 mb-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {videos.map((video, index) => (
                  <div key={index} className="flex justify-center">
                     <YouTubeEmbed key={index} title={video.title} url={video.url} />
                  </div>
                ))}
                </div>
            </div>
            <Graphics />
            <Footer />
        </div>
    )
}