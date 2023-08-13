import styles from './availability-calendar.module.css';

/* eslint-disable-next-line */
export interface AvailabilityCalendarProps {}

export function AvailabilityCalendar(props: AvailabilityCalendarProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to AvailabilityCalendar!</h1>
    </div>
  );
}

export default AvailabilityCalendar;
