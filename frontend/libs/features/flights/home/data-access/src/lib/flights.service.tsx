import { BaseURL, SearchFlightsDTO } from "@frontend/models";
import axios from "axios";

export async function GetAllFlights(){
  let flights;

  await axios.get(BaseURL.URL + "/flight/all").then((response) => {
    flights = response.data });

  return flights;
}

export async function BuyFlightTickets(flightId : string, amount : number){
  console.log(amount);
  //await (await axios.put(BaseURL.URL + "/flight/buyticket/" + flightId));
}

export async function SearchFlights(dto: SearchFlightsDTO){
  let flights;

  await axios.put(BaseURL.URL + "/flight/search", dto).then((response) => {
    flights = response.data });

  return flights;
}

