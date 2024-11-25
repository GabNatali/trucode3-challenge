import { api } from "@/api/api"
import type { IUserPayLoad } from "../interfaces"


export default {
  register(data:IUserPayLoad) {
    return api.post('/users', data)
  }
}
