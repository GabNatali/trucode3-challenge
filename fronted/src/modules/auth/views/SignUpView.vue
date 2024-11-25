<script lang="ts" setup>
import { reactive, ref } from 'vue';

import type { IUserFormRegister } from '../interfaces';
import authApi from '../actions/authApi';
import { useToast } from 'vue-toastification';
import { isAxiosError } from 'axios';


const toast =  useToast();
const isLoading = ref(false);
const errorMessage = ref('');
const formRegister = reactive<IUserFormRegister>({
    username:'',
    email:'',
    password:'',
    repeat: '',
})

const validePassword = ():void => {
  if (formRegister.password !== formRegister.repeat) {
    errorMessage.value = 'Passwords do not match';
  }else {
    errorMessage.value = '';
  }
}

const onRegister = async() => {
  isLoading.value = true
  try {

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { repeat, ...payload } = formRegister;
    const { data } = await authApi.register(payload);
    toast.success(data.message)

    // Resete form
    formRegister.username = '';
    formRegister.email = '';
    formRegister.password = '';
    formRegister.repeat = '';

  } catch (error) {

    if (isAxiosError(error)) {
      const message = error?.response?.data?.message || 'Error server'
      toast.error(message)
      return
    }

    toast.error(error)

  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto h-screen lg:py-0 md:bg-gray-50">
    <div class="w-full md:mt-0 sm:max-w-md xl:p-0 md:shadow bg-white rounded-lg ">
      <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl mx-auto max-w-sm py-4">
        Sing Up
      </h1>
      <form class="max-w-sm mx-auto" @submit.prevent="onRegister">

        <div class="mb-5">
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900 ">Username</label>
          <input v-model="formRegister.username"
            type="text" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Doe" required />
        </div>
        <div class="mb-5">
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900 ">Your email</label>
          <input v-model="formRegister.email"
            type="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="name@flowbite.com" required />
        </div>
        <div class="mb-5">
          <label for="password" class="block mb-2 text-sm font-medium text-gray-900 ">Your password</label>
          <input v-model="formRegister.password"

            type="password" id="password" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="*******" required />
          <span v-if="errorMessage" class="text-red-500 text-sm mx-2 ">{{ errorMessage }}</span>
        </div>

        <div class="mb-5">
          <label for="password" class="block mb-2 text-sm font-medium text-gray-900 ">Your password</label>
          <input v-model="formRegister.repeat"
            @input="validePassword"
            type="password" id="password" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="*******" required />
          <span v-if="errorMessage" class="text-red-500 text-sm mx-2 ">{{ errorMessage }}</span>
        </div>


        <button
          :disabled="isLoading || formRegister.password !== formRegister.repeat"
          type="submit"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center ">
          <span v-if="isLoading">...cargando</span>
          <span v-else>Register</span>
        </button>
      </form>

      <div class="w-full px-6 py-4">
        <p class="text-sm font-light text-black-500 text-left ">
            Already have an account?
            <RouterLink :to="{ name: 'sign-in' }" class="font-medium text-blue-600 hover:underline">Sign in</RouterLink
        ></p>
      </div>
    </div>
  </div>
</template>



