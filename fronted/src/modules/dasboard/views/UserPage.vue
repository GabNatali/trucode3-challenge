<script setup lang="ts">
import { onMounted, reactive } from "vue";

import { useAFilterStore } from "../stores/filter.store"
import { useAuthStore } from '../../../stores/auth.store';
import * as utils from "@/utils/utils";
import type { IFilterString } from "../interfaces";
import authApi from "@/modules/auth/actions/authApi";
import { useToast } from 'vue-toastification';
import { isAxiosError } from "axios";

const Toast = useToast();
const authStore = useAuthStore();
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const useFilter = useAFilterStore();
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const selectedFilters = reactive<Record<string, any>>({
  education: [],
  marital_status: [],
  occupation: [],
  income: "",
  order_by:"",
  order_dir:"",
  max_age:0,
  min_age:0,
  user_id: authStore.userDataResponse!.id
});

onMounted(async() => {
  utils.convertConfig(authStore.userConfig, selectedFilters);
})

const saveConfig = async() => {
  Object.keys(selectedFilters).forEach((key) => {
    if (["education", "marital_status", "occupation"].includes(key)) {
      selectedFilters[key] = selectedFilters[key].join(",");
    }
  })

  const params: IFilterString = {
    education: selectedFilters.education,
    marital_status: selectedFilters.marital_status,
    occupation: selectedFilters.occupation,
    income: selectedFilters.income,
    order_by: selectedFilters.order_by,
    order_dir: selectedFilters.order_dir,
    max_age: selectedFilters.max_age,
    min_age: selectedFilters.min_age,
    user_id: authStore.userDataResponse!.id
  }


  try {
    const { data } = await authApi.updateConfig(params);
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { id , ...config} = data;
    authStore.userDataResponse!.config = config;
    utils.convertConfig(config, selectedFilters);
    Toast.success("Configuraciones guardadas");
  } catch (error) {

    if (isAxiosError(error)) {
      Toast.error("error network");
      return
    }
  }
}
</script>

<template>
  <div class="flex flex-col p-6 justify-center md:px-12 gap-4 w-full overflow-hidden">
    <div class="flex items-center justify-between">
        <h1 class="font-semibold text-base text-gray-700">Configuraciones</h1>
        <button
          @click.prevent="saveConfig"
          class="bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 shadow-sm">Guardar</button>
      </div>
    <div class="relative overflow-x-auto px-4 overflow-auto h-[calc(100vh-20rem)] md:h-[calc(100vh-15rem)] w-[calc(100vw-2rem)]">


    <div class="w-full flex flex-col gap-4 flex-wrap">
        <!-- filter -->
      <div>
        <h1 class="font-bold text-base text-gray-700 mb-4">Filtros:</h1>
        <div class="space-y-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
          <div class="flex flex-col"
            v-for:="(options, index) in useFilter.filterData.filters" :key="index"
            >
            <h2 class="text-gray-900 text-base font-semibold pb-2">{{options.label}}</h2>
            <div class="space-y-2 pl-4">
              <label class="flex items-center text-gray-700" v-for="(filter, index) in options.options" :key="index">
                <input
                  v-model="selectedFilters[options.key]"
                  :value="filter.value"
                  :type="options.type" class="form-checkbox" />
                <span class="ml-2">{{filter.label}}</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- order -->
      <div>
        <h1 class="font-bold text-base text-gray-700 mb-4">Order:</h1>
        <div class="space-y-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 ">
          <div class="flex flex-col"
            v-for:="(options, index) in useFilter.filterData.order" :key="index"
            >
            <h2 class="text-gray-900 text-base font-semibold pb-2">{{options.label}}</h2>
            <div class="space-y-1 pl-2">
              <label class="flex items-center text-gray-700" v-for="(filter, index) in options.options" :key="index">
                <input
                  v-model="selectedFilters[options.key as string]"
                  :value="filter.value"
                  :type="options.type" class="form-checkbox" />
                <span class="ml-2">{{filter.label}}</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>
    </div>

  </div>

</template>


