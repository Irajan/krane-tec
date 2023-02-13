import React, { ReactNode } from "react";
import Loading from "./Loading";

type BtnProps = {
  children: ReactNode;
  isLoading?: boolean;
  className?: string;
};

function Button(props: BtnProps) {
  const { children, isLoading, className } = props;

  return (
    <button disabled={isLoading} className={className}>
      {children}
    </button>
  );
}

export default Button;
