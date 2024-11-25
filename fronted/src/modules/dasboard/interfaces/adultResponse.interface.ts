export interface IAdultsResponse {
  count:  number;
  result: IAdult[];
}

export interface IAdult {
  age:            number;
  workclass:      string;
  fnlwgt:         number;
  education:      string;
  education_num:  number;
  marital_status: string;
  occupation:     string;
  relationship:   string;
  race:           string;
  sex:            string;
  capital_gain:   number;
  capital_loss:   number;
  hours_per_week: number;
  native_country: string;
  income:         string;
}
