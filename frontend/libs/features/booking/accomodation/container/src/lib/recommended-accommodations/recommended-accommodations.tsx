import { GetAllRecommendedAccomodation } from '@frontend/features/booking/accomodation/data';
import { AccommodationInfo } from '@frontend/models';
import { Typography } from '@mui/material';
import { useState, useEffect } from 'react';
import AccommodationCard from '../accomodation-card/accomodation-card';
import styles from './recommended-accommodations.module.css';

/* eslint-disable-next-line */
export interface RecommendedAccommodationsProps {}

export function RecommendedAccommodations(props: RecommendedAccommodationsProps) {
  const [accomodationInfo, setAccomodationInfo] = useState<AccommodationInfo[]>([]);

  useEffect(() => {
    getRecommendedAccomodation();
  }, []);

  const getRecommendedAccomodation = async () => {
    const recommended = await GetAllRecommendedAccomodation();
    setAccomodationInfo(recommended);
  };

  return (
    <>
      <Typography variant="h3" align="left" margin={'0 0 2rem 2rem'}>
        Recommended Results
      </Typography>
      <div className={styles.cardsContainer}>
        {accomodationInfo?.map((accomodation, key) => (
          <AccommodationCard accomodationInfo={accomodation} isForHost={false} />
        ))}
      </div>
    </>
  );
}

export default RecommendedAccommodations;
