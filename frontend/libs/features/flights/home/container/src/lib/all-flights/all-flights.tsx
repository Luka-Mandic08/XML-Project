import { GetAllFlights } from '@frontend/features/flights/home/data-access';
import { AppRoutes, Flight } from '@frontend/models';
import { Button, Grid, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import FlightItem from '../flight-item/flight-item';
import styles from './all-flights.module.css';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface AllFlightsProps {}

export function AllFlights(props: AllFlightsProps) {
  const [flights, setFlights] = useState<Flight[]>();

  /*useEffect(() => {
    const fetchData = async () => {
      const data = await GetAllFlights();
      setFlights(data)
    }
  
    // call the function
    fetchData()
      // make sure to catch any error
      .catch(console.error);
      .catch(console.error);
  }, [flights])*/
  useEffect(() => {
    GetAllFlights()
      .then((result) => {
        setFlights(result);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  const navigate = useNavigate();
  const goToAddFlight = () => {
    navigate(AppRoutes.AddFlight);
  };

  let welcomeText;
  let addFlightButton;

  if (localStorage.getItem('role') === 'USER') {
    welcomeText = <Typography variant="h4">Welcome to AllFlights! as User</Typography>;
  } else if (localStorage.getItem('role') === 'ADMIN') {
    addFlightButton = (
      <Button variant="contained" onClick={goToAddFlight} sx={{ backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}>
        Add new flight
      </Button>
    );
    welcomeText = <Typography variant="h4">Welcome to AllFlights! as Admin</Typography>;
  } else {
    welcomeText = <Typography variant="h4">Welcome to AllFlights! as None</Typography>;
  }

  return (
    <Grid container direction="row" justifyContent="center" sx={{ border: '3px solid #212121', margin: '0', padding: '1.25rem' }}>
      <Grid item xs={5}>
        {welcomeText}
      </Grid>
      <Grid item container justifyContent="flex-end" xs={5}>
        {addFlightButton}
      </Grid>
      <Grid item xs={10}>
        {flights?.map((flight, index) => (
          <FlightItem key={index} flight={flight} />
        ))}
      </Grid>
    </Grid>
  );
}

export default AllFlights;
