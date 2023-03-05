import { SelfWhite } from "./logos/Self";
import "../styles/Footer.css";

const Footer: React.FC<{}> = ({}) => {
  const year = new Date().getFullYear();

  return (
    <footer className="bg-primary-700 py-16 text-white">
      <div className="container px-4">
        <div className="flex justify-between items-center">
          <div className="space-y-2">
            <SelfWhite />
            <p className="text-sm">&copy; {year} Cham Manganti</p>
          </div>
          <div>
            <ul className="flex items-center space-x-8">
              <li>
                <a
                  className="link"
                  href="mailto:chammanganti@gmail.com"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Email
                </a>
              </li>
              <li>
                <a
                  className="link"
                  href="https://www.linkedin.com/in/chammanganti/"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  LinkedIn
                </a>
              </li>
              <li>
                <a
                  className="link"
                  href="https://github.com/chammanganti"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  GitHub
                </a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
