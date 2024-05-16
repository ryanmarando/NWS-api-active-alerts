import Link from "next/link";
import Image from "next/image";
import Gradient from "../../public/assets/Gradient.svg";
import { motion } from "framer-motion";
import { TypeAnimation } from "react-type-animation";

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

      <div className="absolute  w-full flex text-center font-semibold text-[32px] lg:text-[64px]  text-[#172026] items-center justify-center mt-28 lg:mt-56 z-10">
        <TypingAnimation></TypingAnimation>
      </div>

      <div className="relative flex h-full w-full justify-center">
        <Image
          src={Gradient}
          alt="Gradient"
          className="min-h-[500px] w-full object-cover"
        />
      </div>
    </div>
  );
}
