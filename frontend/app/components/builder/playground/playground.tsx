import React from "react";
import { usePlayground } from "./provider";
import Connector from "./connector";

export default function Playground() {
  const {
    canvasRef,
    handleMouseMove,
    handleMouseDown,
    handleMouseUp,
    pan,
    zoom,
  } = usePlayground();

  return (
    <div className="w-full h-full overflow-y-auto flex-1 pointer-events-auto scrollbar-none">
      <div className="h-full w-full transition-all">
        <div className="flex w-full h-full relative p-4 pt-10 justify-center transition-all overflow-x-hidden">
          <div
            ref={canvasRef}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
            onMouseLeave={handleMouseUp}
            onContextMenu={(e) => e.preventDefault()}
            className="flex select-none w-full h-full absolute inset-0 flex-nowrap flex-col justify-start"
            // style={{
            //   transform: `translate(${pan.x}px, ${pan.y}px) scale(${zoom})`,
            //   transformOrigin: "0 0",
            // }}
          >
            <div className="mt-5 ml-20 flex flex-col justify-center">
              <Connector />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
