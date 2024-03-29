export class SearchFlightsDTO {
  StartDate = new Date(1 / 1 / 1970);
  Start = '';
  Destination = '';
  RemainingTickets = 1;

  setFields(date: string, start: string, destination: string, remainingTickets: number) {
    if (date !== '') this.StartDate = new Date(date);
    this.Start = start;
    this.Destination = destination;
    if (!Number.isNaN(remainingTickets)) this.RemainingTickets = remainingTickets;
  }
}
