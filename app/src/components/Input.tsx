import clsx from "clsx";
import React from "react";

import styles from "$/styles/input.module.css";

interface InputProps {
  label: string;
  value?: string;
  required?: boolean;
  error?: string | null;
  id?: string;
  type?: "text" | "email" | "number" | "password" | "url" | "search"; // only the types with plain view
  name?: string;
  className?: string;
  onChange: (event: any) => void;
  onBlur: (event: any) => void;
}

const Input = React.forwardRef((props: InputProps, ref: React.Ref<any>) => {
  const {
    id,
    label,
    value,
    name,
    error,
    required = false,
    type = "text",
    className = "",
    ...rest
  } = props;

  return (
    <div
      className={clsx(styles.input, {
        [className]: className,
      })}
    >
      <label htmlFor={id} className={styles.input__label}>
        {label}
        {required && <span className={styles.input__required}>*</span>}
      </label>
      <input
        ref={ref}
        type={type}
        id={id}
        name={name || id}
        className={clsx(styles.input__field, {
          [styles["input__field--error"]]: error,
        })}
        required={required}
        value={value}
        {...rest}
      />
      {error && <span className={styles.input__error}>{error}</span>}
    </div>
  );
});

export default Input;
