import { SearchedAccommodationInfo } from '@frontend/models';
import styles from './searched-accommodation-card.module.css';
import { Paper, Grid, Typography, Divider, Button } from '@mui/material';
import { MakeReservationDialog } from '@frontend/features/booking/reservation/container';
import { useState } from 'react';

/* eslint-disable-next-line */
export interface SearchedAccommodationCardProps {
  searchedAccomodationInfo: SearchedAccommodationInfo;
}

export function SearchedAccommodationCard(props: SearchedAccommodationCardProps) {
  const [open, setOpen] = useState<boolean>(false);

  const createReservation = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <Paper elevation={6} sx={{ maxWidth: '450px', padding: '1.5rem 2rem 1.5rem 2rem' }}>
        <Grid container justifyContent={'start'}>
          <Grid item xs={12}>
            <Typography variant="h3" align="center" fontWeight={550}>
              {props.searchedAccomodationInfo.name}
            </Typography>
          </Grid>

          <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
            <Grid item xs={12}>
              <Typography variant="h5" align="left">
                Address
              </Typography>
            </Grid>
            <Grid item xs={12}>
              <div className={styles.lineContainer}>
                <Typography variant="subtitle1" align="left">
                  Street: {props.searchedAccomodationInfo.address.street}
                </Typography>
                <Typography variant="subtitle1" align="left">
                  City: {props.searchedAccomodationInfo.address.city}
                </Typography>
                <Typography variant="subtitle1" align="left">
                  Country: {props.searchedAccomodationInfo.address.country}
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
              {props.searchedAccomodationInfo.amenities.map((amenity, key) => (
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
              {props.searchedAccomodationInfo.images?.map((image, key) => (
                // eslint-disable-next-line jsx-a11y/img-redundant-alt
                <img src={image} alt="Accomodation image" width="100%" />
              ))}
            </Grid>
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
              <div className={styles.lineContainer}>
                <Typography variant="subtitle1" align="left">
                  Price per day: {props.searchedAccomodationInfo.unitPrice}
                </Typography>
                <Typography variant="subtitle1" align="left">
                  Total price: {props.searchedAccomodationInfo.totalPrice}
                </Typography>
              </div>
            </Grid>
          </Grid>

          <Grid item xs={12}>
            <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
          </Grid>

          <Grid item xs={12}>
            <Grid container justifyContent={'flex-end'} marginTop={'auto'}>
              <Button
                variant="contained"
                size="small"
                onClick={createReservation}
                sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
              >
                Make a reservation
              </Button>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
      <MakeReservationDialog open={open} selectedAccommodation={props.searchedAccomodationInfo} onClose={handleClose} />
    </>
  );
}

export default SearchedAccommodationCard;
