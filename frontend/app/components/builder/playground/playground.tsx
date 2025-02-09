import React from "react";
import { usePlayground } from "./provider";
import HttpRequestConnector from "./connector/http-request";
import { cn } from "~/lib/utils";
import { useOptionalEndpoint } from "~/hooks/use-project";
import HttpResponseConnector from "./connector/http-response";
import IfConnector from "./connector/if";
import VariableConnector from "./connector/variable";
import SwitchConnector from "./connector/switch";
import CodeblockConnector from "./connector/codeblock";

export default function Playground() {
  const {
    canvasRef,
    handleMouseMove,
    handleMouseDown,
    handleMouseUp,
    pan,
    isPanning,
    zoom,
    handleDrag,
  } = usePlayground();

  const endpoint = useOptionalEndpoint();

  console.log("endpoint", endpoint?.workflows);

  return (
    <div
      className={cn(
        "w-full h-full overflow-y-auto flex-1 pointer-events-auto scrollbar-none",
        isPanning && "cursor-grabbing"
      )}
    >
      <div className="h-full w-full transition-all">
        <div className="flex w-full h-full relative p-4 pt-10 justify-center transition-all overflow-x-hidden scrollbar-none">
          <div
            ref={canvasRef}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
            onMouseLeave={handleMouseUp}
            onDrag={handleDrag}
            onContextMenu={(e) => e.preventDefault()}
            className="flex select-none w-full h-full absolute inset-0 flex-nowrap flex-col justify-start origin-[0_0] scrollbar-none"
            style={{
              transform: `translate(${pan.x}px, ${pan.y}px) scale(${
                zoom / 100
              })`,
            }}
          >
            <div className="mt-5 ml-20 flex flex-col justify-center ">
              <HttpRequestConnector />
              {endpoint?.workflows?.map((workflow, index) => {
                switch (workflow.type) {
                  case "response":
                    return (
                      <HttpResponseConnector key={index} workflow={workflow} />
                    );
                  case "if":
                    return <IfConnector key={index} workflow={workflow} />;
                  case "variable":
                    return (
                      <VariableConnector key={index} workflow={workflow} />
                    );
                  case "switch":
                    return <SwitchConnector key={index} workflow={workflow} />;
                  case "switch":
                    return (
                      <CodeblockConnector key={index} workflow={workflow} />
                    );
                  case "request":
                  default:
                    return null;
                }
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
