import { GetServerSideProps } from "next";
import React, { useEffect } from "react";
import { useRouter } from "next/router";

interface Error {
  error: string;
}
const ShortURL = ({ error }: Error) => {
  // Add some errorr page when a link that doesn't exist is clicked
  const router = useRouter();

  const shortURL = router.query["short"];

  return <div>{error}</div>;
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const shortURL = context.query["short"];

  const res = await fetch(`http://localhost:4001/${shortURL}`);
  const data = await res.json();

  if (!data.error) {
    return {
      redirect: {
        destination: "https://" + data.to,
        permanent: true,
      },
    };
  } else {
    return {
      props: {
        error: data.error,
      },
    };
  }

  return {
    props: {},
  };
};

export default ShortURL;
