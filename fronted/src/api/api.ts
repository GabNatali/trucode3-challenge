import { isTokenExpired } from '@/modules/auth/auth.utils';
import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL
});

// const authStores = useAuthStore();

api.interceptors.request.use(config => {
  const token = localStorage.getItem('accessToken');

  if(token && !isTokenExpired(token)) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

api.interceptors.response.use( response => response,
  error => {
      if (error.response && error.status === 401) {
        // authStores.logout();
        return Promise.reject(error);
      }
      return Promise.reject(error);
  });


export{ api }
