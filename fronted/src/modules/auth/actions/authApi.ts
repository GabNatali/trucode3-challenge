import { api } from "@/api/api"
import type { ILoginPayload, ILoginResponse, IResponse, IUser, IUserPayLoad, IUserResponse } from "../interfaces"
import type { IFilterString } from "@/modules/dasboard/interfaces"

export default {
  register(data:IUserPayLoad) {
    return api.post<IResponse<IUser>>('/users', data)
  },
  signIn(data:ILoginPayload) {
    return api.post<ILoginResponse>('/auth/sign-in', data)
  },
  logout() {
    localStorage.removeItem('accessToken');
  },
  getMe(){
    return api.get<IUserResponse>('/users/me')
  },
  updateConfig(config){
    return api.patch('/users/config', config)
  }
}
