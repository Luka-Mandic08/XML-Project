import { GetAllFlights, SearchFlights } from '@frontend/features/flights/home/data-access';
import { AppRoutes, Flight, SearchFlightsDTO } from '@frontend/models';
import { Button, Grid, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import FlightItem from '../flight-item/flight-item';
import styles from './all-flights.module.css';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';

/* eslint-disable-next-line */
export interface AllFlightsProps {}

export function AllFlights(props: AllFlightsProps) {
  const [flights, setFlights] = useState<Flight[]>();
  const [selectedDate, setSelectedDate] = useState('');
  const [start, setStart] = useState('');
  const [destination, setDestination] = useState('');
  const [tickets, setTickets] = useState('');

  useEffect(() => {
    GetAllFlights()
      .then((result) => {
        setFlights(result);
      })
      .catch((error) => {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Something went wrong, please try again',
          showConfirmButton: false,
          position: 'bottom-right',
          timer: 3000,
          timerProgressBar: true,
          backdrop: 'none',
          width: 300,
          background: '#212121',
          color: 'white',
        });
      });
  }, []);

  const navigate = useNavigate();
  const goToAddFlight = () => {
    navigate(AppRoutes.AddFlight);
  };

  function search() {
    const dto = new SearchFlightsDTO();
    dto.setFields(selectedDate, start, destination, parseInt(tickets));
    SearchFlights(dto)
      .then((result) => {
        if (result !== '') setFlights(result);
        else setFlights([]);
      })
      .catch((error) => {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Something went wrong, please try again',
          showConfirmButton: false,
          position: 'bottom-right',
          timer: 3000,
          timerProgressBar: true,
          backdrop: 'none',
          width: 300,
          background: '#212121',
          color: 'white',
        });
      });
  }

  function clear() {
    setSelectedDate('');
    setStart('');
    setDestination('');
    setTickets('');
    GetAllFlights()
      .then((result) => {
        setFlights(result);
      })
      .catch((error) => {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: 'Something went wrong, please try again',
          showConfirmButton: false,
          position: 'bottom-right',
          timer: 3000,
          timerProgressBar: true,
          backdrop: 'none',
          width: 300,
          background: '#212121',
          color: 'white',
        });
      });
  }

  let welcomeText;
  let addFlightButton;

  if (localStorage.getItem('role') === 'USER') {
    welcomeText = <Typography variant="h4">Buy flight tickets below</Typography>;
  } else if (localStorage.getItem('role') === 'ADMIN') {
    addFlightButton = (
      <Button variant="contained" onClick={goToAddFlight} sx={{ backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}>
        Add new flight
      </Button>
    );
    welcomeText = <Typography variant="h4">Manage flights</Typography>;
  } else {
    welcomeText = <Typography variant="h4">Flight tickets</Typography>;
  }

  const currentDate = new Date();
  const currentDateString = currentDate.toISOString().split('T')[0];

  return (
    <Grid container direction="row" justifyContent="center" sx={{ margin: '0', padding: '1.25rem' }}>
      <Grid item xs={5}>
        {welcomeText}
      </Grid>
      <Grid item container justifyContent="flex-end" xs={5}>
        {addFlightButton}
      </Grid>
      <Grid item container xs={10} sx={{ padding: '1.25rem' }}>
        <Grid item xs>
          <div className={styles.inputContainer}>
            <input type="date" id="startdate" min={currentDateString} value={selectedDate} onChange={(e) => setSelectedDate(e.target.value)} />
            <label className={styles.label} htmlFor="startdate" id="label-startdate">
              <div className={styles.text}>Starting Date</div>
            </label>
          </div>
        </Grid>
        <Grid item xs>
          <div className={styles.inputContainer}>
            <input type="text" id="startLocation" value={start} onChange={(e) => setStart(e.target.value)} />
            <label className={styles.label} htmlFor="startLocation" id="label-startLocation">
              <div className={styles.text}>Start location</div>
            </label>
          </div>
        </Grid>
        <Grid item xs>
          <div className={styles.inputContainer}>
            <input type="text" id="destination" value={destination} onChange={(e) => setDestination(e.target.value)} />
            <label className={styles.label} htmlFor="destination" id="label-destination">
              <div className={styles.text}>Destination</div>
            </label>
          </div>
        </Grid>
        <Grid item xs>
          <div className={styles.inputContainer}>
            <input type="number" id="tickets" min={'1'} value={tickets} onChange={(e) => setTickets(e.target.value)} />
            <label className={styles.label} htmlFor="tickets" id="label-tickets">
              <div className={styles.text}>Number of tickets</div>
            </label>
          </div>
        </Grid>
        <Grid item>
          <Button
            variant="contained"
            onClick={search}
            sx={{ ml: '0.75rem', height: '48px', backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}
          >
            Search
          </Button>
          <Button
            variant="contained"
            onClick={clear}
            sx={{ ml: '0.75rem', height: '48px', backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}
          >
            Clear
          </Button>
        </Grid>
      </Grid>
      <Grid item xs={10}>
        {flights?.map((flight, index) => (
          <FlightItem key={index} flight={flight} ticketAmount={1} />
        ))}
      </Grid>
    </Grid>
  );
}

export default AllFlights;
