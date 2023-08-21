import { AccommodationInfo } from '@frontend/models';
import create from 'zustand';

interface SelectedAccommodationState {
  selectedAccommodation: AccommodationInfo;
  setSelectedAccommodation: (AccommodationDTO: AccommodationInfo) => void;
}

export const useSelectedAccommodationStore = create<SelectedAccommodationState>((set) => ({
  selectedAccommodation: {
    id: '',
    hostId: '',
    name: '',
    address: {
      country: '',
      city: '',
      street: '',
    },
    amenities: [],
    images: [],
    minGuests: 0,
    maxGuests: 0,
    priceIsPerGuest: false,
    hasAutomaticReservations: false,
    rating: 0,
  },
  setSelectedAccommodation: (newAccommodation: AccommodationInfo) => set({ selectedAccommodation: newAccommodation }),
}));
