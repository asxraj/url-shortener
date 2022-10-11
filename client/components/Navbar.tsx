import Link from "next/link";
import React, { useContext, useState, useRef } from "react";
import { AiOutlineCaretDown, AiOutlineLink } from "react-icons/ai";
import { FiLogOut, FiSettings } from "react-icons/fi";
import { MdAutoGraph } from "react-icons/md";
import { AuthContext } from "../context/AuthContext";
import ModalConfirm from "./ModalConfirm";
import Section from "./Section";

const Navbar = () => {
  const [toggle, setToggle] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const bar = useRef<any>();
  const { jwt, user, logout } = useContext(AuthContext);

  const closeOpenMenus = (e: Event) => {
    if (bar.current && toggle && !bar.current.contains(e.target)) {
      setToggle(false);
    }
  };

  if (typeof window !== "undefined") {
    document.addEventListener("mousedown", closeOpenMenus);
  }
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

        {!jwt ? (
          <>
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
          </>
        ) : (
          <>
            <div className="flex gap-4 items-center relative">
              <button className="text-sm font-medium text-white bg-blue-600 px-3 py-2 rounded-lg select-none transition-all hover:bg-blue-800">
                Upgrade
              </button>
              <div
                onClick={() => {
                  if (toggle) {
                  } else {
                    setToggle(true);
                  }
                }}
                className="flex items-center gap-2 transition-all hover:bg-gray-300 rounded-md p-1 px-3 cursor-pointer"
              >
                <p className="flex items-center justify-center p-2 rounded-full bg-white select-none w-10 h-10 font-bold">
                  {user.first_name[0] + user.last_name[0]}
                </p>
                <p className="md:flex hidden font-semibold tracking-wide">
                  {user.first_name + " " + user.last_name}
                </p>
                <AiOutlineCaretDown />
              </div>
              {toggle && (
                <div
                  ref={bar}
                  className="font-semibold bg-gray-100 flex flex-col absolute p-3 py-5 z-50 rounded-md bg-opacity-80 top-14 right-0  min-w-[170px] transition-all"
                >
                  <div className="flex flex-col gap-2">
                    <div className="flex justce items-center gap-2 cursor-pointer list-none hover:bg-gray-400 px-3 py-2 rounded-md transition-all">
                      <MdAutoGraph className="w-5 h-5" />
                      Dashboard
                    </div>
                    <div className="flex justce items-center gap-2 cursor-pointer list-none hover:bg-gray-400 px-3 py-2 rounded-md transition-all">
                      <FiSettings className="w-5 h-5" />
                      Settings
                    </div>
                    <div
                      className="flex justce items-center gap-2 cursor-pointer list-none hover:bg-gray-400 px-3 py-2 rounded-md transition-all"
                      onClick={() => {
                        setIsOpen(true);
                        setToggle(false);
                      }}
                    >
                      <FiLogOut className="w-5 h-5" />
                      Logout
                    </div>
                  </div>
                </div>
              )}
            </div>
          </>
        )}
      </Section>
      <ModalConfirm
        question="Are you sure you want to logout?"
        open={isOpen}
        onClose={() => setIsOpen(false)}
        confirm="Logout"
        deny="Cancel"
        action={() => logout()}
      />
    </nav>
  );
};

export default Navbar;
