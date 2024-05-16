"use client";
import { useState } from "react";
import { IoIosArrowBack } from "react-icons/io";

export function AlertForm() {
  // Initialize state for checkboxes
  const [warningList, setWarningList] = useState([]);
  const [checkboxes, setCheckboxes] = useState({
    option1: false,
    option2: false,
    option3: false,
  });

  // Handle checkbox change
  const handleCheckboxChange = (event) => {
    const option = event.target.name;
    console.log(option);
    if (!checkboxes.option)
      setWarningList([...warningList, event.target.value]);
    setCheckboxes(!checkboxes.option);
  };
  //const { name, checked } = event.target;
  //setCheckboxes((prevState) => ({
  //  ...prevState,
  //  [name]: checked,
  //}));
  //};

  // Handle form submission
  const handleSubmit = (event) => {
    event.preventDefault();
    console.log("Selected checkboxes:", checkboxes);
    console.log(warningList);
  };

  return (
    <div className="w-full items-center justify-center">
      <form onSubmit={handleSubmit} className="p-4 w-full mx-auto">
        <div className="grid grid-cols-6 grid-row-10">
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                value="Tsunami Warning"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Tsunami Warning
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Tornado Warning
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 3
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option2"
                checked={checkboxes.option2}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Tornado Warning
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option1"
                checked={checkboxes.option1}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 1
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option2"
                checked={checkboxes.option2}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 2
            </label>
          </div>
          <div className="">
            <label className="block">
              <input
                type="checkbox"
                name="option3"
                checked={checkboxes.option3}
                onChange={handleCheckboxChange}
                className="mr-2"
              />
              Option 3
            </label>
          </div>
        </div>
      </form>
      <div className="flex items-center justify-center mb-4">
        <button
          type="submit"
          onClick={handleSubmit}
          className="bg-[#4328EB] flex hover:text-gray-500 py-1 px-2 w-50 mr-0 mt-0 w-50 rounded-[8px] text-white "
        >
          Submit
        </button>
      </div>
    </div>
  );
}
