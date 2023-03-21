import styles from './helpers.module.css';

/* eslint-disable-next-line */
export interface HelpersProps {}

export function Helpers(props: HelpersProps) {
  return (
    <div className={styles['container']}>
      <h1>Welcome to Helpers!</h1>
    </div>
  );
}

export default Helpers;
