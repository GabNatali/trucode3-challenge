
<script setup lang="ts">
import { onMounted, reactive } from "vue";
import { useAFilterStore } from "../stores/filter.store";
import OptBlock from "./OptBlock.vue";
import { useToast } from "vue-toastification";
import { useAuthStore } from "@/stores/auth.store";
import * as utils from "@/utils/utils";

const authStore = useAuthStore();
const useFilter = useAFilterStore();
const selectedFilters = reactive({
  education: [],
  marital_status: [],
  occupation: [],
  income: "",
  order_by:"",
  order_dir:"",
  max_age:0,
  min_age:0
});


const toast = useToast();
const emit = defineEmits(["filters"]);
const sendFilters = () => {
  if (selectedFilters.min_age > selectedFilters.max_age) {
    toast.error("La edad mínima no puede ser mayor que la edad máxima.");
    return;
  }

  const filters = {...selectedFilters};
  emit("filters", filters);
}

onMounted(() => {
    utils.convertConfig(authStore.userConfig, selectedFilters);
    emit("filters", selectedFilters);
})
</script>

<template>
  <form
    @submit.prevent="sendFilters"
    class="w-full max-w-sm bg-white">
    <div class="flex items-center justify-between border-b pb-1">
      <h2 class="text-lg font-bold">Filtrar y ordenar</h2>
    </div>
    <!-- Filters -->
    <div class="mt-4 h-[calc(100vh-10rem)] overflow-y-auto">


     <!--order  -->
    <OptBlock :open="true" title="ORDENAR POR">
      <div class="space-y-2 pl-4">
        <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.orderBy.options" :key="index">
          <input
            v-model="selectedFilters.order_by"
            :value="filter.value"
            type="radio" class="form-checkbox" />
          <span class="ml-2">{{filter.label}}</span>
        </label>
      </div>
    </OptBlock>

    <OptBlock :open="true" title="ORDENAR">
      <div class="space-y-2 pl-4">
        <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.orderDir.options" :key="index">
          <input
            v-model="selectedFilters.order_dir"
            :value="filter.value"
            type="radio" class="form-checkbox" />
          <span class="ml-2">{{ filter.label}}</span>
        </label>
      </div>
    </OptBlock>

      <!-- Education -->
      <OptBlock :open="false" title="EDUCACION">
        <div class="space-y-2 pl-4">
          <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.education.options" :key="index">
            <input
              v-model="selectedFilters.education"
              :value="filter.value"
              type="checkbox" class="form-checkbox" />
            <span class="ml-2">{{ filter.label}}</span>
          </label>
        </div>
      </OptBlock>

      <!-- Maritial status -->
      <OptBlock :open="false" title="MARITIAL STATUS">
        <div class="space-y-2 pl-4">
          <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.marital_status.options" :key="index">
            <input
              v-model="selectedFilters.marital_status"
              :value="filter.value"
              type="checkbox" class="form-checkbox" />
            <span class="ml-2">{{ filter.label}}</span>
          </label>
        </div>
      </OptBlock>


      <OptBlock :open="false" title="OCUPATION">
        <div class="space-y-2 pl-4">
          <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.occupation.options" :key="index">
            <input
              v-model="selectedFilters.occupation"
              :value="filter.value"
              type="checkbox" class="form-checkbox" />
            <span class="ml-2">{{ filter.label}}</span>
          </label>
        </div>
      </OptBlock>

      <OptBlock :open="false" title="INCOME">
        <div class="space-y-2 pl-4">
          <div class="space-y-2 pl-4">
          <label class="flex items-center text-gray-700" v-for="(filter, index) in useFilter.filtersData.income.options" :key="index">
            <input
              v-model="selectedFilters.income"
              :value="filter.value"
              type="radio"
              name="income"/>
            <span class="ml-2">{{filter.label}}</span>
          </label>
        </div>
        </div>
      </OptBlock>

      <OptBlock :open="false" title="EDAD">
        <div class="space-y-2 pl-4">
          <div class="flex w-max items-center">
            <span class="ml-2">Min</span>
            <input
              v-model="selectedFilters.min_age"
              type="number" class="w-20 border border-gray-300 ml-2"/>
          </div>
          <div class="flex w-max items-center">
            <span class="ml-2">Max</span>
            <input
              v-model="selectedFilters.max_age"
              max="100"
              type="number" class="w-20 border border-gray-300 ml-2"/>
          </div>
        </div>
      </OptBlock>
    </div>

    <div class="mt-6">
      <button
        @click.prevent="sendFilters"
        type="submit"
        class="w-full bg-black text-white py-3 rounded-md font-bold hover:bg-gray-800"
      >
        APLICAR
      </button>
    </div>
  </form>
</template>
