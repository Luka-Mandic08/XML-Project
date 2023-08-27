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
    register: registerStartFlight,
    handleSubmit: handleSubmitStartFlight,
    watch: watchStartFlight,
    formState: { errors: errorsStartFlight },
  } = useForm({
    defaultValues: {
      startDate: recommendedFlightsProps.startDate.toISOString().split('T')[0],
      startLocation: '',
      destination: recommendedFlightsProps.accommodationLocation.city,
      numberOfTickets: recommendedFlightsProps.numberOfGuests,
    },
  });

  const onSubmitStartFlight = async (data: any) => {
    const dto = new SearchFlightsDTO();
    dto.setFields(data.startDate, data.startLocation, data.destination, parseInt(data.numberOfTickets));
    console.log(dto);
    SearchFlights(dto)
      .then((result) => {
        if (result !== '') setStartFlights(result);
        else setStartFlights([]);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  const {
    register: registerEndFlight,
    handleSubmit: handleSubmitEndFlight,
    watch: watchEndFlight,
    formState: { errors: errorsEndFlight },
  } = useForm({
    defaultValues: {
      startDate: recommendedFlightsProps.endDate.toISOString().split('T')[0],
      startLocation: recommendedFlightsProps.accommodationLocation.city,
      destination: '',
      numberOfTickets: recommendedFlightsProps.numberOfGuests,
    },
  });

  const onSubmitEndFlight = async (data: any) => {
    const dto = new SearchFlightsDTO();
    dto.setFields(data.startDate, data.startLocation, data.destination, parseInt(data.numberOfTickets));
    console.log(dto);
    SearchFlights(dto)
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
      <form onSubmit={handleSubmitStartFlight(onSubmitStartFlight)}>
        <div className={styles.lineContainer}>
          <div className={styles.inputContainer}>
            <input
              type="date"
              id="startDate"
              value={watchStartFlight('startDate')}
              {...registerStartFlight('startDate', {
                required: 'This field is required.',
                min: { value: new Date().toISOString().split('T')[0], message: 'Minimum date is today.' },
              })}
            />
            <label className={styles.label} htmlFor="startDate" id="label-startDate">
              <div className={styles.text}>Flight's date</div>
            </label>
            <label className={styles.errorLabel}>{errorsStartFlight.startDate?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="text"
              id="startLocation"
              value={watchStartFlight('startLocation')}
              {...registerStartFlight('startLocation', { required: 'This field is required.' })}
            />
            <label className={styles.label} htmlFor="startLocation" id="label-startLocation">
              <div className={styles.text}>Start location</div>
            </label>
            <label className={styles.errorLabel}>{errorsStartFlight.startLocation?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="text"
              id="destination"
              value={watchStartFlight('destination')}
              {...registerStartFlight('destination', { required: 'This field is required.' })}
            />
            <label className={styles.label} htmlFor="destination" id="label-destination">
              <div className={styles.text}>Destination</div>
            </label>
            <label className={styles.errorLabel}>{errorsStartFlight.destination?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="number"
              id="numberOfTickets"
              value={watchStartFlight('numberOfTickets')}
              {...registerStartFlight('numberOfTickets', {
                required: 'This field is required.',
                min: { value: 1, message: 'Minimum number of tickets is 1.' },
              })}
            />
            <label className={styles.label} htmlFor="numberOfTickets" id="label-numberOfTickets">
              <div className={styles.text}>Number of tickets</div>
            </label>
            <label className={styles.errorLabel}>{errorsStartFlight.numberOfTickets?.message}</label>
          </div>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Search
          </Button>
        </div>
      </form>
      <Grid container direction="column" alignItems="center" sx={{ mt: '2rem' }}>
        <Grid item>
          {startFlights?.map((flight, index) => (
            <FlightItem key={index} flight={flight} />
          ))}
          {startFlights?.length === 0 && <Typography variant="h5">No flights found</Typography>}
        </Grid>
      </Grid>

      <form onSubmit={handleSubmitEndFlight(onSubmitEndFlight)}>
        <div className={styles.lineContainer}>
          <div className={styles.inputContainer}>
            <input
              type="date"
              id="startDate"
              value={watchEndFlight('startDate')}
              {...registerEndFlight('startDate', {
                required: 'This field is required.',
                min: { value: new Date().toISOString().split('T')[0], message: 'Minimum date is today.' },
              })}
            />
            <label className={styles.label} htmlFor="startDate" id="label-startDate">
              <div className={styles.text}>Flight's date</div>
            </label>
            <label className={styles.errorLabel}>{errorsEndFlight.startDate?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="text"
              id="startLocation"
              value={watchEndFlight('startLocation')}
              {...registerEndFlight('startLocation', { required: 'This field is required.' })}
            />
            <label className={styles.label} htmlFor="startLocation" id="label-startLocation">
              <div className={styles.text}>Start location</div>
            </label>
            <label className={styles.errorLabel}>{errorsEndFlight.startLocation?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="text"
              id="destination"
              value={watchEndFlight('destination')}
              {...registerEndFlight('destination', { required: 'This field is required.' })}
            />
            <label className={styles.label} htmlFor="destination" id="label-destination">
              <div className={styles.text}>Destination</div>
            </label>
            <label className={styles.errorLabel}>{errorsEndFlight.destination?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="number"
              id="numberOfTickets"
              value={watchEndFlight('numberOfTickets')}
              {...registerEndFlight('numberOfTickets', {
                required: 'This field is required.',
                min: { value: 1, message: 'Minimum number of tickets is 1.' },
              })}
            />
            <label className={styles.label} htmlFor="numberOfTickets" id="label-numberOfTickets">
              <div className={styles.text}>Number of tickets</div>
            </label>
            <label className={styles.errorLabel}>{errorsEndFlight.numberOfTickets?.message}</label>
          </div>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Search
          </Button>
        </div>

        <Grid container direction="column" alignItems="center" sx={{ mt: '2rem' }}>
          <Grid item>
            {endFlights?.map((flight, index) => (
              <FlightItem key={index} flight={flight} />
            ))}
            {endFlights?.length === 0 && <Typography variant="h5">No flights found</Typography>}
          </Grid>
        </Grid>
      </form>
    </>
  );
}

export default RecommendedFlights;
