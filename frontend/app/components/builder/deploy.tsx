import React from "react";
import { Button } from "../ui/button";
import { RocketLaunch, TestTube } from "@phosphor-icons/react";

export default function DeployAndTestBtn() {
  return (
    <div className="flex items-center gap-2">
      <Button variant="outline" className="h-8 shadow-lg">
        <TestTube size={20} className="h-4 w-4" />
        Test
      </Button>
      <Button className="h-8 shadow-lg">
        <RocketLaunch size={20} className="h-4 w-4" />
        Deploy
      </Button>
    </div>
  );
}
