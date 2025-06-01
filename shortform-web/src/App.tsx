import "./App.css";
import { Field, Form, Formik } from "formik";
import * as Yup from "yup";
import {
  AppButton,
  FeatureCard,
  Footer,
  Header,
  HeaderMobile,
  PrimaryLink,
  URLInfo,
} from "./components";
import { useEffect } from "react";
import { useAtom } from "jotai";
import { urlsAtom } from "./state";

const CreateURLSchema = Yup.object().shape({
  dest: Yup.string()
    .required("Please add a link")
    .url("Please add a valid URL"),
});

const featuresData = [
  {
    iconsrc: "/icon-brand-recognition.svg",
    subheading: "Brand Recognition",
    text: "Boost your brand recognition with each click. Generic links don’t mean a thing. Branded links help instil confidence in your content.",
  },
  {
    iconsrc: "/icon-detailed-records.svg",
    subheading: "Detailed Records",
    text: "Gain insights into who is clicking your links. Knowing when and where people engage with your content helps inform better decisions.",
  },
  {
    iconsrc: "/icon-fully-customizable.svg",
    subheading: "Fully Customizable",
    text: "Improve brand awareness and content discoverability through customizable links, supercharging audience engagement.",
  },
];

function App() {
  const [URLs, setURLs] = useAtom(urlsAtom)

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
      const newURL = {
        id: data.url.id,
        src: data.url.dest,
      }
    } catch (err) {
      // TODO: Handle errors
      console.log(err);
    }
  };

  useEffect(() => {
    // Retrieve URLs from local storage
    const storedURLs = localStorage.getItem("shortform_urls")
    if (!storedURLs) return;
    setURLs(JSON.parse(storedURLs))
  }, [])

  return (
    <>
      <Header />
      <HeaderMobile />
      <section className="hero-section">
        <div className="hero-img-container">
          <img src="/hero-splash-bg.svg" className="hero-img" alt="" />
        </div>
        <div className="hero-container">
          <div className="hero-text-container">
            <h1 className="hero-heading">More than just shorter links</h1>
            <h2 className="hero-subheading">
              Build your brand’s recognition and get detailed insights on how
              your links are performing.
            </h2>
          </div>
          <PrimaryLink text="Get Started" href="/" />
        </div>
      </section>
      <section className="shorten-container">
        <div className="shorten-inner-container">
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
                  <Field
                    className={
                      errors.dest && touched.dest
                        ? "shorten-form-input-error"
                        : "shorten-form-input"
                    }
                    placeholder="Shorten a link here..."
                    type="text"
                    id="dest"
                    name="dest"
                  />
                  {errors.dest && touched.dest && (
                    <span className="input-error-msg">{errors.dest}</span>
                  )}
                </div>
                <AppButton
                  type="submit"
                  text={isSubmitting ? "Shortening..." : "Shorten It!"}
                />
              </Form>
            )}
          </Formik>
          <div className="urls-info-container">
            <URLInfo
              ogURL="https://www.frontendmentor.io"
              shortenedURL="https://rel.ink/k4lKyk"
            />
            <URLInfo
              ogURL="https://twitter.com/frontendmentor"
              shortenedURL="https://rel.ink/gxOXp9"
            />
            <URLInfo
              ogURL="https://www.reddit.com/r/FromSeries/comments/1if3rka/my_plan_to_capture_jasmine/"
              shortenedURL="https://rel.ink/gob3X9"
            />
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
          {featuresData.map((feature) => (
            <FeatureCard
              key={feature.subheading}
              iconsrc={feature.iconsrc}
              subheading={feature.subheading}
              text={feature.text}
            />
          ))}
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
      <Footer />
    </>
  );
}

export default App;
