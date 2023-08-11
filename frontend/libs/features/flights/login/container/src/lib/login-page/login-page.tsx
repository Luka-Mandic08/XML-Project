import { LoginToBookingApp, LoginToFlightsApp } from '@frontend/features/flights/login/data-access';
import { AppRoutes, BookingAppRoutes } from '@frontend/models';
import { Container, Typography } from '@mui/material';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import styles from './login-page.module.css';

/* eslint-disable-next-line */
export interface LoginPageProps {
  isBookingApp?: boolean;
}

export function LoginPage(props: LoginPageProps) {
  const [error, setError] = useState<string>('');

  const navigate = useNavigate();
  const {
    register,
    getValues,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      username: '',
      password: '',
    },
  });

  const onSubmit = async (data: any) => {
    if (props.isBookingApp === true) {
      await LoginToBookingApp(data.username, data.password);
      if (localStorage.getItem('role') === 'Guest') {
        navigate(BookingAppRoutes.HomeGuest);
      } else if (localStorage.getItem('role') === 'Host') {
        navigate(BookingAppRoutes.HomeHost);
      } else {
        setError('Wrong credentials');
      }
    } else {
      const rsp = await LoginToFlightsApp(data.username, data.password);
      if (rsp === undefined) {
        setError('Wrong credentials');
      } else {
        navigate(AppRoutes.Home);
      }
    }
  };

  return (
    <Container maxWidth="sm" className={styles.loginContainer}>
      <Typography variant="h4" sx={{ mb: '2rem' }} align="center">
        Log in
      </Typography>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.inputContainer}>
          <input
            type="text"
            id="username"
            value={watch('username')}
            {...register('username', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="username" id="label-username">
            <div className={styles.text}>Username</div>
          </label>
          <label className={styles.errorLabel}>{errors.username?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="password"
            id="password"
            value={watch('password')}
            {...register('password', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="password" id="label-password">
            <div className={styles.text}>Password</div>
          </label>
          <label className={styles.errorLabel}>{errors.password?.message}</label>
        </div>

        <input
          style={{ width: '50%', marginLeft: '25%', marginRight: '25%', marginTop: '1rem' }}
          type="submit"
          value={props.isBookingApp ? 'Login to Booking App' : 'Login to Flights App'}
        />
      </form>

      <Typography variant="subtitle1" color={'red'} align="center" sx={{ mt: '0.5rem' }}>
        {error}
      </Typography>
    </Container>
  );
}

export default LoginPage;
