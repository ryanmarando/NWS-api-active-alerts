"use client";
import { useUser } from "@auth0/nextjs-auth0/client";
import Link from "next/link";
import User from "../../public/assets/User.svg";
import Image from "next/image";

export function UserInfo() {
  const { user, error, isLoading } = useUser();
  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>{error.message}</div>;
  //console.log(user);
  if (!user)
    return (
      <div className="flex items-center">
        <Link href="/api/auth/login">
          <div className="flex items-center">
            <Image
              className="mr-[15px] w-8 h-8"
              src={User}
              alt="User Profile"
            />
            <span className="hidden font-medium text-[#36485C] lg:block">
              Sign in
            </span>
          </div>
        </Link>
      </div>
    );
  return (
    user && (
      <div className="flex items-center">
        <div className="flex flex-col text-right">
          <h2>Hi, {user.given_name}!</h2>
          <Link
            className="bg-[#4328EB] text-center w-33 py-1 px-2 ml-3 mt-1 text-white rounded-[4px]"
            href="/api/auth/logout"
          >
            <button className="hover:text-gray-500">Logout</button>
          </Link>
        </div>
      </div>
    )
  );
}
