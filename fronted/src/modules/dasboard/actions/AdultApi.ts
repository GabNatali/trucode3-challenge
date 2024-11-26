import { api } from "@/api/api"
import type { IAdultsResponse, IParams } from "../interfaces"
import type { IParamsStats, IStastResponse } from "../interfaces/stats.interface"

export default {
  getAll(params:IParams) {
    return api.get<IAdultsResponse>('/adults', {params: params})
  },
  downloadFile(params:Partial<IParams>) {
    return api.get('/adults/export', {params: params, responseType: 'blob'})
  },
  getStats(params:IParamsStats){
    return api.get<IStastResponse[]>('/stats', {params: params})
  }
}
