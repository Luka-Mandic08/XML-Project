// eslint-disable-next-line @nrwl/nx/enforce-module-boundaries
import { AccommodationDetails, AvailabilityCalendar } from '@frontend/features/booking/accomodation/container';
import styles from './make-reservation.module.css';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { Paper, Typography, Button, Grid } from '@mui/material';
import { useForm } from 'react-hook-form';
import { RateAccommodationOrHostForm } from '@frontend/features/booking/rating/container';
import { HostDetails } from '@frontend/features/booking/profile/container';

/* eslint-disable-next-line */
export interface MakeReservationProps {}

export function MakeReservation(props: MakeReservationProps) {
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      accommodationId: selectedAccommodation.id,
      dateFrom: '',
      dateTo: '',
      numberOfGuests: 0,
    },
  });

  const onSubmitAvailabilityDates = async (data: any) => {
    data.dateFrom = new Date(data.dateFrom);
    data.dateTo = new Date(data.dateTo);
    //const res = await UpdateAvailableDatesForAccommodation(data);
  };

  return (
    <>
      {selectedAccommodation.id !== '' && (
        <>
          <div className={styles.flexRow}>
            <AccommodationDetails hasMargin={'0 !important'} />
            <div className={styles.flexColumn}>
              <HostDetails hostId={selectedAccommodation.hostId} />
              <RateAccommodationOrHostForm hostId={selectedAccommodation.hostId} />
              <RateAccommodationOrHostForm accommodationId={selectedAccommodation.id} />
            </div>
          </div>
          <div className={styles.inlineGrid}>
            <AvailabilityCalendar shouldRenderCalendar={true} />
            <Paper elevation={6} className={styles.updateAvailabilityForm}>
              <Typography variant="h5" align="left">
                Make a reservation
              </Typography>
              <form onSubmit={handleSubmit(onSubmitAvailabilityDates)}>
                <div className={styles.inputContainer}>
                  <input
                    type="date"
                    id="dateFrom"
                    value={watch('dateFrom')}
                    {...register('dateFrom', {
                      required: 'This field is required.',
                      min: { value: new Date().toISOString(), message: 'Date must be in the future.' },
                    })}
                  />
                  <label className={styles.label} htmlFor="dateFrom" id="label-dateFrom">
                    <div className={styles.text}>From</div>
                  </label>
                  <label className={styles.errorLabel}>{errors.dateFrom?.message}</label>
                </div>

                <div className={styles.inputContainer}>
                  <input
                    type="date"
                    id="dateTo"
                    value={watch('dateTo')}
                    {...register('dateTo', {
                      required: 'This field is required.',
                      min: { value: watch('dateFrom'), message: 'Selected date is before starting date.' },
                    })}
                  />
                  <label className={styles.label} htmlFor="dateTo" id="label-dateTo">
                    <div className={styles.text}>To</div>
                  </label>
                  <label className={styles.errorLabel}>{errors.dateTo?.message}</label>
                </div>

                <div className={styles.inputContainer}>
                  <input
                    type="number"
                    id="numberOfGuests"
                    value={watch('numberOfGuests')}
                    {...register('numberOfGuests', {
                      required: 'This field is required.',
                      min: { value: 1, message: 'Number of guests must be at least 1.' },
                      max: { value: selectedAccommodation.maxGuests, message: 'Number of guests must be less or equal than maximum number of guests.' },
                    })}
                  />
                  <label className={styles.label} htmlFor="numberOfGuests" id="label-numberOfGuests">
                    <div className={styles.text}>Number of guests</div>
                  </label>
                  <label className={styles.errorLabel}>{errors.numberOfGuests?.message}</label>
                </div>
                <div className={styles.inputContainer} style={{ justifyContent: 'center' }}>
                  <Button
                    variant="contained"
                    size="large"
                    type="submit"
                    sx={{ color: 'white', background: '#212121', height: '48px', ':hover': { background: 'white', color: '#212121' } }}
                  >
                    Reserve
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
    </>
  );
}

export default MakeReservation;
