import { useAuthStore } from '@/modules/auth/stores/auth1.store';
import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL
});


api.interceptors.request.use(config => {

  const token = localStorage.getItem('token')

  if(token){
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

api.interceptors.response.use( response => response,
  error => {
      if (error.response && error.response.status === 401) {
          const authStore = useAuthStore();
          authStore.logout();

      }
      return Promise.reject(error.message);


  });


export{ api }
