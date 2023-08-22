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

export async function UpdateAccountInformation(data: UpdateCredentials): Promise<UpdateCredentials> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/auth/update',
    data: {
      userid: localStorage.getItem('userId'),
      username: data.username,
      password: data.password,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: response.data.message,
        showConfirmButton: false,
        position: 'bottom-right',
        timer: 3000,
        timerProgressBar: true,
        backdrop: 'none',
        width: 300,
        background: '#212121',
        color: 'white',
      });
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
    url: BookingBaseURL.URL + '/users/' + localStorage.getItem('userId'),
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

export async function GetHostInformation(hostId: string): Promise<UpdatePersonalData> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/users/' + hostId,
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

export async function UpdateProfileInformation(data: UpdatePersonalData): Promise<UpdatePersonalData> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/users/update',
    data: {
      id: localStorage.getItem('userId'),
      name: data.name,
      surname: data.surname,
      email: data.email,
      address: {
        street: data.address.street,
        city: data.address.city,
        country: data.address.country,
      },
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Your profile information has been updated',
        showConfirmButton: false,
        position: 'bottom-right',
        timer: 3000,
        timerProgressBar: true,
        backdrop: 'none',
        width: 300,
        background: '#212121',
        color: 'white',
      });
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

export async function DeleteAccount(): Promise<void> {
  await axios({
    method: 'delete',
    url: BookingBaseURL.URL + '/auth/delete/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Your account has been deleted',
        showConfirmButton: false,
        position: 'bottom-right',
        timer: 3000,
        timerProgressBar: true,
        backdrop: 'none',
        width: 300,
        background: '#212121',
        color: 'white',
      });
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
