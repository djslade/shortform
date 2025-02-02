import "../styles/AppButton.css";

interface AppButtonProps {
  text: string;
  type?: "button" | "submit" | "reset";
  tinted?: boolean;
  onClick?: () => void;
}

export const AppButton = ({ text, type, tinted, onClick }: AppButtonProps) => {
  return (
    <button
      className={`app-btn ${tinted ? "app-btn-tint" : ""}`}
      type={type}
      onClick={onClick}
    >
      {text}
    </button>
  );
};
