export interface Flight {
  id: string;
  startdate: Date;
  arrivaldate: Date;
  destination: string;
  start: string;
  price: number;
  remainingtickets: number;
  totaltickets: number;
}

export interface NewFlight {
  startdate: number;
  arrivaldate: number;
  destination: string;
  start: string;
  price: number;
  totaltickets: number;
}
