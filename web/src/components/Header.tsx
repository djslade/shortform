import "../styles/Header.css";

export const Header = () => {
  return (
    <header className="header-big">
      <div className="header-left">
        <a href="/">
          <img src="/text-logo.svg" alt="Shortform" />
        </a>
        <nav className="header-nav">
          <a href="/" className="header-link">
            Features
          </a>
          <a href="/" className="header-link">
            About
          </a>
          <a href="/" className="header-link">
            Developers
          </a>
        </nav>
      </div>
      <div className="header-right">
        <a href="/" className="header-link">
          Login
        </a>
        <a href="/" className="header-primary-link">
          Sign Up
        </a>
      </div>
    </header>
  );
};
