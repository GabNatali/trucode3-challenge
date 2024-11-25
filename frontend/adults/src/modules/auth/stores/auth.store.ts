import { api } from "@/api/api";
import { defineStore } from "pinia";
import { computed, ref } from 'vue';
import type { IUser } from '../interfaces';


export const  useAuthStore = defineStore('auth', ()=> {

 const authenticated = ref(false);
 const accessToken = ref(localStorage.getItem('accessToken') || '')
 const user = ref<IUser| undefined>();



 const setAccessToken = (token:string) => {
    accessToken.value = token;
    localStorage.setItem('accessToken', token);
}

const signIn = async (credentials: { email:string ; password:string })=>{

  if (authenticated.value) {
      return Promise.reject('User is already logged in.');
  }

  const { data: { data }} = await api.post('/auth/login', credentials)

  user.value = data
  authenticated.value = true;
  setAccessToken(data.token);
  return Promise.resolve(data)
}

const logout = ()=> {
  localStorage.removeItem('token')
  authenticated.value = false
  return false;
}


 return {


  isAuthenticated: computed(()=> authenticated.value === true),
  accessToken: computed(()=> accessToken.value),

  signIn,
  logout,
 }

});
