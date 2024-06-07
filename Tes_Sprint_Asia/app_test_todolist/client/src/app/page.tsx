// src/pages/index.tsx
import AddTaskModal from "@/components/organism/AddTaskModal";
import Menu from "@/components/organism/Menu";
import TaskList from "@/components/organism/TaskList";
import React from "react";

const Home: React.FC = () => {
  return (
    <div style={{ width: "80vw" }}>
      <h1 className="title">Todo List</h1>
      <AddTaskModal />
      <Menu />
      <TaskList />
    </div>
  );
};

export default Home;
