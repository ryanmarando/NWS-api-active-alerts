import Link from "next/link";
import { useState } from "react";

export function MobileNav() {
  const [isOpen, setIsOpen] = useState(false);

  const toggleNav = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className="lg:hidden">
      <button
        onClick={toggleNav}
        className="block text-gray-500 hover:text-white focus:text-white focus:outline-none"
      >
        <svg
          className="h-6 w-6 fill-current text-right"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          {isOpen ? (
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M4.5 7C4.22386 7 4 7.22386 4 7.5C4 7.77614 4.22386 8 4.5 8H19.5C19.7761 8 20 7.77614 20 7.5C20 7.22386 19.7761 7 19.5 7H4.5ZM4.5 12C4.22386 12 4 12.2239 4 12.5C4 12.7761 4.22386 13 4.5 13H19.5C19.7761 13 20 12.7761 20 12.5C20 12.2239 19.7761 12 19.5 12H4.5ZM4 17.5C4 17.2239 4.22386 17 4.5 17H19.5C19.7761 17 20 17.2239 20 17.5C20 17.7761 19.7761 18 19.5 18H4.5C4.22386 18 4 17.7761 4 17.5Z"
            />
          ) : (
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M4.5 7C4.22386 7 4 7.22386 4 7.5C4 7.77614 4.22386 8 4.5 8H19.5C19.7761 8 20 7.77614 20 7.5C20 7.22386 19.7761 7 19.5 7H4.5ZM4.5 12C4.22386 12 4 12.2239 4 12.5C4 12.7761 4.22386 13 4.5 13H19.5C19.7761 13 20 12.7761 20 12.5C20 12.2239 19.7761 12 19.5 12H4.5ZM4 17.5C4 17.2239 4.22386 17 4.5 17H19.5C19.7761 17 20 17.2239 20 17.5C20 17.7761 19.7761 18 19.5 18H4.5C4.22386 18 4 17.7761 4 17.5Z"
            />
          )}
        </svg>
      </button>

      <div
        className={`${
          isOpen ? "block" : "hidden"
        } lg:hidden bg-[#4328EB] rounded-md`}
      >
        <div className="px-2 pt-2 pb-3">
          <Link
            className="block px-3 py-2 rounded-md text-base font-medium text-white hover:text-white hover:bg-gray-700"
            href="/"
          >
            Home
          </Link>
          <Link
            className="block px-3 py-2 rounded-md text-base font-medium text-white hover:text-white hover:bg-gray-700"
            href="/alert-system"
          >
            Weather Alert System
          </Link>
          <Link
            className="block px-3 py-2 rounded-md text-base font-medium text-white hover:text-white hover:bg-gray-700"
            href="/"
          >
            Contact
          </Link>
        </div>
      </div>
    </div>
  );
}
