<script setup lang="ts">
import * as echarts from 'echarts';
import {watch } from 'vue';
import type { IDataStats } from '../interfaces';

interface Props {
  data: IDataStats;
  pos: string;
  colors:string[];
}

const props = defineProps<Props>();

const render = () => {
  const chartDom = document.getElementById(props.pos);
  if (chartDom) {
    const myChart = echarts.init(chartDom);

    const option = {
      color:[props.colors[0], props.colors[1]],
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'shadow',
        },
      },
      legend: {},
      grid: {
        left: '1%',
        right: '5%',
        bottom: '10%',
        containLabel: true,
      },
      xAxis: [
        {
          type: 'value',
        },
      ],
      yAxis: [
        {
          type: 'category',
          data: props.data.option,
        },
      ],
      series: [
        {
          name: '<=50',
          type: 'bar',
          barWidth: '40%',
          emphasis: {
            focus: 'series',
          },
          data: props.data.low_income_count,
        },
        {
          name: '>=50',
          type: 'bar',
          barWidth: '40%',
          emphasis: {
            focus: 'series',
          },
          data: props.data.high_income_count,
        },
      ],
    };

    myChart.setOption(option);
  }
};

watch(
  () => props.data,
  (val) => {
    if (val && val.option.length > 0) {
      render();
    }
  },
  { deep: true, immediate: true }
);
</script>

<template>
  <div :id="pos" class="w-full h-[450px]"></div>
</template>
