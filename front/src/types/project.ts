export type Project = {
  id: number;
  name: string;
  description: string;
  url: string;
  github_repository: string;
  statuses: ProjectStatus[];
};

export type ProjectStatus = {
  id: number;
  name: string;
  is_healthy: boolean;
}
