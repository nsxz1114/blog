"use client";
import React from "react";
import { motion, useScroll, useTransform } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import { MarqueeDemo } from "../hero_parallax/marquee/marquee_demo";

export const HeroParallax = ({
  products,
}: {
  products: {
    title: string;
    description?: string;
    link: string;
    thumbnail: string;
  }[];
}) => {
  const ref = React.useRef(null);
  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ["start start", "end start"],
  });

  const rotateX = useTransform(scrollYProgress, [0, 0.1], [1, 0]);
  const opacity = useTransform(scrollYProgress, [0, 0.2], [1, 1]);
  const rotateZ = useTransform(scrollYProgress, [0, 0.2], [10, 0]);
  const translateY = useTransform(scrollYProgress, [0, 0.2], [-100, 500]);

  return (
    <div
      ref={ref}
      className="h-[200vh] py-40 overflow-hidden antialiased relative flex flex-col self-auto [perspective:1000px] [transform-style:preserve-3d]"
    >
      <motion.div
        style={{ rotateX, rotateZ, translateY, opacity }}
        className="will-change-transform"
      >
        <motion.div className="flex flex-row mb-20 space-x-20 justify-center">
          <MarqueeDemo reviews={products}></MarqueeDemo>
        </motion.div>
      </motion.div>
    </div>
  );
};

export const ProductCard = ({
  product,
}: {
  product: {
    title?: string;
    link: string;
    thumbnail: string;
  };
}) => {
  return (
    <motion.div
      whileHover={{ y: -10 }}
      key={product.title}
      className="group/product h-[15rem] w-[20rem] relative flex-shrink-0"
    >
      <Link href={product.link}>
        <Image
          src={product.thumbnail}
          height={600}
          width={600}
          className="object-cover object-left-top absolute h-full w-full inset-0"
          alt={product.title ?? product.link}
          priority={false} // 优化加载优先级，除非真的需要高优先级
          unoptimized={false} // 不必要时移除 unoptimized
        />
      </Link>
    </motion.div>
  );
};
