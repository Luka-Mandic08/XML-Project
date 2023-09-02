// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { LoginPage } from '@frontend/features/flights/login/container';
import { BookingAppRoutes, BookingBaseURL, SharedRoutes } from '@frontend/models';
import { Routes, Route } from 'react-router-dom';
import { RegisterPage } from '@frontend/features/booking/register/container';
import { BookingNavBar } from '@frontend/features/booking/navigation/container';
import { ProfileInfo } from '@frontend/features/booking/profile/container';
import { NotificationContainer } from '@frontend/features/booking/notification/container';
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
import { useEffect } from 'react';
import axios from 'axios';
import Swal from 'sweetalert2';

export function App() {
  useEffect(() => {
    const intervalId = setInterval(() => {
      const userId = localStorage.getItem('userId');
      axios
        .get(BookingBaseURL.URL + '/notification/all/' + userId)
        .then((response) => {
          const criticalEvents = response.data;
          if (criticalEvents.length !== 0) {
            console.warn(criticalEvents);
            criticalEvents.forEach((notification: any, index: number) => {
              setTimeout(() => {
                console.log(notification);
                Swal.fire({
                  icon: 'info',
                  title: 'Notification',
                  html: '<div style="max-height: 400px; overflow: auto;">' + notification.notificationText + '</div>',
                  showConfirmButton: false,
                  position: 'bottom-right',
                  timer: 4000,
                  timerProgressBar: true,
                  backdrop: 'none',
                  width: 300,
                  background: '#212121',
                  color: 'white',
                });

                axios
                  .put(BookingBaseURL.URL + '/notification/acknowledge', { id: notification.id })
                  .then((ackResponse) => {
                    console.log('Notification acknowledged:', ackResponse.data);
                  })
                  .catch((ackError) => {
                    console.error('Error acknowledging notification:', ackError);
                  });
              }, index * 5000);
            });
          }
        })
        .catch((error) => {
          console.error(error);
        });
    }, 20000);

    return () => {
      clearInterval(intervalId);
    };
  }, []);

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
        <Route path={BookingAppRoutes.AvailabilityCalendar} element={<UpdateAccommodationAvailability />} />
        <Route path={BookingAppRoutes.GuestReservations} element={<GuestReservations />} />
        <Route path={BookingAppRoutes.AccommodationReservations} element={<AccommodationReservations />} />
        <Route path={BookingAppRoutes.AccommodationDetails} element={<AccommodationDetails />} />
        <Route path={BookingAppRoutes.MakeReservation} element={<MakeReservation />} />
        <Route path={BookingAppRoutes.HostComments} element={<AccomodationComments showHostComments={true} showAccommodationComments={false} />} />
        <Route path={BookingAppRoutes.RecommendedFlights} element={<RecommendedFlights />} />
        <Route path={BookingAppRoutes.Notifications} element={<NotificationContainer />} />
      </Routes>
    </>
  );
}

export default App;
