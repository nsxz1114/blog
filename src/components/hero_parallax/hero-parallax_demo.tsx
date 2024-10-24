"use client";
import { projectList, projectListType } from "@/api/system";
import { HeroParallax } from "../ui/hero-parallax";
import { useEffect, useState } from "react";

export function HeroParallaxDemo() {
  const [projectItems, setProjectItems] = useState<projectListType[]>([]);
  useEffect(() => {
    const fetchNavList = async () => {
      const res = await projectList();
      setProjectItems(res.data.list);
    };
    fetchNavList();
  }, []);
  return <HeroParallax products={projectItems} />;
}
// const products = [
//   {
//     title: "blog",
//     description: "react+gin",
//     link: "https://github.com/As1114/blog",
//     thumbnail: "public/images/image.png",
//   },
//   {
//     title: "blog2",
//     description: "react+gin",
//     link: "https://github.com/As1114/blog",
//     thumbnail: "public/images/image.png",
//   },
//   {
//     title: "blog3",
//     description: "react+gin",
//     link: "https://github.com/As1114/blog",
//     thumbnail: "public/images/image.png",
//   },
//   {
//     title: "blog4",
//     description: "react+gin",
//     link: "https://github.com/As1114/blog",
//     thumbnail: "public/images/image.png",
//   },
//   {
//     title: "blog5",
//     description: "react+gin",
//     link: "https://github.com/As1114/blog",
//     thumbnail: "public/images/image.png",
//   },
// ];
