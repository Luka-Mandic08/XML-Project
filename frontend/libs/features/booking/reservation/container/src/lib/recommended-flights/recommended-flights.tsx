import { useRecomendedFlightdPropsStore } from '@frontend/features/booking/store/container';
import styles from './recommended-flights.module.css';
import { useForm } from 'react-hook-form';
import { Button, Grid, Typography } from '@mui/material';
import { Flight, SearchFlightsDTO } from '@frontend/models';
import { useState } from 'react';
import { FlightItem } from '@frontend/features/flights/home/container';
import { SearchFlights } from '@frontend/features/flights/home/data-access';
import { start } from 'repl';

/* eslint-disable-next-line */
export interface RecommendedFlightsProps {}

export function RecommendedFlights(props: RecommendedFlightsProps) {
  const recommendedFlightsProps = useRecomendedFlightdPropsStore((state) => state.recomendedFlightdProps);
  const [startFlights, setStartFlights] = useState<Flight[]>();
  const [endFlights, setEndFlights] = useState<Flight[]>();

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      startDate: recommendedFlightsProps.startDate.toISOString().split('T')[0],
      startLocation: '',
      startDestination: recommendedFlightsProps.accommodationLocation.city,
      startNumberOfTickets: recommendedFlightsProps.numberOfGuests,

      endDate: recommendedFlightsProps.endDate.toISOString().split('T')[0],
      endLocation: recommendedFlightsProps.accommodationLocation.city,
      endDestination: '',
      endNumberOfTickets: recommendedFlightsProps.numberOfGuests,
    },
  });

  const onSubmit = async (data: any) => {
    const startDto = new SearchFlightsDTO();
    const endDto = new SearchFlightsDTO();

    startDto.setFields(data.startDate, data.startLocation, data.startDestination, parseInt(data.startNumberOfTickets));
    endDto.setFields(data.endDate, data.endLocation, data.endDestination, parseInt(data.endNumberOfTickets));

    console.log(startDto, endDto);

    SearchFlights(startDto)
      .then((result) => {
        if (result !== '') setStartFlights(result);
        else setStartFlights([]);
      })
      .catch((error) => {
        console.error(error);
      });

    SearchFlights(endDto)
      .then((result) => {
        if (result !== '') setEndFlights(result);
        else setEndFlights([]);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.lineContainer}>
          <Typography variant="h4" marginBottom={'2rem'}>
            Departing flight
          </Typography>
          <Typography variant="h6" marginBottom={'2rem'} paddingTop={'8px'}>
            Departing date: {recommendedFlightsProps.startDate.toDateString()}
          </Typography>

          <div className={styles.inputContainer}>
            <input type="text" id="startLocation" value={watch('startLocation')} {...register('startLocation', { required: 'This field is required.' })} />
            <label className={styles.label} htmlFor="startLocation" id="label-startLocation">
              <div className={styles.text}>Departing from</div>
            </label>
            <label className={styles.errorLabel}>{errors.startLocation?.message}</label>
          </div>
          <Typography variant="h6" marginBottom={'2rem'} paddingTop={'8px'}>
            Landing at: {recommendedFlightsProps.accommodationLocation.city}, {recommendedFlightsProps.accommodationLocation.country}
          </Typography>
          <div className={styles.inputContainer}>
            <input
              type="number"
              id="startNumberOfTickets"
              value={watch('startNumberOfTickets')}
              {...register('startNumberOfTickets', {
                required: 'This field is required.',
                min: { value: 1, message: 'Minimum number of tickets is 1.' },
              })}
            />
            <label className={styles.label} htmlFor="startNumberOfTickets" id="label-startNumberOfTickets">
              <div className={styles.text}>Number of tickets</div>
            </label>
            <label className={styles.errorLabel}>{errors.startNumberOfTickets?.message}</label>
          </div>
        </div>
        <div className={styles.lineContainer} style={{ marginTop: 0 }}>
          <Typography variant="h4">Returning flight</Typography>
          <Typography variant="h6" marginBottom={'2rem'} paddingTop={'8px'}>
            Returning date: {recommendedFlightsProps.endDate.toDateString()}
          </Typography>
          <Typography variant="h6" marginBottom={'2rem'} paddingTop={'8px'}>
            Departing from: {recommendedFlightsProps.accommodationLocation.city}, {recommendedFlightsProps.accommodationLocation.country}
          </Typography>
          <div className={styles.inputContainer}>
            <input type="text" id="endDestination" value={watch('endDestination')} {...register('endDestination', { required: 'This field is required.' })} />
            <label className={styles.label} htmlFor="endDestination" id="label-endDestination">
              <div className={styles.text}>Landing at</div>
            </label>
            <label className={styles.errorLabel}>{errors.endDestination?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="number"
              id="endNumberOfTickets"
              value={watch('endNumberOfTickets')}
              {...register('endNumberOfTickets', {
                required: 'This field is required.',
                min: { value: 1, message: 'Minimum number of tickets is 1.' },
              })}
            />
            <label className={styles.label} htmlFor="endNumberOfTickets" id="label-endNumberOfTickets">
              <div className={styles.text}>Number of tickets</div>
            </label>
            <label className={styles.errorLabel}>{errors.endNumberOfTickets?.message}</label>
          </div>
        </div>
        <Button
          variant="contained"
          size="large"
          type="submit"
          sx={{
            marginLeft: '2rem',
            color: 'white',
            background: '#212121',
            height: '48px',
            minWidth: '200px',
            ':hover': { background: 'white', color: '#212121' },
          }}
        >
          Search
        </Button>
      </form>
      <Grid container direction="row" justifyContent={'center'} marginTop={'2rem'}>
        {startFlights?.length !== 0 && (
          <Grid item xs={10}>
            <Typography variant="h5" align="center">
              Departing flights
            </Typography>
          </Grid>
        )}
        {startFlights?.map((flight, index) => (
          <Grid item xs={10}>
            <FlightItem key={index} flight={flight} ticketAmount={recommendedFlightsProps.numberOfGuests} isBookingApp={true} />
          </Grid>
        ))}
        {startFlights?.length === 0 && <Typography variant="h5">No departing flights found</Typography>}
      </Grid>

      <Grid container direction="row" justifyContent={'center'} marginTop={'2rem'}>
        {endFlights?.length !== 0 && (
          <Grid item xs={10}>
            <Typography variant="h5" align="center">
              Returning flights
            </Typography>
          </Grid>
        )}
        {endFlights?.map((flight, index) => (
          <Grid item xs={10}>
            <FlightItem key={index} flight={flight} ticketAmount={recommendedFlightsProps.numberOfGuests} isBookingApp={true} />
          </Grid>
        ))}
        {endFlights?.length === 0 && <Typography variant="h5">No returning flights found</Typography>}
      </Grid>
    </>
  );
}

export default RecommendedFlights;
