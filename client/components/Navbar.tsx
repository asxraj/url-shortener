import Link from "next/link";
import React from "react";
import { AiOutlineLink } from "react-icons/ai";
import Section from "./Section";

const Navbar = () => {
  return (
    <div className="flex items-center w-full py-2 bg-[#F9FAFE]">
      <Section className="justify-between">
        <Link href="/" className="">
          <a href="/">
            <div className="flex items-center gap-2">
              <AiOutlineLink className="text-5xl text-gray-600" />
              <span className="text-5xl text-gray-600 font-semibold font-bungee tracking-wider ">
                URL
              </span>
            </div>
          </a>
        </Link>
        <div className="flex items-center gap-4">
          <a
            href="/login"
            className="text-md font-medium rounded-md text-[#24335A] select-none px-5 py-2 transition-all hover:bg-gray-200 hover:underline "
          >
            Log in
          </a>
          <a
            href="/signup"
            className="text-md font-medium text-white bg-blue-600 px-5 py-2 rounded-lg select-none transition-all hover:bg-blue-800 hover:underline"
          >
            Sign up
          </a>
        </div>
      </Section>
    </div>
  );
};

export default Navbar;
