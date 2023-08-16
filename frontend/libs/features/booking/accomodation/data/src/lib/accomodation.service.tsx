import { AccommodationCreateUpdateDTO, AccommodationInfo, AvailabilityDate, BookingBaseURL } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function GetAccomodationForHost(): Promise<AccommodationInfo[]> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/accommodation/all/host/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data.accommodations;
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

export async function CreateUpdateAccommodationFunction(data: AccommodationCreateUpdateDTO, amenities: string[], images: File[]): Promise<void> {
  await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/accommodation/create',
    data: {
      hostId: data.hostId,
      name: data.name,
      address: data.address,
      amenities: amenities,
      images: [],
      minGuests: 1 * data.minGuests,
      maxGuests: 1 * data.maxGuests,
      priceIsPerGuest: data.priceIsPerGuest,
      hasAutomaticReservations: data.hasAutomaticReservations,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Accommodation created successfully',
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

export async function GetAvailableDatesForAccommodation(data: any): Promise<AvailabilityDate[]> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/accommodation/availability',
    data: {
      accommodationid: data.accommodationId,
      dateFrom: data.dateFrom,
      dateTo: data.dateTo,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Accommodation created successfully',
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

export async function UpdateAvailableDatesForAccommodation(data: any): Promise<string> {
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/accommodation/updateAvailability',
    data: {
      accommodationid: data.accommodationId,
      dateFrom: data.dateFrom,
      dateTo: data.dateTo,
      price: data.price,
    },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Availability updated successfully\n' + response.data,
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
