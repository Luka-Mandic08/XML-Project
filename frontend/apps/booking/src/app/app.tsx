// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { LoginPage } from '@frontend/features/flights/login/container';
import { BookingAppRoutes, SharedRoutes } from '@frontend/models';
import { Routes, Route } from 'react-router-dom';
import { RegisterPage } from '@frontend/features/booking/register/container';
import { BookingNavBar } from '@frontend/features/booking/navigation/container';
import { ProfileInfo } from '@frontend/features/booking/profile/container';
// eslint-disable-next-line @nrwl/nx/enforce-module-boundaries
import { AllAccommodation, AvailabilityCalendar, CreateUpdateAccommodation, HostAccomodation } from '@frontend/features/booking/accomodation/container';
import { AccommodationReservations, GuestReservations } from '@frontend/features/booking/reservation/container';

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
        <Route path={BookingAppRoutes.GuestReservations} element={<GuestReservations />} />
        <Route path={BookingAppRoutes.AccommodationReservations} element={<AccommodationReservations />} />
      </Routes>
    </>
  );
}

export default App;