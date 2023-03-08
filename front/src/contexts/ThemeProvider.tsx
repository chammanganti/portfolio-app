import {
  createContext,
  Dispatch,
  ReactNode,
  SetStateAction,
  useEffect,
  useState,
} from "react";
import { DARK_SCHEME_QUERY } from "../constants";
import { Theme } from "../enums/theme";

interface ThemeProviderInterface {
  children: ReactNode;
}

type ThemeContextType = {
  theme: Theme;
  setTheme: Dispatch<SetStateAction<Theme>>;
};

const THEME_KEY: string = "theme";

export const ThemeContext = createContext({} as ThemeContextType);

const getStoredTheme = (): Theme => {
  const storedTheme = localStorage.getItem(THEME_KEY) as Theme;
  if (typeof storedTheme === "string") {
    return storedTheme;
  }

  return Theme.System;
};

export const ThemeProvider: React.FC<ThemeProviderInterface> = ({
  children,
}) => {
  const [theme, setTheme] = useState<Theme>(getStoredTheme());

  const handleSetTheme = (theme: Theme) => {
    const root = window.document.documentElement;

    switch (theme) {
      case Theme.Dark:
        root.classList.add(Theme.Dark);
        localStorage.setItem(THEME_KEY, Theme.Dark);
        break;
      case Theme.Light:
        root.classList.remove(Theme.Dark);
        localStorage.setItem(THEME_KEY, Theme.Light);
        break;
      default:
        const systemTheme = window.matchMedia(DARK_SCHEME_QUERY).matches
          ? Theme.Dark
          : Theme.Light;
        const isDark = systemTheme === Theme.Dark;
        root.classList.remove(isDark === true ? Theme.Light : Theme.Dark);
        root.classList.add(systemTheme);
        localStorage.removeItem(THEME_KEY);
    }
  };

  useEffect(() => {
    window.matchMedia(DARK_SCHEME_QUERY).addEventListener("change", (_) => {
      handleSetTheme(Theme.System);
    });
    return () =>
      window
        .matchMedia(DARK_SCHEME_QUERY)
        .removeEventListener("change", (_) => {});
  }, []);

  useEffect(() => {
    handleSetTheme(theme);
  }, [theme]);

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};
