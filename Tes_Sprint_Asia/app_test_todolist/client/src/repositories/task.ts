import { SubTask, SubTaskRequest } from "@/models/subTask";
import { Task, TaskRequest } from "@/models/task";

const API_URL = "http://localhost:8080"; // Pastikan skema URL adalah http:// atau https://

export class TaskRepository {
  async fetchTasks(status: string): Promise<any[]> {
    const response = await fetch(`${API_URL}/task?status=${status}`);
    const data = await response.json();
    return data.map((item: any) => ({
      description: item.Description,
      status: item.Status,
      subTasks: item.subTasks,
      deadline: item.Deadline,
      id: item.ID,
    }));
  }

  async addTask(task: TaskRequest): Promise<any> {
    const response = await fetch(`${API_URL}/task`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(task),
    });

    const data = await response.json();

    return {
      description: data.Description,
      status: data.Status,
      subTasks: data.subTasks,
      deadline: data.Deadline,
      id: data.ID,
    };
  }

  async updateTask(task: TaskRequest, taskId: number): Promise<any> {
    const response = await fetch(`${API_URL}/task/${taskId}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(task),
    });

    const data = await response.json();

    return {
      description: data.Description,
      status: data.Status,
      subTasks: data.subTasks,
      deadline: data.Deadline,
      id: data.ID,
    };
  }

  async deleteTask(id: number): Promise<void> {
    await fetch(`${API_URL}/task/${id}`, {
      method: "DELETE",
    });

    return;
  }

  async addSubTask(taskId: number, subTask: SubTask): Promise<any> {
    // Ubah return type ke any
    const response = await fetch(`${API_URL}/${taskId}/subtasks`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(subTask),
    });
    const data = await response.json();
    return {
      taskId: data.taskId,
      description: data.description,
      status: data.status,
      id: data.id,
    };
  }

  async updateSubTask(
    subTask: SubTaskRequest,
    subTaskId: number
  ): Promise<any> {
    console.log(subTask, subTaskId);

    const response = await fetch(`${API_URL}/subtask/${subTaskId}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(subTask),
    });

    const data = await response.json();

    return {
      TaskID: data.TaskID,
      Description: data.Description,
      Status: data.Status,
      ID: data.ID,
    };
  }

  async deleteSubTask(taskId: number, subTaskId: number): Promise<void> {
    await fetch(`${API_URL}/${taskId}/subtasks/${subTaskId}`, {
      method: "DELETE",
    });
  }
}
