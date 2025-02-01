import { useState } from "react";
import "./App.css";

function App() {
  const [dest, setDest] = useState<string>("");

  const changeDest = (evt: any) => {
    setDest(evt.target.value);
  };

  const handleCreateURL = async (dest: string) => {
    try {
      const res = await fetch("", {
        method: "POST",
        body: JSON.stringify({
          dest,
        }),
      });
      const data = await res.json();
      if (!data.message) {
        throw new Error("response is malformed");
      }
      if (res.status !== 201) {
        throw new Error(data.message);
      }
    } catch (err) {
      // TODO: Handle errors
      console.log(err);
    }
  };

  return (
    <>
      <header className="header">
        <a href="/">
          <img src="/text-logo.svg" alt="Shortform" className="text-logo" />
        </a>
        <button className="header-mobile-menu-btn">
          <img src="/mobile-menu-icon.svg" className="mobile-menu-icon" />
        </button>
      </header>
      <div className="hero-img-container">
        <img src="/hero-splash-bg.svg" className="hero-img" alt="" />
      </div>
      <div className="hero-container">
        <div className="hero-text-container">
          <h1 className="hero-heading">More than just shorter links</h1>
          <h2 className="hero-subheading">
            Build your brand’s recognition and get detailed insights on how your
            links are performing.
          </h2>
        </div>
        <a href="/" className="hero-btn">
          Get Started
        </a>
      </div>
      <section className="shorten-container">
        <form className="shorten-form">
          <div className="">
            <label htmlFor="dest"></label>
            <input
              className="shorten-form-input"
              placeholder="Shorten a link here..."
              type="url"
              id="dest"
              name="dest"
              value={dest}
              onChange={changeDest}
            />
          </div>
          <button className="shorten-form-btn">Shorten It!</button>
        </form>
      </section>
    </>
  );
}

export default App;
