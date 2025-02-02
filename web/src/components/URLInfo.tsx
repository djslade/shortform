import { AppButton } from "./AppButton";
import { copyToClipboard } from "../utils";
import { useState } from "react";
import "../styles/URLInfo.css";

interface URLInfoProps {
  ogURL: string;
  shortenedURL: string;
}

export const URLInfo = ({ ogURL, shortenedURL }: URLInfoProps) => {
  const [copied, setCopied] = useState<boolean>(false);

  const handleCopy = async () => {
    if (copied) return;
    await copyToClipboard(shortenedURL);
    setCopied(true);
    setTimeout(() => {
      setCopied(false);
    }, 5000);
  };

  return (
    <div className="urls-info">
      <div className="urls-info-top-container">
        <a className="urls-info-og-url" href={ogURL}>
          {ogURL}
        </a>
      </div>
      <div className="urls-info-bot-container">
        <a className="urls-info-shortened-url" href={shortenedURL}>
          {shortenedURL}
        </a>
        <AppButton
          text={copied ? "Copied!" : "Copy"}
          type="button"
          tinted={copied}
          onClick={handleCopy}
        />
      </div>
    </div>
  );
};
