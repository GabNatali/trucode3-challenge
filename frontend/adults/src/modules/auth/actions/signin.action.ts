import { api } from "@/api/api"
import type { IAuthResponse, IUser } from "../interfaces"
import { isAxiosError } from "axios"

interface ILoginError {
  ok: false;
  message:string
}

interface ILoginSuccess {
  ok:true;
  user: IUser;
}
export const signInAction = async (email: string , password: string): Promise<ILoginError | ILoginSuccess> => {

  try {

    const { data: {data} } = await api.post<IAuthResponse>('/auth/login' , {
      email,
      password,
    })

    return {
      ok:true,
      user: data,
    }

  } catch (error) {

    if(isAxiosError(error) && error.response?.status === 401 ){
      return {
        ok:false,
        message: error.message
      }
    }

    throw new Error('Error server')

  }
}
