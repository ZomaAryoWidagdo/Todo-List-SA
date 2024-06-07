import React from "react";
import { useDispatch } from "react-redux";
import Button from "../atoms/Button";
import { deleteSubTask } from "@/redux/slicer";
import { updateSubTaskAsync } from "@/redux/thunks";
import { AppDispatch } from "@/redux/store";

interface SubTaskItemProps {
  id: number;
  taskId: number;
  description: string;
  status: string;
}

const SubTaskItem: React.FC<SubTaskItemProps> = ({
  taskId,
  id,
  description,
  status,
}) => {
  const dispatch = useDispatch<AppDispatch>();

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        marginBottom: "4px",
        marginTop: "4px",
      }}
    >
      <div style={{ width: "10%" }}>
        <input
          type="checkbox"
          checked={status == "finished"}
          onChange={() =>
            dispatch(updateSubTaskAsync({ subTaskId: id, status: "finished" }))
          }
        />
      </div>
      <p
        style={{
          textDecoration: status == "finished" ? "line-through" : "none",
          width: "80%",
        }}
      >
        {description}
      </p>
      <div style={{ width: "10%" }}>
        <Button
          onClick={() => dispatch(deleteSubTask({ taskId, subTaskId: id }))}
        >
          X
        </Button>
      </div>
    </div>
  );
};

export default SubTaskItem;
