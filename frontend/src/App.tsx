import { useEffect, useMemo } from "react";
import { getRandomBgImage } from "@/mocks/backgrounds";
import { Header } from "@/components/Header";
import { LaunchForm } from "@/components/LaunchForm";
import { useAppStore } from "./stores/useAppStore";

const App: React.FC = () => {
  const bgUrl = useMemo(() => getRandomBgImage(), [])

  useEffect(() => {
    return useAppStore.getState()
      .initialized()
  }, [])

  return (
    <div className="overflow-hidden flex">
      <img 
        className="object-cover h-screen w-screen select-none"
        src={bgUrl} 
      />
      <div className="flex flex-col absolute h-screen w-screen bg-linear-to-t from-[#0f1428] to-transparent z-10 text-white">
        <Header />
        <LaunchForm />
      </div>
    </div>
  )
}

export default App;