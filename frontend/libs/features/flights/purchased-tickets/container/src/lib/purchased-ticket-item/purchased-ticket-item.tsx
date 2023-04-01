import { AppRoutes, Flight } from '@frontend/models';
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Divider, Grid, TextField, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import FlightIcon from '@mui/icons-material/Flight';
import { BuyFlightTickets } from '@frontend/features/flights/home/data-access';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface PurchasedTicketProps {
  flight: Flight;
}

export function PurchasedTicketItem(props: PurchasedTicketProps) {
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
            <Typography variant="caption">Number of purchased tickets:</Typography>
          </Grid>
          <Grid item>
            <Typography variant="h5">{props.flight.remainingtickets}</Typography>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default PurchasedTicketItem;
