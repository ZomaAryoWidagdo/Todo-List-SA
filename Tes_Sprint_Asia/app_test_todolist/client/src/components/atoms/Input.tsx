import React, { ChangeEvent } from "react";

interface InputProps {
  value: string;
  type: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
}

const Input: React.FC<InputProps> = ({ value, type = "text", onChange }) => {
  return <input value={value} type={type} onChange={onChange} />;
};

export default Input;
