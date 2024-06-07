// src/redux/thunks/taskThunks.ts
import { createAsyncThunk } from "@reduxjs/toolkit";
import {
  addTask,
  setTasks,
  updateTask,
  deleteTask,
  addSubTask,
  updateSubTask,
  deleteSubTask,
} from "./slicer";
import { Task, TaskRequest } from "@/models/task";
import { SubTask, SubTaskRequest } from "@/models/subTask";
import { TaskRepository } from "@/repositories/task";

const taskRepository = new TaskRepository();

export const fetchTasks = createAsyncThunk(
  "tasks/fetchTasks",
  async (
    {
      status,
    }: {
      status: string;
    },
    { dispatch }
  ) => {
    const allTasks = await taskRepository.fetchTasks(status);
    dispatch(setTasks(allTasks));
  }
);

export const addTaskAsync = createAsyncThunk(
  "tasks/addTask",
  async (
    {
      description,
      deadline,
      subTasks,
    }: {
      description?: string;
      deadline?: string;
      subTasks?: Array<object>;
    },
    { dispatch }
  ) => {
    const taskRequest = new TaskRequest(
      description,
      undefined,
      deadline,
      subTasks
    );

    const newTask = await taskRepository.addTask(taskRequest);

    dispatch(addTask(newTask));
  }
);

export const updateTaskAsync = createAsyncThunk(
  "tasks/updateTask",
  async (
    {
      taskId,
      status,
      description,
      deadline,
    }: {
      taskId: number;
      status?: string;
      description?: string;
      deadline?: string;
    },
    { dispatch }
  ) => {
    const taskRequest = new TaskRequest(description, status, deadline);

    const task = await taskRepository.updateTask(taskRequest, taskId);

    dispatch(updateTask(task));
  }
);

export const deleteTaskAsync = createAsyncThunk(
  "tasks/deleteTask",
  async (taskId: number, { dispatch }) => {
    await taskRepository.deleteTask(taskId);
    dispatch(deleteTask(taskId));
  }
);

export const addSubTaskAsync = createAsyncThunk(
  "tasks/addSubTask",
  async (
    {
      taskId,
      description,
      status,
    }: { taskId: number; description: string; status: string },
    { dispatch }
  ) => {
    const id = 0;

    const newSubTask = await taskRepository.addSubTask(
      taskId,
      new SubTask(taskId, description, status, id)
    );
    dispatch(addSubTask(newSubTask));
  }
);

export const updateSubTaskAsync = createAsyncThunk(
  "tasks/updateSubTask",
  async (
    {
      subTaskId,
      status,
      description,
    }: {
      subTaskId: number;
      status?: string;
      description?: string;
    },
    { dispatch }
  ) => {
    const subTaskRequest = new SubTaskRequest(description, status);

    const subTask = await taskRepository.updateSubTask(
      subTaskRequest,
      subTaskId
    );

    dispatch(updateSubTask(subTask));
  }
);

export const deleteSubTaskAsync = createAsyncThunk(
  "tasks/deleteSubTask",
  async (
    { taskId, subTaskId }: { taskId: number; subTaskId: number },
    { dispatch }
  ) => {
    await taskRepository.deleteSubTask(taskId, subTaskId);
    dispatch(deleteSubTask({ taskId, subTaskId }));
  }
);
