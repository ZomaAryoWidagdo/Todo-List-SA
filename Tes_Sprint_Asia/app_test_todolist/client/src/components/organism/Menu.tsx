"use client";
import Label from "../atoms/Label";
import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { AppDispatch } from "../../redux/store";
import { fetchTasks } from "@/redux/thunks";

const Menu: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [status, setStatus] = useState("active");

  useEffect(() => {
    dispatch(fetchTasks({ status }));
  }, [dispatch, status]);

  const goToOnGoingTask = () => {
    setStatus("active");
  };

  const goToCompletedTask = () => {
    setStatus("finished");
  };

  return (
    <div
      style={{
        marginTop: "16px",
        display: "flex",
        flexDirection: "column",
        flexWrap: "wrap",
        background: "aliceblue",
        borderBottom: "black solid 1px",
        height: "40px",
      }}
    >
      <Label
        onClick={goToOnGoingTask}
        classname={`mb-0-important border-right w-40 h-100-percent text-center pointer ${
          status == "active" ? "bg-active" : ""
        }`}
        value="ACTIVE TASK"
      />

      <Label
        onClick={goToCompletedTask}
        classname={`mb-0-important w-40 h-100-percent text-center pointer ${
          status == "finished" ? "bg-active" : ""
        }`}
        value="COMPLETED TASK"
      />
    </div>
  );
};

export default Menu;
