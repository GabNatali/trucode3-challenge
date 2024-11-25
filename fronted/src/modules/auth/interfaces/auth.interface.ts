
export interface ILoginPayload {
  email: string,
  password: string
}

export interface IUserFormRegister {
  username: string,
  password:string,
  email: string,
  repeat: string,
}

export interface IUserPayLoad {
  username: string,
  password:string,
  email: string,
}


export interface IResponseError {
  status: number;
  message: string
}

export interface IResponse<T> {
  data: T;
  status: number;
  message: string
}
