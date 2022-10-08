import React, { useEffect, useState } from "react";

import Head from "next/head";
import Navbar from "../components/Navbar";
import Input from "../components//Input";
import logo from "../public/assets/google.svg";
import Image from "next/image";

import Section from "../components/Section";
import { FormErrors } from "../utils/types";
import Link from "next/link";
import { AiOutlineLink } from "react-icons/ai";
import Footer from "../components/Footer";

export default function Login() {
  const [errors, setErrors] = useState<FormErrors>();

  const handleSubmit = () => {};

  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-r from-sky-400 to-blue-500">
      <Head>
        <title>SHORTURL | Login</title>
        <meta name="description" content="SHORTUTL project made by asxraj" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Section className="min-h-screen justify-center ">
        <div className="flex flex-col mt-14 gap-7">
          <Link href="/" className="">
            <a href="/">
              <div className="flex items-center gap-2 justify-center ">
                <AiOutlineLink className="text-5xl text-gray-600" />
                <span className="text-4xl text-gray-600 font-semibold font-bungee tracking-wider text-center">
                  shortURL
                </span>
              </div>
            </a>
          </Link>
          <div className="flex justify-center items-center rounded-xl bg-white">
            <form onSubmit={handleSubmit} className="flex flex-col gap-4 p-5">
              <h1 className="text-center text-3xl font-epilogue tracking-wider font-bold">
                Reset password
              </h1>
              <p className="text-center text-xs text-gray-700 flex-wrap">
                <span>Enter your account's email address and we will</span>
                <br />
                <span>send you a password reset link.</span>
              </p>
              <Input
                name="email"
                title="Email"
                type="email"
                placeholder="gavin@hooli.com"
                hasError={errors?.email ? true : false}
                errorMsg={errors?.email}
              />

              <button
                type="submit"
                className="mt-10 text-md font-medium text-white bg-blue-600 py-3 rounded-lg select-none transition-all hover:bg-blue-800"
              >
                Send reset Email
              </button>
              <p className="text-center font-medium text-gray-700">
                <span className="text-blue-400">
                  <a href="/signup">Return to Login</a>
                </span>
              </p>
            </form>
          </div>
        </div>
      </Section>
      <Footer />
    </div>
  );
}
