import { atom } from "jotai";

interface IURL {
  id: string;
  dest: string;
}

// Comments
export const urlsAtom = atom<IURL[]>([]);
