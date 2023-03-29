// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { Box, Typography } from '@mui/material';
import { Route, Routes } from 'react-router-dom';
import { AppRoutes } from '@frontend/models';
import { HomeContainer } from '@frontend/features/flights/home/container';
import { AddFlightContainer } from '@frontend/features/flights/add-flight/container';

export function App() {
  return (
    <>
      <Typography>Hello world!</Typography>
      <Routes>
        <Route path={AppRoutes.Home} element={<HomeContainer />} />
        <Route path={AppRoutes.AddFlight} element={<AddFlightContainer />} />
      </Routes>
    </>
  );
}

export default App;
