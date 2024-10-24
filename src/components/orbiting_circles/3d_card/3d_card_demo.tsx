"use client";

import { BoxRevealDemo } from "@/components/orbiting_circles/3d_card/box_reveal/box_reveal_demo";
import { CardBody, CardContainer, CardItem } from "../../ui/3d-card";

export function ThreeDCardDemo() {
  return (
    <CardContainer className="inter-var">
      <CardBody className="bg-gray-50 relative group/card border-black/[0.1] w-auto h-auto rounded-xl p-6 border">
        <CardItem translateZ="40">
          <BoxRevealDemo></BoxRevealDemo>
        </CardItem>
      </CardBody>
    </CardContainer>
  );
}
