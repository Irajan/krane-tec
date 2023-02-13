import React from "react";

import styles from "$/styles/form.module.css";
import Input from "$/components/Input";

function register() {
  return (
    <form className={styles.formContainer}>
      <h1 className={styles.formTitle}>Registration Form</h1>
      <p className={styles.formDescription}>
        Provide necessary details to continue
      </p>
      <div className={styles.formGroup}>
        <Input label="Email" id="email" name="email" required />
      </div>
      <div className={styles.formGroup}>
        <Input
          label="Password"
          id="password"
          name="password"
          type="password"
          required
        />
      </div>
      <button type="submit" className={styles.formButton}>
        Submit
      </button>
    </form>
  );
}

export default register;
