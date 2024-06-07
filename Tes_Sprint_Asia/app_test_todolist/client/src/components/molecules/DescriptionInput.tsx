"use client";
import React from "react";
import Input from "../atoms/Input";
import Label from "../atoms/Label";

interface DescriptionInput {
  description: string;
  setDescription: Function;
}
const DescriptionInput: React.FC<DescriptionInput> = ({
  description,
  setDescription,
}) => {
  return (
    <div className="flex-column mb-8">
      <Label value="Task Name" />
      <Input
        value={description}
        type="text"
        onChange={(e) => setDescription(e.target.value)}
      />
    </div>
  );
};

export default DescriptionInput;
