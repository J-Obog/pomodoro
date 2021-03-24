var taskList = [];
var timerOn = false;
var currTask = -1;
const PROGRESS_COLORS = ["red", "green", "blue"];


function removeTask(num) {
    taskList.splice(num, 1);
    renderTaskList();
}

function renderTaskList() {
    $("#task-container").empty(); 
    taskList.forEach((task, idx) => {
        $("#task-container").append(`
        <div>
            <span>${idx}</span>
            <p>${task.name}</p>
            <button>x</button>
        </div>
        `)
    })
}


$("#add-task-btn").click(function(e) {
    taskList.push({
        name: $("#add-task-input").val(),
        status: 0
    })
    renderTaskList(); 
})



