export interface IUser {
  id:        number;
  firstName: string;
  lastName:  string;
  email:     string;
  token?:     string;
}


export interface IUserPayLoad {
  firstName: string,
  lastName: string,
  password:string,
  email: string
}
