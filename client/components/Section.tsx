import React from "react";

const Section = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <div
      className={`flex h-full w-11/12  md:w-10/12 lg:w-9/12 xl:w-8/12 mx-auto gap-8 ${className}`}
    >
      {children}
    </div>
  );
};

export default Section;
