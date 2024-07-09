"use client";
import React from "react";
import { useState } from "react";
import { MdEmail } from "react-icons/md";
import { FaPhoneAlt } from "react-icons/fa";
import { FaMessage } from "react-icons/fa6";
import { Navbar } from "../../components/Navbar";
import { Footer } from "../../components/Footer";

export default function ContactPage() {
  const [userEmail, setUserEmail] = useState("");
  const [userName, setUserName] = useState("");
  const [userCommentary, setUserCommentary] = useState("");

  const SubmitButton = async () => {
    if (userEmail == "" || userName == "" || userCommentary == "")
      return alert("Please enter all fields to send a message.");
    const inputs = {
      userEmail,
      userName,
      userCommentary,
    };
    try {
      const response = await fetch("https://formspree.io/f/mkndeywo", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          data: inputs,
        }),
      });
      if (response.ok) {
        alert("Message sent successfully!");
        setUserEmail("");
        setUserName("");
        setUserCommentary("");
      } else {
        const errorData = await response.json();
        alert(`Failed to send message: ${errorData.error}`);
      }
    } catch {
      const errorData = await response.json();
      alert(`Failed to send message: ${errorData.error}`);
    }
  };

  const ContactInfo = [
    {
      title: "Support",
      email: "marandoryan@gmail.com",
    },
  ];
  return (
    <>
      <Navbar />
      <div className="relative mt-20">
        <div className="absolute top-1/2 left-1/2 mx-auto flex w-11/12 max-w-[1200px] -translate-x-1/2 -translate-y-1/2 flex-col text-center text-white lg:ml-5">
          <h1 className="text-4xl font-bold sm:text-5xl">Contact me</h1>
          <p className="mx-auto pt-3 text-s mb-12 lg:w-3/5 lg:pt-5 lg:text-base text-black">
            Have a question? Need assistance? Want something specifically done
            for you?
          </p>
        </div>
      </div>
      <section className="w-full flex-grow">
        <section className="mx-auto w-[95%] lg:w-[50%] items-center justify-center my-6  px-5  pt-16">
          {ContactInfo.map((info, index) => (
            <div key={index}>
              <div className="border py-5 shadow-md bg-[#DDDDDD] rounded-[8px]">
                <div className="flex justify-between px-4 pb-5">
                  <p className="text-xl font-bold">{info.title}</p>
                </div>
                <div className="flex flex-col px-4">
                  <a
                    className="flex items-center"
                    href={`mailto:${info.email}`}
                  >
                    <MdEmail className="mr-3 h-4 w-4 text-violet-900" />
                    {info.email}
                  </a>
                </div>
              </div>
            </div>
          ))}
        </section>
        <section className="mx-auto my-5 text-center">
          <h2 className="text-3xl font-bold">Have another question?</h2>
          <p>Complete the form below</p>
        </section>
        <form className="mx-auto my-5 max-w-[600px] px-5" action="">
          <div className="mx-auto">
            <div className="my-3 flex w-full gap-2">
              <input
                required
                className="w-1/2 border px-4 py-2"
                type="email"
                placeholder="E-mail"
                value={userEmail}
                onChange={(e) => setUserEmail(e.target.value)}
              />
              <input
                required
                className="w-1/2 border px-4 py-2"
                type="text"
                placeholder="Full Name"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
              />
            </div>
          </div>
          <textarea
            required
            className="w-full border px-4 py-2"
            placeholder="Tell me what's going on..."
            name=""
            id=""
            defaultValue={""}
            value={userCommentary}
            onChange={(e) => setUserCommentary(e.target.value)}
          />
        </form>
        <div className="flex items-center justify-center w-full">
          <button onClick={SubmitButton} className=" bg-amber-400 px-4 py-2">
            Send Message
          </button>
        </div>
      </section>
      <Footer />
    </>
  );
}
