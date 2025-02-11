import "../styles/AppButton.css";

interface AppButtonProps {
  text: string;
  type?: "button" | "submit" | "reset";
  variant?: "small";
  tinted?: boolean;
  onClick?: () => void;
}

export const AppButton = ({
  text,
  type,
  variant,
  tinted,
  onClick,
}: AppButtonProps) => {
  if (variant == "small") {
    return (
      <button
        className={`app-btn-small ${tinted ? "app-btn-tint" : ""}`}
        type={type}
        onClick={onClick}
      >
        {text}
      </button>
    );
  }
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
