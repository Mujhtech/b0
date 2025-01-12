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
      <div className="h-[1500px] w-full transition-all">
        <div className="translate-x-[-5160px] flex w-[10000px] h-[10000px] left-1/2 relative p-4 pt-10 justify-center transition-all overflow-x-hidden">
          <div
            ref={canvasRef}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
            onMouseLeave={handleMouseUp}
            onContextMenu={(e) => e.preventDefault()}
            className="flex select-none w-[10000px] h-[10000px]  absolute inset-0 flex-nowrap flex-col justify-start"
            style={{
              transform: `translate(${pan.x}px, ${pan.y}px) scale(${zoom})`,
              transformOrigin: "0 0",
            }}
          >
            <Connector />
          </div>
        </div>
      </div>
    </div>
  );
}
