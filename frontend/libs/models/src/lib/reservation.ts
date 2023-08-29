import { Address } from './bookingUser';

export interface ReservationInfo {
  id: string;
  accommodationId: string;
  userId: string;
  start: Date;
  end: Date;
  numberOfGuests: number;
  price: number;
  status: string;
  guestName: string;
  guestSurname: string;
  guestEmail: string;
  numberOfPreviousCancellations: number;
}

export interface CreateReservation {
  accommodationId: string;
  userId: string;
  start: Date;
  end: Date;
  numberOfGuests: number;
}

export interface RecommendedFlightsProps {
  startDate: Date;
  endDate: Date;
  numberOfGuests: number;
  accommodationLocation: Address;
}
