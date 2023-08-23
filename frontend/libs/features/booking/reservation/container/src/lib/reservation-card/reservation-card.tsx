import { AccommodationInfo, ReservationInfo } from '@frontend/models';
import styles from './reservation-card.module.css';
import { Button, Divider, Paper, Typography } from '@mui/material';
import { useState, useEffect } from 'react';
import { GetAccommodationById } from '@frontend/features/booking/accomodation/data';
import { CancelReservation, ApproveReservation, DenyReservation } from '@frontend/features/booking/reservation/data-access';

/* eslint-disable-next-line */
export interface ReservationItemProps {
  reservation: ReservationInfo;
  accommodationInfo: AccommodationInfo | undefined;
  isForGuest: boolean;
  isForHost: boolean;
}

export function ReservationCard(props: ReservationItemProps) {
  const [accommodationInfo, setAccommodationInfo] = useState<AccommodationInfo>();
  const [canCancel, setCanCancel] = useState<boolean>(true);

  useEffect(() => {
    if (props.isForGuest) {
      getAccommodationInfo();
      const today = new Date();
      const checkInDate = new Date(props.reservation.start);
      if (today > checkInDate || props.reservation.status === 'Canceled' || props.reservation.status === 'Denied') {
        setCanCancel(false);
      }
    }
    if (props.isForHost) {
      setAccommodationInfo(props.accommodationInfo);
    }
  }, []);

  const getBgColor = () => {
    if (props.reservation.status === 'Pending') {
      return 'yellow';
    }
    if (props.reservation.status === 'Approved') {
      return 'green';
    }
    if (props.reservation.status === 'Denied') {
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
    await CancelReservation(props.reservation.id);
  };

  const acceptReservation = async () => {
    await ApproveReservation(props.reservation.id);
  };

  const denyReservation = async () => {
    await DenyReservation(props.reservation.id);
  };

  return (
    <Paper elevation={6} className={styles.reservationCard} style={{ border: `3px solid ${getBgColor()}`, borderRadius: '8px' }}>
      <div className={styles.reservationCardContent}>
        {props.isForGuest && <Typography variant="h4">Reservation at: {accommodationInfo?.name}</Typography>}
        <div>
          <Typography variant="h6">Check in: {props.reservation.start.split('T')[0]}</Typography>
          <Typography variant="h6">Check out: {props.reservation.end.split('T')[0]}</Typography>
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
          <Typography variant="h5">Number of guests: {props.reservation.numberOfGuests}</Typography>
        </div>
        <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

        {props.reservation.price && <Typography variant="h5">Price: {props.reservation.price}</Typography>}

        <div>
          <Typography variant="h6">Status: {props.reservation.status}</Typography>
        </div>
      </div>
      {props.isForGuest && canCancel && (
        <div className={styles.reservationCardFooter}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginY: '1rem' }} />
          <Button
            variant="contained"
            size="large"
            onClick={cancelReservation}
            sx={{ color: 'white', background: '#212121', width: 'fit-content', alignSelf: 'center', ':hover': { background: 'white', color: '#212121' } }}
          >
            Cancel reservation
          </Button>
        </div>
      )}

      {props.isForHost && (
        <div className={styles.reservationCardFooter}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginY: '1rem' }} />
          <Button
            variant="contained"
            size="large"
            onClick={acceptReservation}
            sx={{ color: 'white', background: '#212121', width: 'fit-content', alignSelf: 'center', ':hover': { background: 'white', color: '#212121' } }}
          >
            Accept reservation
          </Button>
          <Button
            variant="contained"
            size="large"
            onClick={denyReservation}
            sx={{ color: 'white', background: '#212121', width: 'fit-content', alignSelf: 'center', ':hover': { background: 'white', color: '#212121' } }}
          >
            Deny reservation
          </Button>
        </div>
      )}
    </Paper>
  );
}

export default ReservationCard;
