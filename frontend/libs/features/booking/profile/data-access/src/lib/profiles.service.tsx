import { ApiKey, BaseURL, BookingBaseURL, UpdateCredentials, UpdatePersonalData } from '@frontend/models';
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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
        text: 'Your account information has been updated',
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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

export async function GetProfileInformation(userId: string): Promise<UpdatePersonalData> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/users/' + userId,
  })
    .then((response) => {
      return response.data;
    })
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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
    url: BookingBaseURL.URL + '/auth/delete/' + localStorage.getItem('role') + '/' + localStorage.getItem('userId'),
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
      localStorage.clear();
      window.location.href = '/';
    })
    .catch((err) => {
      Swal.fire({
        icon: 'info',
        title: 'Delete account',
        text: 'You cannot delete your account because you have active bookings.',
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

export async function GetApiKey(): Promise<any> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/auth/apikey/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data;
    })
    .catch((err) => {
      // Swal.fire({
      //   icon: 'error',
      //   title: 'Error',
      //   text: 'Something went wrong, please try again\n' + err.message,
      //   showConfirmButton: false,
      //   position: 'bottom-right',
      //   timer: 3000,
      //   timerProgressBar: true,
      //   backdrop: 'none',
      //   width: 300,
      //   background: '#212121',
      //   color: 'white',
      // });
    });
}

export async function CreateApiKey(isPermanent: boolean): Promise<string> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/auth/apikey/create',
    data: {
      userId: localStorage.getItem('userId'),
      isPermanent: isPermanent,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Your API key has been created',
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.message,
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

export async function LinkToFlightsApp(username: string, password: string, apiKey: ApiKey): Promise<string> {
  return await axios({
    method: 'put',
    url: BaseURL.URL + '/user/link',
    data: {
      apiKey: apiKey,
      username: username,
      password: password,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Your account has been linked to the Flights App',
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
    .catch((err) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Something went wrong, please try again\n' + err.response.data,
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
