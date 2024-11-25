import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

import { signInAction,  } from '../actions';
import { AuthStatus } from '../interfaces/auth-status.enum';
import type { IUser } from '../interfaces';


export const useAuthStore = defineStore('auth', () => {


  const authStatus = ref<AuthStatus>(AuthStatus.Checking);
  const user = ref<IUser| undefined>();



  const login = async (email:string , password:string ) => {

    try {
      const loginResponse = await signInAction(email , password)
      if (!loginResponse.ok){
        logout()
        return false;
      }

      user.value = loginResponse.user;
      authStatus.value = AuthStatus.Authenticated;
      localStorage.setItem('token', user.value.token || '' )
      return true;

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch(error) {

      return logout();
    }
  }

  const logout = ()=> {
    authStatus.value = AuthStatus.Unauthenticated;
    user.value = undefined;
    localStorage.removeItem('token')
    return false;
  }


  // const register = async (payload: IUserPayLoad )=> {
  //   try {
  //     const registerRespnse = await signUpAction(payload)
  //     if (!registerRespnse.ok){
  //       return {
  //         ok:false,
  //         message: registerRespnse.message
  //       }
  //     }

  //     return{
  //       ok:true,
  //       message:"user registrado"
  //     }

  //   } catch (error) {
  //     return {
  //       ok:false,
  //       message:error,
  //     }
  //   }
  // }
  return {
    user,
    authStatus,

    //Getters
    isChecking: computed(()=> authStatus.value === AuthStatus.Checking),
    isAuthenticated : computed(()=> authStatus.value === AuthStatus.Authenticated),
    email: computed( ()=> user.value?.email),

    //actions
    login,
    logout,

  }
})
