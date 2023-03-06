import { ReactNode, useState } from "react";
import Pill from "./Label";
import ChevronDown from "./icons/ChevronDown";
import RocketLaunch from "./icons/RocketLaunch";
import { Project as ProjectType } from "../types/project";
import CheckCircle from "./icons/CheckCircle";

interface ProjectContainerInterface {
  children: ReactNode;
}

interface ProjectInterface {
  project: ProjectType;
}

export const ProjectContainer: React.FC<ProjectContainerInterface> = ({
  children,
}) => {
  return <ul className="bg-slate-50 p-4">{children}</ul>;
};

export const Project: React.FC<ProjectInterface> = ({ project }) => {
  const [toggle, setToggle] = useState<boolean>(false);

  const countHealth = (): number => {
    let count = 0;
    for (let status of project.statuses) {
      if (status.is_healthy === true) {
        count++;
      }
    }
    return count;
  };

  return (
    <li
      className={`flex flex-col group bg-white ring-1 ring-slate-200 rounded-md shadow-sm p-4 mb-4 last:mb-0 cursor-pointer ${
        toggle === true ? "is-open" : ""
      }`}
      onClick={() => setToggle(!toggle)}
    >
      <section className="flex justify-between items-center">
        <div>
          <p className="text-slate-900 font-semibold">{project.name}</p>
          <p className="text-slate-400 font-light">{project.description}</p>
        </div>
        <div className="flex items-center space-x-8">
          {countHealth() === project.statuses.length && (
            <Pill color="bg-green-200 text-green-600">Healthy</Pill>
          )}
          {countHealth() < project.statuses.length && countHealth() > 0 && (
            <Pill color="bg-orange-200 text-orange-600">Unhealthy</Pill>
          )}
          {countHealth() === 0 && (
            <Pill color="bg-gray-200 text-gray-600">Inactive</Pill>
          )}
          <p
            className="text-slate-400 text-sm font-light"
            hidden={
              countHealth() === 0 || countHealth() === project.statuses.length
            }
          >
            <span>{countHealth()}</span>
            <span> / </span>
            <span>{project.statuses.length}</span>
          </p>
          <button className="text-slate-400 rotate-0 group-[.is-open]:rotate-180 transition-transform ease-in-out duration-500">
            <ChevronDown />
          </button>
        </div>
      </section>
      <section className="max-h-0 group-[.is-open]:max-h-60 overflow-hidden transition-all ease-in-out duration-500">
        <div>
          <span className="text-slate-400 font-light">URL: </span>
          <a
            className="text-primary-300 hover:text-primary cursor-pointer"
            href={project.url}
            target="_blank"
            rel="noopener noreferrer"
          >
            {project.url}
          </a>
        </div>
        <div>
          <span className="text-slate-400 font-light">Repository: </span>
          <a
            className="text-primary-300 hover:text-primary cursor-pointer"
            href={project.github_repository}
            target="_blank"
            rel="noopener noreferrer"
          >
            {project.github_repository}
          </a>
        </div>
        <div className="mt-6" hidden={countHealth() === 0}>
          <h3 className="text-slate-900 font-semibold">Health checks</h3>
          {project.statuses.map((status) => (
            <div key={status.id} className="flex justify-between items-center">
              <p className="text-slate-400 font-light">{status.name}</p>
              <span
                className={`${
                  status.is_healthy === true ? "text-green-500" : "text-red-500"
                }`}
              >
                <CheckCircle />
              </span>
            </div>
          ))}
        </div>
        <div hidden={countHealth() !== 0}>
          <div className="flex justify-end mt-6">
            <button className="flex items-center bg-primary hover:bg-primary-400 px-3 py-2 rounded-md text-sm font-semibold text-white transition-colors ease-in-out duration-100">
              <span className="mr-2">Launch</span>
              <RocketLaunch />
            </button>
          </div>
        </div>
      </section>
    </li>
  );
};
