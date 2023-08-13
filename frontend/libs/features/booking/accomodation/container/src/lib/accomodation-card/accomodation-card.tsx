import { AccommodationInfo } from '@frontend/models';
import styles from './accomodation-card.module.css';
import { Paper, Grid, Typography, Divider, Button } from '@mui/material';

/* eslint-disable-next-line */
export interface AccommodationCardProps {
  accomodationInfo: AccommodationInfo;
}

const updateAccommodation = async () => {
  //await UpdateAccommodation(props.accomodationInfo);
};

const checkAvailability = async () => {
  //await CheckAvailability(props.accomodationInfo.id);
};

const deleteAccommodation = async () => {
  //await DeleteAccommodation(props.accomodationInfo.id);
};

export function AccommodationCard(props: AccommodationCardProps) {
  return (
    <Paper elevation={6} sx={{ maxWidth: '450px', margin: '1rem', padding: '1.5rem 2rem 1.5rem 2rem' }}>
      <Grid container justifyContent={'start'}>
        <Grid item xs={12}>
          <Typography variant="h3" align="center" fontWeight={550}>
            {props.accomodationInfo.name}
          </Typography>
        </Grid>

        <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Address
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Typography variant="subtitle1" align="left">
              Street: {props.accomodationInfo.address.street}
            </Typography>
            <Typography variant="subtitle1" align="left">
              City: {props.accomodationInfo.address.city}
            </Typography>
            <Typography variant="subtitle1" align="left">
              Country: {props.accomodationInfo.address.country}
            </Typography>
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

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid container direction={'row'} xs={12}>
          <Grid item xs={12}>
            <Typography variant="h5">Images</Typography>
          </Grid>
          <Grid container direction={'row'} xs={12} gap={2}>
            {props.accomodationInfo.images?.map((image, key) => (
              // eslint-disable-next-line jsx-a11y/img-redundant-alt
              <img src={image} alt="Accomodation image" width="100%" />
            ))}
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid container direction={'row'} xs={12} gap={2} justifyContent={'center'}>
          <Typography variant="h5" width={'45%'}>
            Minimum number of guests: {props.accomodationInfo.minGuests}
          </Typography>
          <Typography variant="h5" width={'45%'}>
            Maximum number of guests: {props.accomodationInfo.maxGuests}
          </Typography>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid item direction={'row'} xs={12}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Price
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Typography variant="subtitle1" align="left">
              {props.accomodationInfo.priceIsPerGuest ? 'Price is per guest' : 'Price is for whole accommodation'}
            </Typography>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid item direction={'row'} xs={12}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Automatic reservation
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Typography variant="subtitle1" align="left">
              {props.accomodationInfo.hasAutomaticReservations ? 'Accommodation has automatic reservation' : 'Accomodation does not have automatic reservation'}{' '}
            </Typography>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid item xs={12}>
          <Grid container justifyContent={'space-between'}>
            <Button
              variant="contained"
              size="small"
              onClick={updateAccommodation}
              sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
            >
              Update info
            </Button>
            <Button
              variant="contained"
              size="small"
              onClick={checkAvailability}
              sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
            >
              Check Availability
            </Button>
            <Button
              variant="contained"
              size="small"
              onClick={deleteAccommodation}
              sx={{ color: 'white', background: 'red', ':hover': { background: 'white', color: 'red' } }}
            >
              Delete
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Paper>
  );
}

export default AccommodationCard;
