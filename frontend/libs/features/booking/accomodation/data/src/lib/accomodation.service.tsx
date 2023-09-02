import { AccommodationCreateUpdateDTO, AccommodationInfo, AvailabilityDate, BookingBaseURL, SearchedAccommodationInfo } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function GetAccomodationRatings(accomodationId: any): Promise<any> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/rating/accommodation/all/' + accomodationId,
  })
    .then((response) => {
      return response.data.accommodationRatings;
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

export async function GetHostRatings(hostId: any): Promise<any> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/rating/host/all/' + hostId,
  })
    .then((response) => {
      return response.data.hostRatings;
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

export async function DeleteAccommodationRating(ratingId: string, guestId: string): Promise<any> {
  return await axios({
    method: 'delete',
    url: BookingBaseURL.URL + '/rating/accommodation/delete',
    data: { GuestId: guestId, RatingId: ratingId },
  })
    .then((response) => {
      return response.data.hostRatings;
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

export async function DeleteHostRating(ratingId: string, guestId: string): Promise<any> {
  return await axios({
    method: 'delete',
    url: BookingBaseURL.URL + '/rating/host/delete',
    data: { GuestId: guestId, RatingId: ratingId },
  })
    .then((response) => {
      return response.data.hostRatings;
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

export async function EditHostRating(comment: any): Promise<any> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/rating/host/update',
    data: {
      id: comment.id,
      hostId: comment.hostId,
      guestId: comment.guestId,
      date: new Date(),
      score: comment.score,
      comment: comment.comment,
    },
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

export async function EditAccommodationRating(comment: any): Promise<any> {
  return await axios({
    method: 'put',
    url: BookingBaseURL.URL + '/rating/accommodation/update',
    data: {
      id: comment.id,
      accommodationId: comment.accommodationId,
      guestId: comment.guestId,
      date: new Date(),
      score: comment.score,
      comment: comment.comment,
    },
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

export async function GetAccomodationForHost(): Promise<AccommodationInfo[]> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/accommodation/all/host/' + localStorage.getItem('userId'),
  })
    .then((response) => {
      return response.data.accommodations;
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

export async function GetAllAccomodation(pageNumber: number): Promise<AccommodationInfo[]> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/accommodation/all/' + pageNumber,
  })
    .then((response) => {
      return response.data.accommodations;
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

export async function CreateUpdateAccommodationFunction(data: AccommodationCreateUpdateDTO, amenities: string[], images: string[]): Promise<void> {
  await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/accommodation/create',
    data: {
      hostId: data.hostId,
      name: data.name,
      address: data.address,
      amenities: amenities,
      images: images,
      minGuests: 1 * data.minGuests,
      maxGuests: 1 * data.maxGuests,
      priceIsPerGuest: data.priceIsPerGuest,
      hasAutomaticReservations: data.hasAutomaticReservations,
    },
  })
    .then(() => {
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
        text: 'Availability updated successfully',
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

export async function SearchAccommodation(data: any, pageNumber: number): Promise<SearchedAccommodationInfo[]> {
  const newData = {
    city: data.city,
    country: data.country,
    dateFrom: new Date(data.dateFrom),
    dateTo: new Date(data.dateTo),
    numberOfGuests: 1 * data.numberOfGuests,
    maxPrice: 1 * data.maxPrice,
    amenities: data.amenities !== '' ? data.amenities.split(',') : [],
    ownedByProminentHost: data.ownedByProminentHost,
  };
  return await axios({
    method: 'post',
    url: BookingBaseURL.URL + '/accommodation/search',
    data: {
      city: newData.city,
      country: newData.country,
      dateFrom: newData.dateFrom,
      dateTo: newData.dateTo,
      numberOfGuests: newData.numberOfGuests,
      maxPrice: newData.maxPrice,
      amenities: newData.amenities,
      ownedByProminentHost: newData.ownedByProminentHost,
      pageNumber: pageNumber,
    },
  })
    .then((response) => {
      return response.data.accommodations;
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

export async function GetAccommodationById(id: string): Promise<AccommodationInfo> {
  return await axios({
    method: 'get',
    url: BookingBaseURL.URL + '/accommodation/' + id,
  })
    .then((response) => {
      return response.data.accommodation;
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
