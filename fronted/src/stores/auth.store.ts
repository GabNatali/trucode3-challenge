import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import type { IUserResponse } from "../modules/auth/interfaces";



export const useAuthStore = defineStore('auth', ()=> {

  const authenticated = ref(false);
  const accessToken = ref(localStorage.getItem('accessToken') || '')

  const router = useRouter();
  const setToken = (token: string) => {
    accessToken.value = token;
    localStorage.setItem('accessToken', token);
  };

const userDataResponse = ref<IUserResponse | null>(null);

const userData = computed(() => {
  return {
    id: userDataResponse.value?.id,
    username: userDataResponse.value?.username,
    email: userDataResponse.value?.email
  }
})


const userConfig = computed(() => {
  return {
    ...userDataResponse.value?.config
  }

})
  const logout = () => {
    localStorage.removeItem('accessToken');
    authenticated.value = false;
    router.push('/sign-in');
  };

  return {
    authenticated,
    accessToken,
    setToken,
    logout,
    userData,
    userDataResponse,
    userConfig
  }

});
