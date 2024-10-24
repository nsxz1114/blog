"use client";
import { PinContainer } from "../../../ui/3d-pin";
import { Image } from "antd";
import { Typography } from "antd";

export function AnimatedPinDemo({
  title,
  description,
  link,
  thumbnail,
}: {
  title: string;
  description?: string;
  link: string;
  thumbnail: string;
}) {
  const { Paragraph } = Typography;
  return (
    <div className="h-[250px] w-[350px]">
      <PinContainer href={link}>
        <div>
          <h3>{title}</h3>
          <Paragraph ellipsis={true ? { rows: 2 } : false}>
            {description}
          </Paragraph>
          <div style={{ height: 200, width: "100%", overflow: "hidden" }}>
            <Image
              style={{ objectFit: "cover", height: "100%", width: "100%" }}
              preview={false}
              src={thumbnail}
            />
          </div>
        </div>
      </PinContainer>
    </div>
  );
}
