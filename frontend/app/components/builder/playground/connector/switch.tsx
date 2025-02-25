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

export default function SwitchConnector({ workflow }: ConnectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);

  console.log(workflow);

  const switchCases = useMemo(() => {
    // check if then is array
    if (workflow.cases && Array.isArray(workflow.cases)) {
      return workflow.cases;
    }

    return [];
  }, []);

  const firstCase = useMemo(() => {
    if (switchCases.length > 0) {
      return switchCases[0];
    }

    return null;
  }, [switchCases]);

  const lastCase = useMemo(() => {
    if (switchCases.length > 0) {
      return switchCases[switchCases.length - 1];
    }

    return null;
  }, [switchCases]);

  const remaingCases = useMemo(() => {
    if (switchCases.length > 2) {
      return switchCases.slice(1, switchCases.length - 1);
    }

    return switchCases;
  }, [switchCases]);

  return (
    <div className="flex flex-col group h-min">
      <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
        <div className="border-b border-input flex items-center justify-center">
          <h3 className="text-xs font-mono text-muted-foreground p-2">
            Switch
          </h3>
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
        <div className="absolute text-xs font-mono text-muted-foreground self-center py-1 px-3 mt-[5px] bg-input">
          Cases
        </div>
      </div>
      <div className="relative flex self-center min-h-[50px] border-solid border-x border-input cursor-pointer"></div>
      <div className="flex justify-center items-stretch min-h-[140px]">
        {switchCases.length > 1 && firstCase && (
          <RightCase value={firstCase.value} body={firstCase.body} />
        )}
        {remaingCases.map((i, index) => (
          <CenterCase key={index} value={i.value} body={i.body} />
        ))}
        {switchCases.length > 1 && lastCase && (
          <LeftCase value={lastCase.value} body={lastCase.body} />
        )}
      </div>
      <div className="flex flex-col">
        <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div>
      </div>
    </div>
  );
}

const CenterCase = ({ value, body }: { value?: string; body: unknown }) => {
  const workflows = useMemo(() => {
    // check if then is array
    if (body && Array.isArray(body)) {
      return body as EndpointWorkflow[];
    }

    return [];
  }, []);

  return (
    <div className="flex flex-col flex-nowrap h-auto">
      <div className="flex flex-col group min-w-[295px]">
        <div className="flex relative justify-center items-start min-h-[30px] border-base border-t-2 border-input w-full self-center after:content-[ ] after:absolute after:min-h-[30px] after:border-l-2 after:border-input after:-z-10">
          <div className="relative self-center py-1 px-3 -mt-[30px] max-w-[140px] text-xs font-mono text-muted-foreground  bg-input whitespace-nowrap text-ellipsis overflow-hidden">
            {value}
          </div>
          <div className="absolute top-2/4 -mt-[12px] cursor-pointer inset-x-1/2 -ml-[14px]">
            <div
              data-v-ba3c3b62=""
              className="relative select-none w-min"
              // tab="F503C95F1E094765858ACBEC960B5F7B"
            >
              <div data-v-ba3c3b62="">
                <i className="opacity-0 absolute p-2 fa-solid fa-circle-plus text-controlBase text-small-plus group-hover:opacity-100 transition-all delay-50 duration-250 ease-in-out"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
      {workflows.map((workflow, index) => {
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
      <div className="flex relative justify-center items-start min-h-[10px] border-b-2 border-input w-full self-center after:content-[ ] after:absolute after:min-h-[10px] after:border-l-2 after:border-input after:-z-10"></div>
    </div>
  );
};

const RightCase = ({ value, body }: { value?: string; body: unknown }) => {
  const workflows = useMemo(() => {
    // check if then is array
    if (body && Array.isArray(body)) {
      return body as EndpointWorkflow[];
    }

    return [];
  }, []);

  return (
    <div className="flex flex-col flex-nowrap h-auto">
      <div className="flex flex-col group min-w-[295px]">
        <div className="flex relative justify-center items-start min-h-[30px] border-base border-t-2 border-input w-[calc(50%+1px)] border-l-2 rounded-tl-lg self-end">
          <div className="relative self-center py-1 px-3 -mt-[30px] max-w-[140px] text-xs font-mono text-muted-foreground  bg-input whitespace-nowrap text-ellipsis overflow-hidden">
            {value}
          </div>
          <div className="absolute top-2/4 -mt-[12px] cursor-pointer left-0 -ml-[15px]">
            <div
              data-v-ba3c3b62=""
              className="relative select-none w-min"
              // tab="F503C95F1E094765858ACBEC960B5F7B"
            >
              <div data-v-ba3c3b62="">
                <i className="opacity-0 absolute p-2 fa-solid fa-circle-plus text-controlBase text-small-plus group-hover:opacity-100 transition-all delay-50 duration-250 ease-in-out"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
      {workflows.map((workflow, index) => {
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
      <div className="flex relative justify-center items-start min-h-[10px] border-input w-[calc(50%+1px)] self-end border-b-2 border-l-2 rounded-bl-lg"></div>
    </div>
  );
};

const LeftCase = ({ value, body }: { value?: string; body: unknown }) => {
  const workflows = useMemo(() => {
    // check if then is array
    if (body && Array.isArray(body)) {
      return body as EndpointWorkflow[];
    }

    return [];
  }, []);

  return (
    <div className="flex flex-col flex-nowrap h-auto">
      <div className="flex flex-col group min-w-[295px]">
        <div className="flex relative justify-center items-start min-h-[30px] w-[calc(50%+1px)] border-r-2 rounded-tr-lg self-start border-solid border-t-2 border-input">
          <div className="relative self-center py-1 px-3 -mt-[30px] max-w-[140px] text-xs font-mono text-muted-foreground  bg-input whitespace-nowrap text-ellipsis overflow-hidden">
            {value}
          </div>
          <div className="absolute top-2/4 -mt-[12px] right-0 -mr-[14px] cursor-pointer">
            <div
              data-v-ba3c3b62=""
              className="relative select-none w-min"
              // tab="F503C95F1E094765858ACBEC960B5F7B"
            >
              <div data-v-ba3c3b62="">
                <i className="opacity-0 p-2 fa-solid fa-circle-plus text-controlBase text-small-plus group-hover:opacity-100 transition-all delay-50 duration-250 ease-in-out"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
      {workflows.map((workflow, index) => {
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
      <div className="relative flex grow self-center min-h-[0px] border-solid border-x border-input"></div>
      <div className="flex relative justify-center items-start min-h-[10px] w-[calc(50%+1px)] self-start border-r-2 rounded-br-lg border-solid border-b-2 border-input"></div>
    </div>
  );
};
