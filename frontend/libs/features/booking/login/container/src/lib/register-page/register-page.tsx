import { useNavigate } from 'react-router-dom';
import styles from './register-page.module.css';
import { AppRoutes, RegisterUser } from '@frontend/models';
import { Button, Paper, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { RegisterNewUser } from '@frontend/features/booking/login/data-access';

/* eslint-disable-next-line */
export interface RegisterPageProps {}

export function RegisterPage(props: RegisterPageProps) {
  const navigate = useNavigate();
  const {
    register,
    getValues,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      name: '',
      surname: '',
      email: '',
      address: {
        street: '',
        city: '',
        country: '',
      },
      username: '',
      password: '',
      confirmPassword: '',
      role: '',
    },
  });

  const onSubmit = (data: RegisterUser) => {
    RegisterNewUser(data);
    navigate(AppRoutes.Login);
  };

  return (
    <Paper elevation={6} className={styles.registerContainer}>
      <Typography variant="h4" sx={{ mb: '2rem' }} align={'center'}>
        Registration form
      </Typography>

      <form onSubmit={handleSubmit(onSubmit)}>
        <Typography variant="h6" marginBottom={'0.5rem'}>
          Peronal information
        </Typography>
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
            type="email"
            id="email"
            value={watch('email')}
            {...register('email', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="email" id="label-email">
            <div className={styles.text}>Email</div>
          </label>
          <label className={styles.errorLabel}>{errors.email?.message}</label>
        </div>
        <Typography variant="h6" marginBottom={'0.5rem'}>
          Address information
        </Typography>
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

        <Typography variant="h6" marginBottom={'0.5rem'}>
          Account information
        </Typography>
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

        <div className={styles.inputContainer}>
          <input
            type="password"
            id="confirmPassword"
            value={watch('confirmPassword')}
            {...register('confirmPassword', {
              required: 'This field is required.',
              validate: {
                isSameAsPassword: (v) => v === getValues('password') || 'Passwords do not match',
              },
            })}
          />
          <label className={styles.label} htmlFor="confirmPassword" id="label-confirmPassword">
            <div className={styles.text}>Confirm password</div>
          </label>
          <label className={styles.errorLabel}>{errors.confirmPassword?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <select
            id="role"
            value={watch('role')}
            {...register('role', {
              required: 'This field is required.',
            })}
          >
            <option value=""></option>
            <option value="Host">Host - Plans to rent accomodation</option>
            <option value="Guest">Guest - Plans to use accomodation</option>
          </select>
          <label className={styles.label} htmlFor="role" id="label-role">
            <div className={styles.text}>Usertype</div>
          </label>
          <label className={styles.errorLabel}>{errors.role?.message}</label>
        </div>

        <div className={styles.inputContainer} style={{ justifyContent: 'center' }}>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Register
          </Button>
        </div>
      </form>
    </Paper>
  );
}

export default RegisterPage;
