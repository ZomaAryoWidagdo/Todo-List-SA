import React from "react";

interface ButtonProps {
  onClick: () => void;
  children: React.ReactNode;
  classname?: string;
}

const Button: React.FC<ButtonProps> = ({ onClick, children, classname }) => {
  return (
    <button className={`rad-8 ${classname}`} onClick={onClick}>
      {children}
    </button>
  );
};

export default Button;
