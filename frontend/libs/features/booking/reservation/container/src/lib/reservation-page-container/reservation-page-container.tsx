import styles from './reservation-page-container.module.css';

/* eslint-disable-next-line */
export interface ReservationPageContainerProps {}

export function ReservationPageContainer(props: ReservationPageContainerProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to ReservationPageContainer!</h1>
    </div>
  );
}

export default ReservationPageContainer;
