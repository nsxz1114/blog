import React from "react";
import "./nav_right.css";
import { DockDemo } from "./dock_demo";

export type IconProps = React.HTMLAttributes<SVGElement>;

export function Navright() {
  return (
    <div className="nav_right">
      <DockDemo></DockDemo>
    </div>
  );
}
