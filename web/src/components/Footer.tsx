import { FaGithub } from "react-icons/fa";
import { FooterLinks } from "./FooterLinks";
import "../styles/Footer.css";

const footerLinksData = [
  {
    heading: "Features",
    links: [
      { name: "Link Shortening", url: "/" },
      { name: "Branded Links", url: "/" },
      { name: "Analytics", url: "/" },
    ],
  },
  {
    heading: "Resources",
    links: [
      {
        name: "Developers",
        url: "/",
      },
    ],
  },
  {
    heading: "Other",
    links: [
      {
        name: "About",
        url: "/",
      },
      {
        name: "Contact",
        url: "/",
      },
    ],
  },
];

export const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-inner-container">
        <div className="footer-left">
          <img
            src="/text-logo-footer.svg"
            alt="Shortform"
            className="footer-logo"
          />
        </div>
        <div className="footer-right">
          {footerLinksData.map((fl) => (
            <FooterLinks
              key={fl.heading}
              heading={fl.heading}
              links={fl.links}
            />
          ))}
          <div className="footer-social-container">
            <a href="/" className="footer-social">
              <FaGithub className="footer-social-icon" />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};
