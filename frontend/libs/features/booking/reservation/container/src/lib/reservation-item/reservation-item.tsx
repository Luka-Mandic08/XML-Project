import { AccomodationDtoTry, MyReservation } from '@frontend/models';
import { Button, Grid, Typography } from '@mui/material';

const accomodation: AccomodationDtoTry = {
  name: 'Cozy Cabin Retreat',
  address: '1234 Mountain View Road, Exampleville',
  amenities: 'Wi-Fi, Parking, Air Conditioning',
  images: ['image1.jpg', 'image2.jpg', 'image3.jpg'],
  minGuests: '2',
  maxGuests: '6',
  priceIsPerGuest: 'true',
  hostId: 'user123',
  hasAutomaticReservations: 'true',
};

/* eslint-disable-next-line */
export interface ReservationItemProps {
  reservation: MyReservation;
}

export function ReservationItem(props: ReservationItemProps) {
  const deleteFlight = () => {
    //CancelReservation(props.reservation.id);
    window.location.reload();
  };

  return (
    <Grid
      container
      direction="row"
      justifyContent="space-evenly"
      sx={{
        marginY: '1rem',
        marginBottom: '1.75rem',
        maxWidth: '90vw',
        padding: '1rem',
        borderRadius: '6px',
        backgroundColor: 'white',
        boxShadow: '0px 7px 7px 5px lightgray',
      }}
    >
      <Grid container direction="column" justifyContent="center" alignItems="center" xs={5}>
        <Grid item>
          <Typography variant="caption">Name:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">Name (Adress)</Typography>
        </Grid>
      </Grid>

      <Grid container direction="column" justifyContent="center" alignItems="center" xs={3}>
        <Grid item>
          <Typography variant="caption">From/To:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">
            {props.reservation.dateFrom.toString().split('T')[0]} / {props.reservation.dateTo.toString().split('T')[0]}
          </Typography>
        </Grid>
      </Grid>

      <Grid container direction="column" justifyContent="center" alignItems="center" xs={1}>
        <Grid item>
          <Typography variant="caption">Price:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">{props.reservation.price}</Typography>
        </Grid>
      </Grid>

      <Grid container direction="row" justifyContent="flex-end" alignItems="center" xs={3}>
        <Grid container direction="column" justifyContent="center" alignItems="center" xs>
          <Grid item>
            <Typography variant="caption">Status:</Typography>
          </Grid>
          <Grid item>
            <Typography variant="h5" style={{ color: props.reservation.status === 'pending' ? 'orange' : 'green' }}>
              {props.reservation.status}
            </Typography>
          </Grid>
        </Grid>
        <Grid container direction="column" justifyContent="center" alignItems="center" xs>
          <Grid item>
            <Button
              onClick={deleteFlight}
              sx={{
                backgroundColor: '#e12e2e',
                color: '#ffffff',
                margin: '0.5rem',
                '&:hover': { backgroundColor: '#ffffff', color: '#e12e2e', borderColor: '#e12e2e', borderWidth: '2px' },
              }}
            >
              Cancel
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default ReservationItem;
