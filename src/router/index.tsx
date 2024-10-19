import { Admin } from "@/views/admin";
import { Web } from "@/views/web/index";
import { Home } from "@/views/web/home/home";
import Notfound from "@/views/web/not_found/not_found";
import { createBrowserRouter } from "react-router-dom";
import { Timeline } from "@/views/web/timeline/timeline";
import { Blog } from "@/views/web/blog/blog";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Web />,
    children: [
      {
        path: "",
        element: <Home />,
      },
      {
        path: "timeline",
        element: <Timeline />,
      },
      {
        path: "blog",
        element: <Blog />,
      },
    ],
  },
  {
    path: "/admin",
    element: <Admin />,
  },
  {
    path: "*",
    element: <Notfound />,
  },
]);

export default router;
