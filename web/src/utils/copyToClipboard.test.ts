import { copyToClipboard } from "./copyToClipboard";
import { beforeEach, describe, expect, it, vi } from "vitest";

beforeEach(() => {
  Object.assign(navigator, {
    clipboard: {
      writeText: vi.fn(),
    },
  });
});

describe("Happy path", () => {
  it("should call the write to clipboard method", () => {
    const spy = vi.spyOn(navigator.clipboard, "writeText");
    copyToClipboard("https://example.com");
    expect(spy).toHaveBeenCalledTimes(1);
  });
});
