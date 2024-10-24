import { Meteors } from "../ui/meteors";

export function MeteorsDemo() {
  return (
    <div className=" w-full relative">
      <div className="absolute inset-0 h-full w-full blur-3xl" />
      <div className="relative px-4 py-8 h-full overflow-hidden flex flex-col justify-end items-start">
        <Meteors number={20} />
      </div>
    </div>
  );
}
