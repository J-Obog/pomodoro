import React from "react";
import TaskList from "./components/TaskList"; 

function App() {
  return (
    <div className="timer-app flex m-3">
      <div className="w-4/5">
        {/* Timer */}
      </div>
      <div className="w-1/5">
        <TaskList/>
      </div>
    </div>
  );
}

export default App;
