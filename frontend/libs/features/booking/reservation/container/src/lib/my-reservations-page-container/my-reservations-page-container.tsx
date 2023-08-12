import { Grid } from '@mui/material';
import { useEffect, useState } from 'react';
import { MyReservation } from '@frontend/models';
import ReservationItem from '../reservation-item/reservation-item';

const reservationsData: MyReservation[] = [
  {
    id: '0',
    dateFrom: '2023-08-15',
    dateTo: '2023-08-20',
    status: 'confirmed',
    accommodationId: 'abc123',
    price: 350,
  },
  {
    id: '1',
    dateFrom: '2023-09-10',
    dateTo: '2023-09-15',
    status: 'pending',
    accommodationId: 'def456',
    price: 280,
  },
];

/* eslint-disable-next-line */
export interface MyReservationsPageContainerProps {}

export function MyReservationsPageContainer(props: MyReservationsPageContainerProps) {
  const [reservations, setReservations] = useState<MyReservation[]>();
  let welcomeText;

  useEffect(() => {
    welcomeText = 'My Reservations';
    setReservations(reservationsData);
  }, []);

  return (
    <Grid container direction="row" justifyContent="center" sx={{ margin: '0', padding: '1.25rem' }}>
      <Grid item xs={5}>
        {welcomeText}
      </Grid>
      <Grid item xs={10}>
        {reservations?.map((reservation, index) => (
          <ReservationItem key={index} reservation={reservation} />
        ))}
      </Grid>
    </Grid>
  );
}

export default MyReservationsPageContainer;
