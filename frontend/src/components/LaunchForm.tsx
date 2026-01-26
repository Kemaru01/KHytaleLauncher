import { useCallback, useEffect } from "react"
import { SubmitButton } from "./SubmitButton";

import { useAppStore } from "@/stores/useAppStore"
import { LaunchTheGame } from "../../wailsjs/go/app/App";

import hytaleLogo from "@/assets/icons/hytale-logo.png"

export const LaunchForm: React.FC = () => {
  const { 
    username, 
    brenchs,
    versions,
    setUsername
  } = useAppStore()

  const handleSubmit = useCallback((e: React.FormEvent) => {
    e.preventDefault()
    
    if(username.length < 4)
      return 

    LaunchTheGame(username, "4")
  }, [username])
  
  return (
    <form 
      className="flex flex-col items-center flex-1"
      onSubmit={handleSubmit}>
      <img 
        src={hytaleLogo}
        className="h-50 my-10 select-none"
      />
      <div className="w-100 flex flex-col justify-center flex-end flex-1">
        <input 
          type="text"
          pattern="[a-zA-Z0-9_]+"
          className="h-12.5 px-4 bg-[#203658] border-[#4e5765] border-2 outline-none gold-focus rounded-t-lg w-full"
          onChange={(e) => setUsername(e.target.value.trim())}
          value={username}
          placeholder="Kullanıcı Adınızı Girin"
        />

        <div className="flex mb-6 text-[#a5b9c6] select-none">
          <select
            className="flex-1 h-12.5 px-4 bg-[#203658] border-[#4e5765] border-2 outline-none gold-focus rounded-bl-lg w-full cursor-pointer">
            {brenchs.map(item => (
              <option key={item.toLowerCase()}>{item}</option>
            ))}
          </select>
          <select
            className="flex-1 h-12.5 px-4 bg-[#203658] border-[#4e5765] border-2 outline-none gold-focus rounded-br-lg w-full cursor-pointer">
            {versions.map(item => (
              <option key={item}>Versiyon: {item}</option>
            ))} 
          </select>
        </div>

        <SubmitButton />
      </div>
    </form>
  )
}
