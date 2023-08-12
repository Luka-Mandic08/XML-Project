import styles from './host-accomodation.module.css';

/* eslint-disable-next-line */
export interface HostAccomodationProps {}

export function HostAccomodation(props: HostAccomodationProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to HostAccomodation!</h1>
    </div>
  );
}

export default HostAccomodation;
