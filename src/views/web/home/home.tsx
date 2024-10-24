import "./home.css";
import { OrbitingCirclesDemo } from "@/components/orbiting_circles/orbiting_circles_demo";
import { HeroParallaxDemo } from "@/components/hero_parallax/hero-parallax_demo";
import { FadeComponent } from "@/components/ui/fade-component";
import BackgroundMedia from "@/components/ui/backgroundmedia";

export function Home() {
  return (
    <div className="home">
      <BackgroundMedia
        type="video"
        variant="none"
        src="/public/media/185096-874643413.mp4"
      ></BackgroundMedia>
      <FadeComponent
        direction="left"
        framerProps={{
          show: { transition: { delay: 0.5 } },
        }}
        className="circle_fade"
      >
        <OrbitingCirclesDemo></OrbitingCirclesDemo>
      </FadeComponent>
      <HeroParallaxDemo></HeroParallaxDemo>
    </div>
  );
}
