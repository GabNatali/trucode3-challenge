<script setup lang="ts">
import { onMounted, ref } from 'vue';
import  apiAdult  from '../actions/AdultApi';
import BarCharts from '../components/BarCharts.vue';
import type { IDataStats, IParamsStats } from '../interfaces';
import { useToast } from 'vue-toastification';
import { isAxiosError } from 'axios';

const statsByEducation = ref<IDataStats>({
  option: [],
  high_income_count: [],
  low_income_count: []
})

const statsByOccupation = ref<IDataStats>({
  option: [],
  high_income_count: [],
  low_income_count: []
})
const toast = useToast();
const getDataByEducation = async(params:IParamsStats, stats:IDataStats) => {

  try {
    const { data } = await apiAdult.getStats(params)

    data.forEach((key) => {
      stats.option.push(key.option)
      stats.high_income_count.push(key.high_income_count)
      stats.low_income_count.push(key.low_income_count)
    })

  }catch (error) {
    if (isAxiosError(error)) {
      toast.error(error.response?.statusText);
      return
    }
  }
}

onMounted(async () => {
  console.log('mounted container', statsByEducation.value)
  await getDataByEducation({option: 'education'}, statsByEducation.value)
  await getDataByEducation({option: 'occupation'}, statsByOccupation.value)
  console.log(statsByEducation.value)
})
</script>


<template>
  <div class="flex w-full h-[calc(100vh-15rem)] overflow-auto gap-4 flex-col md:flex-row">
    <div class="flex flex-col lg:shadow-lg rounded-md w-[calc(100vw-2rem)] md:w-[400px] h-max lg:px-4" >
      <h1 class="text-base text-gray-700 font-semibold text-center py-4">Nivel de Incresos por Educación</h1>
      <BarCharts :data="statsByEducation" pos="ed" :colors="['#FFA94D', '#74C0FC']"/>
    </div>

    <div class="flex flex-col lg:shadow-lg rounded-md w-[calc(100vw-2rem)] md:w-[400px] h-max lg:px-8" >
      <h1 class="text-base text-gray-700 font-semibold text-center py-4">Nivel de Incresos por Educación</h1>
      <BarCharts :colors="['#6A89CC ', '#F8C471 ']" :data="statsByOccupation" pos="oc" />
    </div>


  </div>



</template>

