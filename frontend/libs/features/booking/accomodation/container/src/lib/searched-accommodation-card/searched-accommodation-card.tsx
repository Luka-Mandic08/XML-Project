import { SearchedAccommodationInfo } from '@frontend/models';
import styles from './searched-accommodation-card.module.css';
import { Paper, Grid, Typography, Divider, Button } from '@mui/material';
// eslint-disable-next-line @nrwl/nx/enforce-module-boundaries
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

          <Grid container justifyContent={'center'}>
            {props.searchedAccomodationInfo.images !== undefined && (
              // eslint-disable-next-line jsx-a11y/img-redundant-alt
              <img src={props.searchedAccomodationInfo.images[0]} alt="Accomodation image" className={styles.imageContainer} />
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
                  {props.searchedAccomodationInfo.address.street}, {props.searchedAccomodationInfo.address.city},{' '}
                  {props.searchedAccomodationInfo.address.country}
                </Typography>
              </div>
            </Grid>
          </Grid>

          <Grid item xs={12}>
            <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
          </Grid>

          <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
            <Grid item marginBottom={'1rem'}>
              <Typography variant="h4" align="left">
                Amenities
              </Typography>
            </Grid>
            <div className={styles.amenitiesContainer}>
              {props.searchedAccomodationInfo.amenities?.map((amenity, idx) => (
                <div className={styles.amenityCard}>
                  <Typography>
                    {idx + 1}. {amenity}
                  </Typography>
                </div>
              ))}
            </div>
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
                Price per day: {props.searchedAccomodationInfo.unitPrice}
              </Typography>
              <Typography variant="subtitle1" align="left">
                Total price: {props.searchedAccomodationInfo.totalPrice}
              </Typography>
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
