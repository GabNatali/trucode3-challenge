import type { NavigationGuardNext, RouteLocationNormalized } from "vue-router";
import { useAuthStore } from "../stores/auth.store";


const isNotAuthenticatedGurard = async(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext,
)=> {

  const authStore = useAuthStore();



  if (to.name === 'sign-in' && authStore.isAuthenticated) {
    return next({ name: 'home' });

  } else {

    if (to.name !== 'sign-in') {
      return next({ name: 'sign-in' });
    }
  }

  next();
}


export default isNotAuthenticatedGurard;
