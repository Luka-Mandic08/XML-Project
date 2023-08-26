import { BookingBaseURL } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function RateAccommodation(data: any): Promise<string> {
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/rating/accommodation/create',
    data: {
      accommodationId: data.ratedId,
      guestId: data.guestId,
      date: new Date(data.date),
      score: data.score,
      comment: data.comment,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Accommodation rated successfully',
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

export async function RateHost(data: any): Promise<string> {
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/rating/host/create',
    data: {
      hostId: data.ratedId,
      guestId: data.guestId,
      date: new Date(data.date),
      score: data.score,
      comment: data.comment,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Host rated successfully',
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
