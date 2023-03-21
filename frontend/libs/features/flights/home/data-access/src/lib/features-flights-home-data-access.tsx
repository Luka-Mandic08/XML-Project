import styles from './features-flights-home-data-access.module.css';

/* eslint-disable-next-line */
export interface FeaturesFlightsHomeDataAccessProps {}

export function FeaturesFlightsHomeDataAccess(
  props: FeaturesFlightsHomeDataAccessProps
) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to FeaturesFlightsHomeDataAccess!</h1>
    </div>
  );
}

export default FeaturesFlightsHomeDataAccess;
