// src/redux/slices/taskSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Task } from "@/models/task";
import { SubTask } from "@/models/subTask";

interface TaskState {
  tasks: Task[];
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: TaskState = {
  tasks: [],
  status: "idle",
  error: null,
};

const taskSlice = createSlice({
  name: "tasks",
  initialState,
  reducers: {
    setTasks: (state, action: PayloadAction<Task[]>) => {
      state.tasks = action.payload;
    },
    addTask: (state, action: PayloadAction<Task>) => {
      state.tasks.unshift(action.payload);
    },
    updateTask: (state, action: PayloadAction<Task>) => {
      const index = state.tasks.findIndex((t) => t.id === action.payload.id);

      const task = (state.tasks[index] = action.payload);

      if (task.status == "finished") {
        state.tasks = state.tasks.filter((t) => t.id !== task.id);
      }
    },
    deleteTask: (state, action: PayloadAction<number>) => {
      state.tasks = state.tasks.filter((task) => task.id !== action.payload);
    },
    addSubTask: (state, action: PayloadAction<SubTask>) => {
      const { TaskID, Description, Status, ID } = action.payload;
      const task = state.tasks.find((task) => task.id === TaskID);
      if (task) {
        task.subTasks.push({ TaskID, Description, Status, ID } as SubTask);
      }
    },
    updateSubTask: (state, action: PayloadAction<SubTask>) => {
      const newTasks = [...state.tasks];

      const index = newTasks.findIndex((t) => t.id === action.payload.TaskID);

      const task = newTasks[index];

      let subTaskIndex = 0;

      const totalSubtask = task.subTasks.length;
      let counterFinishedSubTask = 0;

      task.subTasks.forEach((st, i) => {
        if (st.ID == action.payload.ID) {
          subTaskIndex = i;
        }

        if (st.Status == "finished") {
          counterFinishedSubTask++;
        }
      });

      if (counterFinishedSubTask == totalSubtask) {
        newTasks.filter((nt) => nt.id != action.payload.TaskID);
      } else {
        newTasks[index].subTasks[subTaskIndex] = action.payload;
      }

      state.tasks = newTasks;
    },
    deleteSubTask: (
      state,
      action: PayloadAction<{ taskId: number; subTaskId: number }>
    ) => {
      const { taskId, subTaskId } = action.payload;
      const task = state.tasks.find((task) => task.id === taskId);
      if (task) {
        task.subTasks = task.subTasks.filter((st) => st.ID !== subTaskId);
      }
    },
  },
  extraReducers: (builder) => {
    // builder.addCase(fetchTasks.pending, (state) => {
    //   state.status = "loading";
    // });
    // .addCase(fetchTasks.fulfilled, (state, action) => {
    //   action.payload.map(
    //     (item: any) =>
    //       new Task(
    //         item.id,
    //         item.description,
    //         item.status,
    //         item.subTasks,
    //         item.deadline
    //       )
    //   );
    //   state.status = "succeeded";
    //   state.tasks = action.payload;
    // })
    // .addCase(fetchTasks.rejected, (state, action) => {
    //   state.status = "failed";
    //   state.error = action.error.message || null;
    // });
  },
});

export const {
  setTasks,
  addTask,
  updateTask,
  deleteTask,
  addSubTask,
  updateSubTask,
  deleteSubTask,
} = taskSlice.actions;

export default taskSlice.reducer;
