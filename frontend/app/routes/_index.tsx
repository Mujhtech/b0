import {
  ArrowUpRight,
  Folder,
  GlobeHemisphereWest,
  PaperPlaneTilt,
} from "@phosphor-icons/react";
import type { MetaFunction } from "@remix-run/node";
import AnimatedGradient from "~/components/animated-bg";
import { Button } from "~/components/ui/button";
import { Textarea } from "~/components/ui/textarea";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <main className="relative h-full w-full flex flex-col flex-1">
      {/* <AnimatedGradient
        speed={0.05}
        blur="medium"
        colors={["#0E0E0E", "#0E0E0E"]}
      /> */}
      <div className="z-10 flex flex-col w-full max-w-[50rem] items-center mx-auto pt-56">
        <h1 className="font-sans text-3xl font-black">
          b0, Your AI backend builder
        </h1>
        <div className="mt-4 flex flex-col gap-4 w-full">
          <form className="w-full border border-input bg-background shadow-lg">
            <Textarea
              placeholder="What are you building today?"
              className=" resize-none border-none focus-visible:ring-0 min-h-[100px] shadow-none"
            ></Textarea>
            <div className="flex items-center px-3 pb-2">
              <div className="flex gap-2">
                <button
                  type="button"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <Folder className="h-5 w-5" />
                </button>
                <button
                  type="button"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <GlobeHemisphereWest className="h-5 w-5" />
                </button>
              </div>
              <div className="ml-auto flex items-center gap-2">
                <button
                  type="button"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <PaperPlaneTilt className="h-5 w-5" />
                </button>
              </div>
            </div>
          </form>
        </div>
        <div className="mt-12 flex flex-wrap justify-center gap-3 w-full">
          {Array.from({ length: 10 }).map((_, i) => (
            <TemplateCard key={i} />
          ))}
        </div>
      </div>
    </main>
  );
}

const TemplateCard = () => {
  return (
    <Button
      variant="outline"
      type="button"
      className="border border-input shadow-lg"
    >
      TemplateCard <ArrowUpRight className="!h-4 !w-4" />
    </Button>
  );
};
