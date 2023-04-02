// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { Route, Routes } from 'react-router-dom';
import { AppRoutes } from '@frontend/models';
import { HomeContainer } from '@frontend/features/flights/home/container';
import { LoginPage, RegistrationPage } from '@frontend/features/flights/login/container';
import { NavBar } from '@frontend/features/flights/nav-bar/container';
import { AddFlightContainer } from '@frontend/features/flights/add-flight/container';
import { PurchasedTicketsContainer } from '@frontend/features/flights/purchased-tickets/container';

export function App() {
  return (
    <>
      <NavBar></NavBar>
      <Routes>
        <Route path={AppRoutes.Login} element={<LoginPage />} />
        <Route path={AppRoutes.Home} element={<HomeContainer />} />
        <Route path={AppRoutes.Home} element={<HomeContainer />} />
        <Route path={AppRoutes.AddFlight} element={<AddFlightContainer />} />
        <Route path={AppRoutes.PurchasedTickets} element={<PurchasedTicketsContainer />} />
        <Route path={AppRoutes.Register} element={<RegistrationPage />} />
      </Routes>
    </>
  );
}

export default App;
