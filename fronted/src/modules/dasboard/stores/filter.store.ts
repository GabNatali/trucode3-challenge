import { defineStore } from "pinia";
import { ref } from "vue";



export const useAFilterStore = defineStore('filter', ()=> {

  const filtersData = ref({
    education:{
        label: "Education",
        key: "education",
        options: [
          { label: "7th-8th", value: "7th-8th" },
          { label: "Doctorate", value: "Doctorate" },
          { label: "Bachelors", value: "Bachelors" },
          { label: "Some-college", value: "Some-college" },
          { label: "11th", value: "11th" },
          { label: "HS-grad", value: "HS-grad" },
          { label: "Prof-school", value: "Prof-school" },
          { label: "Assoc-acdm", value: "Assoc-acdm" },
          { label: "Assoc-voc", value: "Assoc-voc" },
          { label: "9th", value: "9th" },
          { label: "7th", value: "7th" },
          { label: "12th", value: "12th" },
          { label: "Masters", value: "Masters" },
          { label: "1st-4th", value: "1st-4th" },
          { label: "10th", value: "10th" },
          { label: "5th-6th", value: "5th-6th" },
          { label: "Preschool", value: "Preschool" },
        ],
    },
    marital_status:{
      label: "Marital Status",
      key: "marital_status",
      options: [
        { label: "Married Civil Spouse", value: "Married-civ-spouse" },
        { label: "Divorced", value: "Divorced" },
        { label: "Never Married", value: "Never-married" },
        { label: "Separated", value: "Separated" },
        { label: "Widowed", value: "Widowed" },
        { label: "Married Spouse Absent", value: "Married-spouse-absent" },
        { label: "Married Armed Forces Spouse", value: "Married-AF-spouse" },
      ],
    },
    occupation:{
      label: "Occupation",
      key: "occupation",
      options: [
        { label: "Tech Support", value: "Tech-support" },
        { label: "Craft Repair", value: "Craft-repair" },
        { label: "Other Service", value: "Other-service" },
        { label: "Sales", value: "Sales" },
        { label: "Exec Managerial", value: "Exec-managerial" },
        { label: "Prof Specialty", value: "Prof-specialty" },
        { label: "Handlers Cleaners", value: "Handlers-cleaners" },
        { label: "Machine Operator Inspector", value: "Machine-op-inspct" },
        { label: "Administrative Clerical", value: "Adm-clerical" },
        { label: "Farming Fishing", value: "Farming-fishing" },
        { label: "Transport Moving", value: "Transport-moving" },
        { label: "Private Household Service", value: "Priv-house-serv" },
        { label: "Protective Service", value: "Protective-serv" },
        { label: "Armed Forces", value: "Armed-Forces" },
      ],
    },
    income:{
      label: "Income",
      key: "income",
      options: [
        { label: "Less than 50K", value: "<=50K" },
        { label: "More than 50K", value: ">50K" },
      ],
    },
    orderBy:{
      label: "Order By",
      key: "order_by",
      options: [
        { label: "Age", value: "age" },
        { label: "Income", value: "income" },
        { label: "Education", value: "education" },
        { label: "Occupation", value: "occupation" },
        { label: "Marital Status", value: "marital_status" },
      ],
    },
    orderDir:{
      label: "Order",
      key: "order_dir",
      options: [
        { label: "De mayor a menor", value: "desc" },
        { label: "De menor a mayor", value: "asc" },
      ],
    },
  });


  const filterData = ref({
    filters: [
      {
          label: "Education",
          key: "education",
          type: "checkbox",
          options: [
            { label: "7th-8th", value: "7th-8th" },
            { label: "Doctorate", value: "Doctorate" },
            { label: "Bachelors", value: "Bachelors" },
            { label: "Some-college", value: "Some-college" },
            { label: "11th", value: "11th" },
            { label: "HS-grad", value: "HS-grad" },
            { label: "Prof-school", value: "Prof-school" },
            { label: "Assoc-acdm", value: "Assoc-acdm" },
            { label: "Assoc-voc", value: "Assoc-voc" },
            { label: "9th", value: "9th" },
            { label: "7th", value: "7th" },
            { label: "12th", value: "12th" },
            { label: "Masters", value: "Masters" },
            { label: "1st-4th", value: "1st-4th" },
            { label: "10th", value: "10th" },
            { label: "5th-6th", value: "5th-6th" },
            { label: "Preschool", value: "Preschool" },
          ],
      },
      {
        label: "Occupation",
        key: "occupation",
        type: "checkbox",
        options: [
          { label: "Tech Support", value: "Tech-support" },
          { label: "Craft Repair", value: "Craft-repair" },
          { label: "Other Service", value: "Other-service" },
          { label: "Sales", value: "Sales" },
          { label: "Exec Managerial", value: "Exec-managerial" },
          { label: "Prof Specialty", value: "Prof-specialty" },
          { label: "Handlers Cleaners", value: "Handlers-cleaners" },
          { label: "Machine Operator Inspector", value: "Machine-op-inspct" },
          { label: "Administrative Clerical", value: "Adm-clerical" },
          { label: "Farming Fishing", value: "Farming-fishing" },
          { label: "Transport Moving", value: "Transport-moving" },
          { label: "Private Household Service", value: "Priv-house-serv" },
          { label: "Protective Service", value: "Protective-serv" },
          { label: "Armed Forces", value: "Armed-Forces" },
        ],
      },
      {
        label: "Marital Status",
        key: "marital_status",
        type: "checkbox",
        options: [
          { label: "Married Civil Spouse", value: "Married-civ-spouse" },
          { label: "Divorced", value: "Divorced" },
          { label: "Never Married", value: "Never-married" },
          { label: "Separated", value: "Separated" },
          { label: "Widowed", value: "Widowed" },
          { label: "Married Spouse Absent", value: "Married-spouse-absent" },
          { label: "Married Armed Forces Spouse", value: "Married-AF-spouse" },
        ],
      },
      {
        label: "Income",
        key: "income",
        type: "radio",
        options: [
          { label: "Less than 50K", value: "<=50K" },
          { label: "More than 50K", value: ">50K" },
        ],
      },
    ],
    order:[
      {
        label: "Order By",
        key: "order_by",
        type: "radio",
        options: [
          { label: "Age", value: "age" },
          { label: "Income", value: "income" },
          { label: "Education", value: "education" },
          { label: "Occupation", value: "occupation" },
          { label: "Marital Status", value: "marital_status" },
        ],
      },
      {
        label: "Order",
        key: "order_dir",
        type: "radio",
        options: [
          { label: "De mayor a menor", value: "desc" },
          { label: "De menor a mayor", value: "asc" },
        ],
      },
    ],
  });

  return {
    filtersData,
    filterData
  }

});
