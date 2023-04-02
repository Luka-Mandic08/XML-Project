import { BaseURL, SearchFlightsDTO } from '@frontend/models';
import axios from 'axios';

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

export async function SearchFlights(dto: SearchFlightsDTO) {
  let flights;

  await axios.put(BaseURL.URL + '/flight/search', dto).then((response) => {
    flights = response.data;
  });

  return flights;
}
