import React from "react";
import { EndpointWorkflow } from "~/models/endpoint";
import { ConnectorProps } from ".";

export default function IfConnector({ workflow }: ConnectorProps) {
  return (
    <div className="flex flex-col group h-min">
      <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
        <div className="border-b border-input flex items-center justify-center">
          <h3 className="text-xs font-mono text-muted-foreground p-2">If</h3>
        </div>
        <div className="p-2">
          {workflow && (
            <div>
              <p className="text-xs font-mono text-muted-foreground">
                {workflow.instruction}
              </p>
            </div>
          )}
        </div>
      </div>
      <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div>
    </div>
  );
}
