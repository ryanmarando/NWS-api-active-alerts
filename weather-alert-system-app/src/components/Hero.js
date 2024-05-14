import Link from "next/link";
import Image from "next/image";
import Gradient from "../../public/assets/Gradient.svg";
import HeroAlertImage from "../../public/assets/alert-system.png";

export function Hero() {
  return (
    <div className="pt-4">
      <div className="pt-5">
        <h1 className="text-center font-semibold text-[32px] leading-[40px] text-[#172026]">
          Weather. But from the future.
        </h1>
      </div>

      <div className="flex w-full items-center justify-center mt-8">
        <Link
          className="bg-[#4328EB] text-center lg:w-1/4 sm:w-1/2 py-4 px-8 ml-3 mt-4 text-white rounded-[4px]"
          href="/alert-system"
        >
          <button>Try the Weather Alert System</button>
        </Link>
        <button className="text-[#4328EB] text-center font-semibold lg:w-1/4 sm:w-1/2  py-4 px-8 ml-3 mt-4">
          Want something custom made?
        </button>
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
