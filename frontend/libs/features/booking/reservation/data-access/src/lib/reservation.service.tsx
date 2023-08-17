import { BookingBaseURL, ReservationInfo } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function MakeReservation(data: any): Promise<string> {
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/reservation/request',
    data: {
      accommodationId: data.accommodationId,
      start: data.startDate,
      end: data.endDate,
      userId: data.userId,
      numberOfGuests: 1 * data.numberOfGuests,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Reservation request sent.\n' + response.data,
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

export async function GetReservationsForGuest(): Promise<ReservationInfo[]> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/reservation/getAllByUserId/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data.reservation;
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

export async function GetAccommodationReservations(id: string): Promise<ReservationInfo[]> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/reservation/all/accommodation/' + id,
  })
    .then((response) => {
      return response.data.reservation;
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
