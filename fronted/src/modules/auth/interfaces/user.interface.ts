export interface IUser {
  id:        number;
  username:  string;
  email:     string;
}

export interface IUserResponse {
  id:         number;
  username:   string;
  email:      string;
  config: ConfigUser;
}

export interface ConfigUser {
  education:      string;
  income:         string;
  marital_status: string;
  max_age:        number;
  min_age:        number;
  occupation:     string;
  order_by:       string;
  order_dir:      string;
}
