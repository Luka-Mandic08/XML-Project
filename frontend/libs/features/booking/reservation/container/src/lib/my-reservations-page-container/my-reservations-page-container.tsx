import styles from './my-reservations-page-container.module.css';

/* eslint-disable-next-line */
export interface MyReservationsPageContainerProps {}

export function MyReservationsPageContainer(props: MyReservationsPageContainerProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to MyReservationsPageContainer!</h1>
    </div>
  );
}

export default MyReservationsPageContainer;
