interface FeatureCardProps {
  iconsrc: string;
  subheading: string;
  text: string;
}

export const FeatureCard = ({
  iconsrc,
  subheading,
  text,
}: FeatureCardProps) => {
  return (
    <div className="features-info-container">
      <div className="features-icon-container">
        <img src={iconsrc} alt={subheading} className="features-icon" />
      </div>
      <h2 className="features-subeading">{subheading}</h2>
      <h3 className="features-text-small">{text}</h3>
    </div>
  );
};
