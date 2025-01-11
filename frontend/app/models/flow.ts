export interface Position {
  x: number;
  y: number;
}

export interface Node {
  id: string;
  type: "start" | "end" | "process" | "condition" | "loop" | "integration";
  position: Position;
  data: {
    label: string;
    [key: string]: any;
  };
}

export interface Connection {
  id: string;
  source: string;
  target: string;
  sourceHandle?: string;
  targetHandle?: string;
}

export interface FlowState {
  nodes: Node[];
  connections: Connection[];
}
