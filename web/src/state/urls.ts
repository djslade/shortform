import { atom } from "jotai";

interface IURL {
  id: string;
  dest: string;
}

export const urlsAtom = atom<IURL[]>([]);
