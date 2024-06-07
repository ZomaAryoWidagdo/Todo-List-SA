import { SubTask } from "./subTask";

export class Task {
  constructor(
    public id: number,
    public description: string,
    public status: string,
    public subTasks: SubTask[] = [],
    public deadline: string
  ) {}
}

export class TaskRequest {
  constructor(
    public description?: string,
    public status?: string,
    public deadline?: string,
    public subTasks?: object[]
  ) {}
}
