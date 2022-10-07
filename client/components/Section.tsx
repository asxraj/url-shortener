import React from "react";

const Section = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <div className={`flex h-full w-8/12 mx-auto items-center ${className}`}>
      {children}
    </div>
  );
};

export default Section;
