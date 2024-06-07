import React from "react";

interface LabelProps {
  value: string;
  classname?: string;
  onClick?: () => void;
}

const Label: React.FC<LabelProps> = ({ value, classname, onClick }) => {
  return (
    <label className={`${classname}`} onClick={onClick}>
      {value}
    </label>
  );
};

export default Label;
