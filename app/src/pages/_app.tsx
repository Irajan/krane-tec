import type { AppProps } from "next/app";
import { ApolloProvider } from "@apollo/client";

import "$/styles/globals.css";
import "$/styles/helper.css";

import Layout from "$/components/Layout";
import Header from "$/components/Header";
import { LayoutBackgrounds } from "$/enums/ui";

import client from "$/lib/apolloClient";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={client}>
      <Header {...pageProps} />
      <Layout backgroundType={LayoutBackgrounds.LIGHT}>
        <Component {...pageProps} />
      </Layout>
    </ApolloProvider>
  );
}
