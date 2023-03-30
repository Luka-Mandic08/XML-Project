import { Flight } from '@frontend/models';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Divider, Grid, TextField, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import FlightIcon from '@mui/icons-material/Flight';
import { BuyFlightTickets } from '@frontend/features/flights/home/data-access';

/* eslint-disable-next-line */
export interface FlightItemProps {
  flight:Flight;
}

export function FlightItem(props: FlightItemProps) {

  const [amount, setAmount] = useState<number>(1)
  const [amountError, setAmountError] = useState<boolean>()

  useEffect(() => {
    setAmountError(false);

    if (props.flight.remainingtickets < amount || amount < 1) {
      setAmountError(true);
    }

  }, [amount, props.flight.remainingtickets, setAmount]);

  const [isBuyTicketDialogOpen, setIsBuyTicketDialogOpen] = useState(false);
  const handleBuyTicketClick = () => {
    setIsBuyTicketDialogOpen(true);
    setAmount(1);
  };

  const handleBuyTicketClose = () => {
    setIsBuyTicketDialogOpen(false);
  };

  function buyTickets(amount : number) {
    setIsBuyTicketDialogOpen(false);
    BuyFlightTickets(props.flight.id, amount)
  };

  let customButton;
  if (localStorage.getItem('role') === "USER") {
    customButton = <Button variant="contained" onClick={handleBuyTicketClick} sx={{backgroundColor: "#212121", '&:hover': { backgroundColor: '#ffffff', color:  "#212121" }}}>Buy now: {props.flight.price}$</Button>
  } else if(localStorage.getItem('role') === "ADMIN") {
    customButton = <Button variant="contained" sx={{backgroundColor: "#212121", '&:hover': { backgroundColor: '#ffffff', color:  "#212121" }}}>Delete flight</Button>
  } else {
    customButton = <Button variant="contained"  disabled={true} sx={{backgroundColor: "#212121"}}>Buy now: {props.flight.price}$</Button>
  }

  return (
    <Grid container direction="row" justifyContent="space-evenly" sx={{ marginY: "1rem", marginBottom : "1.75rem", maxWidth: "90vw", padding: "1rem", borderRadius: "6px", backgroundColor: "white", boxShadow: "0px 7px 7px 5px lightgray"}}>
      <Grid container direction="column" justifyContent="center" alignItems="center" xs={3}>
        <Grid item>
          <Typography variant="caption">From:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">{props.flight.start}</Typography>
        </Grid>
        <Divider sx={{backgroundColor: "#212121", mb: 1, mt: 1}}/>
        <Grid item>
          <Typography variant="h5">{props.flight.startdate.toString().split('T')[0]} at: {props.flight.startdate.toString().split('T')[1].substring(0, 5)}</Typography>
        </Grid>
      </Grid>

      <Grid container direction="row" justifyContent="center" alignItems="center" xs={1}>
        <Grid item>
          <FlightIcon sx={{ fontSize: 35, transform:'rotate(90deg)'}}></FlightIcon>
        </Grid>
      </Grid>

      <Grid container direction="column" justifyContent="center" alignItems="center" xs={3}>
        <Grid item>
          <Typography variant="caption">To:</Typography>
        </Grid>
        <Grid item>
          <Typography variant="h5">{props.flight.destination}</Typography>
        </Grid>
        <Divider sx={{backgroundColor: "#212121", mb: 1, mt: 1}}/>
        <Grid item>
          <Typography variant="h5">{props.flight.arrivaldate.toString().split('T')[0]} at: {props.flight.arrivaldate.toString().split('T')[1].substring(0, 5)}</Typography>
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
            <DialogTitle>Buy tickets</DialogTitle>
            <DialogContent>
              <DialogContentText>
                To subscribe to this website, please enter your email address here. We
                will send updates occasionally.
              </DialogContentText>
              <TextField
                autoFocus
                margin="dense"
                id="amount"
                label="Amount of tickets"
                type="number"
                defaultValue="1"
                fullWidth
                variant="standard"
                error={amountError}
                helperText="Exceded the number of remaining tickets"
                onChange={event => setAmount(parseInt(event.target.value))}
              />
            </DialogContent>
            <DialogActions>
              <Button onClick={handleBuyTicketClose} sx={{backgroundColor: '#ffffff', color:  "#212121", margin:'0.1rem', '&:hover': { backgroundColor: '#FF6666', color:  "#ffffff"}}}>Cancel</Button>
              <Button disabled={amountError} onClick={() => buyTickets(amount)} sx={{backgroundColor: "#212121", color:'#ffffff', margin:'0.1rem', '&:hover': { backgroundColor: '#ffffff', color:  "#212121", outline:'1px solid #212121' }, '&:disabled': { backgroundColor: '#ffffff', color:  "#FF6666", outline:'1px solid #FF6666' }}}>Buy now {amount*props.flight.price}$</Button>
            </DialogActions>
          </Dialog>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default FlightItem;
