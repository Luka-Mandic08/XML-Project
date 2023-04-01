import { BaseURL, NewFlight } from '@frontend/models';
import axios from 'axios';

export async function AddNewFlight(flight: NewFlight) {
  console.log(flight);
  await axios({
    method: 'post',
    url: BaseURL.URL + '/flight/add',
    data: {
      startdate: flight.startdate.toString() + ':00Z',
      arrivaldate: flight.arrivaldate.toString() + ':00Z',
      destination: flight.destination,
      start: flight.start,
      price: flight.price,
      remainingtickets: flight.totaltickets,
      totaltickets: flight.totaltickets,
    },
  });
}

export async function DeleteFlight(flightId: string) {
  await axios({
    method: 'delete',
    url: BaseURL.URL + '/flight/delete/' + flightId,
    data: {},
  });
}
