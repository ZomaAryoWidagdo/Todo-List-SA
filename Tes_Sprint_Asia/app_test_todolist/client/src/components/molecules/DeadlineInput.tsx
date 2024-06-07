"use client";
import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { addTask } from "@/redux/slicer";
import Input from "../atoms/Input";
import Label from "../atoms/Label";

interface DeadlineInput {
  date: string;
  setDate: Function;
}

const DeadlineInput: React.FC<DeadlineInput> = ({ date, setDate }) => {
  return (
    <div className="flex-column mb-8">
      <Label value="Deadline" />
      <Input
        value={date}
        type="datetime-local"
        onChange={(e) => setDate(e.target.value)}
      />
    </div>
  );
};

export default DeadlineInput;
