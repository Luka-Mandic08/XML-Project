import { ReservationInfo } from '@frontend/models';
import { useState, useEffect } from 'react';
import styles from './guest-reservations.module.css';
import ReservationCard from '../reservation-card/reservation-card';
import { GetReservationsForGuest } from '@frontend/features/booking/reservation/data-access';
import { Button, Typography } from '@mui/material';

/* eslint-disable-next-line */
export interface GuestReservationsProps {}

export function GuestReservations(props: GuestReservationsProps) {
  const [view, setView] = useState<string>('upcoming'); // ['upcoming', 'past'
  const [upcomingReservations, setUpcomingReservations] = useState<ReservationInfo[]>([]);
  const [pastReservations, setPastReservations] = useState<ReservationInfo[]>([]);

  useEffect(() => {
    getGuestReservations();
  }, []);

  const getGuestReservations = async () => {
    const res = await GetReservationsForGuest();
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
        guestEmail: '',
        guestName: '',
        guestSurname: '',
        numberOfPreviousCancellations: -1,
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
        guestEmail: '',
        guestName: '',
        guestSurname: '',
        numberOfPreviousCancellations: -1,
      });
    });

    setPastReservations(newPastReservations);
  };

  const switchView = () => {
    if (view === 'upcoming') {
      setView('past');
    } else {
      setView('upcoming');
    }
  };

  return (
    <>
      <div className={styles.lineContainer}>
        <Typography variant="h3">Your {view} reservations</Typography>
        <Button
          variant="contained"
          size="large"
          onClick={switchView}
          sx={{ color: 'white', background: '#212121', width: 'fit-content', alignSelf: 'center', ':hover': { background: 'white', color: '#212121' } }}
        >
          View {view === 'upcoming' ? 'past' : 'upcoming'} reservations
        </Button>
      </div>
      {view === 'upcoming' && (
        <>
          <div className={styles.cardsContainer}>
            {upcomingReservations?.map((reservation, key) => (
              <ReservationCard reservation={reservation} isForGuest={true} isForHost={false} accommodationInfo={undefined} />
            ))}
          </div>
          {upcomingReservations.length === 0 && (
            <Typography variant="h4" sx={{ margin: '2rem' }}>
              You have no upcoming reservations
            </Typography>
          )}
        </>
      )}

      {view === 'past' && (
        <>
          <div className={styles.cardsContainer}>
            {pastReservations?.map((reservation, key) => (
              <ReservationCard reservation={reservation} isForGuest={true} isForHost={false} accommodationInfo={undefined} />
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

export default GuestReservations;
