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

export default function SwitchConnector({ workflow }: ConnectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);

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
