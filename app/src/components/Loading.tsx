import React from "react";
import clsx from "clsx";
import styles from "$/styles/loading.module.css";

interface LoadingProps {
  isLoading?: boolean;
  children?: React.ReactNode;
  className?: string;
}

const Loading: React.FC<LoadingProps> = (props) => {
  const { children, className, isLoading = true } = props;

  return (
    <div className={clsx(styles["loading-container"], className)}>
      {isLoading && (
        <>
          <div className={styles["bouncing-ball"]}></div>
          <div className={styles["bouncing-ball"]}></div>
          <div className={styles["bouncing-ball"]}></div>
        </>
      )}
      {
        <div style={{ visibility: isLoading ? "hidden" : "visible" }}>
          {children}
        </div>
      }
    </div>
  );
};

export default Loading;
