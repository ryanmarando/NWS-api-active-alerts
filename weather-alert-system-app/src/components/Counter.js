import { useState, useEffect } from "react";

const CountdownTimer = ({ initialTime }) => {
  const [timeLeft, setTimeLeft] = useState(initialTime);

  useEffect(() => {
    if (timeLeft === 0) return;

    const timer = setInterval(() => {
      setTimeLeft((prevTime) => prevTime - 1);
    }, 1000);

    return () => clearInterval(timer);
  }, [timeLeft]);

  // Convert remaining time to hours, minutes, and seconds
  const hours = Math.floor(timeLeft / 3600);
  const minutes = Math.floor((timeLeft % 3600) / 60);
  const seconds = timeLeft % 60;

  return (
    <div>
      {timeLeft === 0 ? (
        <p>Countdown Over!</p>
      ) : (
        <p>
          Refresh in: {seconds.toString().padStart(2, "0")} s<br></br>
          Click To Stop
        </p>
      )}
    </div>
  );
};

export default CountdownTimer;
