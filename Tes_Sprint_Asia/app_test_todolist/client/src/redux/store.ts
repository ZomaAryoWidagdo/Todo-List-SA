import { configureStore } from "@reduxjs/toolkit";
import taskReducer from "./reducer";
import { thunk } from "redux-thunk";

const store = configureStore({
  reducer: {
    tasks: taskReducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(thunk),
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;

export default store;
