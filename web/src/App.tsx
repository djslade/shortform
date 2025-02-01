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
        <img src="/text-logo.svg" alt="Shortform" />
        <button className="header-mobile-menu-btn">
          <img src="/mobile-menu-icon.svg" />
        </button>
      </header>
      <img src="/hero-splash-bg.svg" alt="" />
      <h1>Shortform Link Shortener</h1>
      <p>
        This website allows you to condense larger URLs into simpler, bite-sized
        links.
      </p>
      <form action="">
        <div className="">
          <label htmlFor="dest">Destination URL</label>
          <input
            type="url"
            id="dest"
            name="dest"
            value={dest}
            onChange={changeDest}
          />
        </div>
        <button>Create Link</button>
      </form>
    </>
  );
}

export default App;
