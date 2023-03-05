import { ReactNode } from "react";

interface PillInterface {
  children: ReactNode;
  color: string;
}

const Pill: React.FC<PillInterface> = ({ children, color }) => {
  return (
    <span className={`${color} py-1 px-2 rounded-full text-xs`}>
      {children}
    </span>
  );
};

export default Pill;
