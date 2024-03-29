export const AppRoutes = {
  Login: '/login',
  Home: '/',
  Flights: '/flights',
  AddFlight: '/add/flight',
  PurchasedTickets: '/mytickets',
  Register: '/register',
};

export const BookingAppRoutes = {
  Register: '/register',
  HomeGuest: '/',
  HomeHost: '/myaccomodations',
  Profile: '/profile',
  CreateAccommodation: '/accommodation/create',
  AvailabilityCalendar: '/accommodation/availability',
  GuestReservations: '/reservations/guest',
  AccommodationReservations: '/reservations/accommodation',
  CreateReservation: '/reservations/create',
  MakeReservation: '/accommodation/make-reservation',
  AccommodationDetails: '/accommodation/details',
  HostComments: '/host-comments',
  RecommendedFlights: '/flights/recommended',
  Notifications: '/notifications',
  RecommendedAccommodations: '/accommodation/recommended',
};

export const SharedRoutes = {
  Login: '/login',
};

export const BaseURL = {
  URL: 'http://localhost:8082',
};

export const BookingBaseURL = {
  URL: 'http://localhost:8000',
};
