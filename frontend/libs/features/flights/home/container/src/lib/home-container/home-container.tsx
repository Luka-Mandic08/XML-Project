import { Typography } from '@mui/material';
import AllFlights from '../all-flights/all-flights';

/* eslint-disable-next-line */
export interface FeaturesFlightsHomeContainerProps {}

export function HomeContainer():React.ReactElement {

  return (
    <>
    <AllFlights></AllFlights>
    </>
  );
};

export default HomeContainer;
