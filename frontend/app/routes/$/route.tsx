import React from "react";
import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";

export default function Page() {
  return (
    <main className="w-full h-full relative">
      <div className="flex w-full h-full absolute inset-0 canvas-bg"></div>
      <div className="absolute bottom-4 w-full">
        <div className="flex justify-between">
          <div className="flex items-center gap-2 ml-4">
            <UserMenu />
            <ZoomInAndOut />
          </div>
          <div>
            <AskB0 />
          </div>
          <div className="mr-4">
            <DeployAndTestBtn />
          </div>
        </div>
      </div>
      <BuilderTools />
    </main>
  );
}
