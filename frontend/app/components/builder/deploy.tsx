import React from "react";
import { Button } from "../ui/button";
import {
  RocketLaunch,
  TestTube,
  CaretDown,
  Export,
} from "@phosphor-icons/react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";
import { usePlaygroundBuilder } from "./provider";

export default function DeployAndTestBtn({
  isThinking,
}: {
  isThinking?: boolean;
}) {
  const [isOpen, setIsOpen] = React.useState(false);
  const [isLoading, setIsLoading] = React.useState(false);

  const { handleProjectAction, defaultAction, setDefaultAction } =
    usePlaygroundBuilder();

  const handleActionClick = async () => {
    try {
      setIsLoading(true);
      await handleProjectAction(defaultAction);
    } catch (error) {
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex items-center gap-2 ">
      <Button
        variant="outline"
        className="h-8 shadow-lg"
        disabled={isLoading || isThinking == true}
      >
        <TestTube size={20} className="h-4 w-4" />
        Test
      </Button>
      <div className="h-8 flex items-center whitespace-nowrap rounded-md text-sm font-medium transition-colors hover:text-accent-foreground shadow-lg bg-primary text-primary-foreground hover:bg-primary/90">
        <Button
          className="h-full shadow-none capitalize"
          onClick={handleActionClick}
          disabled={isLoading || isThinking == true}
        >
          {defaultAction == "deploy" ? (
            <RocketLaunch size={20} className="h-4 w-4" />
          ) : (
            <Export size={20} className="h-4 w-4" />
          )}
          {defaultAction}
        </Button>
        <DropdownMenu onOpenChange={setIsOpen} open={isOpen}>
          <DropdownMenuTrigger
            className="h-full px-4 focus-visible:outline-none focus-visible:ring-0 focus-visible:ring-ring disabled:cursor-not-allowed"
            disabled={isLoading || isThinking == true}
          >
            <CaretDown size={20} className="h-4 w-4" />
          </DropdownMenuTrigger>
          <DropdownMenuContent side="top" align="end">
            <DropdownMenuItem
              onSelect={(e) => e.preventDefault()}
              onClick={() => {
                setDefaultAction("deploy");
                setIsOpen(false);
              }}
            >
              Deploy
            </DropdownMenuItem>
            <DropdownMenuItem
              onSelect={(e) => e.preventDefault()}
              onClick={() => {
                setDefaultAction("export");
                setIsOpen(false);
              }}
            >
              Export
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </div>
  );
}
