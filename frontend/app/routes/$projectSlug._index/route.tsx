import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import { PlaygroundProvider } from "~/components/builder/playground/provider";
import Playground from "~/components/builder/playground/playground";
import BuilderMenu from "~/components/builder/builder-menu";
import { PlaygroundBuilderProvider } from "~/components/builder/provider";

export default function Page() {
  return (
    <PlaygroundBuilderProvider>
      <PlaygroundProvider>
        <main className="h-full w-full relative canvas-bg">
          <Playground />

          <div className="absolute bottom-4 w-full z-10">
            <div className="flex justify-between">
              <div className=" ml-4">
                <div className="flex items-center gap-2">
                  <UserMenu />
                  <ZoomInAndOut />
                </div>
              </div>
              <div className="flex flex-col gap-0.5">
                <AskB0 />
                <div className="flex items-center justify-center">
                  <p className="font-mono text-[10px] text-muted-foreground text-center">
                    b0 can make mistakes. Please double-check it.
                  </p>
                </div>
              </div>
              <div className="mr-4">
                <DeployAndTestBtn />
              </div>
            </div>
          </div>
          <BuilderMenu />
          <BuilderTools />
        </main>
      </PlaygroundProvider>
    </PlaygroundBuilderProvider>
  );
}
