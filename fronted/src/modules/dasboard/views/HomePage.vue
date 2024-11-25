<script setup lang="ts" >
import { computed, onMounted, ref } from 'vue';
import AdultsTable from '../components/AdultsTable.vue';
import FiltersTable from '../components/FiltersTable.vue';
import Paginator from '../components/PaginatorTable.vue';
import adultApi from '../actions/AdultApi';
import type { IAdultsResponse, IFilter } from '../interfaces';
import { useToast } from 'vue-toastification';
import { initDrawers } from 'flowbite';
import axios from 'axios';
import { useAuthStore } from '@/stores/auth.store';

const authStore = useAuthStore();
// filters
const filters = ref<Partial<Record<keyof IFilter, string | number>>>({
  education: "",
  marital_status: "",
  occupation: "",
  income: "",
  order_by: "",
  order_dir: "",
  max_age: 0,
  min_age: 0,
});
const getFilters = async (options:IFilter) => {

  Object.keys(options).forEach((key) => {

    const typedKey = key as keyof IFilter;

    if (Array.isArray(options[typedKey])) {
      filters.value[typedKey] = options[typedKey].join(",");
    }
    else {
      filters.value[typedKey] = options[typedKey];
    }
  })

  currentPage.value = 1;
  await loadData();
};

// paginator
const currentPage = ref(1);
const pageSize = ref(15);
const handlePageChange = async(page: number) => {
  currentPage.value = page;
  await loadData();
};
const handlePageSizeChange = async(newSize: number) => {
  pageSize.value = newSize;
  currentPage.value = 1;
  await loadData();
};

const params = computed(() => {
  return {
    ...filters.value,
    page: currentPage.value,
    page_size: pageSize.value,
  };
});

const response = ref<IAdultsResponse>({
  count: 0,
  result: [],
});

const toast = useToast();

const loadData = async () => {
  try {
    const { data } = await adultApi.getAll(params.value);
    response.value = data
  }catch (error) {
    if (axios.isAxiosError(error)) {
      toast.error(error.message);
    }
  }
}

const donwloadFileCSV = async () => {
  const params = {...filters.value};
  try {
    const { data } = await adultApi.downloadFile(params);
    const blob = new Blob([data], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = 'adults.csv';
    link.click();
    window.URL.revokeObjectURL(url);
  }catch (error) {
    if (axios.isAxiosError(error)) {
      toast.error(error.response?.data.message);
    }
  }
}
onMounted(async() => {
  initDrawers();
  await loadData();

})
</script>

<template>
  <div class="w-full h-full flex flex-col flex-auto p-4 overflow-hidden">
    <div class="p-4 flex justify-between flex-wrap gap-4 items-center">
        <h1 class="font-medium text-lg text-gray-900">Datos</h1>
        <div class="flex gap-2 flex-wrap">

          <button
                @click.prevent="donwloadFileCSV"
                type="button" class="text-black-500 border border-black-500 hover:bg-gray-100 rounded-lg text-sm p-2.5 text-center inline-flex items-center">
                <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="2"
                stroke="currentColor"
                class="w-8 h-8"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 16v-8m0 8l-4-4m4 4l4-4m-4 4v-8m0 8H6m6 0h6"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M6 18h12m0 0v2H6v-2m12 0H6"
                />
              </svg>
          </button>
          <button class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mb-2 "
            type="button" data-drawer-target="drawer-right-example" data-drawer-show="drawer-right-example" data-drawer-placement="right" aria-controls="drawer-right-example">
          Filtrar y Ordenar
        </button>
        </div>

    </div>
    <div class="flex ">
      <div
        id="drawer-right-example" class="fixed top-0 right-0 z-40 h-screen p-4 overflow-y-auto transition-transform translate-x-full bg-white w-80" tabindex="-1" aria-labelledby="drawer-right-label">
        <FiltersTable @filters="getFilters"></FiltersTable>
      </div>
    <div class="flex flex-auto flex-col">
      <div >
        <AdultsTable :data=response.result></AdultsTable>
      </div>

      <Paginator
        :total-items=response?.count
        :current-page=currentPage
        :page-size=pageSize
        :page-size-options="[15,50, 100, 150, 200]"
        @page-change="handlePageChange"
        @page-size-change="handlePageSizeChange"
      />
    </div>
    </div>


  </div>
</template>
