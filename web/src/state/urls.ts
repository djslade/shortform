import { atom } from "jotai";

interface IURL {
    id: string;
    src: string;
    dest: string;
}

export const urlsAtom = atom<IURL[]>([])