export interface FormErrors {
  first_name?: string;
  last_name?: string;
  username?: string;
  email?: string;
  password?: string;
}

export interface User {
  id?: number;
  username?: string;
  email?: string;
}
