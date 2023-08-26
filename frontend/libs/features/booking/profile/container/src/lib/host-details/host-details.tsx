import { UpdatePersonalData, UpdateCredentials } from '@frontend/models';
import { Button, Divider, Grid, Paper, Rating, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import styles from './host-details.module.css';
import { useState, useEffect } from 'react';
import { GetHostInformation } from '@frontend/features/booking/profile/data-access';

/* eslint-disable-next-line */
export interface HostDetailsProps {
  hostId: string;
}

export function HostDetails(props: HostDetailsProps) {
  const [hostInfo, setHostInfo] = useState<UpdatePersonalData>({
    name: '',
    surname: '',
    email: '',
    address: {
      street: '',
      city: '',
      country: '',
    },
    rating: 0,
    isOutstanding: false,
  });

  useEffect(() => {
    GetHostInformation(props.hostId).then((data) => {
      setHostInfo(data);
    });
  }, []);

  return (
    <Paper elevation={6} sx={{ padding: '1.5rem 2rem 1.5rem 2rem' }}>
      <Grid container justifyContent={'start'}>
        <Grid container justifyContent={'center'} direction={'column'}>
          <Typography variant="h3" align="center" fontWeight={550}>
            Host information
          </Typography>
          <Rating name="half-rating-read" value={hostInfo.rating} precision={0.1} readOnly size="large" sx={{ marginX: 'auto' }} />
          <Typography variant="subtitle1" align="center">
            {hostInfo.isOutstanding ? 'Outstanding host' : 'Regular host'}
          </Typography>
        </Grid>

        <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Full Name
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <div className={styles.lineContainer}>
              <Typography variant="subtitle1" align="left">
                {hostInfo.name} {hostInfo.surname}
              </Typography>
            </div>
          </Grid>
        </Grid>

        <Grid item xs={12}>
          <Divider sx={{ backgroundColor: 'grey', width: '100%', marginTop: '1.25rem', marginBottom: '1.25rem' }} />
        </Grid>

        <Grid item direction={'row'} xs={12}>
          <Grid item xs={12}>
            <Typography variant="h5" align="left">
              Contact
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <div className={styles.lineContainer}>
              <Typography variant="subtitle1" align="left">
                Email: {hostInfo.email}
              </Typography>
            </div>
          </Grid>
        </Grid>
      </Grid>
    </Paper>
  );
}

export default HostDetails;
