import { BaseURL } from '@frontend/models';
import axios from 'axios';

export async function GetAllPurchasedTickets() {
  let flights;

  await axios.get(BaseURL.URL + '/flight/user/' + localStorage.getItem('id')).then((response) => {
    flights = response.data;
  });

  return flights;
}
