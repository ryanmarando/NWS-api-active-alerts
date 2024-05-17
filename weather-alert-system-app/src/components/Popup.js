import { useEffect } from "react";
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
    <div className="absolute w-[50%] h-[50%] inset-0 m-auto items-center justify-center z-50">
      <div className="flex w-full h-full items-center justify-center border-2 border-black rounded-md bg-[#DDDDDD]">
        <p className="w-full text-center">Want the weather alert system?</p>
      </div>
    </div>
  );
}
