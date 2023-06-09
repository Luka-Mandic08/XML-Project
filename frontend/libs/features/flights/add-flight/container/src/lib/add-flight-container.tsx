import styles from './add-flight-container.module.css';
import { useForm } from 'react-hook-form';
import { Container, Typography } from '@mui/material';
import { AppRoutes, NewFlight } from '@frontend/models';
import { AddNewFlight } from '@frontend/features/flights/add-flight/data-access';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface AddFlightContainerProps {}

export function AddFlightContainer(props: AddFlightContainerProps) {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      startdate: Date.now(),
      arrivaldate: Date.now(),
      destination: '',
      start: '',
      price: 1,
      totaltickets: 1,
    },
  });

  const navigate = useNavigate();
  const goToHome = () => {
    navigate(AppRoutes.Home);
  };

  const onSubmit = (data: NewFlight) => {
    AddNewFlight(data);
    goToHome();
  };

  const currentDate = new Date().toISOString();

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" sx={{ mb: '2rem' }}>
        Schedule new flight
      </Typography>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.inputContainer}>
          <input
            type="datetime-local"
            id="startdate"
            value={watch('startdate')}
            {...register('startdate', {
              required: 'This field is required.',
              min: { value: currentDate, message: 'Selected date is in past.' },
            })}
          />
          <label className={styles.label} htmlFor="startdate" id="label-startdate">
            <div className={styles.text}>Starting Date</div>
          </label>
          <label className={styles.errorLabel}>{errors.startdate?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="datetime-local"
            id="arrivaldate"
            value={watch('arrivaldate')}
            {...register('arrivaldate', {
              required: 'This field is required.',
              min: { value: watch('startdate'), message: 'Selected date is before starting date.' },
            })}
          />
          <label className={styles.label} htmlFor="arrivaldate" id="label-arrivaldate">
            <div className={styles.text}>Arraval Date</div>
          </label>
          <label className={styles.errorLabel}>{errors.arrivaldate?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="destination"
            value={watch('destination')}
            {...register('destination', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="destination" id="label-destination">
            <div className={styles.text}>Destination</div>
          </label>
          <label className={styles.errorLabel}>{errors.destination?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input type="text" id="start" value={watch('start')} {...register('start', { required: 'This field is required.' })} />
          <label className={styles.label} htmlFor="start" id="label-start">
            <div className={styles.text}>Starting location</div>
          </label>
          <label className={styles.errorLabel}>{errors.start?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="price"
            value={watch('price')}
            {...register('price', {
              required: 'This field is required.',
              min: {
                value: 1,
                message: 'Minimal price is 1$.',
              },
            })}
          />
          <label className={styles.label} htmlFor="price" id="label-price">
            <div className={styles.text}>Price</div>
          </label>
          <label className={styles.errorLabel}>{errors.price?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="totaltickets"
            value={watch('totaltickets')}
            {...register('totaltickets', {
              required: 'This field is required.',
              min: {
                value: 1,
                message: 'Minimal number of passengers is 1.',
              },
            })}
          />
          <label className={styles.label} htmlFor="totaltickets" id="label-totaltickets">
            <div className={styles.text}>Number of passengers</div>
          </label>
          <label className={styles.errorLabel}>{errors.totaltickets?.message}</label>
        </div>

        <input style={{ width: '50%', marginLeft: '25%', marginRight: '25%' }} type="submit" value="Schedule flight" />
      </form>
    </Container>
  );
}

export default AddFlightContainer;
