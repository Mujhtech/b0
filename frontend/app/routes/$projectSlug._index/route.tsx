import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools, { ToolCard, TOOLS } from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import { PlaygroundProvider } from "~/components/builder/playground/provider";
import Playground from "~/components/builder/playground/playground";
import BuilderMenu from "~/components/builder/builder-menu";
import {
  PlaygroundBuilderProvider,
  usePlaygroundBuilder,
} from "~/components/builder/provider";
import { useSSE } from "~/hooks/use-ses";
import { useAuthToken, usePlatformUrl, useUser } from "~/hooks/use-user";
import { useOptionalEndpoint, useProject } from "~/hooks/use-project";
import { useCallback, useMemo, useState } from "react";
import { Spinner } from "@phosphor-icons/react";
import { useNavigate } from "@remix-run/react";
import { LogPreviewDialog } from "~/components/builder/log-preview-dialog";
import {
  DndContext,
  DragOverlay,
  useSensors,
  useSensor,
  PointerSensor,
  closestCenter,
  useDroppable,
  type DragStartEvent,
  type DragEndEvent,
  type DragOverEvent,
} from "@dnd-kit/core";
import { EndpointWorkflow, EndpointWorkflowType } from "~/models/endpoint";
import { toast } from "sonner";
import { ToastUI } from "~/components/custom-toast";
import Connector from "~/components/builder/playground/connector";

export default function Page() {
  return (
    <PlaygroundBuilderProvider>
      <PlaygroundProvider>
        <PageContent />
      </PlaygroundProvider>
    </PlaygroundBuilderProvider>
  );
}

const PageContent = () => {
  const project = useProject();
  const {
    isThinking,
    contextMessage,
    error,
    isSaving,
    handleUpdateEndpointWorkflow,
    setIsSaving,
  } = usePlaygroundBuilder();
  const user = useUser();

  const [activeCard, setActiveCard] = useState<string | null>(null);
  const [activeTool, setActiveTool] = useState<string | null>(null);
  const [activeSlot, setActiveSlot] = useState<string | null>(null);
  const endpoint = useOptionalEndpoint();

  const [newWorkflows, setNewWorkflows] = useState<
    Array<EndpointWorkflow> | undefined
  >(undefined);

  const workflows = useMemo(() => {
    if (newWorkflows) {
      return newWorkflows;
    }

    return endpoint?.workflows || [];
  }, [newWorkflows, endpoint, setNewWorkflows]);

  const handleUpdateEndpoint = useCallback(
    (updatedWorkflows: EndpointWorkflow[]) => {
      if (!endpoint) return;

      setTimeout(() => {
        setIsSaving(true);
        handleUpdateEndpointWorkflow(endpoint.id, updatedWorkflows)
          .catch((err) => {
            if (err.error) {
              toast.custom(
                (t) => (
                  <ToastUI
                    variant={"error"}
                    message={err.error}
                    t={t as string}
                  />
                ),
                {}
              );
            }
          })
          .finally(() => {
            setIsSaving(false);
          });
      }, 2500);
    },
    [endpoint, handleUpdateEndpointWorkflow, setIsSaving]
  );

  const handleUpdateWorkflows = (index: number, workflow: EndpointWorkflow) => {
    const newWorkflows = workflows;

    newWorkflows[index] = workflow;

    setNewWorkflows(newWorkflows);

    if (!endpoint) {
      return;
    }

    setTimeout(() => {
      setIsSaving(true);
      handleUpdateEndpointWorkflow(endpoint.id, newWorkflows)
        .catch((err) => {
          if (err.error) {
            toast.custom(
              (t) => (
                <ToastUI
                  variant={"error"}
                  message={err.error}
                  t={t as string}
                />
              ),
              {}
            );
          }
        })
        .finally(() => {
          setIsSaving(false);
        });
    }, 2500);
  };

  const handleRemoveWorkflow = useCallback(
    (index: number) => {
      console.log(index);

      setNewWorkflows((old) => {
        const prev = old || workflows;

        if (!prev) {
          return undefined;
        }

        const updatedWorkflows = [...prev];

        updatedWorkflows.splice(index, 1);

        return updatedWorkflows;
      });
    },
    [newWorkflows, workflows]
  );

  const messageToDisplay = useMemo(() => {
    if (contextMessage != undefined) {
      return contextMessage;
    }

    if (error != undefined) {
      return error;
    }

    if (isThinking) {
      return "Thinking...";
    }

    if (isSaving) {
      return "Saving...";
    }

    return undefined;
  }, [contextMessage, error, isThinking, isSaving]);

  const showLoaderIcon = useMemo(() => {
    return isThinking || isSaving;
  }, [isSaving, isThinking]);

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        distance: 5,
      },
    })
  );

  const activeWorkflow = useMemo(() => {
    if (activeCard) {
      return workflows.find((workflow) => workflow.action_id === activeCard);
    }

    return null;
  }, [activeCard, workflows]);

  // Handle drag start
  function handleDragStart(event: DragStartEvent) {
    const { active } = event;
    const cardId = active.id as string;

    console.log("cardId", cardId);

    if (cardId.startsWith("tool-")) {
      setActiveTool(cardId);
    } else {
      setActiveCard(cardId);
    }
  }

  function handleDragOver(event: DragOverEvent) {
    const { over } = event;
    if (over) {
      const overId = over.id as string;
      console.log("overId", overId);
      setActiveSlot(overId);
    } else {
      setActiveSlot(null);
    }
  }

  const handleDragEnd = useCallback(
    (event: DragEndEvent) => {
      const { active, over } = event;

      if (over && active.id !== over.id) {
        const cardId = active.id as string;
        const targetSlotId = over.id as string;

        const targetWorkflow = workflows.find(
          (workflow) => workflow.action_id == targetSlotId
        );

        console.log(targetWorkflow);

        if (!targetWorkflow) {
          return;
        }

        // Extract the index from the slot ID (e.g., "slot-1" -> 1)
        const targetIndex = parseInt(targetSlotId.split("-")[1]);

        // Handle reordering of existing workflows
        if (!cardId.startsWith("tool-")) {
          setNewWorkflows((old) => {
            const prev = workflows;
            const sourceIndex = parseInt(cardId.split("-")[1]);
            const updatedWorkflows = [...prev];

            // Remove the workflow from its original position
            const [movedWorkflow] = updatedWorkflows.splice(sourceIndex, 1);
            // Insert it at the new position
            updatedWorkflows.splice(targetIndex, 0, movedWorkflow);

            handleUpdateEndpoint(updatedWorkflows);
            return updatedWorkflows;
          });
        }
        // Handle adding new workflow from tools
        else {
          const toolType = cardId.replace("tool-", "");

          // Create a new workflow based on the tool type
          const newWorkflow: EndpointWorkflow = {
            action_id: new Date().getTime().toString(),
            type: toolType as EndpointWorkflowType,
            name: TOOLS[cardId as keyof typeof TOOLS].name,
            instruction: "",
            method: "GET",
            url: "",
            value: undefined,
            cases: [],
            then: undefined,
            else: undefined,
            body: undefined,
            model: undefined,
            provider: undefined,
            status: undefined,
          };

          // Add the new workflow to the list after the target card
          setNewWorkflows((old) => {
            const prev = workflows;
            const updatedWorkflows = [...prev];
            // If targetIndex is 0, insert at index 1 to preserve request tool
            const insertIndex = targetIndex === 0 ? 1 : targetIndex + 1;
            updatedWorkflows.splice(insertIndex, 0, newWorkflow);

            handleUpdateEndpoint(updatedWorkflows);
            return updatedWorkflows;
          });
        }
      }

      // Reset active states
      setActiveTool(null);
      setActiveCard(null);
      setActiveSlot(null);
    },
    [
      workflows,
      handleUpdateEndpointWorkflow,
      endpoint,
      setNewWorkflows,
      newWorkflows,
      handleUpdateEndpoint,
    ]
  );

  return (
    <main className="h-full w-full relative canvas-bg">
      <DndContext
        sensors={sensors}
        collisionDetection={closestCenter}
        onDragStart={handleDragStart}
        onDragOver={handleDragOver}
        onDragEnd={handleDragEnd}
      >
        <Playground
          workflows={workflows}
          handleRemoveWorkflow={handleRemoveWorkflow}
          activeSlot={activeSlot}
          activeCard={activeCard}
          handleUpdateWorkflows={handleUpdateWorkflows}
        />

        <div className="absolute bottom-4 w-full z-10">
          {messageToDisplay != undefined && (
            <div className="flex items-center justify-center gap-2 mb-1">
              {showLoaderIcon && (
                <Spinner
                  size={18}
                  className="text-muted-foreground animate-spin"
                />
              )}
              <p className="font-mono text-[10px] text-muted-foreground text-center">
                {messageToDisplay}
              </p>
            </div>
          )}
          <div className="flex justify-between">
            <div className=" ml-4">
              <div className="flex items-center gap-2">
                <UserMenu user={user} showHomepage={true} />
                <ZoomInAndOut />
              </div>
            </div>
            <div className="flex flex-col gap-0.5">
              <AskB0 isThinking={isThinking} projectModel={project.model} />
              <div className="flex items-center justify-center">
                <p className="font-mono text-[10px] text-muted-foreground text-center">
                  b0 can make mistakes. Please double-check it.
                </p>
              </div>
            </div>
            <div className="mr-4">
              <DeployAndTestBtn isThinking={isThinking} />
            </div>
          </div>
        </div>
        <BuilderMenu />
        <BuilderTools />
        <LogPreviewDialog />

        <DragOverlay>
          {activeTool ? (
            <ToolCard id={activeTool}>
              {TOOLS[activeTool as keyof typeof TOOLS].icon}
              {TOOLS[activeTool as keyof typeof TOOLS].name}
            </ToolCard>
          ) : null}
          {activeCard && activeWorkflow ? (
            <Connector
              workflow={activeWorkflow}
              index={0}
              key={-1}
              onUpdate={(workflow) => {
                // handleUpdateWorkflows(index, workflow);
              }}
              onRemove={() => {}}
              draggable={{
                isActive: false,
              }}
            />
          ) : null}
        </DragOverlay>
      </DndContext>
    </main>
  );
};
