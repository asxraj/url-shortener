import Link from "next/link";
import React from "react";
import { AiOutlineLink } from "react-icons/ai";
import Section from "./Section";

const Navbar = () => {
  return (
    <nav className="flex items-center w-full py-2">
      <Section className="justify-between">
        <div className="flex items-center">
          <Link href="/" className="">
            <a href="/">
              <div className="flex items-center gap-2">
                <AiOutlineLink className="text-2xl md:text-5xl text-gray-600" />
                <span className="text-2xl md:text-4xl text-gray-600 font-semibold font-bungee tracking-wider ">
                  shortURL
                </span>
              </div>
            </a>
          </Link>
        </div>
        <div className="flex items-center justify-between">
          <a
            href="/pricing"
            className="text-md font-medium rounded-md text-[#24335A] select-none px-5 py-2 transition-all hover:underline hover:decoration-white "
          >
            Products
          </a>
          <a
            href="/pricing"
            className="text-md font-medium rounded-md text-[#24335A] select-none px-5 py-2 transition-all hover:underline hover:decoration-white "
          >
            Pricing
          </a>
          <a
            href="/pricing"
            className="text-md font-medium rounded-md text-[#24335A] select-none px-5 py-2 transition-all hover:underline hover:decoration-white "
          >
            Pricing
          </a>
        </div>
        <div className="flex items-center gap-4">
          <a
            href="/login"
            className="text-md font-medium rounded-md text-[#24335A] select-none px-5 py-2 transition-all hover:bg-gray-200 "
          >
            Log in
          </a>
          <a
            href="/signup"
            className="text-md font-medium text-white bg-blue-600 px-5 py-2 rounded-lg select-none transition-all hover:bg-blue-800"
          >
            Sign up
          </a>
        </div>
      </Section>
    </nav>
  );
};

export default Navbar;
