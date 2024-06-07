"use client";
import React, { useState } from "react";
import Input from "../atoms/Input";
import Label from "../atoms/Label";

interface SubTaskInputProps {
  index: number;
  handleSetSubTask: Function;
}

const SubTaskInput: React.FC<SubTaskInputProps> = ({
  index,
  handleSetSubTask,
}) => {
  const [subTaskDescription, setSubTaskDescription] = useState("");

  const handleSetSubTaskDescription = (e: any) => {
    setSubTaskDescription(e.target.value);
    handleSetSubTask(e.target.value, index);
  };

  return (
    <div className="flex-column mb-8">
      <Label value="Sub Task Name" />
      <Input
        value={subTaskDescription}
        type="text"
        onChange={handleSetSubTaskDescription}
      />
    </div>
  );
};

export default SubTaskInput;
