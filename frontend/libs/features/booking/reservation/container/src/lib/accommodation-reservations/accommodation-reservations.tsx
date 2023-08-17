import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { ReservationInfo } from '@frontend/models';
import { Typography } from '@mui/material';
import { useState, useEffect } from 'react';
import ReservationCard from '../reservation-card/reservation-card';
import styles from './accommodation-reservations.module.css';
import { GetAccommodationReservations } from '@frontend/features/booking/reservation/data-access';

/* eslint-disable-next-line */
export interface AccommodationReservationsProps {}

export function AccommodationReservations(props: AccommodationReservationsProps) {
  const [reservations, setReservations] = useState<ReservationInfo[]>([]);
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);

  useEffect(() => {
    getAccommodationReservations();
  }, []);

  const getAccommodationReservations = async () => {
    setReservations(await GetAccommodationReservations(selectedAccommodation.id));
  };

  return (
    <>
      <Typography variant="h3" sx={{ margin: '2rem' }}>
        Reservations at: {selectedAccommodation.name}
      </Typography>

      <div className={styles.cardsContainer}>
        {reservations?.map((reservation, key) => (
          <ReservationCard reservation={reservation} isForGuest={false} isForHost={true} accommodationInfo={selectedAccommodation} />
        ))}
      </div>
    </>
  );
}

export default AccommodationReservations;
