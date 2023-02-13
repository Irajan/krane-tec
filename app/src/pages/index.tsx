import Head from "next/head";
import {
  GetServerSideProps,
  GetServerSidePropsContext,
  GetServerSidePropsResult,
} from "next";
import jwt from "jsonwebtoken";

import { User } from "$/types/user";
import Dashboard from "./_components/Dashboard";
import LoginForm from "./_components/LoginForm";
import { useState } from "react";

type AuthenticatedProps = {
  isLoggedIn: true;
  userData: User;
};

type UnAuthenticatedProps = {
  isLoggedIn: false;
  userData: null;
};

export type HomeProps = AuthenticatedProps | UnAuthenticatedProps;

export default function Home(props: HomeProps) {
  const [isLoggedIn, setIsLoggedIn] = useState(props.isLoggedIn);

  return (
    <>
      <Head>
        <title>Allergy Less</title>
        <meta
          name="description"
          content="App for tracking allergies and making them managed"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      {isLoggedIn ? <Dashboard /> : <LoginForm setIsLoggedIn={setIsLoggedIn} />}
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<HomeProps>> => {
  const { cookies } = context.req;

  // Check if access token exist on cookies or not

  const accessToken = cookies.accessToken;

  if (!accessToken) {
    return {
      props: {
        isLoggedIn: false,
        userData: null,
      },
    };
  }

  // Decode access token and set the values in access token and set the value of access token to user data

  const payload = jwt.decode(accessToken) as User;

  const user = {
    id: payload.id,
    fullName: payload.fullName,
    role: payload.role,
  };

  return {
    props: {
      isLoggedIn: true,
      userData: user,
    },
  };
};
