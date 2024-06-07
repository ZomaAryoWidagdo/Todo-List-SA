import React, { useState } from "react";
import { useDispatch } from "react-redux";
import SubTaskItem from "./SubTaskItem";
import Button from "../atoms/Button";
import { SubTask } from "@/models/subTask";
import { addTaskAsync, deleteTaskAsync, updateTaskAsync } from "@/redux/thunks";
import { AppDispatch } from "@/redux/store";

interface TaskItemProps {
  id: number;
  description: string;
  status: string;
  deadline: string;
  subTasks: SubTask[];
}

const TaskItem: React.FC<TaskItemProps> = ({
  id,
  description,
  status,
  deadline = "",
  subTasks = [],
}) => {
  const dispatch = useDispatch<AppDispatch>();
  const [subTodoText, setSubTodoText] = useState("");

  const handleEditTask = () => {
    if (subTodoText.trim()) {
      dispatch(
        addTaskAsync({
          subTasks,
          description,
          deadline,
        })
      );
      setSubTodoText("");
    }
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        marginBottom: "16px",
        marginTop: "16px",
        height: "100%",
        width: "250px",
        border: "black solid 1px",
        padding: "12px",
        borderRadius: "8px",
        backgroundColor: "floralwhite",
      }}
    >
      <div>
        <input
          type="checkbox"
          checked={status == "finished"}
          onChange={() =>
            dispatch(updateTaskAsync({ taskId: id, status: "finished" }))
          }
        />
        <span
          style={{
            textDecoration: status == "finished" ? "line-through" : "none",
          }}
        >
          {description}
        </span>
      </div>
      <div>
        {subTasks.map((subTask) => (
          <SubTaskItem
            key={subTask.ID}
            taskId={subTask.TaskID}
            id={subTask.ID}
            description={subTask.Description}
            status={subTask.Status}
          />
        ))}
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "4px",
          marginTop: "8px",
        }}
      >
        <Button onClick={() => dispatch(deleteTaskAsync(id))}>Delete</Button>
        <Button onClick={handleEditTask}>Edit Task</Button>
      </div>
    </div>
  );
};

export default TaskItem;
