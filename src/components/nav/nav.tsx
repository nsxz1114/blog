import { Navleft } from "./nav_left/nav_left";
import { Navright } from "./nav_right/nav_right";
import { FloatingNav } from "../ui/floating-navbar";

import "./nav.css";
export function Nav() {
  return (
    <FloatingNav>
      <div className="nav">
        <Navleft></Navleft>
        <Navright></Navright>
      </div>
    </FloatingNav>
  );
}
