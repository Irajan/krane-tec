import React, { SetStateAction } from "react";
import { useForm } from "react-hook-form";

import Input from "$/components/Input";
import styles from "$/styles/form.module.css";
import Link from "next/link";
import { useLazyQuery, useMutation, useQuery } from "@apollo/client";
import { LOG_IN_QUERY } from "$/queries/login";
import Button from "$/components/Button";
import { setAccessToken, setRefreshToken } from "$/utils/token";

interface FormData {
  email: string;
  password: string;
}

type LoginFormProps = {
  setIsLoggedIn: (isLoggedIn: boolean) => void;
};

const LoginForm = (props: LoginFormProps) => {
  const { setIsLoggedIn } = props;

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<FormData>();

  const [logIn, { loading, error, data }] = useMutation(LOG_IN_QUERY);

  const onSubmit = async (formData: FormData) => {
    await logIn({
      variables: {
        email: formData.email,
        password: formData.password,
      },
    });

    console.log(data);

    if (data) {
      setAccessToken(data.accessToken);
      setRefreshToken(data.refreshToken);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className={styles.formContainer}>
      <h1 className={styles.formTitle}>Login Form</h1>
      <p className={styles.formDescription}>
        Provide necessary details to continue
      </p>
      <div className={styles.formGroup}>
        <Input
          label="Email"
          {...register("email")}
          required
          error={errors.email?.message}
        />
      </div>
      <div className={styles.formGroup}>
        <Input
          label="Password"
          {...register("password")}
          type="password"
          required
          error={errors.password?.message}
        />
      </div>
      <Button isLoading={loading} className={styles.formButton}>
        Log In
      </Button>
      <div className={styles.formGroup}>
        <Link href={"/register"} className={styles.formDescription}>
          Don't have account
        </Link>
      </div>
    </form>
  );
};

export default LoginForm;
