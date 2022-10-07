import type { NextPage } from "next";
import Head from "next/head";
import Navbar from "../components/Navbar";
import Section from "../components/Section";
import { BsLink } from "react-icons/bs";

const Home: NextPage = () => {
  const submitHandler = (e: any) => {
    e.preventDefault();

    const data: FormData = new FormData(e.target);
    const payload = Object.fromEntries(data.entries());
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const request = {
      url: payload.url,
      short: "",
      expiry: 24,
    };

    console.log(request);

    fetch("http://localhost:4001/v1/shorten", {
      method: "POST",
      body: JSON.stringify(request),
      headers: headers,
    })
      .then((response) =>
        response.json().then((data) => {
          console.log(data);
        })
      )
      .catch((err) => console.log(err));

    console.log("form was submitted");
  };
  return (
    <div className="min-h-screen flex flex-col bg-blue-600">
      <Head>
        <title>URL</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Navbar />
      <Section>
        <div className="flex mt-10 justify-center items-center p-4">
          <div className=" bg-gray-100 p-5 rounded-xl py-8">
            <form onSubmit={submitHandler} className="flex flex-col gap-6">
              {/* URL */}
              <div className="flex flex-col gap-2">
                <label className="text-xl  font-bold">
                  Enter URL to Shorten
                </label>

                <div className="flex items-center w-[375px] rounded-lg bg-white p-4 gap-2">
                  <label htmlFor="url" className="">
                    <BsLink className="text-2xl font-medium text-gray-400" />
                  </label>

                  <input
                    className="flex-1 outline-none rounded-r-lg  text-gray-400 font-medium"
                    type="text"
                    name="url"
                    id="url"
                    autoFocus={true}
                  />
                </div>
              </div>

              {/* ALIAS */}
              <div className="flex flex-col gap-2">
                <label className="text-xl font-bold">Customize your link</label>

                <div className="flex flex-col w-[375px] rounded-lg bg-white  gap-2">
                  <div className="flex gap-2">
                    <div className="border-r-2 border-gray-300 p-4">
                      <label className="text-lg font-extralight">
                        shortenURL.com
                      </label>
                    </div>

                    <div className="flex items-center">
                      <input
                        className="outline-none  rounded-r-lg font-medium text-gray-400"
                        placeholder="alias"
                        type="text"
                        name="url"
                      />
                    </div>
                  </div>
                </div>
              </div>

              <button
                type="submit"
                className="px-5 py-3 rounded-md bg-blue-600 text-white font-semibold transition-all hover:bg-blue-700"
              >
                Shorten URL
              </button>
            </form>
          </div>
        </div>

        <div className=""> sdaasdsddas</div>
      </Section>
    </div>
  );
};

export default Home;
