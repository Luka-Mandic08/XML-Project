import { AvailabilityDate } from '@frontend/models';
import styles from './availability-calendar.module.css';
import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { Grid, Typography, Button, Paper } from '@mui/material';
import { GetAvailableDatesForAccommodation, UpdateAvailableDatesForAccommodation } from '@frontend/features/booking/accomodation/data';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';

/* eslint-disable-next-line */
export interface AvailabilityCalendarProps {
  shouldRenderCalendar: boolean;
}

export function AvailabilityCalendar(props: AvailabilityCalendarProps) {
  const [availabilityDates, setAvailabilityDates] = useState<AvailabilityDate[]>([]);
  const [month, setMonth] = useState<string>(new Date().toLocaleString('default', { month: 'long' }));
  const [year, setYear] = useState<number>(new Date().getFullYear());
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);

  const months = ['January', 'February', 'March', 'April', 'May', 'June', 'Jully', 'August', 'September', 'October', 'November', 'December'];
  const weekDays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];

  useEffect(() => {
    if (selectedAccommodation.id !== '') renderCalendar();
  }, [month, year, props.shouldRenderCalendar]);

  const renderCalendar = async () => {
    const availabilityDates: AvailabilityDate[] = [];
    const numberOfDaysInMonth = getNumberOfDaysInMonth(month, year);

    const firstDayOfMonth = new Date(year, months.indexOf(month), 0).getDay();

    const startOfCalendar = new Date(year, months.indexOf(month) - 1, getNumberOfDaysInMonth(months[months.indexOf(month) - 1], year) - firstDayOfMonth);
    const endOfCalendar = new Date(year, months.indexOf(month) + 1, 7 - new Date(year, months.indexOf(month), numberOfDaysInMonth).getDay());

    const availabilityDatesFromBackend: AvailabilityDate[] = new Array<AvailabilityDate>();
    const res: any = await GetAvailableDatesForAccommodation({
      accommodationId: selectedAccommodation.id,
      dateFrom: startOfCalendar,
      dateTo: endOfCalendar,
    });

    if (res) {
      res.availabilityDates?.forEach((availabilityDate: any) => {
        availabilityDatesFromBackend.push({
          date: new Date(availabilityDate.date.seconds * 1000),
          isAvailable: availabilityDate.isAvailable,
          price: availabilityDate.price,
        });
      });
    }

    let i = 0;
    while (startOfCalendar < endOfCalendar) {
      const tempDate = new Date(startOfCalendar);
      if (availabilityDatesFromBackend[i]?.date?.toDateString() === tempDate.toDateString()) {
        availabilityDates.push(availabilityDatesFromBackend[i]);
        i++;
      } else {
        availabilityDates.push({
          date: tempDate,
          isAvailable: undefined,
          price: 0,
        });
      }
      startOfCalendar.setDate(startOfCalendar.getDate() + 1);
    }

    setAvailabilityDates(availabilityDates);
  };

  const setColor = (isAvailable: boolean | undefined) => {
    if (isAvailable === false) {
      return 'lightcoral';
    }
    if (isAvailable === true) {
      return 'lightgreen';
    }
    return 'lightgrey';
    //return isAvailable ? (isAvailable === true ? 'lightgreen' : 'lightgrey') : isAvailable === false ? 'lightcoral' : 'lightgrey';
  };

  const getNumberOfDaysInMonth = (month: string, year: number) => {
    if (month === 'February') {
      if (year % 4 === 0) {
        return 29;
      } else {
        return 28;
      }
    }
    if (month === 'April' || month === 'June' || month === 'September' || month === 'November') {
      return 30;
    }
    return 31;
  };

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      month: month,
      year: year,
    },
  });

  const onSubmit = (data: any) => {
    setMonth(data.month);
    setYear(data.year);
  };

  return (
    <div style={{ width: '100%' }}>
      {selectedAccommodation.id !== '' && (
        <>
          <Grid container alignItems={'left'} direction={'column'}>
            <Grid item marginBottom={'0.5rem'}>
              <Typography variant="h5" align="left">
                Change month and year below
              </Typography>
            </Grid>
          </Grid>

          <form onSubmit={handleSubmit(onSubmit)}>
            <div className={styles.lineContainer}>
              <div className={styles.inputContainer}>
                <select
                  id="month"
                  value={watch('month')}
                  {...register('month', {
                    required: 'This field is required.',
                  })}
                >
                  {months.map((month) => (
                    <option value={month}>{month}</option>
                  ))}
                </select>
                <label className={styles.label} htmlFor="month" id="label-month">
                  <div className={styles.text}>Month</div>
                </label>
                <label className={styles.errorLabel}>{errors.month?.message}</label>
              </div>

              <div className={styles.inputContainer}>
                <input
                  type="number"
                  id="year"
                  value={watch('year')}
                  {...register('year', {
                    required: 'This field is required.',
                    min: 2000,
                    max: 2500,
                  })}
                />
                <label className={styles.label} htmlFor="year" id="label-year">
                  <div className={styles.text}>Year</div>
                </label>
                <label className={styles.errorLabel}>{errors.year?.message}</label>
              </div>

              <Button
                variant="contained"
                size="large"
                type="submit"
                sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
              >
                Change
              </Button>
            </div>
          </form>

          <div className={styles.inlineGrid}>
            <Paper elevation={6} className={styles.calendarContainer}>
              {weekDays.map((weekDay) => (
                <div className={styles.calendarHeaderItem}>{weekDay}</div>
              ))}
              {availabilityDates.map((availabilityDate) => (
                <div className={styles.calendarItem} style={{ backgroundColor: setColor(availabilityDate.isAvailable) }}>
                  <div className={styles.calendarItemDate}>{availabilityDate.date.toDateString()}</div>
                  <div className={styles.calendarItemPrice}>{availabilityDate.price ? `Price:  ${availabilityDate.price}rsd` : ''}</div>
                  <div className={styles.calendarItemAvailability}>
                    {availabilityDate.isAvailable
                      ? availabilityDate.isAvailable === true
                        ? 'Available'
                        : ''
                      : availabilityDate.isAvailable === false
                      ? 'Not available'
                      : ''}
                  </div>
                </div>
              ))}
            </Paper>
          </div>
        </>
      )}
      {selectedAccommodation.id === '' && (
        <Typography variant="h4" align="left">
          Please select accommodation to see availability.
        </Typography>
      )}
    </div>
  );
}

export default AvailabilityCalendar;
