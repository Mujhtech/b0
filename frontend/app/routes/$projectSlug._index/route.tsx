import React, { useCallback, useEffect, useRef, useState } from "react";
import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import invariant from "tiny-invariant";
import { typedjson, useTypedLoaderData } from "remix-typedjson";
import { z } from "zod";
import { ActionFunction, LoaderFunctionArgs, redirect } from "@remix-run/node";
import { requireUser } from "~/services/user.server";
import { Project } from "~/models/project";
import { getProject } from "~/services/project.server";
import { RouteErrorDisplay } from "~/components/error-boundary";
import { Position } from "~/models/flow";
import Playground from "~/components/builder/playground";

export const ProjectSlugParamSchema = z.object({
  projectSlug: z.string(),
});

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { projectSlug } = ProjectSlugParamSchema.parse(params);

  invariant(projectSlug, "No project found in request.");

  const user = await requireUser(request);

  let project: Project | null = null;

  try {
    project = await getProject(request, projectSlug);
  } catch (err) {
    // redirect or throw not found
    throw new Response(null, {
      status: 404,
      statusText: "Not Found",
    });
  }

  // if (!appId) {
  //   return redirect(appsPath());
  // }

  return typedjson({
    user,
    projectSlug,
    project,
  });
};

export default function Page() {
  const canvasRef = useRef<HTMLDivElement>(null);
  const [draggingNode, setDraggingNode] = useState<string | null>(null);
  const [selectedNode, setSelectedNode] = useState<Node | null>(null);
  const [connecting, setConnecting] = useState<{
    sourceId: string;
    sourceHandle?: string;
    startPos: Position;
  } | null>(null);
  const [mousePosition, setMousePosition] = useState<Position>({ x: 0, y: 0 });
  const [zoom, setZoom] = useState(25);
  const [pan, setPan] = useState<Position>({ x: 0, y: 0 });
  const [isPanning, setIsPanning] = useState(false);
  const [lastMousePos, setLastMousePos] = useState<Position>({ x: 0, y: 0 });

  const handleWheel = useCallback((e: WheelEvent) => {
    if (e.ctrlKey) {
      e.preventDefault();
      const delta = e.deltaY > 0 ? -25 : 25;
      setZoom((z) => Math.min(Math.max(z + delta, 25), 300));
    }
  }, []);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (canvas) {
      canvas.addEventListener("wheel", handleWheel, { passive: false });
      return () => canvas.removeEventListener("wheel", handleWheel);
    }
  }, [handleWheel]);

  const handleMouseDown = useCallback((e: React.MouseEvent) => {
    if (e.button === 1 || e.button === 2) {
      setIsPanning(true);
      setLastMousePos({ x: e.clientX, y: e.clientY });
    }
  }, []);

  const handleMouseMove = useCallback(
    (e: React.MouseEvent) => {
      setMousePosition({ x: e.clientX, y: e.clientY });

      if (isPanning) {
        const dx = e.clientX - lastMousePos.x;
        const dy = e.clientY - lastMousePos.y;
        setPan((p) => ({ x: p.x + dx, y: p.y + dy }));
        setLastMousePos({ x: e.clientX, y: e.clientY });
      }
    },
    [isPanning, lastMousePos]
  );

  const handleMouseUp = useCallback(() => {
    setIsPanning(false);
    if (connecting) {
      setConnecting(null);
    }
  }, [connecting]);

  const handleNodeDragStart = useCallback((nodeId: string) => {
    setDraggingNode(nodeId);
  }, []);

  const handleNodeDrag = useCallback(
    (e: React.MouseEvent, nodeId: string) => {
      if (!draggingNode || !canvasRef.current) return;

      // const canvas = canvasRef.current.getBoundingClientRect();
      // const newNodes = nodes.map((node) => {
      //   if (node.id === nodeId) {
      //     return {
      //       ...node,
      //       position: {
      //         x: (e.clientX - canvas.left - pan.x) / zoom,
      //         y: (e.clientY - canvas.top - pan.y) / zoom,
      //       },
      //     };
      //   }
      //   return node;
      // });

      // onNodesChange(newNodes);
    },
    [draggingNode, pan, zoom]
  );

  const handleNodeDragEnd = useCallback(() => {
    setDraggingNode(null);
  }, []);

  const handleConnectionStart = useCallback(
    (nodeId: string, handle: string, position: Position) => {
      setConnecting({
        sourceId: nodeId,
        sourceHandle: handle,
        startPos: position,
      });
    },
    []
  );

  const handleConnectionEnd = useCallback(
    (targetId: string, handle?: string) => {
      // if (!connecting || connecting.sourceId === targetId) return;
      // const existingConnection = connections.find(
      //   (conn) =>
      //     conn.source === connecting.sourceId && conn.target === targetId
      // );
      // if (!existingConnection) {
      //   const newConnection: Connection = {
      //     id: `${connecting.sourceId}-${targetId}`,
      //     source: connecting.sourceId,
      //     target: targetId,
      //     sourceHandle: connecting.sourceHandle,
      //     targetHandle: handle,
      //   };
      //   onConnectionsChange([...connections, newConnection]);
      // }
      // setConnecting(null);
    },
    [connecting]
  );

  const handleNodeSelect = useCallback((node: Node) => {
    setSelectedNode(node);
  }, []);

  const handleNodeConfigChange = useCallback((updatedNode: Node) => {
    // const newNodes = nodes.map((node) =>
    //   node.id === updatedNode.id ? updatedNode : node
    // );
    // onNodesChange(newNodes);
    // setSelectedNode(updatedNode);
  }, []);

  const handleDeleteConnection = useCallback((connectionId: string) => {
    // onConnectionsChange(
    //   connections.filter((conn) => conn.id !== connectionId)
    // );
  }, []);

  const handleZoomOut = useCallback(() => {
    setZoom((z) => {
      if (z <= 25) return z;

      return z - 25;
    });
  }, [zoom, setZoom]);

  const handleZoomIn = useCallback(() => {
    setZoom((z) => {
      if (z >= 300) return z;

      return z + 25;
    });
  }, [zoom, setZoom]);

  const handleResetZoom = useCallback(() => {
    setZoom(100);
    setPan({ x: 0, y: 0 });
  }, [zoom, setZoom, pan, setPan]);

  const handleZoomInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setZoom(parseFloat(e.target.value));
    },
    []
  );

  return (
    <main className="h-full w-full relative canvas-bg">
      <div className="w-full h-full overflow-y-auto flex-1 pointer-events-auto scrollbar-none">
        <div className="h-[1500px] w-full transition-all">
          <div className="translate-x-[-5160px] flex w-[10000px] h-[10000px] left-1/2 relative p-4 pt-10 justify-center transition-all overflow-x-hidden">
            <div
              ref={canvasRef}
              onMouseDown={handleMouseDown}
              onMouseMove={handleMouseMove}
              onMouseUp={handleMouseUp}
              onMouseLeave={handleMouseUp}
              onContextMenu={(e) => e.preventDefault()}
              className="flex select-none w-[10000px] h-[10000px]  absolute inset-0 flex-nowrap flex-col justify-start"
              style={{
                transform: `translate(${pan.x}px, ${pan.y}px) scale(${zoom})`,
                transformOrigin: "0 0",
              }}
            >
              <Playground />
            </div>
          </div>
        </div>
      </div>

      <div className="absolute bottom-4 w-full">
        <div className="flex justify-between">
          <div className="flex items-center gap-2 ml-4">
            <UserMenu />
            <ZoomInAndOut
              zoom={zoom}
              handleZoomIn={handleZoomIn}
              handleZoomOut={handleZoomOut}
              handleResetZoom={handleResetZoom}
              handleZoomInputChange={handleZoomInputChange}
            />
          </div>
          <div>
            <AskB0 />
          </div>
          <div className="mr-4">
            <DeployAndTestBtn />
          </div>
        </div>
      </div>
      <BuilderTools />
    </main>
  );
}

export function ErrorBoundary() {
  return <RouteErrorDisplay className="canvas-bg" />;
}
