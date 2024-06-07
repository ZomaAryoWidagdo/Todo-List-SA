export class SubTask {
  constructor(
    public TaskID: number,
    public Description: string,
    public Status: string,
    public ID: number
  ) {}
}

export class SubTaskRequest {
  constructor(public description?: string, public status?: string) {}
}
