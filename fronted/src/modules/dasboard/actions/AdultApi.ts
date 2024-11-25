import { api } from "@/api/api"
import type { IAdultsResponse, IParams } from "../interfaces"

export default {
  getAll(params:IParams) {
    return api.get<IAdultsResponse>('/adults', {params: params})
  },
  downloadFile(params:Partial<IParams>) {
    return api.get('/adults/export', {params: params, responseType: 'blob'})
  }
}
