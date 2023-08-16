import { GetAllAccomodation, SearchAccommodation } from '@frontend/features/booking/accomodation/data';
import { AccommodationInfo, SearchedAccommodationInfo } from '@frontend/models';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import AccommodationCard from '../accomodation-card/accomodation-card';
import SearchedAccommodationCard from '../searched-accommodation-card/searched-accommodation-card';
import styles from './all-accommodation.module.css';
import { Grid, Button, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';

/* eslint-disable-next-line */
export interface AllAccommodationProps {}

export function AllAccommodation(props: AllAccommodationProps) {
  const [accomodationInfo, setAccomodationInfo] = useState<AccommodationInfo[]>([]);
  const [searchedAccomodationInfo, setSearchedAccomodationInfo] = useState<SearchedAccommodationInfo[]>([]);
  const [searched, setSearched] = useState<boolean>(false);

  const navigate = useNavigate();

  useEffect(() => {
    getAllAccomodation();
  }, []);

  const getAllAccomodation = async () => {
    setAccomodationInfo(await GetAllAccomodation());
  };

  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors },
  } = useForm({
    defaultValues: {
      city: '',
      country: '',
      dateFrom: '',
      dateTo: '',
      numberOfGuests: '',
    },
  });

  const onSubmit = async (data: any) => {
    setSearched(true);
    const res = await SearchAccommodation(data);
    setSearchedAccomodationInfo(res);
  };

  const resetSearch = () => {
    reset();
    setSearched(false);
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.lineContainer}>
          <div className={styles.inputContainer}>
            <input
              type="date"
              id="dateFrom"
              value={watch('dateFrom')}
              {...register('dateFrom', {
                required: 'This field is required.',
              })}
            />
            <label className={styles.label} htmlFor="dateFrom" id="label-dateFrom">
              <div className={styles.text}>From</div>
            </label>
            <label className={styles.errorLabel}>{errors.dateFrom?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="date"
              id="dateTo"
              value={watch('dateTo')}
              {...register('dateTo', {
                required: 'This field is required.',
              })}
            />
            <label className={styles.label} htmlFor="dateTo" id="label-dateTo">
              <div className={styles.text}>Until</div>
            </label>
            <label className={styles.errorLabel}>{errors.dateTo?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input type="text" id="city" value={watch('city')} {...register('city')} />
            <label className={styles.label} htmlFor="city" id="label-city">
              <div className={styles.text}>City</div>
            </label>
            <label className={styles.errorLabel}>{errors.city?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input type="text" id="country" value={watch('country')} {...register('country')} />
            <label className={styles.label} htmlFor="country" id="label-country">
              <div className={styles.text}>Country</div>
            </label>
            <label className={styles.errorLabel}>{errors.country?.message}</label>
          </div>

          <div className={styles.inputContainer}>
            <input
              type="number"
              id="numberOfGuests"
              value={watch('numberOfGuests')}
              {...register('numberOfGuests', {
                required: 'This field is required.',
              })}
            />
            <label className={styles.label} htmlFor="numberOfGuests" id="label-numberOfGuests">
              <div className={styles.text}>Number of Guests</div>
            </label>
            <label className={styles.errorLabel}>{errors.numberOfGuests?.message}</label>
          </div>

          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Search
          </Button>
          <Button
            variant="contained"
            size="large"
            type="reset"
            onClick={resetSearch}
            sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Reset
          </Button>
        </div>
      </form>

      {!searched && (
        <div className={styles.cardsContainer}>
          {accomodationInfo?.map((accomodation, key) => (
            <AccommodationCard accomodationInfo={accomodation} isForHost={false} />
          ))}
        </div>
      )}

      {searched && (
        <>
          <Typography variant="h3" align="left" margin={'0 0 2rem 2rem'}>
            Searched Results
          </Typography>
          <div className={styles.cardsContainer}>
            {searchedAccomodationInfo?.map((accomodation, key) => (
              <SearchedAccommodationCard searchedAccomodationInfo={accomodation} />
            ))}
          </div>
        </>
      )}
    </>
  );
}

export default AllAccommodation;
