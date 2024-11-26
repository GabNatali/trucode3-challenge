import { isTokenExpired } from '@/modules/auth/auth.utils';
import axios from 'axios';
import { useRouter } from 'vue-router';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL
});

const router = useRouter();

api.interceptors.request.use(config => {
  const token = localStorage.getItem('accessToken');

  if(token && !isTokenExpired(token)) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

api.interceptors.response.use( response => response,
  error => {
    if (error.response && error.response.status === 401 || error.response.status === 403) {
      localStorage.removeItem('accessToken');
      window.location.reload();
      router.push('/sign-in');
    }
    return Promise.reject(error);
  });


export{ api }
