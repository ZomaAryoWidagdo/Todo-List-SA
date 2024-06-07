"use client";
import React, { useState } from "react";
import { useDispatch } from "react-redux";
import Button from "../atoms/Button";
import DescriptionInput from "../molecules/DescriptionInput";
import DeadlineInput from "../molecules/DeadlineInput";
import { addTaskAsync } from "@/redux/thunks";
import SubTaskInput from "../molecules/SubTaskInput";
import { dateRequestFormatter } from "@/utils/dateFormatter";
import { AppDispatch } from "@/redux/store";

const AddTaskModal: React.FC = () => {
  const [description, setDescription] = useState("");
  const [date, setDate] = useState("");
  const [totalSubTask, setTotalSubtask] = useState(0);
  const [subTasks, setSubTasks] = useState<object[]>([]);
  const dispatch = useDispatch<AppDispatch>();

  const handleAddTask = () => {
    if (description.trim()) {
      const deadline = dateRequestFormatter(date);

      dispatch(
        addTaskAsync({
          description,
          deadline,
          subTasks,
        })
      );

      setDescription("");
      setDate("");
      setSubTasks([]);
      setTotalSubtask(0);
    }
  };

  const handleAddSubTask = () => {
    setTotalSubtask(totalSubTask + 1);
  };

  const handleRemoveSubTask = () => {
    if (totalSubTask == subTasks.length) {
      const newSubTask = [...subTasks];

      newSubTask.pop();

      setSubTasks(newSubTask);
    }

    setTotalSubtask(totalSubTask - 1);
  };

  const handleSetSubTask = (subTaskDescription: string, index: number) => {
    const newSubTask = [...subTasks];

    newSubTask[index] = { description: subTaskDescription };

    setSubTasks(newSubTask);
  };

  return (
    <div
      className="flex-column add-task-modal"
      style={{ background: "white", padding: "8px", borderRadius: "8px" }}
    >
      <DescriptionInput
        description={description}
        setDescription={setDescription}
      />
      {Array.from({ length: totalSubTask }, (_, i) => (
        <SubTaskInput key={i} index={i} handleSetSubTask={handleSetSubTask} />
      ))}
      <DeadlineInput date={date} setDate={setDate} />
      <Button onClick={handleAddSubTask} classname="my-8">
        Add Sub Task
      </Button>
      {totalSubTask ? (
        <Button onClick={handleRemoveSubTask} classname="my-8">
          Remove Sub Task
        </Button>
      ) : (
        <></>
      )}
      <Button onClick={handleAddTask} classname="my-8">
        Save Task
      </Button>
    </div>
  );
};

export default AddTaskModal;
