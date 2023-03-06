import { useEffect, useState } from "react";
import { Self } from "./logos/Self";
import SocialMedia from "./SocialMedia";

interface NavbarContentInterface {
  invisible?: boolean;
}

const NavbarContent: React.FC<NavbarContentInterface> = ({ invisible }) => {
  return (
    <div
      className={`container px-4 flex justify-between ${
        invisible === true ? "invisible" : ""
      }`}
    >
      <div
        className={`flex items-center space-x-4 ${
          invisible === true ? "invisible" : ""
        }`}
      >
        <a href="/">
          <Self />
        </a>
        <a href="/">
          <p className="font-semibold text-slate-900 hidden sm:block">
            Cham Manganti
          </p>
        </a>
      </div>
      <SocialMedia />
    </div>
  );
};

const Navbar: React.FC<{}> = ({}) => {
  const [hasScrolled, setHasScrolled] = useState<boolean>(false);

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  });

  const handleScroll = () => {
    if (window.scrollY > 30) {
      setHasScrolled(true);
    } else {
      setHasScrolled(false);
    }
  };

  return (
    <>
      <nav
        className={`fixed w-full bg-white py-4 z-50 ${
          hasScrolled === true ? "border-b" : "border-none"
        }`}
      >
        <NavbarContent />
      </nav>
      <nav className="w-full py-6">
        <NavbarContent invisible={true} />
      </nav>
    </>
  );
};

export default Navbar;
