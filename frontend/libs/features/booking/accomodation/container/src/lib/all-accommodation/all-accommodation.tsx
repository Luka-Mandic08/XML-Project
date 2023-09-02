import { GetAllAccomodation, SearchAccommodation } from '@frontend/features/booking/accomodation/data';
import { AccommodationInfo, SearchedAccommodationInfo } from '@frontend/models';
import { useState, useEffect } from 'react';
import AccommodationCard from '../accomodation-card/accomodation-card';
import SearchedAccommodationCard from '../searched-accommodation-card/searched-accommodation-card';
import styles from './all-accommodation.module.css';
import { Grid, Button, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { useSearchParametersStore } from '@frontend/features/booking/store/container';

/* eslint-disable-next-line */
export interface AllAccommodationProps {}

export function AllAccommodation(props: AllAccommodationProps) {
  const [accomodationInfo, setAccomodationInfo] = useState<AccommodationInfo[]>([]);
  const [pageNumber, setPageNumber] = useState<number>(1);
  const [shouldLoadMore, setShouldLoadMore] = useState<boolean>(true);

  const [searchedAccomodationInfo, setSearchedAccomodationInfo] = useState<SearchedAccommodationInfo[]>([]);
  const [searched, setSearched] = useState<boolean>(false);
  const [searchPageNumber, setSearchPageNumber] = useState<number>(1);
  const [searchShouldLoadMore, setSearchShouldLoadMore] = useState<boolean>(true);
  const setSearchParameters = useSearchParametersStore((state) => state.setSearchParameters);

  useEffect(() => {
    getAllAccomodation();
  }, []);

  const getAllAccomodation = async () => {
    const newAccomodations = await GetAllAccomodation(pageNumber);
    if (newAccomodations === undefined) {
      setPageNumber(pageNumber - 1);
      setShouldLoadMore(false);
      return;
    }
    setAccomodationInfo((prevAccomodations) => [...prevAccomodations, ...newAccomodations]);
    setPageNumber(pageNumber + 1);
  };

  const {
    register,
    handleSubmit,
    watch,
    reset,
    getValues,
    formState: { errors },
  } = useForm({
    defaultValues: {
      city: '',
      country: '',
      dateFrom: '',
      dateTo: '',
      numberOfGuests: '',
      maxPrice: 0,
      amenities: '',
      ownedByProminentHost: false,
    },
  });

  const onSubmit = async (data: any) => {
    setSearched(true);
    const res = await SearchAccommodation(data, searchPageNumber);
    if (res === undefined) {
      return;
    }
    setSearchedAccomodationInfo(res);
    setSearchParameters(data);
  };

  const resetSearch = () => {
    reset();
    setSearched(false);
    setSearchedAccomodationInfo([]);
    setSearchPageNumber(1);
    setSearchShouldLoadMore(true);
  };

  useEffect(() => {
    setSearchPageNumber(1);
    setSearchedAccomodationInfo([]);
    setSearchShouldLoadMore(true);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [watch('dateFrom'), watch('dateTo'), watch('city'), watch('country'), watch('numberOfGuests')]);

  const loadMoreForSearch = async () => {
    setSearchPageNumber(searchPageNumber + 1);
    const data = {
      city: getValues('city'),
      country: getValues('country'),
      dateFrom: getValues('dateFrom'),
      dateTo: getValues('dateTo'),
      numberOfGuests: getValues('numberOfGuests'),
    };
    const newAccomodations = await SearchAccommodation(data, searchPageNumber);
    if (newAccomodations === undefined) {
      setSearchPageNumber(searchPageNumber - 1);
      setSearchShouldLoadMore(false);
      return;
    }
    setSearchedAccomodationInfo((prevAccomodations) => [...prevAccomodations, ...newAccomodations]);
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
                min: { value: new Date().toISOString().split('T')[0], message: 'Minimum date is today.' },
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
                min: { value: watch('dateFrom'), message: 'Minimum date is the date from.' },
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
                min: { value: 1, message: 'Minimum number of guests is 1.' },
              })}
            />
            <label className={styles.label} htmlFor="numberOfGuests" id="label-numberOfGuests">
              <div className={styles.text}>Number of Guests</div>
            </label>
            <label className={styles.errorLabel}>{errors.numberOfGuests?.message}</label>
          </div>
          <div className={styles.break} />

          <div className={styles.inputContainer}>
            <input type="number" id="maxPrice" value={watch('maxPrice')} {...register('maxPrice')} />
            <label className={styles.label} htmlFor="maxPrice" id="label-maxPrice">
              <div className={styles.text}>Max Price</div>
            </label>
            <label className={styles.errorLabel}>{errors.maxPrice?.message}</label>
          </div>

          <div className={styles.lineContainer2}>
            <div>
              <Typography variant="h6" paddingTop={'8px'}>
                Owned by outstanding host
              </Typography>
              <Typography variant="h6" color={'red'}>
                <label className={styles.errorLabel}>{errors.ownedByProminentHost?.message}</label>
              </Typography>
            </div>
            <input style={{ width: '48px', height: '48px' }} type="checkbox" id="ownedByProminentHost" {...register('ownedByProminentHost')} />
          </div>

          <div className={styles.inputContainer}>
            <input type="text" id="amenities" value={watch('amenities')} {...register('amenities')} />
            <label className={styles.label} htmlFor="amenities" id="label-amenities">
              <div className={styles.text}>Amenities</div>
            </label>
            <label className={styles.errorLabel}>{errors.amenities?.message}</label>
          </div>

          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Search
          </Button>
          <Button
            variant="contained"
            size="large"
            type="reset"
            onClick={resetSearch}
            sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Reset
          </Button>
        </div>
      </form>

      {!searched && (
        <>
          <div className={styles.cardsContainer}>
            {accomodationInfo?.map((accomodation, key) => (
              <AccommodationCard accomodationInfo={accomodation} isForHost={false} />
            ))}
          </div>
          <Grid container justifyContent={'center'} marginTop={'2rem'}>
            {shouldLoadMore && (
              <Button
                variant="contained"
                size="large"
                onClick={() => getAllAccomodation()}
                sx={{
                  color: 'white',
                  background: '#212121',
                  height: '48px',
                  minWidth: '200px',
                  ':hover': { background: 'white', color: '#212121' },
                }}
              >
                Load More
              </Button>
            )}
            {!shouldLoadMore && (
              <Typography variant="h5" align="center">
                No more accomodations to load
              </Typography>
            )}
          </Grid>
        </>
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
          <Grid container justifyContent={'center'} marginTop={'2rem'}>
            {searchShouldLoadMore && (
              <Button
                variant="contained"
                size="large"
                onClick={() => loadMoreForSearch()}
                sx={{
                  color: 'white',
                  background: '#212121',
                  height: '48px',
                  minWidth: '200px',
                  ':hover': { background: 'white', color: '#212121' },
                }}
              >
                Load More
              </Button>
            )}
            {!searchShouldLoadMore && (
              <Typography variant="h5" align="center">
                No more accomodations to load
              </Typography>
            )}
          </Grid>
        </>
      )}
    </>
  );
}

export default AllAccommodation;
