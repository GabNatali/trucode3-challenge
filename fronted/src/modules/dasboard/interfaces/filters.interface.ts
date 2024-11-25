export interface IFilter {
  education: string | string[];
  marital_status: string | string[];
  occupation: string | string[];
  income: string;
  order_by: string;
  order_dir: string;
  max_age: number;
  min_age: number;
}



export interface IParams {
  education?: string[];
  marital_status?: string[];
  occupation?: string[];
  income?: string;
  order_by?: string;
  order_dir?: string;
  max_age?: number;
  min_age?: number;
  page?: number;
  page_size?: number;
}
