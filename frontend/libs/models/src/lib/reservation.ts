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
