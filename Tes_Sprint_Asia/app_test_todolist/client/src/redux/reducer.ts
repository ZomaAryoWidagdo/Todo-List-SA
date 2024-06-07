import { combineReducers } from "redux";
import taskReducer from "./slicer";

const rootReducer = combineReducers({
  tasks: taskReducer,
});

export default rootReducer;
