import { AccommodationInfo, BookingAppRoutes } from '@frontend/models';
import styles from './accomodation-card.module.css';
import { Paper, Grid, Typography, Divider, Button, Rating } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { set } from 'react-hook-form';

/* eslint-disable-next-line */
export interface AccommodationCardProps {
  accomodationInfo: AccommodationInfo;
  isForHost: boolean;
}

export function AccommodationCard(props: AccommodationCardProps) {
  const setSelectedAccommodation = useSelectedAccommodationStore((state) => state.setSelectedAccommodation);
  const navigate = useNavigate();

  const checkAvailability = async () => {
    setSelectedAccommodation(props.accomodationInfo);
    navigate(BookingAppRoutes.AvailabilityCalendar);
  };

  const accommodationDetails = async () => {
    setSelectedAccommodation(props.accomodationInfo);
    navigate(BookingAppRoutes.AccommodationDetails);
  };

  const accommodationReservations = async () => {
    setSelectedAccommodation(props.accomodationInfo);
    navigate(BookingAppRoutes.AccommodationReservations);
  };

  const makeReservation = async () => {
    setSelectedAccommodation(props.accomodationInfo);
    navigate(BookingAppRoutes.MakeReservation);
  };

  return (
    <Paper elevation={6} sx={{ width: '500px', padding: '1.5rem 2rem 1.5rem 2rem' }}>
      <Grid container justifyContent={'start'}>
        <Grid item xs={12}>
          <Typography variant="h3" align="center" fontWeight={550}>
            {props.accomodationInfo.name}
          </Typography>
        </Grid>

        <Grid container justifyContent={'center'}>
          {props.accomodationInfo.images !== undefined && (
            <img src={props.accomodationInfo.images[0]} alt="Accomodation image" className={styles.imageContainer} />
          )}
        </Grid>

        <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Location
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <div className={styles.lineContainer}>
              <Typography variant="subtitle1" align="left">
                {props.accomodationInfo.address.street}, {props.accomodationInfo.address.city}, {props.accomodationInfo.address.country}
              </Typography>
            </div>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid item direction={'row'} xs={12}>
          <Grid item>
            <Typography variant="h5">Amenities</Typography>
          </Grid>
          <Grid container direction={'row'} xs={12} gap={2}>
            {props.accomodationInfo.amenities.map((amenity, key) => (
              <Typography variant="subtitle1">
                {key + 1}) {amenity}
              </Typography>
            ))}
          </Grid>
        </Grid>

        {!props.isForHost && (
          <>
            <Grid item xs={12}>
              <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
            </Grid>
            <Grid item xs={12}>
              <Grid container justifyContent={'space-between'}>
                <div>
                  <Rating name="half-rating-read" value={props.accomodationInfo.rating} precision={0.1} readOnly />
                  <Typography>Score: {props.accomodationInfo.rating}</Typography>
                </div>
                <Button
                  variant="contained"
                  size="small"
                  onClick={makeReservation}
                  sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
                >
                  Details
                </Button>
              </Grid>
            </Grid>
          </>
        )}

        {props.isForHost && (
          <>
            <Grid item xs={12}>
              <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
            </Grid>

            <Grid item xs={12}>
              <Grid container justifyContent={'space-between'}>
                <div>
                  <Rating name="half-rating-read" value={props.accomodationInfo.rating} precision={0.1} readOnly />
                  <Typography>Score: {props.accomodationInfo.rating ? props.accomodationInfo.rating.toString() : 0}</Typography>
                </div>
                <Button
                  variant="contained"
                  size="small"
                  onClick={accommodationDetails}
                  sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
                >
                  Details
                </Button>
                <Button
                  variant="contained"
                  size="small"
                  onClick={accommodationReservations}
                  sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
                >
                  Reservations
                </Button>
                <Button
                  variant="contained"
                  size="small"
                  onClick={checkAvailability}
                  sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
                >
                  Availability
                </Button>
              </Grid>
            </Grid>
          </>
        )}
      </Grid>
    </Paper>
  );
}

export default AccommodationCard;
