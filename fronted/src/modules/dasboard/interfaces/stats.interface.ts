export interface IStastResponse {
  option:            string;
  total_count:       number;
  high_income_count: number;
  low_income_count:  number;
}


export interface IParamsStats {
  option:            string;
}


export interface IDataStats {
  option: string[],
  high_income_count: number[],
  low_income_count: number[]
}
