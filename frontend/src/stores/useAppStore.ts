import { EventsOff, EventsOn } from "@wailsapp/runtime";
import { create } from "zustand";

interface IAppStore {
  username: string;

  currentBrench?: string;
  currentVersion?: string;

  brenchs: string[];
  versions: string[];

  progressInfo: {
    text: string;
    present: number;
  },

  isGameRunning: boolean;
  isLauncherProcessing: boolean; 

  initialized: () => () => void;
  setUsername: (username: string) => void;
}

export const useAppStore = create<IAppStore>((set) => ({
  username: "",

  brenchs: ["Release"],
  versions: ["4"],

  progressInfo: {
    text: "",
    present: -1,
  },

  isGameRunning: false,
  isLauncherProcessing: false,

  setUsername: (username) => set(() => { 
    localStorage.setItem("launcher:username", username)
    return { username } 
  }),

  initialized: () => {
    set((state) => ({
      username: localStorage.getItem("launcher:username") ?? "",
      currentBrench: state.brenchs[0],
      currentVersion: state.versions[0]
    }))

    EventsOn("progress:status", (text: string, present: number) => {
      set({ isLauncherProcessing: text.length > 0, progressInfo: { text, present } })
    })

    return () => 
      EventsOff("progress:status")
  }
}))