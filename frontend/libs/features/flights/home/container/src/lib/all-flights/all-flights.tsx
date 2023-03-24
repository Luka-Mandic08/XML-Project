import { GetAllFlights } from '@frontend/features/flights/home/data-access';
import { Flight } from '@frontend/models';
import { useEffect, useState } from 'react';
import FlightItem from '../flight-item/flight-item';
import styles from './all-flights.module.css';

/* eslint-disable-next-line */
export interface AllFlightsProps {}

export function AllFlights(props: AllFlightsProps) {

  const [flights, setFlights] = useState<Flight[]>()

  /*useEffect(() => {
    const fetchData = async () => {
      const data = await GetAllFlights();
      setFlights(data)
    }
  
    // call the function
    fetchData()
      // make sure to catch any error
      .catch(console.error);
  }, [flights])*/
  useEffect(() => {
    GetAllFlights()
      .then(result => {
        setFlights(result);
      })
      .catch(error => {
        console.error(error)
      })
  }, [])

  return (
    <div className={styles['container']}>
      <h1>Welcome to AllFlights!</h1>
      {flights?.map((flight, index) =>
        <FlightItem key={index} flight={flight} />
      )}
    </div>
  );
}

export default AllFlights;
