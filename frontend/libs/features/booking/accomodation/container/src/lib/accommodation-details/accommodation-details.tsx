import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import styles from './accommodation-details.module.css';
import Carousel from 'react-material-ui-carousel';
import { Divider, Grid, Paper, Typography } from '@mui/material';

/* eslint-disable-next-line */
export interface AccommodationDetailsProps {}

export function AccommodationDetails(props: AccommodationDetailsProps) {
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);

  return (
    <Paper elevation={6} className={styles.detailsContainer}>
      <Typography variant="h3" align="center" fontWeight={550}>
        {selectedAccommodation.name}
      </Typography>
      <Carousel autoPlay={false} navButtonsAlwaysVisible={true} duration={700} height={700}>
        {selectedAccommodation.images.map((item, i) => (
          <>
            <Typography variant="h4" align="left" marginBottom={'1rem'}>
              Images
            </Typography>
            <div className={styles.imageContainer}>
              <img key={i} src={item} alt="Accomodation image" className={styles.imageStyle} />
            </div>
          </>
        ))}
      </Carousel>

      <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
        <Grid item marginBottom={'1rem'}>
          <Typography variant="h4" align="left">
            Amenities
          </Typography>
        </Grid>
        <div className={styles.amenitiesContainer}>
          {selectedAccommodation.amenities?.map((amenity, idx) => (
            <div className={styles.amenityCard}>
              <Typography>
                {idx + 1}. {amenity}
              </Typography>
            </div>
          ))}
        </div>
      </Grid>

      <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

      <Grid container direction={'row'} xs={12}>
        <Grid item xs={12}>
          <Typography variant="h4" align="left">
            Location
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="h6" align="left">
            {selectedAccommodation.address.street}, {selectedAccommodation.address.city}, {selectedAccommodation.address.country}
          </Typography>
        </Grid>
      </Grid>

      <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

      <Grid container direction={'row'} xs={12}>
        <Grid item xs={12}>
          <Typography variant="h4" align="left">
            Number of guests
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <div className={styles.lineContainer}>
            <Typography variant="h6" align="left">
              Minimun: {selectedAccommodation.minGuests}
            </Typography>
            <Typography variant="h6" align="left">
              Maximum: {selectedAccommodation.maxGuests}
            </Typography>
          </div>
        </Grid>
      </Grid>

      <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

      <Grid item direction={'row'} xs={12}>
        <Grid item xs={12}>
          <Typography variant="h4" align="left">
            Price
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="h6" align="left">
            {selectedAccommodation.priceIsPerGuest ? 'Price is per guest' : 'Price is for whole accommodation'}
          </Typography>
        </Grid>
      </Grid>

      <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />

      <Grid item direction={'row'} xs={12}>
        <Grid item xs={12}>
          <Typography variant="h4" align="left">
            Automatic reservation
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="h6" align="left">
            {selectedAccommodation.hasAutomaticReservations ? 'Accommodation has automatic reservation' : 'Accomodation does not have automatic reservation'}{' '}
          </Typography>
        </Grid>
      </Grid>
    </Paper>
  );
}

export default AccommodationDetails;
