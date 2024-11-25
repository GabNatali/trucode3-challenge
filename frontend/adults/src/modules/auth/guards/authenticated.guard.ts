/* eslint-disable @typescript-eslint/no-unused-expressions */
import type { NavigationGuardNext, RouteLocationNormalized } from "vue-router";
import { useAuthStore } from "../stores/auth.store";


const authenticatedGurard = async(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext,
)=> {

  const authStore = useAuthStore();



  authStore.isAuthenticated === false ? next({name :'sign-in'}) : next({name : 'home'});

}


export default authenticatedGurard;
