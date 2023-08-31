// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { LoginPage } from '@frontend/features/flights/login/container';
import { BookingAppRoutes, SharedRoutes } from '@frontend/models';
import { Routes, Route, Navigate, Outlet } from 'react-router-dom';
import { RegisterPage } from '@frontend/features/booking/register/container';
import { BookingNavBar } from '@frontend/features/booking/navigation/container';
import { ProfileInfo } from '@frontend/features/booking/profile/container';
// eslint-disable-next-line @nrwl/nx/enforce-module-boundaries
import {
  AccommodationDetails,
  AccomodationComments,
  AllAccommodation,
  CreateUpdateAccommodation,
  HostAccomodation,
  UpdateAccommodationAvailability,
} from '@frontend/features/booking/accomodation/container';
import { AccommodationReservations, GuestReservations, MakeReservation, RecommendedFlights } from '@frontend/features/booking/reservation/container';
import jwt from 'jwt-decode';
import Swal from 'sweetalert2';

export function App() {
  return (
    <>
      <BookingNavBar />
      <Routes>
        {/*Unprotected Routes*/}
        <Route path={SharedRoutes.Login} element={<LoginPage isBookingApp={true} />} />
        <Route path={BookingAppRoutes.Register} element={<RegisterPage />} />
        <Route path={BookingAppRoutes.HomeGuest} element={<AllAccommodation />} />

        <Route element={<PrivateRoutes userRole="Guest" />}>
          <Route path={BookingAppRoutes.GuestReservations} element={<GuestReservations />} />
          <Route path={BookingAppRoutes.MakeReservation} element={<MakeReservation />} />
          <Route path={BookingAppRoutes.RecommendedFlights} element={<RecommendedFlights />} />
        </Route>

        <Route element={<PrivateRoutes userRole="Host" />}>
          <Route path={BookingAppRoutes.HomeHost} element={<HostAccomodation />} />
          <Route path={BookingAppRoutes.HostComments} element={<AccomodationComments showHostComments={true} showAccommodationComments={false} />} />
          <Route path={BookingAppRoutes.CreateAccommodation} element={<CreateUpdateAccommodation />} />
          <Route path={BookingAppRoutes.AvailabilityCalendar} element={<UpdateAccommodationAvailability />} />
          <Route path={BookingAppRoutes.AccommodationDetails} element={<AccommodationDetails />} />
          <Route path={BookingAppRoutes.AccommodationReservations} element={<AccommodationReservations />} />
        </Route>

        <Route element={<PrivateRoutes userRole="Guest|Host" />}>
          <Route path={BookingAppRoutes.Profile} element={<ProfileInfo />} />
        </Route>
      </Routes>
    </>
  );
}

function PrivateRoutes(props: any) {
  let isAllowed = false;
  const token = localStorage.getItem('jwt');
  if (token === null) {
    Swal.fire({
      icon: 'error',
      title: 'Error',
      text: 'You need to login first!',
      showConfirmButton: false,
      position: 'bottom-right',
      timer: 3000,
      timerProgressBar: true,
      backdrop: 'none',
      width: 300,
      background: '#212121',
      color: 'white',
    });
    return <Navigate to="/" />;
  } else {
    const decodedToken: any = jwt(token);
    isAllowed = props.userRole.includes(decodedToken['roles']);
    if (!isAllowed) {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'You are not authorized to access this page!',
        showConfirmButton: false,
        position: 'bottom-right',
        timer: 3000,
        timerProgressBar: true,
        backdrop: 'none',
        width: 300,
        background: '#212121',
        color: 'white',
      });
      switch (decodedToken['roles']) {
        case 'Guest':
          return <Navigate to={BookingAppRoutes.HomeGuest} />;
        case 'Host':
          return <Navigate to={BookingAppRoutes.HomeHost} />;
      }
    }
    return <Outlet />;
  }
}

export default App;
