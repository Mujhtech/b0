import { CaretDown, Check, X } from "@phosphor-icons/react";
import React, { useCallback, useState } from "react";
import { Button } from "~/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog";
import { Input } from "~/components/ui/input";

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "~/components/ui/sheet";
import FormField from "~/components/ui/form-field";
import { Label } from "~/components/ui/label";
import FormError from "~/components/ui/form-error";
import { Textarea } from "~/components/ui/textarea";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "~/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "~/components/ui/popover";
import { cn } from "~/lib/utils";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";

const methods = ["GET", "POST", "PUT", "DELETE", "PATCH"];

export default function Connector() {
  const [method, setMethod] = useState<string | undefined>(undefined);
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);
  return (
    <Sheet open={isOpen} onOpenChange={setIsOpen}>
      <SheetTrigger>
        <div className="flex flex-col group h-min">
          <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
            <div className="border-b border-input flex items-center justify-center">
              <h3 className="text-xs font-mono text-muted-foreground p-2">
                HTTP Request
              </h3>
            </div>
            <div className="p-2"></div>
          </div>
          <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div>
        </div>
      </SheetTrigger>
      <SheetContent
        className="bg-transparent border-none p-4"
        hideCloseBtn={true}
      >
        <div className="relative bg-background w-full h-full border border-input">
          <SheetHeader className="flex flex-row items-center justify-between border-b border-input space-y-0 p-2">
            <SheetTitle className="text-xs font-mono text-muted-foreground">
              HTTP Request
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
              <TabsTrigger value="other" className="font-mono text-xs">
                Other
              </TabsTrigger>
            </TabsList>
            <TabsContent value="general">
              <div className="flex flex-col p-2">
                <form className="flex flex-col gap-2">
                  <FormField>
                    <Label>Name</Label>
                    <Input
                      placeholder="Name"
                      // key={fields.name.key}
                      // name={fields.name.name}
                      // defaultValue={fields.name.initialValue || app?.name || ""}
                    />
                    {/* <FormError>fields.name.errors</FormError> */}
                  </FormField>
                  <FormField>
                    <Label>Description (Optional)</Label>
                    <Textarea
                      placeholder="Description (Optional)"
                      className="resize-none"
                      // key={fields.description.key}
                      // name={fields.description.name}
                      // defaultValue={
                      //   fields.description.initialValue || app?.description || ""
                      // }
                    />
                    {/* <FormError>fields.description.errors</FormError> */}
                  </FormField>
                  <FormField>
                    <Label>Url Path</Label>
                    <Input
                      placeholder="/index"
                      // key={fields.name.key}
                      // name={fields.name.name}
                      // defaultValue={fields.name.initialValue || app?.name || ""}
                    />
                    {/* <FormError>fields.path.errors</FormError> */}
                  </FormField>
                  <FormField>
                    <Label>Url Method</Label>

                    <SelectMethodDialog onChange={(val) => {}} />

                    {/* <FormError>fields.method.errors</FormError> */}
                  </FormField>
                </form>
              </div>
              <div className="bg-background border-t border-input absolute w-full bottom-0 p-2">
                <div className="flex items-center justify-end gap-2">
                  <Button
                    variant="outline"
                    className="h-8 shadow-lg border-destructive text-red-500 hover:border-destructive hover:text-red-500"
                  >
                    Delete
                  </Button>
                  <Button className="h-8 shadow-lg">Save</Button>
                </div>
              </div>
            </TabsContent>
            <TabsContent value="other">Change your password here.</TabsContent>
          </Tabs>
        </div>
      </SheetContent>
    </Sheet>
    // <Dialog>
    //   <DialogTrigger asChild>

    //   </DialogTrigger>
    //   <DialogContent className="sm:max-w-[425px]">
    //     <DialogTitle>HTTP Request</DialogTitle>
    //   </DialogContent>
    // </Dialog>
  );
}

export function SelectMethodDialog({
  value,
  onChange,
}: {
  value?: string;
  onChange: (value: string) => void;
}) {
  const [open, setOpen] = useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value ? methods.find((item) => item === value) : "Select method..."}
          <CaretDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput placeholder="Search method..." />
          <CommandList>
            <CommandEmpty>No method found.</CommandEmpty>
            <CommandGroup>
              {methods.map((method) => (
                <CommandItem
                  key={method}
                  value={method}
                  onSelect={(currentValue) => {
                    onChange(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      value === method ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {method}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
