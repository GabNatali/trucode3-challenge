<script setup lang="ts">
import type { IAdult } from '../interfaces';


  interface Props {
    data : IAdult[]
  }

  const props = defineProps<Props>();
  const columns:(keyof IAdult)[] = ['age', 'workclass', 'fnlwgt', 'education', 'education_num', 'marital_status', 'occupation', 'relationship', 'race', 'sex', 'capital_gain', 'capital_loss', 'hours_per_week', 'native_country', 'income'];
  const displayColumn = ['Age', 'workclass', 'fnlwgt', 'education', 'education num', 'marital status', 'occupation', 'relationship', 'race', 'sex', 'capital gain', 'capital loss', 'hours per week', 'native country', 'income'];
</script>
<template>
<div class="flex flex-col w-full gap-4 flex-auto">
  <div class="relative overflow-x-auto px-4 overflow-auto h-[calc(100vh-20rem)] md:h-[calc(100vh-17rem)] w-[calc(100vw-2rem)]">
      <table class="w-full text-sm text-left rtl:text-right">
          <thead class="text-xs text-gray-700 border-b-2 border-b-gray-300 uppercase ">
              <tr>
                  <th
                    v-for="(column , index) in displayColumn" :key="index"
                    scope="col" class="px-2 py-3 sticky top-0  border-b bg-gray-100">
                      {{ column }}
                  </th>
              </tr>
          </thead>
          <tbody>
              <tr
                v-for="(adult , index) in props.data" :key="index"
                class="bg-white border-b border-b-gray-200">

                  <th
                    v-for="(column , index) in columns" :key="index"
                    scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                    {{ adult[column] }}
                  </th>

              </tr>
          </tbody>
      </table>
  </div>
</div>
</template>



