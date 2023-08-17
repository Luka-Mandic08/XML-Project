import create from 'zustand';

interface SearchParametersState {
  searchParameters: any;
  setSearchParameters: (SearchParametersDTO: any) => void;
}

export const useSearchParametersStore = create<SearchParametersState>((set) => ({
  searchParameters: {
    city: '',
    country: '',
    dateFrom: '',
    dateTo: '',
    numberOfGuests: '',
  },
  setSearchParameters: (newParameters: any) => set({ searchParameters: newParameters }),
}));
