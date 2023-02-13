import { gql } from "@apollo/client";

export const LOG_IN_QUERY = gql`
  mutation logIn($email: String!, $password: String!) {
    login(email: $email, password: $password) {
      userId
      accessToken
      refreshToken
    }
  }
`;
