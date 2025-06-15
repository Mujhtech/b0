import React, { useCallback, useState } from "react";
import { EndpointWorkflow } from "~/models/endpoint";
import { ConnectorProps } from ".";
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

export default function GithubConnector({
  workflow,
  draggable,
}: ConnectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);
  return (
    <Sheet open={isOpen} onOpenChange={setIsOpen}>
      <SheetTrigger>
        <div
          className="flex flex-col group h-min"
          ref={draggable?.setNodeRef}
          style={draggable?.style}
          {...draggable?.attributes}
          {...draggable?.listeners}
        >
          <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
            <div className="border-b border-input flex items-center justify-center">
              <h3 className="text-xs font-mono text-muted-foreground p-2">
                Github
              </h3>
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
      </SheetTrigger>
      <SheetContent
        className="bg-transparent border-none p-4"
        hideCloseBtn={true}
      >
        <SheetDescription className="hidden"></SheetDescription>
        <div className="relative bg-background w-full h-full border border-input">
          <SheetHeader className="flex flex-row items-center justify-between border-b border-input space-y-0 p-2">
            <SheetTitle className="text-xs font-mono text-muted-foreground">
              Github
            </SheetTitle>
            <Button
              variant={"outline"}
              size={"icon"}
              onClick={handleClose}
              className="h-5 w-5 [&_svg]:size-4"
            >
              <X size={16} className="h-4 w-4" />
            </Button>
          </SheetHeader>
          <Tabs defaultValue="general" className="w-full">
            <TabsList className="w-full bg-transparent justify-start border-b border-input">
              <TabsTrigger value="general" className="font-mono text-xs">
                General
              </TabsTrigger>
            </TabsList>
            <TabsContent value="general"></TabsContent>
          </Tabs>
        </div>
      </SheetContent>
    </Sheet>
  );
}
