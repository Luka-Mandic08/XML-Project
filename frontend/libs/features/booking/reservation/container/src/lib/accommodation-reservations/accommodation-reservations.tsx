import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { ReservationInfo } from '@frontend/models';
import { Button, Typography } from '@mui/material';
import { useState, useEffect } from 'react';
import ReservationCard from '../reservation-card/reservation-card';
import styles from './accommodation-reservations.module.css';
import { GetAccommodationReservations } from '@frontend/features/booking/reservation/data-access';

/* eslint-disable-next-line */
export interface AccommodationReservationsProps {}

export function AccommodationReservations(props: AccommodationReservationsProps) {
  const [view, setView] = useState<string>('Upcoming'); // ['Upcoming', 'Past'
  const [upcomingReservations, setUpcomingReservations] = useState<ReservationInfo[]>([]);
  const [pastReservations, setPastReservations] = useState<ReservationInfo[]>([]);
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);

  useEffect(() => {
    getAccommodationReservations();
  }, []);

  const getAccommodationReservations = async () => {
    const res = await GetAccommodationReservations(selectedAccommodation.id);
    const newUpcomingReservations: ReservationInfo[] = [];
    const newPastReservations: ReservationInfo[] = [];

    res.futureReservations?.forEach((reservation: any) => {
      newUpcomingReservations.push({
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
    setUpcomingReservations(newUpcomingReservations);

    res.pastReservations?.forEach((reservation: any) => {
      newPastReservations.push({
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

    setPastReservations(newPastReservations);
  };
  const switchView = () => {
    if (view === 'Upcoming') {
      setView('Past');
    } else {
      setView('Upcoming');
    }
  };

  return (
    <>
      <div className={styles.lineContainer}>
        <Typography variant="h3">
          {view} reservations at: {selectedAccommodation.name}
        </Typography>{' '}
        <Button
          variant="contained"
          size="large"
          onClick={switchView}
          sx={{ color: 'white', background: '#212121', width: 'fit-content', alignSelf: 'center', ':hover': { background: 'white', color: '#212121' } }}
        >
          View {view === 'Upcoming' ? 'Past' : 'Upcoming'} reservations
        </Button>
      </div>

      {view === 'Upcoming' && (
        <>
          <div className={styles.cardsContainer}>
            {upcomingReservations?.map((reservation, key) => (
              <ReservationCard reservation={reservation} isForGuest={false} isForHost={true} accommodationInfo={undefined} />
            ))}
          </div>
          {upcomingReservations.length === 0 && (
            <Typography variant="h4" sx={{ margin: '2rem' }}>
              You have no upcoming reservations
            </Typography>
          )}
        </>
      )}

      {view === 'Past' && (
        <>
          <div className={styles.cardsContainer}>
            {pastReservations?.map((reservation, key) => (
              <ReservationCard reservation={reservation} isForGuest={false} isForHost={true} accommodationInfo={undefined} />
            ))}
          </div>
          {pastReservations.length === 0 && (
            <Typography variant="h4" sx={{ margin: '2rem' }}>
              You have no past reservations
            </Typography>
          )}
        </>
      )}
    </>
  );
}

export default AccommodationReservations;
