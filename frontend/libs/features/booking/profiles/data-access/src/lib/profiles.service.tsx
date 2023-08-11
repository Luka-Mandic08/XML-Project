import { BookingBaseURL, UpdateCredentials, UpdatePersonalData } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function GetAccountInformation(): Promise<UpdateCredentials> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/auth/get/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data;
    })
    .catch(() => {
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
    });
}

export async function GetProfileInformation(): Promise<UpdatePersonalData> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/user/get' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data;
    })
    .catch(() => {
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
    });
}
