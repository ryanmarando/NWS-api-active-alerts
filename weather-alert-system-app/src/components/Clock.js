import React, { useState, useEffect } from "react";

const Clock = () => {
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const intervalId = setInterval(() => {
      setTime(new Date());
    }, 1000);

    // Clear the interval when the component unmounts or before re-rendering
    return () => clearInterval(intervalId);
  }, []); // Empty dependency array ensures this effect runs only once

  return (
    <div>
      <p>{time.toLocaleTimeString()}</p>
    </div>
  );
};

export default Clock;
