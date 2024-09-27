import Link from "next/link";
import Image from "next/image";
import Gradient from "../../public/assets/Gradient.svg";
import { motion } from "framer-motion";
import { TypeAnimation } from "react-type-animation";
import YouTubeEmbed from "./Youtube";

const TypingAnimation = () => {
  return (
    <TypeAnimation
      sequence={[
        // Same substring at the start will only be typed out once, initially
        "Stay weather aware.",
        1000, // wait 1s before replacing "Mice" with "Hamsters"
        "Rise above the storm.",
        1000,
        "It's always darkest before dawn.",
        1000,
        "Clear Moon, frost soon.",
        1000,
      ]}
      wrapper="span"
      speed={50}
      repeat={Infinity}
    />
  );
};

export function Hero() {
  return (
    <div className="pt-4">
      <motion.div
        initial="hidden"
        animate="visible"
        variants={{
          hidden: { scale: 0.8, opacity: 0 },
          visible: {
            scale: 1,
            opacity: 1,
            transition: {
              delay: 0.4,
            },
          },
        }}
      >
        <div className="pt-5">
          <h1 className="text-center font-semibold text-[32px] leading-[40px] text-[#172026]">
            Weather. But from the future.
          </h1>
        </div>
      </motion.div>
      <motion.div
        initial="hidden"
        animate="visible"
        variants={{
          hidden: { scale: 0.8, opacity: 0 },
          visible: {
            scale: 1,
            opacity: 1,
            transition: {
              delay: 0.4,
            },
          },
        }}
      >
        <div className="flex w-full items-center justify-center mt-8">
          <Link
            className="bg-[#4328EB] text-center lg:w-1/4 sm:w-1/2 py-4 px-8 ml-3 mt-4 text-white rounded-[4px]"
            href="/alert-system"
          >
            <button>Try the Weather Alert System</button>
          </Link>
          <Link
            className="text-[#4328EB] text-center font-semibold lg:w-1/4 sm:w-1/2  py-4 px-8 ml-3 mt-4"
            href="/contact"
          >
            <button>Want something custom made?</button>
          </Link>
        </div>
      </motion.div>
      <motion.div
        initial="hidden"
        animate="visible"
        variants={{
          hidden: { scale: 0.8, opacity: 0 },
          visible: {
            scale: 1,
            opacity: 1,
            transition: {
              delay: 0.4,
            },
          },
        }}
      >
      <div className="relative flex h-full w-full justify-center">
      <div className="absolute  w-full flex text-center font-semibold text-[32px] lg:text-[64px]  text-[#172026] items-center justify-center mt-14 lg:mt-20 z-10">
        <TypingAnimation></TypingAnimation>
      </div>
      <div className="absolute flex mt-60 mr-[50%]">
        <YouTubeEmbed url={"https://youtu.be/g8IAGWvhe-s"} autoplay={true}/>
      </div>
      <div className="absolute flex ml-[45%] mt-48 items-center">
        <h2 className="font-semibold text-[20px] text-center">About Ryan Marando</h2>
      </div>
      <div className="absolute flex ml-[45%] mt-60 w-[50%] text-center">
        <p className="border-2 border-gray-800 rounded-md p-2">Colleagues call me an engaging and compelling meteorologist who blends storytelling and science with ease. As a Meteorologist for WHIO in Dayton, Ohio, <strong>I shape my forecasts to help the community understand the impacts of everyday to dangerous, severe weather.</strong></p>
        <p className="border-2 border-gray-800 rounded-md p-2 ml-2">What I enjoy most is serving my community by producing accurate, digestible, and refreshing forecasts using creativity with augmented 3D graphics and enhancing them with my Python certification and programming experience. <strong>To me, it's important to not only tell the weather story but also the context and perspective behind it has so viewers truly understand the forecast.</strong></p>
        <p className="border-2 border-gray-800 rounded-md p-2 ml-2">Weather is my passion! <strong>My constant curiosity about our world and the impact science has started when I was young growing up in Bear, Delaware.</strong> That love fueled my education graduating Summa Cum Laude as Embry-Riddle Aeronautical University's Meteorology Student of the Year.</p>
      </div>
        <Image
          src={Gradient}
          alt="Gradient"
          className="min-h-[500px] w-full object-cover"
        />
      </div>
      </motion.div>
    </div>
  );
}
