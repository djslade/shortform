import "../styles/Header.css";

export const Header = () => {
  return (
    <header className="header">
      <a href="/">
        <img src="/text-logo.svg" alt="Shortform" className="text-logo" />
      </a>
      <button className="header-mobile-menu-btn">
        <img src="/mobile-menu-icon.svg" className="mobile-menu-icon" />
      </button>
    </header>
  );
};
