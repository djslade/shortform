import "../styles/PrimaryLink.css";

interface PrimaryLinkProps {
  text: string;
  href: string;
}

export const PrimaryLink = ({ text, href }: PrimaryLinkProps) => {
  return (
    <a href={href} className="primary-link">
      {text}
    </a>
  );
};
