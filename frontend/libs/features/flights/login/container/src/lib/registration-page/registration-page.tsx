import { useForm } from 'react-hook-form';
import styles from './registration-page.module.css';
import { AppRoutes, NewUser } from '@frontend/models';
import { AddNewUser } from '@frontend/features/flights/login/data-access';
import { Container, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface RegistrationPageProps {}

export function RegistrationPage(props: RegistrationPageProps) {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      name: '',
      surname: '',
      phoneNumber: '',
      address: {
        street: '',
        city: '',
        country: '',
      },
      credentials: {
        username: '',
        password: '',
      },
      role: 'USER',
    },
  });

  const onSubmit = (data: NewUser) => {
    AddNewUser(data);
    navigate(AppRoutes.Login);
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" sx={{ mb: '2rem' }}>
        Register
      </Typography>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.inputContainer}>
          <input
            type="text"
            id="name"
            value={watch('name')}
            {...register('name', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="name" id="label-name">
            <div className={styles.text}>Name</div>
          </label>
          <label className={styles.errorLabel}>{errors.name?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="surname"
            value={watch('surname')}
            {...register('surname', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="surname" id="label-surname">
            <div className={styles.text}>Surname</div>
          </label>
          <label className={styles.errorLabel}>{errors.surname?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="phoneNumber"
            value={watch('phoneNumber')}
            {...register('phoneNumber', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="phoneNumber" id="label-phoneNumber">
            <div className={styles.text}>Phone number</div>
          </label>
          <label className={styles.errorLabel}>{errors.phoneNumber?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input type="text" id="address.street" value={watch('address.street')} {...register('address.street', { required: 'This field is required.' })} />
          <label className={styles.label} htmlFor="address.street" id="label-address.street">
            <div className={styles.text}>Street</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.street?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="address.city"
            value={watch('address.city')}
            {...register('address.city', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="address.city" id="label-praddress.city">
            <div className={styles.text}>City</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.city?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="address.country"
            value={watch('address.country')}
            {...register('address.country', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="address.country" id="label-address.country">
            <div className={styles.text}>Country</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.country?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="credentials.username"
            value={watch('credentials.username')}
            {...register('credentials.username', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="credentials.username" id="label-credentials.username">
            <div className={styles.text}>Username</div>
          </label>
          <label className={styles.errorLabel}>{errors.credentials?.username?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="password"
            id="credentials.password"
            value={watch('credentials.password')}
            {...register('credentials.password', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="credentials.password" id="label-credentials.password">
            <div className={styles.text}>Password</div>
          </label>
          <label className={styles.errorLabel}>{errors.credentials?.password?.message}</label>
        </div>

        <input style={{ width: '50%', marginLeft: '25%', marginRight: '25%' }} type="submit" value="Register" />
      </form>
    </Container>
  );
}

export default RegistrationPage;
