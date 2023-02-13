import { UserRole } from "$/enums/user";

export interface User {
  id: number;
  fullName: string;
  role: UserRole;
}
