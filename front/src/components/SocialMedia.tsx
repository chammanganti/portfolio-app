import "./../styles/SocialMedia.css"

const SocialMedia: React.FC<{}> = ({}) => {
  return (
    <div className="flex items-center space-x-8">
      <a
        className="social-media-link"
        href="mailto:me@chammanganti.dev"
        target="_blank"
        rel="noopener noreferrer"
      >
        Email
      </a>
      <a
        className="social-media-link"
        href="https://www.linkedin.com/in/chammanganti/"
        target="_blank"
        rel="noopener noreferrer"
      >
        LinkedIn
      </a>
      <a
        className="social-media-link"
        href="https://github.com/chammanganti"
        target="_blank"
        rel="noopener noreferrer"
      >
        GitHub
      </a>
    </div>
  );
};

export default SocialMedia;
