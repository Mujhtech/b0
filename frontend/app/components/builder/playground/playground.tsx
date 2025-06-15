import { usePlayground } from "./provider";
import HttpRequestConnector from "./connector/http-request";
import { cn } from "~/lib/utils";
import Connector from "./connector";
import { EndpointWorkflow } from "~/models/endpoint";

export default function Playground({
  handleRemoveWorkflow,
  handleUpdateWorkflows,
  workflows,
  activeSlot,
  activeCard,
}: {
  handleRemoveWorkflow: (index: number) => void;
  handleUpdateWorkflows: (index: number, workflow: EndpointWorkflow) => void;
  workflows: EndpointWorkflow[];
  activeSlot: string | null;
  activeCard: string | null;
}) {
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
              {workflows.map((workflow, index) => {
                const id = workflow.action_id ?? index.toString();
                return (
                  <Connector
                    workflow={workflow}
                    index={index}
                    key={index}
                    onUpdate={(workflow) => {
                      handleUpdateWorkflows(index, workflow);
                    }}
                    onRemove={() => handleRemoveWorkflow(index)}
                    draggable={{
                      isActive: activeSlot === id,
                    }}
                  />
                );
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
