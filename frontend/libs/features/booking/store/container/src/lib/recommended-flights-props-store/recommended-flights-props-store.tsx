import { RecommendedFlightsProps } from '@frontend/models';
import create from 'zustand';

interface RecomendedFlightdPropsState {
  recomendedFlightdProps: RecommendedFlightsProps;
  setRecomendedFlightdProps: (RecomendedFlightdPropsDTO: RecommendedFlightsProps) => void;
}

export const useRecomendedFlightdPropsStore = create<RecomendedFlightdPropsState>((set) => ({
  recomendedFlightdProps: {
    startDate: new Date(),
    endDate: new Date(),
    numberOfGuests: 0,
    accommodationLocation: {
      city: '',
      country: '',
      street: '',
    },
  },
  setRecomendedFlightdProps: (newProps: RecommendedFlightsProps) => set({ recomendedFlightdProps: newProps }),
}));
