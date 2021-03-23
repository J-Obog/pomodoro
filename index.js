var taskList = []
const PROGRESS_COLORS = ["red", "green", "blue"]

function removeTask(taskNum) {
    console.log(taskNum); 
}

function renderTaskList() {
    taskList.forEach((task, idx)=> {
        let div = $('<div></div>') 
        .append($(`<p>${task.name}</p>`)
        )

        $("#task-container").append(div);
    })
}

$("#add-task-btn").click(function(e) {
    taskList.push({
        name: $("#add-task-input").val(),
        status: 0
    })
    renderTaskList(); 
})



