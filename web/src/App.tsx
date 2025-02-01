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
      <section className="features-container">
        <div className="features-heading-container">
          <h1 className="features-heading">Advanced Statistics</h1>
          <h3 className="features-text">
            Track how your links are performing across the web with our advanced
            statistics dashboard.
          </h3>
        </div>
        <div className="features-cards-container">
          <div className="features-info-container">
            <div className="features-icon-container">
              <img
                src="/icon-brand-recognition.svg"
                alt=""
                className="features-icon"
              />
            </div>
            <h2 className="features-subeading">Brand Recognition</h2>
            <h3 className="features-text">
              Boost your brand recognition with each click. Generic links don’t
              mean a thing. Branded links help instil confidence in your
              content.
            </h3>
          </div>
          <div className="features-info-container">
            <div className="features-icon-container">
              <img
                src="/icon-detailed-records.svg"
                alt=""
                className="features-icon"
              />
            </div>
            <h2 className="features-subeading">Detailed Records</h2>
            <h3 className="features-text-small">
              Gain insights into who is clicking your links. Knowing when and
              where people engage with your content helps inform better
              decisions.
            </h3>
          </div>
          <div className="features-info-container">
            <div className="features-icon-container">
              <img
                src="/icon-fully-customizable.svg"
                alt=""
                className="features-icon"
              />
            </div>
            <h2 className="features-subeading">Fully Customizable</h2>
            <h3 className="features-text-small">
              Improve brand awareness and content discoverability through
              customizable links, supercharging audience engagement.
            </h3>
          </div>
        </div>
      </section>
    </>
  );
}

export default App;
