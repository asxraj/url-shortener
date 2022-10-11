import React from "react";
import Modal from "./Modal";

const ModalConfirm = ({
  open,
  onClose,
  question,
  confirm,
  deny,
  action,
}: {
  open: boolean;
  onClose: any;
  question: string;
  confirm: string;
  deny: string;
  action: Function;
}) => {
  return (
    <Modal open={open} onClose={() => onClose(false)}>
      <>
        <div className="md:flex flex-col gap-4 hidden">
          <h1 className="text-gray-700 font-bold text-center">{question}</h1>
          <div className="flex my-5 mx-auto gap-8">
            <button
              className="px-10 py-3 bg-red-500 rounded-md font-semibold text-algo hover:bg-red-700"
              onClick={() => {
                action();
                onClose(false);
              }}
            >
              {confirm}
            </button>
            <button
              className="px-10 py-3 bg-blue-500 rounded-md font-semibold text-algo hover:bg-blue-700"
              onClick={() => onClose(false)}
            >
              {deny}
            </button>
          </div>
        </div>

        <div className="flex md:hidden flex-col ">
          <h5 className="text-gray-300 font-semibold text-center text-sm">
            {question}
          </h5>
          <div className="flex my-5 mx-auto gap-8">
            <button
              className="px-7 py-2 bg-red-500 rounded-md font-semibold text-algo hover:bg-red-700"
              onClick={() => {
                onClose(false);
              }}
            >
              {confirm}
            </button>
            <button
              className="px-7 py-2 rounded-md font-semibold bg-blue-500 text-algo hover:bg-blue-700 "
              onClick={() => onClose(false)}
            >
              {deny}
            </button>
          </div>
        </div>
      </>
    </Modal>
  );
};

export default ModalConfirm;
