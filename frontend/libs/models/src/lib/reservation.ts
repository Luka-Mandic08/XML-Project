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
