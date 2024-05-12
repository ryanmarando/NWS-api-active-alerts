import Image from "next/image";
import Logo from "@/app/icon.ico";
import Menu from "../../public/assets/Menu.svg";
import Facebook from "../../public/assets/Facebook.svg";
import Twitter from "../../public/assets/X.svg";

export function Navbar() {
  return (
    <nav className="flex w-full items-center justify-between px-[20px] py-[16px] lg:container lg:mx-auto lg:px-22">
      <div className="flex gap-x-5">
        <Image alt="Logo" src={Logo} />
        <h1 className="py-[16px]">Weather Alert System</h1>
      </div>
      <div className="flex items-center">
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
        <Image alt="Menu" src={Menu} className="lg:hidden" />
      </div>
    </nav>
  );
}
