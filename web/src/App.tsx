import { useState } from "react";
import "./App.css";
import { FaGithub } from "react-icons/fa";
import { Form, Formik } from "formik";
import * as Yup from "yup";

const CreateURLSchema = Yup.object().shape({
  dest: Yup.string()
    .required("Please add a link")
    .url("Please add a valid URL"),
});

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
        <Formik
          initialValues={{ dest: "" }}
          validationSchema={CreateURLSchema}
          onSubmit={async (values) => await handleCreateURL(values.dest)}
        >
          {({ errors, touched, isSubmitting }) => (
            <Form className="shorten-form">
              <div className="shorten-form-control">
                <label
                  htmlFor="dest"
                  aria-hidden="false"
                  className="shorten-form-label"
                >
                  Dest
                </label>
                <input
                  className={
                    errors.dest && touched.dest
                      ? "shorten-form-input-error"
                      : "shorten-form-input"
                  }
                  placeholder="Shorten a link here..."
                  type="text"
                  id="dest"
                  name="dest"
                  value={dest}
                  onChange={changeDest}
                />
                {errors.dest && touched.dest && (
                  <span className="input-error-msg">{errors.dest}</span>
                )}
              </div>
              <button className="shorten-form-btn" disabled={isSubmitting}>
                {isSubmitting ? "Shortening..." : "Shorten It!"}
              </button>
            </Form>
          )}
        </Formik>
        <div className="urls-info-container">
          <div className="urls-info">
            <div className="urls-info-top-container">
              <a className="urls-info-og-url">https://www.frontendmentor.io</a>
            </div>
            <div className="urls-info-bot-container">
              <a className="urls-info-shortened-url">https://rel.ink/k4lKyk</a>
              <button className="shorten-form-btn">Copy</button>
            </div>
          </div>
          <div className="urls-info">
            <div className="urls-info-top-container">
              <a className="urls-info-og-url">
                https://twitter.com/frontendmentor
              </a>
            </div>
            <div className="urls-info-bot-container">
              <a className="urls-info-shortened-url">https://rel.ink/gxOXp9</a>
              <button className="shorten-form-btn">Copy</button>
            </div>
          </div>
          <div className="urls-info">
            <div className="urls-info-top-container">
              <a className="urls-info-og-url">
                https://www.reddit.com/r/FromSeries/comments/1if3rka/my_plan_to_capture_jasmine/
              </a>
            </div>
            <div className="urls-info-bot-container">
              <a className="urls-info-shortened-url">https://rel.ink/gob3X9</a>
              <button className="shorten-form-btn">Copy</button>
            </div>
          </div>
        </div>
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
      <section className="boost-container">
        <div className="boost-cta-container">
          <h1 className="boost-cta-heading">Boost your links today</h1>
          <a href="/" className="hero-btn">
            Get Started
          </a>
        </div>
      </section>
      <footer className="footer">
        <img
          src="/text-logo-footer.svg"
          alt="Shortform"
          className="footer-logo"
        />
        <div className="footer-links-container">
          <h3 className="footer-link-heading">Features</h3>
          <div className="inner-links-container">
            <a href="/" className="footer-link">
              Link Shortening
            </a>
            <a href="/" className="footer-link">
              Branded Links
            </a>
            <a href="/" className="footer-link">
              Analytics
            </a>
          </div>
        </div>
        <div className="footer-links-container">
          <h3 className="footer-link-heading">Resources</h3>
          <div className="inner-links-container">
            <a href="/" className="footer-link">
              Developers
            </a>
          </div>
        </div>
        <div className="footer-links-container">
          <h3 className="footer-link-heading">Other</h3>
          <div className="inner-links-container">
            <a href="/" className="footer-link">
              About
            </a>
            <a href="/" className="footer-link">
              Contact
            </a>
          </div>
        </div>
        <div className="footer-social-container">
          <a href="/" className="footer-social">
            <FaGithub className="footer-social-icon" />
          </a>
        </div>
      </footer>
    </>
  );
}

export default App;
