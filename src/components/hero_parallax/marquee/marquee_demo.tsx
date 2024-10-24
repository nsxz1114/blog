import React, { useMemo } from "react";
import Marquee from "@/components/ui/marquee";
import { AnimatedPinDemo } from "../marquee/3d_animated_pin/3d_animated_pin_demo";

// ReviewCard 使用 React.memo 包装，避免不必要的重新渲染
const ReviewCard = React.memo(
  ({
    title,
    description,
    link,
    thumbnail,
  }: {
    title: string;
    description?: string;
    link: string;
    thumbnail: string;
  }) => {
    return (
      <AnimatedPinDemo
        title={title}
        description={description}
        link={link}
        thumbnail={thumbnail}
      ></AnimatedPinDemo>
    );
  }
);

// MarqueeDemo 使用 React.memo 包装，避免父组件引发不必要的重新渲染
export const MarqueeDemo = React.memo(
  ({
    reviews,
  }: {
    reviews: {
      title: string;
      description?: string;
      link: string;
      thumbnail: string;
    }[];
  }) => {
    // 使用 useMemo 缓存切片操作，防止每次渲染都重新计算
    const [firstRow, secondRow] = useMemo(() => {
      const halfLength = Math.ceil(reviews.length / 2);
      return [reviews.slice(0, halfLength), reviews.slice(halfLength)];
    }, [reviews]);

    return (
      <div className="relative flex h-[550px] w-[1200px] flex-col items-center justify-center overflow-hidden rounded-lg bg-background">
        {/* 第一行 Marquee */}
        <Marquee pauseOnHover className="[--duration:20s]">
          {firstRow.map((review) => (
            <ReviewCard key={review.title} {...review} />
          ))}
        </Marquee>

        {/* 第二行 Marquee (反向滚动) */}
        <Marquee reverse pauseOnHover className="[--duration:20s]">
          {secondRow.map((review) => (
            <ReviewCard key={review.title} {...review} />
          ))}
        </Marquee>
      </div>
    );
  }
);
