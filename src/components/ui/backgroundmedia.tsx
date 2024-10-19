"use client";
import React, { useRef, useEffect, memo } from "react";
import { cva } from "class-variance-authority";
import { cn } from "@/lib/utils";
import { useWebContext } from "@/store/web_context";

// 定义类型
type OverlayVariant = "none" | "light" | "dark";
type MediaType = "image" | "video";

// 样式变体
const backgroundVariants = cva(
  "relative h-screen max-h-[1000px] w-full min-h-[500px] lg:min-h-[600px]",
  {
    variants: {
      overlay: {
        none: "",
        light:
          "before:absolute before:inset-0 before:bg-white before:opacity-30",
        dark: "before:absolute before:inset-0 before:bg-black before:opacity-30",
      },
      type: {
        image: "",
        video: "z-10",
      },
    },
    defaultVariants: {
      overlay: "none",
      type: "image",
    },
  }
);

// 组件属性接口
interface BackgroundMediaProps {
  variant?: OverlayVariant;
  type?: MediaType;
  src: string;
  alt?: string;
}

// 优化：懒加载处理
const useLazyLoadVideo = (mediaRef: React.RefObject<HTMLVideoElement>, src: string) => {
  const observerRef = useRef<IntersectionObserver | null>(null);

  useEffect(() => {
    const videoElement = mediaRef.current;
    if (videoElement) {
      observerRef.current = new IntersectionObserver(
        (entries) => {
          const entry = entries[0];
          if (entry.isIntersecting) {
            videoElement.play();
          } else {
            videoElement.pause();
          }
        },
        { threshold: 0.5 } // 开始播放的阈值
      );
      observerRef.current.observe(videoElement);
    }
    return () => observerRef.current?.disconnect();
  }, [mediaRef, src]);
};

// BackgroundMedia 组件
const BackgroundMedia: React.FC<BackgroundMediaProps> = memo(
  ({ variant = "light", type = "image", src, alt = "" }) => {
    const { mediaRef } = useWebContext();
    const imgRef = useRef<HTMLImageElement>(null);

    // Lazy load video using IntersectionObserver
    useLazyLoadVideo(mediaRef, src);

    // 渲染媒体内容
    const renderMedia = () => {
      const commonProps = {
        className:
          "absolute inset-0 h-full w-full object-cover transition-opacity duration-300",
      };

      if (type === "video") {
        return (
          <video
            ref={mediaRef}
            aria-hidden="true"
            muted
            preload="metadata" // 优化视频加载
            playsInline
            loop
            {...commonProps}
          >
            <source src={src} type="video/mp4" />
            Your browser does not support the video tag.
          </video>
        );
      } else {
        return (
          <img
            ref={imgRef}
            src={src}
            alt={alt}
            loading="lazy" // 懒加载图片
            className={`${commonProps.className}`}
          />
        );
      }
    };

    return (
      <div
        className={cn(
          backgroundVariants({ overlay: variant, type }),
          "overflow-hidden fixed inset-0 z-0"
        )}
      >
        {renderMedia()}
      </div>
    );
  }
);

export default BackgroundMedia;
