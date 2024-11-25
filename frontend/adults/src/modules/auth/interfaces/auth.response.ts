import type { IUser } from ".";

export interface IAuthResponse {
  data:    IUser;
  message: string;
  status:  number;
}

