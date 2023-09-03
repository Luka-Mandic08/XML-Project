import { BookingBaseURL } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function GetNotifications(): Promise<any> {
  const userId = localStorage.getItem('userId');
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/notification/all/' + userId,
  })
    .then((response) => {
      return response.data.notifications;
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
export async function GetNotificationPreferences(): Promise<any> {
  const userId = localStorage.getItem('userId');
  console.log(userId);
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/notification/selectedtypes',
    data: { userId: userId },
  })
    .then((response) => {
      console.log(response);
      return response.data.notifications;
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

export async function UpdateNotificationPreferences(preferences: string[]): Promise<any> {
  const userId = localStorage.getItem('userId');
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/notification/selectedtypes/update',
    data: { userId: userId, selectedTypes: preferences },
  })
    .then((response) => {
      return response.data.notifications;
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
