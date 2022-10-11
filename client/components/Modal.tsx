import React from "react";
import ReactDOM from "react-dom";
import { IoMdClose } from "react-icons/io";

const Modal = ({
  open,
  children,
  onClose,
}: {
  open: boolean;
  children: JSX.Element;
  onClose: React.MouseEventHandler<any> | undefined;
}) => {
  if (!open) return null;
  return ReactDOM.createPortal(
    <>
      <div
        onClick={onClose}
        className="fixed top-0 left-0 right-0 bottom-0 bg-black opacity-70 z-10"
      />
      <div className="fixed w-10/12 sm:w-2/4 lg:w-1/3 2xl:w-1/4 top-1/3 left-2/4 -translate-x-1/2 -translate-y-1/2 px-8 py-8 z-10  bg-gray-100  rounded-md">
        <IoMdClose
          onClick={onClose}
          className="w-7 h-7 absolute right-0 top-0 m-3 cursor-pointer text-red-600 hover:text-red-700"
        />
        {children}
      </div>
    </>,
    document.body
  );
};

export default Modal;
