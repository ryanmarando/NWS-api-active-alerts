"use client";
import Link from "next/link";
import { useUser } from "@clerk/clerk-react";
import { useEffect } from "react";
import Confetti from "react-confetti";

export default function Success({ params }) {
  if (
    params.id !==
    "3Pb0xmpcABuhT8DpZZU7GMLCExF1is9nsfhb3zrWmGTwIGuAGB7X4ODZwMtRM8iSJ5Y1uPt1OUXAuJ0MqJ4LfOC0P4PW22zfy5HhX3eV5cmaQBTVXofyLg4lXCAaE7BIrQsk4tpsFyPlmWRAPWEK73txs6P5g3B4e654TIn0stKv8gzgzGRdfYcko3KlnUrnA0VtLi1c2vmAZhBTEoPBQUcmn3uWf1tohjnEX9AQyEbEhY3EAMx9L9XyIN3GlErwEHl5AQVZbID9yVWyZCzmohFru9kU4yZsc3iEMj7uaRobUyHpSJJrC6BKEWYOVYR7TsQXEtzsCOaKlwrHDbXwgPfxX3k4rQIXbkwjb0jAJqiIqg3vrX1jMcTALt2uuSk1WofpiJo9Mo9HjHmRHHkPxWLZRLrMVXZL47W5kgOQEifizWGkWbQDgBj36bllSuhHbM6jZhSSXRZw9qcbNslHn1uoOm8g2kf8s3HQzq2ocI2Yw1EncUT5qz1SIxEBoC0qNBLCXEKcnDwrbbUxxRafsM9qGwKs90Uid3CUotY23AMRwxXgmIKejUDuZJ9aRCu911PUDd63cKSXvhIE3VSnir0hLDNoRdCFOw7qtyTdGitd0pqQvz5CHTTAljmaexQUBEQqEQZCZt3AvQkda4yCImNRZn9l3vInjgnIGnLDj5edq5yfY1aaK6KhP06L9anqNwIkM0Jvkw6OM9gUU3pYC97e9urdYgtJrIYStTJEB4DGH7t5132kqfFnn67aE6nhWsI0gT2tZZN53DTZYSa3wo4KT"
  ) {
    return (
      <div>
        <h1>Unauthorized</h1>
      </div>
    );
  }
  const { user } = useUser();
  useEffect(() => {
    activateSubscription();
  }, [user?.id]);

  const activateSubscription = async () => {
    try {
      const response = await fetch("/api/updatePrivateMetadata", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ userId: user?.id }),
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
    <div
      onClick={activateSubscription}
      className="flex flex-col items-center justify-center min-h-screen py-2"
    >
      <Confetti
        width={window.innerWidth}
        height={window.innerHeight}
        recycle={false}
        numberOfPieces={1000}
        gravity={0.5}
        friction={0.99}
        initialVelocityX={5}
        initialVelocityY={-5}
        colors={["#ff0000", "#00ff00", "#0000ff"]}
        className="absolute top-0 left-0"
      />
      <h1 className="text-3xl font-bold mb-4">Payment Successful!</h1>
      <p>Thank you for your purchase.</p>
      <Link href="/alert-system">
        <button
          className="bg-[#4328EB] hover:text-gray-500 w-33 py-1 px-2 rounded-[8px] text-[white] my-[5px] mx-[15px] mt-[14px]"
          onClick={activateSubscription}
        >
          Try it out now!
        </button>
      </Link>
    </div>
  );
}
