import { ReactNode, useState } from "react";
import { Theme } from "../enums/theme";
import useTheme from "../hooks/useTheme";
import Moon from "./icons/Moon";
import Sun from "./icons/Sun";
import "./../styles/ThemePicker.css";
import ComputerDesktop from "./icons/ComputerDesktop";

const ThemePicker: React.FC<{}> = ({}) => {
  const { theme, setTheme } = useTheme();

  const [showPicker, setShowPicker] = useState<boolean>(false);

  const themePickerIcon = (): ReactNode => {
    switch (theme) {
      case Theme.Dark:
        return (
          <span className="text-yellow-500">
            <Moon />
          </span>
        );
      case Theme.Light:
        return (
          <span className="text-primary">
            <Sun />
          </span>
        );
      default:
        const systemIsDark = window.matchMedia(
          "(prefers-color-scheme: dark)"
        ).matches;
        return (
          <span className="text-slate-900 dark:text-slate-100">
            {systemIsDark === true ? <Moon /> : <Sun />}
          </span>
        );
    }
  };

  const handleThemeUpdate = (theme: Theme) => {
    setTheme(theme);
    setShowPicker(false);
  };

  return (
    <div className="flex items-center">
      <button className="p-2" onClick={() => setShowPicker(true)}>
        {themePickerIcon()}
      </button>
      <ul
        className="absolute top-full w-32 -translate-x-24 bg-white dark:bg-slate-800 rounded-lg ring-1 dark:ring-0 ring-slate-900/10 shadow-lg overflow-hidden"
        hidden={!showPicker}
      >
        <li
          className={`theme-item ${theme === Theme.Light ? "active" : ""}`}
          onClick={() => handleThemeUpdate(Theme.Light)}
        >
          <Sun />
          Light
        </li>
        <li
          className={`theme-item ${theme === Theme.Dark ? "active" : ""}`}
          onClick={() => handleThemeUpdate(Theme.Dark)}
        >
          <Moon />
          Dark
        </li>
        <li
          className={`theme-item ${theme == Theme.System ? "active" : ""}`}
          onClick={() => handleThemeUpdate(Theme.System)}
        >
          <ComputerDesktop />
          System
        </li>
      </ul>
    </div>
  );
};

export default ThemePicker;
