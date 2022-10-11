import React, { useEffect, useState } from "react";

import Head from "next/head";
import Input from "../components//Input";
import logo from "../public/assets/google.svg";
import Image from "next/image";

import Section from "../components/Section";
import { FormErrors } from "../utils/types";
import Link from "next/link";
import { AiOutlineLink } from "react-icons/ai";
import Footer from "../components/Footer";
import { useRouter } from "next/router";

export default function Login() {
  const [errors, setErrors] = useState<FormErrors>();

  const router = useRouter();

  const handleSubmit = async (e: any) => {
    e.preventDefault();

    const data: FormData = new FormData(e.target);
    const payload = Object.fromEntries(data.entries());
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    console.log(payload);

    await fetch("http://localhost:4001/v1/users/register", {
      method: "POST",
      body: JSON.stringify(payload),
      headers: headers,
    })
      .then((response) =>
        response.json().then((data) => {
          if (data.error) {
            setErrors(data.error);
            console.log(data);
          } else {
            router.push("/login");
          }
        })
      )
      .catch((err) => console.log(err));

    console.log("form was submitted");
  };

  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-r from-sky-400 to-blue-500">
      <Head>
        <title>SHORTURL | Signup</title>
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
              <h1 className="text-2xl lg:text-3xl font-epilogue tracking-wider font-bold mb-2">
                Sign up
              </h1>

              <button
                type="submit"
                className="flex items-center gap-2 justify-center text-md font-medium py-2 border-[1px] text-blue-600 border-blue-600 rounded-lg select-none transition-all hover:bg-blue-100"
              >
                <Image src={logo} alt="google.svg" height={20} width={20} />
                <span>Sign in with Google</span>
              </button>

              <div className="flex items-center gap-2 mt-2 mb-3">
                <div className="flex-1 bg-gray-500 h-[1px] w-20"></div>
                <div className="text-gray-500 text-xs font-epilogue">OR</div>
                <div className="flex-1 bg-gray-500 h-[1px] w-20"></div>
              </div>
              <Input
                name="first_name"
                title="first name"
                type="text"
                placeholder="Gavin"
                hasError={errors?.first_name ? true : false}
                errorMsg={errors?.first_name}
              />
              <Input
                name="last_name"
                title="last name"
                type="text"
                placeholder="Belson"
                hasError={errors?.last_name ? true : false}
                errorMsg={errors?.last_name}
              />
              <Input
                name="email"
                title="Email"
                type="email"
                placeholder="gavin@hooli.com"
                hasError={errors?.email ? true : false}
                errorMsg={errors?.email}
              />
              <Input
                name="password"
                title="Password"
                type="password"
                minLength={8}
                placeholder="••••••••"
                hasError={errors?.password ? true : false}
                errorMsg={errors?.password}
              />

              <button
                type="submit"
                className="text-md font-medium text-white bg-blue-600 py-3 rounded-lg select-none transition-all hover:bg-blue-800 mt-4"
              >
                Sign up
              </button>
              <p className="font-medium text-gray-700">
                Already have an account?{" "}
                <span className="text-blue-400">
                  <a href="/login">Log in</a>
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
