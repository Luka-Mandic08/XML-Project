import { ReservationInfo } from '@frontend/models';
import { useState, useEffect } from 'react';
import styles from './guest-reservations.module.css';
import ReservationCard from '../reservation-card/reservation-card';
import { GetReservationsForGuest } from '@frontend/features/booking/reservation/data-access';
import { Typography } from '@mui/material';

/* eslint-disable-next-line */
export interface GuestReservationsProps {}

export function GuestReservations(props: GuestReservationsProps) {
  const [reservations, setReservations] = useState<ReservationInfo[]>([]);

  useEffect(() => {
    getGuestReservations();
  }, []);

  const getGuestReservations = async () => {
    const res = await GetReservationsForGuest();
    const newReservations: ReservationInfo[] = [];
    res?.forEach((reservation: any) => {
      newReservations.push({
        id: reservation.id,
        accommodationId: reservation.accommodationId,
        userId: reservation.userId,
        numberOfGuests: reservation.numberOfGuests,
        start: new Date(reservation.start.seconds * 1000),
        end: new Date(reservation.end.seconds * 1000),
        status: reservation.status,
        price: reservation.price,
      });
    });
    setReservations(newReservations);
  };

  return (
    <>
      <Typography variant="h3" sx={{ margin: '2rem' }}>
        Your reservations
      </Typography>

      <div className={styles.cardsContainer}>
        {reservations?.map((reservation, key) => (
          <ReservationCard reservation={reservation} isForGuest={true} isForHost={false} accommodationInfo={undefined} />
        ))}
      </div>
    </>
  );
}

export default GuestReservations;
