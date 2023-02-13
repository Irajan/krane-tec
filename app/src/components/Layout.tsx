import React, { ReactNode } from "react";

import styles from "$/styles/layout.module.css";
import { LayoutBackgrounds } from "$/enums/ui";
import clsx from "clsx";

interface LayoutProps {
  children: ReactNode;
  className?: string;
  backgroundType?: LayoutBackgrounds;
}

function Layout(props: LayoutProps) {
  return (
    <div
      className={clsx(styles.container, {
        [props.className]: !!props.className,
        [styles["container--dark"]]:
          props.backgroundType == LayoutBackgrounds.DARK,
        [styles["container--light"]]:
          props.backgroundType == LayoutBackgrounds.LIGHT,
        [styles["container--unset"]]:
          props.backgroundType == LayoutBackgrounds.NONE,
      })}
    >
      {props.children}
    </div>
  );
}

export default Layout;
