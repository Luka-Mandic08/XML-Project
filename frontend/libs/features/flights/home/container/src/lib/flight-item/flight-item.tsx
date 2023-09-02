import { AppRoutes, Flight } from '@frontend/models';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Divider, Grid, TextField, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import FlightIcon from '@mui/icons-material/Flight';
import { BuyFlightTickets, BuyFlightTicketsFromBookingApp } from '@frontend/features/flights/home/data-access';
import { useNavigate } from 'react-router-dom';
import { DeleteFlight } from '@frontend/features/flights/add-flight/data-access';
import { GetApiKey } from '@frontend/features/booking/profile/data-access';

/* eslint-disable-next-line */
export interface FlightItemProps {
  flight: Flight;
  ticketAmount: number;
  isBookingApp?: boolean;
}

export function FlightItem(props: FlightItemProps) {
  const [amount, setAmount] = useState<number>(props.ticketAmount);
  const [amountError, setAmountError] = useState<boolean>();
  const [isBuyTicketDialogOpen, setIsBuyTicketDialogOpen] = useState(false);

  useEffect(() => {
    setAmountError(false);

    if (props.flight.remainingtickets < amount || amount < 1) {
      setAmountError(true);
    }
  }, [amount, props.flight.remainingtickets, setAmount]);

  const handleBuyTicketClick = () => {
    setIsBuyTicketDialogOpen(true);
  };

  const handleBuyTicketClose = () => {
    setIsBuyTicketDialogOpen(false);
  };

  const buyTickets = async (amount: number) => {
    if (!props.isBookingApp) {
      setIsBuyTicketDialogOpen(false);
      BuyFlightTickets(props.flight.id, amount);
      window.location.reload();
    } else {
      const apiKey = await GetApiKey();
      await BuyFlightTicketsFromBookingApp(props.flight.id, amount, apiKey.apiKeyValue);
    }
  };

  const deleteFlight = () => {
    DeleteFlight(props.flight.id);
    window.location.reload();
  };

  let customButton;
  if (localStorage.getItem('role') === 'USER' || localStorage.getItem('role') === 'Guest') {
    customButton = (
      <Button
        variant="contained"
        onClick={handleBuyTicketClick}
        sx={{ backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}
        disabled={props.flight.remainingtickets === 0}
      >
        Buy now: {props.flight.price}$
      </Button>
    );
  } else if (localStorage.getItem('role') === 'ADMIN') {
    customButton = (
      <Button variant="contained" onClick={deleteFlight} sx={{ backgroundColor: '#212121', '&:hover': { backgroundColor: '#ffffff', color: '#212121' } }}>
        Delete flight
      </Button>
    );
  } else {
    customButton = (
      <Button variant="contained" disabled={true} sx={{ backgroundColor: '#212121' }}>
        Buy now: {props.flight.price}$
      </Button>
    );
  }

  let buyBtnText;
  if (amount > 0) {
    buyBtnText = 'Buy now ' + amount * props.flight.price + '$';
  } else {
    buyBtnText = 'Buy now';
  }

  return (
    <Grid
      container
      direction="row"
      justifyContent="space-evenly"
      sx={{
        marginTop: '1rem',
        marginBottom: '1.75rem',
        maxWidth: '90vw',
        padding: '1rem',
        borderRadius: '6px',
        backgroundColor: 'white',
        boxShadow: '0px 7px 7px 5px lightgray',
      }}
    >
      <Grid container direction="column" justifyContent="center" alignItems="center" xs={3}>
        <Grid item>
          <Typography variant="caption">From:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">{props.flight.start}</Typography>
        </Grid>
        <Divider sx={{ backgroundColor: '#212121', mb: 1, mt: 1 }} />
        <Grid item>
          <Typography variant="h5">
            {props.flight.startdate.toString().split('T')[0]} at: {props.flight.startdate.toString().split('T')[1].substring(0, 5)}
          </Typography>
        </Grid>
      </Grid>

      <Grid container direction="row" justifyContent="center" alignItems="center" xs={1}>
        <Grid item>
          <FlightIcon sx={{ fontSize: 35, transform: 'rotate(90deg)' }}></FlightIcon>
        </Grid>
      </Grid>

      <Grid container direction="column" justifyContent="center" alignItems="center" xs={3}>
        <Grid item>
          <Typography variant="caption">To:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">{props.flight.destination}</Typography>
        </Grid>
        <Divider sx={{ backgroundColor: '#212121', mb: 1, mt: 1 }} />
        <Grid item>
          <Typography variant="h5">
            {props.flight.arrivaldate.toString().split('T')[0]} at: {props.flight.arrivaldate.toString().split('T')[1].substring(0, 5)}
          </Typography>
        </Grid>
      </Grid>

      <Grid container direction="row" justifyContent="flex-end" alignItems="center" xs={3}>
        <Grid container direction="column" justifyContent="center" alignItems="center" xs>
          <Grid item>
            <Typography variant="caption">Remaining tickets:</Typography>
          </Grid>
          <Grid item>
            <Typography variant="h5">{props.flight.remainingtickets}</Typography>
          </Grid>
        </Grid>
        <Grid item>
          {customButton}
          <Dialog open={isBuyTicketDialogOpen} onClose={handleBuyTicketClose}>
            <DialogTitle sx={{ minWidth: '20vw' }}>Buy tickets</DialogTitle>
            <DialogContent>
              <DialogContentText>
                Starting Location: {props.flight.start} <br />
                Date: {props.flight.startdate.toString().split('T')[0]} at: {props.flight.startdate.toString().split('T')[1].substring(0, 5)} <br />
                <Divider sx={{ my: '0.5rem' }}></Divider>
                Destination: {props.flight.destination} <br />
                Date: {props.flight.arrivaldate.toString().split('T')[0]} at: {props.flight.arrivaldate.toString().split('T')[1].substring(0, 5)} <br />
                <Divider sx={{ my: '0.5rem' }}></Divider>
                Remaining tickets: {props.flight.remainingtickets} <br />
                <Divider sx={{ mt: '0.5rem' }}></Divider>
              </DialogContentText>
              <TextField
                autoFocus
                margin="dense"
                id="amount"
                label="Amount of tickets"
                type="number"
                defaultValue={amount}
                fullWidth
                variant="standard"
                error={amountError}
                helperText={amountError ? 'Exceded the number of remaining tickets' : ' '}
                onChange={(event) => setAmount(parseInt(event.target.value))}
              />
            </DialogContent>
            <DialogActions>
              <Button
                onClick={handleBuyTicketClose}
                sx={{ backgroundColor: '#ffffff', color: '#212121', margin: '0.1rem', '&:hover': { backgroundColor: '#FF6666', color: '#ffffff' } }}
              >
                Cancel
              </Button>
              <Button
                disabled={amountError}
                onClick={() => buyTickets(amount)}
                sx={{
                  backgroundColor: '#212121',
                  color: '#ffffff',
                  margin: '0.1rem',
                  '&:hover': { backgroundColor: '#ffffff', color: '#212121', outline: '1px solid #212121' },
                  '&:disabled': { backgroundColor: '#ffffff', color: '#FF6666', outline: '1px solid #FF6666' },
                }}
              >
                {buyBtnText}
              </Button>
            </DialogActions>
          </Dialog>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default FlightItem;
