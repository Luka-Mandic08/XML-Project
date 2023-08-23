export interface UpdatePersonalData {
  name: string;
  surname: string;
  email: string;
  address: Address;
  rating: number;
  isOutstanding: boolean;
}

export interface UpdateCredentials {
  username: string;
  password: string;
}

export interface RegisterUser {
  name: string;
  surname: string;
  email: string;
  address: Address;

  username: string;
  password: string;
  role: string;
}

export interface Address {
  street: string;
  city: string;
  country: string;
}
