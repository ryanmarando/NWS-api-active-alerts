"use client";
import Link from "next/link";
import { useUser } from "@clerk/clerk-react";

export default function Success() {
  const { user } = useUser();
  const activateSubscription = async () => {
    try {
      const response = await fetch("/api/updateUserMetadata", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          userId: user?.id,
          metadata: {
            subscription: true,
          },
        }),
      });

      if (response.ok) {
        console.log("Subscription activated!");
      } else {
        const errorData = await response.json();
      }
    } catch (error) {
      console.error("Error updating metadata:", error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <h1 className="text-3xl font-bold mb-4">Payment Successful!</h1>
      <p>Thank you for your purchase.</p>
      <Link href="/">
        <button onClick={activateSubscription}>Try it out now!</button>
      </Link>
    </div>
  );
}
