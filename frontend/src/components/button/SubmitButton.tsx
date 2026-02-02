import { useAppStore } from "@/stores/useAppStore";

export const SubmitButton: React.FC = () => {
  const { 
    isLauncherProcessing, 
    progressInfo, 
    username 
  } = useAppStore()
  
  return (
    <button
      type="submit"
      disabled={username.length < 4 || isLauncherProcessing}
      className="bg-[#203658] disabled:opacity-60 cursor-pointer disabled:cursor-not-allowed border-[#4e5765] 
        border-2 rounded-lg overflow-hidden relative w-full h-12.5 gold-focus">
        <div className={progressInfo.present < 0 ? "" : `bg-white/10 h-full`} style={{ width: `${progressInfo.present}%` }}></div>
        <div className={"absolute top-0 left-0 right-0 bottom-0 font-bold flex items-center justify-center"}>
          {!isLauncherProcessing ? "Oyuna Başla" : progressInfo.text}
        </div>
    </button>
  )
}