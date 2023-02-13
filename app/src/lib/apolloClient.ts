import config from "$/config";
import { getAccessToken } from "$/utils/token";
import { setContext } from "@apollo/client/link/context";
import { ApolloClient, concat, HttpLink, InMemoryCache } from "@apollo/client";

const httpLink = new HttpLink({
  uri: "http://localhost:8080/graphql",
});

const authLink = setContext((_, { headers }) => {
  const accessToken = getAccessToken();

  return {
    headers: {
      ...headers,
      ...(accessToken ? { authorization: `Bearer ${accessToken}}` } : {}),
    },
  };
});

const client = new ApolloClient({
  link: concat(authLink, httpLink),
  cache: new InMemoryCache(),
});

export default client;
