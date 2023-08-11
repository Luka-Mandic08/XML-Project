// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { LoginPage } from '@frontend/features/flights/login/container';
import { BookingAppRoutes, SharedRoutes } from '@frontend/models';
import { Routes, Route } from 'react-router-dom';
import { RegisterPage } from '@frontend/features/booking/login/container';
import { BookingNavBar } from '@frontend/features/booking/navigation/container';
import { ProfileInfo } from '@frontend/features/booking/profiles/container';

export function App() {
  return (
    <>
      <BookingNavBar />
      <Routes>
        <Route path={SharedRoutes.Login} element={<LoginPage isBookingApp={true} />} />
        <Route path={BookingAppRoutes.Register} element={<RegisterPage />} />
        <Route path={BookingAppRoutes.HomeGuest} element={<LoginPage isBookingApp={true} />} />
        <Route path={BookingAppRoutes.HomeHost} element={<LoginPage isBookingApp={true} />} />
        <Route path={BookingAppRoutes.Profile} element={<ProfileInfo />} />
      </Routes>
    </>
  );
}

export default App;
