export interface NewUser {
  name: string;
  surname: string;
  phoneNumber: string;
  address: {
    street: string;
    city: string;
    country: string;
  };
  credentials: {
    username: string;
    password: string;
  };
  role: string;
}
