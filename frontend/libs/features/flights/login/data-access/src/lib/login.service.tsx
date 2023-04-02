import { AppRoutes, BaseURL, NewUser } from '@frontend/models';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export async function login(username: string, password: string) {
  let rsp;
  await axios({
    method: 'post',
    url: BaseURL.URL + '/login',
    data: { username, password },
  })
    .then((response) => {
      rsp = response.data;
      localStorage.setItem('id', rsp.id);
      localStorage.setItem('role', rsp.role);
      // eslint-disable-next-line @typescript-eslint/no-empty-function
    })
    // eslint-disable-next-line @typescript-eslint/no-empty-function
    .catch((error) => {});

  return rsp;
}

export async function AddNewUser(user: NewUser) {
  await axios({
    method: 'post',
    url: BaseURL.URL + '/user/add',
    data: {
      name: user.name,
      surname: user.surname,
      phoneNumber: user.phoneNumber,
      address: user.address,
      credentials: user.credentials,
      role: 'USER',
    },
  })
    .then(() => {
      alert('Successful registration');
    })
    .catch(() => {
      alert('Error while registering, please try again');
    });
}
