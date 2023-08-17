import { BaseURL, NewUser, BookingBaseURL, TokenDTO } from '@frontend/models';
import axios from 'axios';
import jwt from 'jwt-decode';
import Swal from 'sweetalert2';

export async function LoginToFlightsApp(username: string, password: string) {
  let rsp;
  await axios({
    method: 'post',
    url: BaseURL.URL + '/login',
    data: { username, password },
  })
    .then((response) => {
      rsp = response.data;
      localStorage.setItem('id', rsp.id);
      localStorage.setItem('role', rsp.role);
      // eslint-disable-next-line @typescript-eslint/no-empty-function
    })
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    .catch((error) => {});

  return rsp;
}

export async function AddNewUser(user: NewUser) {
  await axios({
    method: 'post',
    url: BaseURL.URL + '/user/add',
    data: {
      name: user.name,
      surname: user.surname,
      phoneNumber: user.phoneNumber,
      address: user.address,
      credentials: user.credentials,
      role: 'USER',
    },
  })
    .then(() => {
      alert('Successful registration');
    })
    .catch(() => {
      alert('Error while registering, please try again');
    });
}

export async function LoginToBookingApp(username: string, password: string): Promise<TokenDTO> {
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/auth/login',
    data: { username, password },
  })
    .then(function (response) {
      const access_token = response.data;
      localStorage.setItem('jwt', access_token);
      const decodedToken: any = jwt(access_token);
      const userId = decodedToken.user_id;
      const username = decodedToken.username;
      const role = decodedToken.roles;
      localStorage.setItem('userId', userId);
      localStorage.setItem('username', username);
      localStorage.setItem('role', role);
    })
    .catch(function (error) {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again',
        showConfirmButton: false,
        position: 'bottom-right',
        timer: 3000,
        timerProgressBar: true,
        backdrop: 'none',
        width: 300,
        background: '#212121',
        color: 'white',
      });
      return error;
    });
}
