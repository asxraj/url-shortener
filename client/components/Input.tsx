import React from "react";

const Input = ({
  name,
  title,
  type,
  minLength,
  placeholder,
  errorMsg,
  hasError,
}: {
  name: string;
  title: string;
  type: string;
  minLength?: number;
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
      <input
        type={type}
        name={name}
        minLength={minLength ? 8 : 0}
        className={`p-2 lg:p-3 border-[1px] bg-darkgray text-sm w-[325px] lg:w-[425px] rounded-lg focus-within:bg-darkgray ${
          hasError
            ? "border-2 border-red-400 bg-red-100 focus:border-red-400"
            : "focus:border-blue-200 focus:border-[1px]"
        }`}
        placeholder={placeholder}
      />
      <label className="text-xs text-red-600 opacity-75 ml-1 mt-1">
        {errorMsg}
      </label>
    </div>
  );
};

export default Input;
