import React, { useCallback, useMemo, useState } from "react";
import { EndpointWorkflow } from "~/models/endpoint";
import Connector, { ConnectorProps } from ".";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "~/components/ui/sheet";
import { Label } from "~/components/ui/label";
import { Textarea } from "~/components/ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "~/components/ui/form";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import { X } from "@phosphor-icons/react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { Input } from "~/components/ui/input";

export default function IfConnector({ workflow, draggable }: ConnectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);

  const thenWorkflows = useMemo(() => {
    // check if then is array
    if (workflow.then && Array.isArray(workflow.then)) {
      return workflow.then as EndpointWorkflow[];
    }

    return [];
  }, []);

  const elseWorkflows = useMemo(() => {
    // check if then is array
    if (workflow.else && Array.isArray(workflow.else)) {
      return workflow.else as EndpointWorkflow[];
    }

    return [];
  }, []);

  return (
    <div
      className="flex flex-col group h-min"
      ref={draggable?.setNodeRef}
      style={draggable?.style}
      {...draggable?.attributes}
      {...draggable?.listeners}
    >
      <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
        <div className="border-b border-input flex items-center justify-center">
          <h3 className="text-xs font-mono text-muted-foreground p-2">If</h3>
        </div>
        <div className="p-2">
          {workflow && (
            <div>
              <p className="text-xs font-mono text-muted-foreground text-center">
                {workflow.instruction}
              </p>
            </div>
          )}
        </div>
      </div>
      <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div>
      <div className="flex justify-center relative z-10">
        <div className="absolute text-xs font-mono text-muted-foreground self-center py-1 px-3 mt-[5px] -ml-[120px] bg-input">
          Then
        </div>
        <div className="absolute text-xs font-mono text-muted-foreground self-center py-1 px-3 mt-[5px] ml-[120px] bg-input">
          Else
        </div>
      </div>
      <div className="flex justify-center items-stretch min-h-[140px]">
        <div className="flex flex-col flex-nowrap">
          <div className="flex flex-col group min-w-[295px]">
            <div className="flex relative justify-center self-end items-start w-[calc(50%+1px)] min-h-[30px] border-solid border-t-2 border-l-2 border-input rounded-tl-lg cursor-pointer"></div>
          </div>
          {thenWorkflows.map((workflow, index) => {
            return (
              <Connector
                workflow={workflow}
                index={index}
                key={index}
                onUpdate={(workflow) => {
                  // handleUpdateWorkflows(index, workflow);
                }}
                // onRemove={() => handleRemoveWorkflow(index)}
              />
            );
          })}
          <div className="relative flex grow self-center min-h-[0px] border-solid border-l-2 border-input"></div>
          <div className="flex relative justify-center self-end items-start w-[calc(50%+1px)] min-h-[10px] border-solid border-b-2 border-l-2 border-input rounded-bl-lg"></div>
        </div>
        <div className="flex flex-col flex-nowrap">
          <div className="flex flex-col group min-w-[295px]">
            <div className="flex relative justify-center self-start items-start w-[calc(50%+1px)] min-h-[30px] border-solid border-t-2 border-r-2 border-input rounded-tr-lg cursor-pointer"></div>
          </div>
          {elseWorkflows.map((workflow, index) => {
            return (
              <Connector
                workflow={workflow}
                index={index}
                key={index}
                onUpdate={(workflow) => {
                  // handleUpdateWorkflows(index, workflow);
                }}
                // onRemove={() => handleRemoveWorkflow(index)}
              />
            );
          })}
          <div className="relative flex grow self-center min-h-[0px] border-solid border-l-2 border-input"></div>
          <div className="flex relative justify-center self-start items-start w-[calc(50%+1px)] min-h-[10px] border-solid border-b-2 border-r-2 border-input rounded-br-lg"></div>
        </div>
      </div>
      <div className="flex flex-col">
        <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer">
          {/* <div
            data-v-ba3c3b62=""
            class="relative select-none w-min"
            tab="9C2DB2BE83844190AF7EC55D9BE38E39"
          >
            <div data-v-ba3c3b62="">
              <i class="absolute -ml-[13.5px] p-2 top-[3px] -mt-1.5 fa-solid fa-circle-plus text-controlBase text-small-plus opacity-0 transition-all delay-0 duration-150 disableDemo group-hover:opacity-100"></i>
            </div>
          </div> */}
        </div>
      </div>
    </div>
  );
}
