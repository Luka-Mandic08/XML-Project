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

  const getBgColor = () => {
    if (props.reservation.status === 'Pending') {
      return 'yellow';
    }
    if (props.reservation.status === 'Accepted') {
      return 'green';
    }
    if (props.reservation.status === 'Rejected') {
      return 'red';
    }
    if (props.reservation.status === 'Canceled') {
      return 'grey';
    }
  };

  const getAccommodationInfo = async () => {
    setAccommodationInfo(await GetAccommodationById(props.reservation.accommodationId));
  };

  const cancelReservation = async () => {
    //await CancelReservation(props.reservation.id);
  };

  return (
    <Paper elevation={6} className={styles.reservationCard} style={{ border: `3px solid ${getBgColor()}`, borderRadius: '8px' }}>
      <div className={styles.reservationCardContent}>
        {props.isForGuest && <Typography variant="h4">Reservation at: {accommodationInfo?.name}</Typography>}
        <div>
          <Typography variant="h6">Check in: {props.reservation.start}</Typography>
          <Typography variant="h6">Check out: {props.reservation.end}</Typography>
        </div>
        <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />
        {props.isForGuest && (
          <>
            <div>
              <Typography variant="h5">Location</Typography>
              <Typography variant="h6">
                {accommodationInfo?.address.street}, {accommodationInfo?.address.city}, {accommodationInfo?.address.country}
              </Typography>
            </div>
            <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />
          </>
        )}
        <div>
          <Typography variant="h5">Guest numbers</Typography>
          <Typography variant="h6">
            Minimun: {accommodationInfo?.minGuests} | Maximum: {accommodationInfo?.maxGuests}
          </Typography>
          <Typography variant="h6">Number of guests: {props.reservation.numberOfGuests}</Typography>
        </div>
        <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

        {props.reservation.price && <Typography variant="h5">Price: {props.reservation.price}</Typography>}

        <div>
          <Typography variant="h6">Status: {props.reservation.status}</Typography>
        </div>
      </div>
      {props.isForGuest && (
        <div className={styles.reservationCardFooter}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginY: '1rem' }} />
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
