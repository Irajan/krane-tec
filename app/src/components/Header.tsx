import React from "react";

import Layout from "./Layout";
import { HomeProps } from "$/pages";
import styles from "$/styles/header.module.css";

type BaseHeaderProps = {
  // Can add base header props later
};

type HeaderProps = BaseHeaderProps & HomeProps;

export default function Header(props: HeaderProps) {
  return (
    <header className={styles.header}>
      <Layout className="auto-height">
        <div className={styles.headerContent}>
          <div className={styles.logoWrapper}>
            <img className={styles.logo} src="logo.png" alt="Logo" />
          </div>
        </div>
      </Layout>
    </header>
  );
}
