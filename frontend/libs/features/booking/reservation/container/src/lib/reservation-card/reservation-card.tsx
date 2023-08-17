import { AccommodationInfo, ReservationInfo } from '@frontend/models';
import styles from './reservation-card.module.css';
import { Button, Divider, Paper, Typography } from '@mui/material';
import { useState, useEffect } from 'react';
import { GetAccommodationById } from '@frontend/features/booking/accomodation/data';

/* eslint-disable-next-line */
export interface ReservationItemProps {
  reservation: ReservationInfo;
  accommodationInfo: AccommodationInfo | undefined;
  isForGuest: boolean;
  isForHost: boolean;
}

export function ReservationCard(props: ReservationItemProps) {
  const [accommodationInfo, setAccommodationInfo] = useState<AccommodationInfo>();

  useEffect(() => {
    if (props.isForGuest) {
      getAccommodationInfo();
    }
    if (props.isForHost) {
      setAccommodationInfo(props.accommodationInfo);
    }
  }, []);

  const getAccommodationInfo = async () => {
    setAccommodationInfo(await GetAccommodationById(props.reservation.accommodationId));
  };

  const cancelReservation = async () => {
    //await CancelReservation(props.reservation.id);
  };

  return (
    <Paper elevation={6} className={styles.reservationCard}>
      <div className={styles.reservationCardContent}>
        {props.isForGuest && <Typography variant="h4">Reservation at: {accommodationInfo?.name}</Typography>}
        <div>
          <Typography variant="h6">Check in: {props.reservation.start}</Typography>
          <Typography variant="h6">Check out: {props.reservation.end}</Typography>
        </div>
        {props.isForGuest && (
          <div>
            <Typography variant="h5">Address</Typography>
            <Typography variant="h6">Street: {accommodationInfo?.address.street}</Typography>
            <Typography variant="h6">City: {accommodationInfo?.address.city}</Typography>
            <Typography variant="h6">Country: {accommodationInfo?.address.country}</Typography>
          </div>
        )}
        <div>
          <Typography variant="h6">Number of guests: {props.reservation.numberOfGuests}</Typography>
          <Typography variant="h6">Price: {props.reservation.price}</Typography>
        </div>
        <div>
          <Typography variant="h6">Status: {props.reservation.status}</Typography>
        </div>
      </div>
      {props.isForGuest && (
        <div className={styles.reservationCardFooter}>
          <Divider sx={{ backgroundColor: 'grey', flexGrow: '1', marginY: '2rem' }} />
          <Button
            variant="contained"
            size="large"
            onClick={cancelReservation}
            sx={{ color: 'white', background: 'red', width: 'fit-content', marginLeft: 'auto', ':hover': { background: 'white', color: 'red' } }}
          >
            Cancel reservation
          </Button>
        </div>
      )}
    </Paper>
  );
}

export default ReservationCard;
