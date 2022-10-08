import React from "react";

const Footer = () => {
  return (
    <div className="relative lg:absolute left-0 bottom-0 w-full flex text-center pb-2 mt-10">
      <div className="w-full text-sm">
        &copy; shortURL Insights, Inc •{" "}
        <span className="text-xs hover:underline cursor-pointer transition-all">
          Terms of Service
        </span>
      </div>
    </div>
  );
};

export default Footer;
