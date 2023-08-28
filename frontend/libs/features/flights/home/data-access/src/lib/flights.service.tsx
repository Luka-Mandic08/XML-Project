import { ApiKey, BaseURL, SearchFlightsDTO } from '@frontend/models';
import axios from 'axios';
import Swal from 'sweetalert2';

export async function GetAllFlights() {
  let flights;

  await axios.get(BaseURL.URL + '/flight/all').then((response) => {
    flights = response.data;
  });

  return flights;
}

export async function BuyFlightTickets(flightId: string, amount: number) {
  /*
  NE KORISTITI OVO
  const buyTicketsDto = {flightId: flightId, amount: amount}
  MORA SE STAVITI DIREKT U DATA
  JER GLUPI AXIOS/GO
  */

  await axios({
    method: 'put',
    url: BaseURL.URL + '/flight/buyticket',
    data: { flightId: flightId, amount: +amount, userId: localStorage.getItem('id') },
  });
}

export async function BuyFlightTicketsFromBookingApp(flightId: string, amount: number, apiKeyValue: string) {
  await axios({
    method: 'put',
    url: BaseURL.URL + '/flight/buyticket/bookingapp',
    headers: {
      'X-API-Key': apiKeyValue,
    },
    data: { flightId: flightId, amount: 1 * amount, userId: '' },
  })
    .then((response) => {
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: 'Your flight has been booked',
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
    .catch((error) => {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: error.response.data,
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

export async function SearchFlights(dto: SearchFlightsDTO) {
  let flights;

  await axios.put(BaseURL.URL + '/flight/search', dto).then((response) => {
    flights = response.data;
  });

  return flights;
}
