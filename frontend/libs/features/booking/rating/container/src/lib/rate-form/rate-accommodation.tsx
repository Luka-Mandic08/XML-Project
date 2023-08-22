import { useForm } from 'react-hook-form';
import styles from './rate-accommodation.module.css';
import { Button, Paper, Rating, Typography } from '@mui/material';
import { RateAccommodation, RateHost } from '@frontend/features/booking/rating/data-access';

/* eslint-disable-next-line */
export interface RateAccommodationProps {
  accommodationId?: string;
  hostId?: string;
}

export function RateAccommodationOrHostForm(props: RateAccommodationProps) {
  const {
    register,
    handleSubmit,
    watch,
    setValue,
    formState: { errors },
  } = useForm({
    defaultValues: {
      ratedId: '',
      guestId: localStorage.getItem('userId'),
      date: '',
      score: 0,
      comment: '',
    },
  });

  const onSubmit = async (data: any) => {
    data.date = new Date().toISOString();
    if (props.accommodationId) {
      data.ratedId = props.accommodationId;
      RateAccommodation(data);
    }
    if (props.hostId) {
      data.ratedId = props.hostId;
      RateHost(data);
    }
  };

  return (
    <Paper elevation={6} className={styles.rateForm}>
      {props.accommodationId && (
        <Typography variant="h5" align="left">
          Rate accommodation
        </Typography>
      )}
      {props.hostId && (
        <Typography variant="h5" align="left">
          Rate host
        </Typography>
      )}
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.inputContainer}>
          <input type="text" id="comment" value={watch('comment')} {...register('comment')} />
          <label className={styles.label} htmlFor="comment" id="label-comment">
            <div className={styles.text}>Comment</div>
          </label>
          <label className={styles.errorLabel}>{errors.comment?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="score"
            value={watch('score')}
            step={1}
            {...register('score', {
              required: 'This field is required.',
              min: { value: 1, message: 'Min score is 1' },
              max: { value: 5, message: 'Max score is 5' },
              pattern: { value: /^[1-5]$/, message: 'Score must be whole number between 1 and 5' },
            })}
          />
          <Rating
            value={watch('score')}
            precision={1}
            size="large"
            sx={{ marginLeft: '1rem', alignItems: 'center' }}
            onChange={(ecent, newValue) => {
              setValue('score', newValue!);
            }}
          />

          <label className={styles.label} htmlFor="score" id="label-score">
            <div className={styles.text}>Score</div>
          </label>
          <label className={styles.errorLabel}>{errors.score?.message}</label>
        </div>

        <div className={styles.inputContainer} style={{ justifyContent: 'center' }}>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Rate
          </Button>
        </div>
      </form>
    </Paper>
  );
}

export default RateAccommodationOrHostForm;
