import "../styles/FooterLinks.css";

interface FooterLinksProps {
  heading: string;
  links: {
    name: string;
    url: string;
  }[];
}

export const FooterLinks = ({ heading, links }: FooterLinksProps) => {
  return (
    <div className="footer-links-container">
      <h3 className="footer-link-heading">{heading}</h3>
      <div className="inner-links-container">
        {links.map((link) => (
          <a key={link.name} href={link.url} className="footer-link">
            {link.name}
          </a>
        ))}
      </div>
    </div>
  );
};
