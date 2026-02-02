import folderIcon from "@/assets/icons/folder-icon.svg";
import { OpenToDir } from "../../../wailsjs/go/app/App";
import { useCallback } from "react";

export default function FilesIcon() {
  const handleOpenAppDir = useCallback(() => OpenToDir(), [])

  return (
    <div className="absolute z-10 -translate-x-14.5">
      <button 
        onClick={handleOpenAppDir}
        className="text-2xl h-12.5 w-12.5 bg-[#203658] disabled:opacity-60 cursor-pointer disabled:cursor-not-allowed border-[#4e5765] 
          border-2 rounded-lg flex justify-center items-center gold-focus outline-none">
        <img
          className="select-none size-8 filter brightness-0 saturate-100 invert sepia"
          src={folderIcon}
        />
      </button>
    </div>
  )
}
