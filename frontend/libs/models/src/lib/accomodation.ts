import { Address } from './bookingUser';

export interface AccommodationInfo {
  id: string;
  hostId: string;
  name: string;
  address: Address;
  amenities: string[];
  images: string[];
  minGuests: number;
  maxGuests: number;
  priceIsPerGuest: boolean;
  hasAutomaticReservations: boolean;
}

export interface SearchedAccommodationInfo {
  id: string;
  name: string;
  address: Address;
  amenities: string[];
  images: string[];
  unitPrice: number;
  totalPrice: number;
}

export interface AccommodationCreateUpdateDTO {
  id: string;
  hostId: string;
  name: string;
  address: Address;
  minGuests: number;
  maxGuests: number;
  priceIsPerGuest: boolean;
  hasAutomaticReservations: boolean;
}

export interface AvailabilityDate {
  date: Date;
  isAvailable: boolean | undefined;
  price: number | undefined;
}
