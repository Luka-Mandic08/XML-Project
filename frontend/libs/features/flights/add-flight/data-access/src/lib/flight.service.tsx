import { BaseURL, Flight } from "@frontend/models";
import axios from "axios";

export async function AddNewFlight(flight: Flight){
  await (axios({
    method: 'post', 
    url: BaseURL.URL + "/flight/add",
    data: { id              :flight.id,               
            startdate       :flight.startdate,        
            arrivaldate     :flight.arrivaldate,     
            destination     :flight.destination,     
            start           :flight.start,   
            price           :flight.price,
            remainingtickets:flight.remainingtickets,
            totaltickets    :flight.totaltickets}
    }))
}

export async function DeleteFlight(flightId: string){
    await (axios({
        method: 'delete', 
        url: BaseURL.URL + "/flight/delete/" + flightId,
        data: { }
    }))
}


