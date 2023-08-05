import styles from './features-booking-login-container.module.css';

/* eslint-disable-next-line */
export interface FeaturesBookingLoginContainerProps {}

export function FeaturesBookingLoginContainer(props: FeaturesBookingLoginContainerProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to FeaturesBookingLoginContainer!</h1>
    </div>
  );
}

export default FeaturesBookingLoginContainer;
