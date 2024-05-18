import Image from "next/image";
import Link from "next/link";
import Logo from "@/app/icon.ico";
import Menu from "../../public/assets/Menu.svg";
import Facebook from "../../public/assets/Facebook.svg";
import Twitter from "../../public/assets/X.svg";
import User from "../../public/assets/User.svg";
import { SignInButton, SignedIn, SignedOut, UserButton } from "@clerk/nextjs";
import { MobileNav } from "@/components/MobileNav.js";

<Image alt="Menu" src={Menu} className="lg:hidden" />;
export function Navbar() {
  return (
    <nav className="flex w-full items-center justify-between px-[20px] py-[16px] lg:container lg:mx-auto lg:px-22">
      <div className="flex gap-x-5">
        <Link href="/">
          <Image alt="Logo" src={Logo} className="object-cover" />
        </Link>
        <Link
          className="hidden lg:block py-[16px] text-center"
          href="/alert-system"
        >
          Weather Alert System
        </Link>
        <Link className="hidden lg:block py-[16px]" href="/contact">
          Contact
        </Link>
        <Link className="hidden lg:block py-[16px]" href="/payment">
          Pricing
        </Link>
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
        <div className="flex items-center mr-[15px]">
          <SignedOut>
            <SignInButton className="font-medium text-[#36485C]" />
          </SignedOut>
          <SignedIn>
            <UserButton />
          </SignedIn>
        </div>
        <div>
          <MobileNav />
        </div>
      </div>
    </nav>
  );
}
