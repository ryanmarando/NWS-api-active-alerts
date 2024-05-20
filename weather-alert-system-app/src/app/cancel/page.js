"use client";
import { useState } from "react";
import { useUser } from "@clerk/clerk-react";
import Link from "next/link";

const CancelSubscription = () => {
  const [isCanceled, setIsCanceled] = useState(false);
  const { user } = useUser();

  const handleCancel = async () => {
    console.log("Here");
    // Your cancellation logic goes here, such as sending a request to your backend
    // and updating the user's subscription status
    /*
    try {
      const response = await fetch("/api/updatePrivateMetadata", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          userId: user?.id,
          metadata: { subscription: false },
        }),
      });

      if (response.ok) {
        console.log("Subscription deactivated!");
      } else {
        const errorData = await response.json();
      }
    } catch (error) {
      console.error("Error updating subscription:", error);
    } */
    try {
      const response = await fetch("/api/cancel-subscription", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          subscriptionId: "sub_1PIXLkBhEzKZHebbHK2eb8aJ",
        }),
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      const data = await response.json();
      console.log("Subscription canceled:", data);
      // Optionally, you can update the UI to reflect the cancellation
    } catch (error) {
      console.error("Failed to cancel subscription:", error);
    }
    setIsCanceled(true);
  };

  return (
    <div className="min-h-screen flex items-center justify-center  py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Cancel Subscription
          </h2>
        </div>
        <div>
          {isCanceled ? (
            <div className="rounded-md bg-green-100 p-4">
              <div className="flex">
                <div className="flex-shrink-0">
                  <svg
                    className="h-5 w-5 text-green-400"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      fillRule="evenodd"
                      d="M9 0C4.03 0 0 4.03 0 9c0 4.97 4.03 9 9 9s9-4.03 9-9c0-4.97-4.03-9-9-9zM8 14.586l-3.293-3.293a1 1 0 011.414-1.414L8 11.172l4.879-4.879a1 1 0 111.414 1.414L8.707 14.586l-.707.707-.707-.707z"
                      clipRule="evenodd"
                    />
                  </svg>
                </div>
                <div className="ml-3">
                  <p className="text-sm text-green-800">
                    Your subscription has been successfully canceled.
                  </p>
                </div>
              </div>
              <Link
                href="/"
                className="flex w-full items-center justify-center"
              >
                <button className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[15px] mt-[14px]">
                  Return Home
                </button>
              </Link>
            </div>
          ) : (
            <div className="rounded-md bg-white shadow-md p-6">
              <p className="text-sm text-gray-700">
                Are you sure you want to cancel your subscription? This action
                cannot be undone.
              </p>
              <button
                onClick={handleCancel}
                className="mt-6 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent text-base font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
              >
                Cancel Subscription
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default CancelSubscription;
