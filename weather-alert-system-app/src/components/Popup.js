import { useEffect } from "react";
import Link from "next/link";

export function Popup({ show, onClose, children }) {
  useEffect(() => {
    if (show) {
      document.body.style.overflow = "hidden"; // Prevent background scrolling
    } else {
      document.body.style.overflow = "auto";
    }
    return () => {
      document.body.style.overflow = "auto"; // Clean up on component unmount
    };
  }, [show]);

  if (!show) {
    return null;
  }

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50">
      <div
        className="absolute inset-0 bg-black bg-opacity-50"
        onClick={onClose}
      ></div>
      <div className="bg-white p-8 rounded-lg shadow-lg relative z-10 max-w-md mx-auto">
        <button
          className="absolute top-2 right-2 mr-2 text-gray-500 hover:text-gray-700"
          onClick={onClose}
        >
          &times;
        </button>
        <div className="text-center">
          <h2 className="text-2xl font-bold mb-4">
            Want to set it and forget it?
          </h2>
          <h2 className="text-xl font-bold mb-4">Upgrade to Pro</h2>
          <p className="text-gray-700 mb-6">
            Get access to automatic 15 second updates, choosing your own
            warnings, and saving data for easy access.
          </p>
          <Link href="/payment">
            <button className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-white my-[5px] mx-[15px] mt-[14px]">
              Upgrade Now
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}
