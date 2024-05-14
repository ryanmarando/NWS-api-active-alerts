import Image from "next/image";
import Logo from "@/app/icon.ico";
import Facebook from "../../public/assets/Facebook.svg";
import Twitter from "../../public/assets/X.svg";

export function Footer() {
  return (
    <footer className="pt-[80px] pb-[40px]">
      <div className="flex items-center justify-center gap-x-[12px]">
        <Image className="h-[35px] w-[35px]" src={Logo} alt="Logo" />
        <p className="font-bold text-[#36485C] text-[17px]">Ryan Marando</p>
      </div>
      <p className="pt-[20px] pb-[14px] text-center text-[14px] font-medium text-[#5F7896]">
        © Copyright 2024. All rights reserved.
      </p>
      <div className="flex items-center justify-center gap-y-[32px] pt-[20px]">
        <a
          href="https://www.facebook.com/ryanmarandowx"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image className="mr-[15px] w-8 h-8" alt="Facebook" src={Facebook} />
        </a>
        <a
          href="https://twitter.com/ryanmarando"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image className="mr-[15px] w-8 h-8" alt="Twitter" src={Twitter} />
        </a>
      </div>
    </footer>
  );
}
