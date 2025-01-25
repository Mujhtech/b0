import {
  createContext,
  ReactNode,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import { Position } from "~/models/flow";

interface PlaygroundProviderProps {
  canvasRef: React.RefObject<HTMLDivElement>;
  draggingNode: string | null;
  selectedNode: Node | null;
  connecting: {
    sourceId: string;
    sourceHandle?: string;
    startPos: Position;
  } | null;
  mousePosition: Position;
  zoom: number;
  pan: Position;
  isPanning: boolean;
  lastMousePos: Position;
  handleWheel: (e: WheelEvent) => void;
  handleMouseDown: (e: React.MouseEvent) => void;
  handleMouseMove: (e: React.MouseEvent) => void;
  handleMouseUp: () => void;
  handleDrag: (e: React.MouseEvent) => void;
  handleNodeDragStart: (nodeId: string) => void;
  handleNodeDrag: (e: React.MouseEvent, nodeId: string) => void;
  handleNodeDragEnd: () => void;
  handleConnectionStart: (
    nodeId: string,
    handle: string,
    position: Position
  ) => void;
  handleConnectionEnd: (targetId: string, handle?: string) => void;
  handleNodeSelect: (node: Node) => void;
  handleNodeConfigChange: (updatedNode: Node) => void;
  handleDeleteConnection: (connectionId: string) => void;
  handleZoomOut: () => void;
  handleZoomIn: () => void;
  handleResetZoom: () => void;
  handleZoomInputChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  expandToolPanel: boolean;
  handleExpandToolPanel: () => void;
  handleSetIsPanning: () => void;
}

const PlaygroundProviderContext = createContext<
  PlaygroundProviderProps | undefined
>(undefined);

export const PlaygroundProvider = ({ children }: { children: ReactNode }) => {
  const canvasRef = useRef<HTMLDivElement>(null);
  const [draggingNode, setDraggingNode] = useState<string | null>(null);
  const [selectedNode, setSelectedNode] = useState<Node | null>(null);
  const [connecting, setConnecting] = useState<{
    sourceId: string;
    sourceHandle?: string;
    startPos: Position;
  } | null>(null);
  const [mousePosition, setMousePosition] = useState<Position>({ x: 0, y: 0 });
  const [zoom, setZoom] = useState(100);
  const [pan, setPan] = useState<Position>({ x: 39.15, y: 55.5222 });
  const [isPanning, setIsPanning] = useState(false);
  const [lastMousePos, setLastMousePos] = useState<Position>({ x: 0, y: 0 });
  const [expandToolPanel, setExpandToolPanel] = useState(false);

  const handleWheel = useCallback(
    (e: WheelEvent) => {
      if (e.ctrlKey) {
        e.preventDefault();
        const delta = e.deltaY > 0 ? -25 : 25;
        setZoom((z) => Math.min(Math.max(z + delta, 25), 300));
      }
    },
    [zoom, setZoom]
  );

  useEffect(() => {
    const canvas = canvasRef.current;
    if (canvas) {
      canvas.addEventListener("wheel", handleWheel, { passive: false });
      return () => canvas.removeEventListener("wheel", handleWheel);
    }
  }, [handleWheel]);

  const handleMouseDown = useCallback((e: React.MouseEvent) => {
    if (e.button === 1 || e.button === 2) {
      //setIsPanning(true);
      setLastMousePos({ x: e.clientX, y: e.clientY });
    }
  }, []);

  const handleMouseMove = useCallback(
    (e: React.MouseEvent) => {
      setMousePosition({ x: e.clientX, y: e.clientY });

      // if (isPanning) {
      //   const dx = e.clientX - lastMousePos.x;
      //   const dy = e.clientY - lastMousePos.y;
      //   setPan((p) => ({ x: p.x + dx, y: p.y + dy }));
      //   setLastMousePos({ x: e.clientX, y: e.clientY });
      // }
    },
    [isPanning, lastMousePos]
  );

  const handleDrag = useCallback(
    (e: React.MouseEvent) => {
      const dx = e.clientX - lastMousePos.x;
      const dy = e.clientY - lastMousePos.y;

      console.log("dx", dx, "dy", dy);
    },
    [pan, zoom]
  );

  const handleMouseUp = useCallback(() => {
    // setIsPanning(false);
    // if (connecting) {
    //   setConnecting(null);
    // }
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
      return Math.max(z - 25, 25);
    });
  }, []);

  const handleZoomIn = useCallback(() => {
    setZoom((z) => {
      return Math.min(z + 25, 300);
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
    [zoom, setZoom]
  );

  const handleExpandToolPanel = useCallback(() => {
    setExpandToolPanel(!expandToolPanel);
  }, [expandToolPanel, setExpandToolPanel]);

  const handleSetIsPanning = useCallback(() => {
    setIsPanning(!isPanning);
  }, [isPanning]);

  const value = useMemo<PlaygroundProviderProps>(
    () => ({
      canvasRef,
      draggingNode,
      selectedNode,
      connecting,
      mousePosition,
      zoom,
      pan,
      isPanning,
      lastMousePos,
      expandToolPanel,
      handleWheel,
      handleMouseDown,
      handleMouseMove,
      handleSetIsPanning,
      handleMouseUp,
      handleNodeDragStart,
      handleNodeDrag,
      handleNodeDragEnd,
      handleConnectionStart,
      handleConnectionEnd,
      handleNodeSelect,
      handleNodeConfigChange,
      handleDeleteConnection,
      handleZoomOut,
      handleZoomIn,
      handleResetZoom,
      handleZoomInputChange,
      handleExpandToolPanel,
      handleDrag,
    }),
    [
      canvasRef,
      draggingNode,
      selectedNode,
      connecting,
      mousePosition,
      zoom,
      pan,
      isPanning,
      lastMousePos,
      expandToolPanel,
      handleWheel,
      handleMouseDown,
      handleDrag,
      handleSetIsPanning,
      handleMouseMove,
      handleMouseUp,
      handleNodeDragStart,
      handleNodeDrag,
      handleNodeDragEnd,
      handleConnectionStart,
      handleConnectionEnd,
      handleNodeSelect,
      handleNodeConfigChange,
      handleDeleteConnection,
      handleZoomOut,
      handleZoomIn,
      handleResetZoom,
      handleZoomInputChange,
      handleExpandToolPanel,
    ]
  );

  return (
    <PlaygroundProviderContext.Provider value={value}>
      {children}
    </PlaygroundProviderContext.Provider>
  );
};

export const usePlayground = () => {
  const context = useContext(PlaygroundProviderContext);
  if (!context) {
    throw new Error("usePlayground must be used within an PlaygroundProvider");
  }
  return context;
};
