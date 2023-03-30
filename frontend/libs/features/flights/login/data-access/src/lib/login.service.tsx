import { BaseURL } from "@frontend/models";
import axios from "axios";

export async function login(username: string, password: string){
  let rsp
  await (axios({
    method: 'post',
    url: BaseURL.URL + "/login",
    data: {username, password},
  }).then((response) => {
    rsp = response.data;
    localStorage.setItem('id', rsp.id)
    localStorage.setItem('role', rsp.role)
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  }).catch((error) => {}));

  return rsp
}
