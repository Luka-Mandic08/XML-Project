import { Button, Grid, Typography } from '@mui/material';
import styles from './host-accomodation.module.css';
import { GetAccomodationForHost } from '@frontend/features/booking/accomodation/data';
import { AccommodationInfo, BookingAppRoutes } from '@frontend/models';
import { useEffect, useState } from 'react';
import AccommodationCard from '../accomodation-card/accomodation-card';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface HostAccomodationProps {}

export function HostAccomodation(props: HostAccomodationProps) {
  const [accomodationInfo, setAccomodationInfo] = useState<AccommodationInfo[]>([]);

  const navigate = useNavigate();

  useEffect(() => {
    getAccomodationForHost();
  }, []);

  const getAccomodationForHost = async () => {
    setAccomodationInfo(await GetAccomodationForHost());
  };

  const newAccommodation = () => {
    navigate(BookingAppRoutes.CreateAccommodation);
  };

  return (
    <>
      <Grid container direction={'row'} justifyContent={'space-between'} marginY={'1rem'} alignItems={'center'}>
        <Grid item marginBottom={'1rem'} paddingX={'1rem'}>
          <Typography variant="h2" align="center">
            Your Accommodations
          </Typography>
        </Grid>
        <Grid item marginBottom={'1rem'} paddingX={'1rem'}>
          <Button
            variant="contained"
            size="large"
            onClick={newAccommodation}
            sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
          >
            Add Accommodation
          </Button>
        </Grid>
      </Grid>
      <div className={styles.cardsContainer}>
        {accomodationInfo?.map((accomodation, key) => (
          <AccommodationCard accomodationInfo={accomodation} isForHost={true} />
        ))}
      </div>
    </>
  );
}

export default HostAccomodation;
