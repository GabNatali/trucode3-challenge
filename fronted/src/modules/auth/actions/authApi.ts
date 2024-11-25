import { api } from "@/api/api"
import type { ILoginPayload, IResponse, IUser, IUserPayLoad, IUserResponse } from "../interfaces"

export default {
  register(data:IUserPayLoad) {
    return api.post<IResponse<IUser>>('/users', data)
  },
  signIn(data:ILoginPayload) {
    return api.post<string>('/auth/sign-in', data)
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
