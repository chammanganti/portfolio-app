import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";
import { Theme } from "../enums/theme";

interface ThemeProviderInterface {
  children: ReactNode;
}

type ThemeContextType = {
  theme: Theme;
  setTheme: Dispatch<SetStateAction<Theme>>;
};

export const ThemeContext = createContext({} as ThemeContextType);

const getInitialTheme = (): Theme => {
  if (
    localStorage?.theme === "dark" ||
    (!("theme" in localStorage) &&
      window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    return Theme.Dark;
  }

  return Theme.Light;
};

export const ThemeProvider: React.FC<ThemeProviderInterface> = ({
  children,
}) => {
  const [theme, setTheme] = useState<Theme>(getInitialTheme());

  const handleSetTheme = (theme: Theme) => {
    const root = window.document.documentElement;
    const isDark = theme === Theme.Dark;

    root.classList.remove(isDark === true ? Theme.Light : Theme.Dark);
    root.classList.add(theme);

    localStorage.setItem("theme", theme.toString());
  };

  useEffect(() => {
    handleSetTheme(theme);
  }, [theme]);

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};
