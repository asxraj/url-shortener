import React from "react";

const Input = ({
  name,
  title,
  type,
  placeholder,
  errorMsg,
  hasError,
}: {
  name: string;
  title: string;
  type: string;
  placeholder?: string;
  errorMsg: string | undefined;
  hasError: boolean | undefined;
}) => {
  return (
    <div className="flex flex-col">
      <label
        htmlFor={name}
        className="font-semibold font-epilogue tracking-wider"
      >
        {title.toUpperCase()}
      </label>
      <label className="text-xs text-red-400 opacity-75">{errorMsg}</label>
      <input
        type={type}
        name={name}
        className={`p-2 lg:p-3 border-[1px] focus:border-[1px] focus:border-blue-200  bg-darkgray text-sm w-[400px] lg:w-[425px] rounded-lg focus-within:bg-darkgray ${
          hasError ? "border-2 border-red-400 bg-red-100" : ""
        }`}
        placeholder={placeholder}
      />
    </div>
  );
};

export default Input;
