import "./nav_left.css";
import { NavbarDemo } from "./navbar";
import { Link } from "react-router-dom";

export function Navleft() {
  return (
    <div className="nav_left">
      <Link to="/" className="slogan">
        <span className="pointer-events-none whitespace-pre-wrap bg-gradient-to-b from-black to-gray-300/80 bg-clip-text text-center font-semibold leading-none text-transparent">
          NSXZ
        </span>
      </Link>
      <div className="function">
        <NavbarDemo></NavbarDemo>
      </div>
    </div >
  );
}
