import closeIcon from "@/assets/icons/close-x.svg";
import { Quit } from "@wailsapp/runtime";
import { useCallback } from "react";

import hytaleIconH from "@/assets/icons/hytale-icon-h.png";

export const Header: React.FC = () => {
  const handleClose = useCallback(() => Quit(), [])

  return (
    <header
      className="flex h-10 bg-[#0f1428]/25 justify-between select-none">
      <div className="h-full flex items-center gap-x-2 px-4 font-light">
        <img 
          className="size-5"
          src={hytaleIconH}
        />
        <span>
          KHytale Launcher
        </span>
        <a
          onClick={() => {}}
          className="hover:text-[#6a86f8] text-sm"
          href="#">
          (by github.com/Kemaru01)
        </a>
      </div>
      <button 
        className="h-full px-4 hover:bg-red-600 cursor-pointer"
        tabIndex={-1}
        onClick={handleClose}>
          <img
            src={closeIcon}
            className="filter brightness-0 saturate-100 invert sepia size-6.5"
          />
      </button>
    </header>
  )
}