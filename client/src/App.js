import React from "react";
import TaskList from "./components/TaskList"; 

function App() {
  return (
    <div className="timer-app flex h-screen">
      <div className="w-4/5 m-3">
        {/* Timer */}
      </div>
      <div className="w-1/5 m-3">
        <TaskList/>
      </div>
    </div>
  );
}

export default App;
