"use client";
import { useEffect, useState } from "react";
import { loadStripe } from "@stripe/stripe-js";
import { useUser } from "@clerk/clerk-react";

export default function Payment() {
  const [loading, setLoading] = useState(true);
  const { isLoaded, isSignedIn, user } = useUser();

  if (!isSignedIn) {
    window.location.href = "/sign-in";
  }

  const testKEY =
    "pk_test_51PHaB5BhEzKZHebbhsYxL7UnvOXHH0IvdnshQweYT05jMRv4YcnyzF3M46qSO2QZUOtuzYols8dq5Y5GKyTkOuap00Ih42UG36";
  const liveKEY =
    "pk_live_51PHaB5BhEzKZHebbyuCxZ2PJRPkoAqZiF2AJzo37vhSYWhSRQcuPxDfZnAJeIQxskT9q4pw5sYblVBfCqhyWndd200vQIXPMFI";

  useEffect(() => {
    const initiateCheckout = async () => {
      const response = await fetch("/api/create-checkout-session", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
      });

      const { id } = await response.json();
      const stripe = await loadStripe(testKEY);

      const { error } = await stripe.redirectToCheckout({ sessionId: id });

      if (error) {
        console.error("Error redirecting to checkout:", error);
        setLoading(false);
      }
    };

    initiateCheckout();
  }, []);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      {loading && (
        <h1 className="text-3xl font-bold mb-4">Redirecting to payment...</h1>
      )}
    </div>
  );
}
