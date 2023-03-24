import { Flight } from '@frontend/models';
import { Typography } from '@mui/material';
import styles from './flight-item.module.css';

/* eslint-disable-next-line */
export interface FlightItemProps {
  flight:Flight;
}

export function FlightItem(props: FlightItemProps) {
  return (
    <div className={styles['container']}>
      <Typography>{props.flight.start}</Typography>
    </div>
  );
}

export default FlightItem;
