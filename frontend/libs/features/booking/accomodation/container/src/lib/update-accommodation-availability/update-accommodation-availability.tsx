import { GetAvailableDatesForAccommodation, UpdateAvailableDatesForAccommodation } from '@frontend/features/booking/accomodation/data';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { AvailabilityDate } from '@frontend/models';
import { Grid, Typography, Button, Paper } from '@mui/material';
import { useState, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import styles from './update-accommodation-availability.module.css';
import AvailabilityCalendar from '../availability-calendar/availability-calendar';

/* eslint-disable-next-line */
export interface UpdateAccommodationAvailabilityProps {}

export function UpdateAccommodationAvailability(props: UpdateAccommodationAvailabilityProps) {
  const [renderCalendar, setRenderCalendar] = useState<boolean>(false);
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);

  const {
    register: registerAvailabilityDates,
    handleSubmit: handleSubmitAvailabilityDates,
    watch: watchAvailabilityDates,
    formState: { errors: errorsAvailabilityDates },
  } = useForm({
    defaultValues: {
      accommodationId: selectedAccommodation.id,
      dateFrom: '',
      dateTo: '',
      price: 0,
    },
  });

  const onSubmitAvailabilityDates = async (data: any) => {
    data.dateFrom = new Date(data.dateFrom);
    data.dateTo = new Date(data.dateTo);
    const res = await UpdateAvailableDatesForAccommodation(data);

    if (res) {
      shouldRenderCalendar();
    }
  };

  const shouldRenderCalendar = () => {
    if (selectedAccommodation.id !== '') setRenderCalendar(!renderCalendar);
  };

  return (
    <div style={{ margin: '2rem' }}>
      {selectedAccommodation.id !== '' && (
        <>
          <Grid container alignItems={'left'} direction={'column'}>
            <Grid item marginBottom={'0.5rem'}>
              <Typography variant="h4" align="left">
                Availability for accommodation: {selectedAccommodation.name}
              </Typography>
            </Grid>
          </Grid>
          <div className={styles.inlineGrid}>
            <AvailabilityCalendar shouldRenderCalendar={renderCalendar} />
            <Paper elevation={6} className={styles.updateAvailabilityForm}>
              <Typography variant="h5" align="left">
                Update availability
              </Typography>
              <form onSubmit={handleSubmitAvailabilityDates(onSubmitAvailabilityDates)}>
                <div className={styles.inputContainer}>
                  <input
                    type="date"
                    id="dateFrom"
                    value={watchAvailabilityDates('dateFrom')}
                    {...registerAvailabilityDates('dateFrom', {
                      required: 'This field is required.',
                      min: { value: new Date().toISOString(), message: 'Date must be in the future.' },
                    })}
                  />
                  <label className={styles.label} htmlFor="dateFrom" id="label-dateFrom">
                    <div className={styles.text}>From</div>
                  </label>
                  <label className={styles.errorLabel}>{errorsAvailabilityDates.dateFrom?.message}</label>
                </div>

                <div className={styles.inputContainer}>
                  <input
                    type="date"
                    id="dateTo"
                    value={watchAvailabilityDates('dateTo')}
                    {...registerAvailabilityDates('dateTo', {
                      required: 'This field is required.',
                      min: { value: watchAvailabilityDates('dateFrom'), message: 'Selected date is before starting date.' },
                    })}
                  />
                  <label className={styles.label} htmlFor="dateTo" id="label-dateTo">
                    <div className={styles.text}>To</div>
                  </label>
                  <label className={styles.errorLabel}>{errorsAvailabilityDates.dateTo?.message}</label>
                </div>

                <div className={styles.inputContainer}>
                  <input
                    type="number"
                    id="price"
                    value={watchAvailabilityDates('price')}
                    {...registerAvailabilityDates('price', {
                      required: 'This field is required.',
                      min: 0,
                    })}
                  />
                  <label className={styles.label} htmlFor="price" id="label-price">
                    <div className={styles.text}>Price</div>
                  </label>
                  <label className={styles.errorLabel}>{errorsAvailabilityDates.price?.message}</label>
                </div>
                <div className={styles.inputContainer} style={{ justifyContent: 'center' }}>
                  <Button
                    variant="contained"
                    size="large"
                    type="submit"
                    sx={{ color: 'white', background: '#212121', height: '48px', ':hover': { background: 'white', color: '#212121' } }}
                  >
                    Update availability
                  </Button>
                </div>
              </form>
            </Paper>
          </div>
        </>
      )}
      {selectedAccommodation.id === '' && (
        <Typography variant="h4" align="left">
          Please select accommodation to see availability.
        </Typography>
      )}
    </div>
  );
}

export default UpdateAccommodationAvailability;
