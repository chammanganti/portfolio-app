import { Project as ProjectType, ProjectStatus as ProjectStatusType } from "../types/project";
import { Project, ProjectContainer } from "./Project";

const Showcase: React.FC<{}> = ({}) => {
  let projects: ProjectType[] = [
    {
      id: 1,
      name: "App 1",
      description: "App 1 description",
      url: "https://app-1.chammanganti.dev",
      github_repository: "https://github.com",
      statuses: [
        {
          id: 1,
          name: "ec2",
          is_healthy: true
        },
        {
          id: 2,
          name: "elb",
          is_healthy: false
        },
        {
          id: 3,
          name: "app",
          is_healthy: false
        }
      ]
    },
    {
      id: 2,
      name: "App 2",
      description: "App 2 description",
      url: "https://app-2.chammanganti.dev",
      github_repository: "https://github.com",
      statuses: [
        {
          id: 4,
          name: "eks",
          is_healthy: true
        },
        {
          id: 5,
          name: "app",
          is_healthy: true
        }
      ]
    },
    {
      id: 3,
      name: "App 3",
      description: "App 3 description",
      url: "https://app-3.chammanganti.dev",
      github_repository: "https://github.com",
      statuses: [
        {
          id: 6,
          name: "ec2",
          is_healthy: true
        },
        {
          id: 7,
          name: "app",
          is_healthy: true
        }
      ]
    }
  ]

  return (
    <>
      <div className="bg-white dark:bg-slate-800 rounded-xl shadow-xl ring-1 ring-slate-900/5 overflow-hidden">
        <header className="md:flex md:items-center md:justify-between p-4">
          <h2 className="mb-4 md:mb-0 font-semibold text-slate-900 dark:text-slate-100">Projects</h2>
          <div className="flex items-center space-x-2">
            <input className="w-full bg-white dark:bg-slate-700 p-2 ring-1 dark:ring-0 ring-slate-200 focus:ring-2 dark:focus:ring-2 focus:ring-primary dark:focus:ring-primary focus:outline-none dark:focus-within:bg-transparent shadow-sm rounded-md text-sm font-light text-slate-900 dark:text-slate-100" type="text" placeholder="Token" />
            <button className="bg-primary hover:bg-primary-400 p-2 rounded-md text-sm font-semibold text-white transition-colors ease-in-out duration-100">Unlock</button>
          </div>
        </header>
        <ProjectContainer>
          {projects.map((project) => (
            <Project key={project.id} project={project} />
          ))}
        </ProjectContainer>
      </div>
    </>
  );
};

export default Showcase;
