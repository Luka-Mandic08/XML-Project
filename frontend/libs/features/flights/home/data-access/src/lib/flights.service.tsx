import { BaseURL } from "@frontend/models";
import axios from "axios";

export async function GetAllFlights(){
  let flights;

  await axios.get(BaseURL.URL + "/flight/all").then((response) => {
    flights = response.data });

  return flights;
}

