import { ArrowsIn, Minus, Plus, ArrowsOut } from "@phosphor-icons/react";
import { Input } from "../ui/input";
import { usePlayground } from "./playground/provider";

export default function ZoomInAndOut() {
  const {
    handleZoomIn,
    handleZoomOut,
    handleResetZoom,
    handleZoomInputChange,
    zoom,
  } = usePlayground();

  return (
    <div className="flex items-center gap-2">
      <div className="flex border border-input h-8 bg-background shadow-lg">
        <button
          type="button"
          onClick={handleZoomOut}
          className="px-2 border-r border-input"
        >
          <Minus size={28} className="h-4 w-4" />
        </button>
        <Input
          className="h-8 border-none focus-visible:ring-0 w-12 max-w-12 px-2 items-center justify-center text-center"
          type="number"
          value={zoom}
          onChange={handleZoomInputChange}
        />
        <button
          type="button"
          onClick={handleZoomIn}
          className="px-2 border-l border-input"
        >
          <Plus size={28} className="h-4 w-4" />
        </button>
      </div>
      <button
        type="button"
        onClick={handleResetZoom}
        disabled={zoom === 100}
        className="px-2 bg-background h-8 w-8 border border-input shadow-lg disabled:cursor-not-allowed"
      >
        <ArrowsIn size={28} className="h-4 w-4" />
      </button>
    </div>
  );
}
