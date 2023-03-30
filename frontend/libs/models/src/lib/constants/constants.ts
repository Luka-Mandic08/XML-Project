export const AppRoutes = {
  Home: '/',
  Flights: '/flights'
};

export const BaseURL = {
  URL : 'http://localhost:8082'
}

export class SearchFlightsDTO {
  StartDate = new Date(1/1/1970);
  Start = "";
  Destination = "";
  RemainingTickets = 1;

  setFields(date : string, start : string, destination: string, remainingTickets: number){
    if(date !== "")
      this.StartDate = new Date(date);
    this.Start = start ;
    this.Destination = destination;
    if(!Number.isNaN(remainingTickets))
      this.RemainingTickets = remainingTickets;
  }
}
