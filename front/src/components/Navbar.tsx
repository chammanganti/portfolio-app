import { useEffect, useState } from "react";
import { Self } from "./logos/Self";
import SocialMedia from "./SocialMedia";
import ThemePicker from "./ThemePicker";

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
          <p className="font-semibold text-slate-900 dark:text-slate-100 hidden sm:block">
            Cham Manganti
          </p>
        </a>
      </div>
      <div className="flex items-center space-x-12">
        <SocialMedia />
        <ThemePicker />
      </div>
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
        className={`fixed w-full bg-white dark:bg-slate-900 py-4 z-50 ${
          hasScrolled === true
            ? "border-b dark:border-b-slate-800"
            : "border-none"
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
