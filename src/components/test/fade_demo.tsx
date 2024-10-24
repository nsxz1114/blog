import { FadeComponent } from "../ui/fade-component";
import { OrbitingCirclesDemo } from "@/components/orbiting_circles/orbiting_circles_demo";

export function FadeComponentDemo() {
  return (
    <FadeComponent
      direction="left"
      framerProps={{
        show: { transition: { delay: 0.8 } },
      }}
    >
      <OrbitingCirclesDemo></OrbitingCirclesDemo>
    </FadeComponent>
  );
}
