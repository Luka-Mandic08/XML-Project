// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { LoginPage } from '@frontend/features/flights/login/container';
import { BookingAppRoutes, SharedRoutes } from '@frontend/models';
import { Routes, Route } from 'react-router-dom';
import { RegisterPage } from '@frontend/features/booking/login/container';
import { BookingNavBar } from '@frontend/features/booking/navigation/container';
import { ProfileInfo } from '@frontend/features/booking/profiles/container';
// eslint-disable-next-line @nrwl/nx/enforce-module-boundaries
import { AllAccommodation, AvailabilityCalendar, CreateUpdateAccommodation, HostAccomodation } from '@frontend/features/booking/accomodation/container';

export function App() {
  return (
    <>
      <BookingNavBar />
      <Routes>
        <Route path={SharedRoutes.Login} element={<LoginPage isBookingApp={true} />} />
        <Route path={BookingAppRoutes.Register} element={<RegisterPage />} />
        <Route path={BookingAppRoutes.HomeGuest} element={<AllAccommodation />} />
        <Route path={BookingAppRoutes.HomeHost} element={<HostAccomodation />} />
        <Route path={BookingAppRoutes.CreateAccommodation} element={<CreateUpdateAccommodation />} />
        <Route path={BookingAppRoutes.Profile} element={<ProfileInfo />} />
        <Route path={BookingAppRoutes.AvailabilityCalendar} element={<AvailabilityCalendar />} />
      </Routes>
    </>
  );
}

export default App;
