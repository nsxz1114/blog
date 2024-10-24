import { Button } from "@/components/ui/button";
import BoxReveal from "@/components/ui/box-reveal";
import WordRotate from "../../../ui/word-rotate";

export function BoxRevealDemo() {
    return (
        <div className="size-full max-w-[400px] items-center justify-center overflow-hidden">
            <BoxReveal boxColor={"#5046e6"} duration={0.5}>
                <p className="text-[1.5rem] font-semibold">
                    你好！我是<WordRotate className="text-[#5046e6]" words={["溺水寻舟", "NSXZ"]} />
                </p>
            </BoxReveal>

            <BoxReveal boxColor={"#5046e6"} duration={0.5}>
                <p className="mt-[.5rem] text-[1rem]">
                </p>
            </BoxReveal>

            <BoxReveal boxColor={"#5046e6"} duration={0.5}>
                <p className="mt-[.5rem] text-[1rem]">
                </p>
            </BoxReveal>

            <BoxReveal boxColor={"#5046e6"} duration={0.5}>
                <Button className="mt-[.5rem] bg-[#5046e6]">Explore</Button>
            </BoxReveal>
        </div>
    );
}

