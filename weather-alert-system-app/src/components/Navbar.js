"use client";
import { useState, useEffect } from "react";
import { useUser } from "@clerk/clerk-react";
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
  const { user } = useUser();
  const [hasSubscription, setHasSubscription] = useState();
  useEffect(() => {
    const fetchPrivateMetadata = async () => {
      if (!user?.id) return;

      try {
        const response = await fetch(
          `/api/updatePrivateMetadata?userId=${user?.id}`
        );
        if (!response.ok) {
          throw new Error("Failed to fetch user data");
        }
        const data = await response.json();
        setHasSubscription(data.privateMetadata.subscription);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    fetchPrivateMetadata();
  }, [user?.id]);
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
        <div className="relative group">
  {/* Main Products Link */}
  <button className="hidden lg:block py-[16px] hover:bg-gray-100 rounded-md">
    Products
  </button>

  {/* Dropdown Content */}
  <div className="absolute hidden group-hover:block bg-white shadow-lg rounded-md mt-2 z-20">
    <Link
      href="/wbgt"
      className="block px-4 py-2 text-gray-700 hover:bg-[rgb(221,221,221)] rounded-md"
    >
      Wet Bulb Globe Temperatures
    </Link>
  </div>
</div>
<Link className="hidden lg:block py-[16px]" href="/broadcasting">
          Broadcasting
        </Link>
        <Link className="hidden lg:block py-[16px]" href="/contact">
          Contact
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
            <Link href="/sign-in">
              <p className="font-medium text-[#36485C]">Sign in</p>
            </Link>
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
