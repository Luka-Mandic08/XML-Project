import { Flight } from '@frontend/models';
import { Grid, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import PurchasedTicketItem from '../purchased-ticket-item/purchased-ticket-item';
import { GetAllPurchasedTickets } from '@frontend/features/flights/purchased-tickets/data-access';
import Swal from 'sweetalert2';

/* eslint-disable-next-line */
export interface AllPurchasedTicketsProps {}

export function AllFlights(props: AllPurchasedTicketsProps) {
  const [flights, setFlights] = useState<Flight[]>();

  useEffect(() => {
    GetAllPurchasedTickets()
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

  return (
    <Grid container direction="row" justifyContent="center" sx={{ margin: '0', padding: '1.25rem' }}>
      <Grid item xs={10}>
        <Typography variant="h4">My tickets</Typography>
      </Grid>
      <Grid item xs={10}>
        {flights?.map((flight, index) => (
          <PurchasedTicketItem key={index} flight={flight} />
        ))}
      </Grid>
    </Grid>
  );
}

export default AllFlights;
