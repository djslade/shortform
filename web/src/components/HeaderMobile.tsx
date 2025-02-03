import { useState } from "react";
import "../styles/HeaderMobile.css";

export const HeaderMobile = () => {
  const [mobileMenuIsOpen, setMobileMenuIsOpen] = useState<boolean>(false);

  const handleOpenMobileMenu = () => {
    setMobileMenuIsOpen(true);
  };

  return (
    <header className="header">
      <a href="/">
        <img src="/text-logo.svg" alt="Shortform" className="text-logo" />
      </a>
      <button className="header-mobile-menu-btn" onClick={handleOpenMobileMenu}>
        <img src="/mobile-menu-icon.svg" className="mobile-menu-icon" />
      </button>
      {mobileMenuIsOpen && (
        <div className="mobile-menu">
          <div className="mobile-menu-links-container">
            <a href="/" className="mobile-menu-link">
              Features
            </a>
            <a href="/" className="mobile-menu-link">
              Pricing
            </a>
            <a href="/" className="mobile-menu-link">
              Resources
            </a>
          </div>
          <div className="mobile-menu-auth-container">
            <a href="/" className="mobile-menu-link">
              Login
            </a>
            <a href="/" className="mobile-menu-signup-link">
              Sign Up
            </a>
          </div>
        </div>
      )}
    </header>
  );
};
