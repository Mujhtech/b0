import React, { useMemo, useState } from "react";
import { usePlayground } from "./provider";
import HttpRequestConnector from "./connector/http-request";
import { cn } from "~/lib/utils";
import { useOptionalEndpoint } from "~/hooks/use-project";
import Connector from "./connector";
import { EndpointWorkflow } from "~/models/endpoint";
import { usePlaygroundBuilder } from "../provider";

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
  const { handleUpdateEndpointWorkflow, setIsSaving } = usePlaygroundBuilder();

  const endpoint = useOptionalEndpoint();

  const [newWorkflows, setNewWorkflows] = useState<
    Array<EndpointWorkflow> | undefined
  >(undefined);

  const workflows = useMemo(() => {
    if (newWorkflows) {
      return newWorkflows;
    }

    return endpoint?.workflows || [];
  }, [newWorkflows, endpoint]);

  const handleUpdateWorkflows = (index: number, workflow: EndpointWorkflow) => {
    setNewWorkflows((old) => {
      const prev = newWorkflows || old;

      if (!prev) {
        return undefined;
      }

      const workflows = [...prev];

      workflows[index] = workflow;

      return workflows;
    });

    if (!endpoint) {
      return;
    }
    // TODO: update endpoint
    setTimeout(() => {
      setIsSaving(true);
      handleUpdateEndpointWorkflow(endpoint.id, workflows).finally(() => {
        setIsSaving(false);
      });
    }, 1000);
  };

  const handleRemoveWorkflow = (index: number) => {
    setNewWorkflows((old) => {
      const prev = newWorkflows || old;

      if (!prev) {
        return undefined;
      }
      const workflows = [...prev];

      workflows.splice(index, 1);

      return workflows;
    });
  };

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
                return (
                  <Connector
                    workflow={workflow}
                    index={index}
                    key={index}
                    onUpdate={(workflow) =>
                      handleUpdateWorkflows(index, workflow)
                    }
                    onRemove={() => handleRemoveWorkflow(index)}
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
