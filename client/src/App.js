import React from "react";
import TaskList from "./components/TaskList"; 
import Timer from "./components/Timer";
import Time from "./context/TimeContext"; 

function App() {
  return (
    <Time>
      <div className="timer-app flex h-screen">
        <div className="w-4/5 m-3 flex flex-col items-center">
          <Timer/>
        </div>
        <div className="w-1/5 m-3">
          <TaskList/>
        </div>
      </div>
    </Time>
  );
}

export default App;
