import { AvailabilityDate } from '@frontend/models';
import styles from './availability-calendar.module.css';
import { useEffect, useState } from 'react';
import { set, useForm } from 'react-hook-form';
import { Grid, Typography, Button, Paper } from '@mui/material';
import { render } from 'react-dom';
import { GetAvailableDatesForAccommodation } from '@frontend/features/booking/accomodation/data';

/* eslint-disable-next-line */
export interface AvailabilityCalendarProps {}

export function AvailabilityCalendar(props: AvailabilityCalendarProps) {
  const [availabilityDates, setAvailabilityDates] = useState<AvailabilityDate[]>([]);
  const [month, setMonth] = useState<string>(new Date().toLocaleString('default', { month: 'long' }));
  const [year, setYear] = useState<number>(new Date().getFullYear());

  const months = ['January', 'February', 'March', 'April', 'May', 'June', 'Jully', 'August', 'September', 'October', 'November', 'December'];
  const weekDays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];

  useEffect(() => {
    renderCalendar();
  }, [month, year]);

  const renderCalendar = async () => {
    const availabilityDates: AvailabilityDate[] = [];
    const numberOfDaysInMonth = getNumberOfDaysInMonth(month, year);

    const firstDayOfMonth = new Date(year, months.indexOf(month), 1).getDay() - 1;

    const startOfCalendar = new Date(year, months.indexOf(month) - 1, getNumberOfDaysInMonth(months[months.indexOf(month) - 1], year) - firstDayOfMonth + 1);
    const endOfCalendar = new Date(year, months.indexOf(month) + 1, 7 - new Date(year, months.indexOf(month), numberOfDaysInMonth).getDay());

    const availabilityDatesFromBackend: AvailabilityDate[] = new Array<AvailabilityDate>();
    availabilityDatesFromBackend.concat(
      await GetAvailableDatesForAccommodation({
        hostId: localStorage.getItem('userId'),
        dateFrom: startOfCalendar,
        dateTo: endOfCalendar,
      })
    );
    availabilityDatesFromBackend.push({
      date: new Date(),
      isAvailable: true,
      price: 1000,
    });
    availabilityDatesFromBackend.push({
      date: new Date(year, months.indexOf(month), numberOfDaysInMonth),
      isAvailable: false,
      price: 1000,
    });
    let i = 0;
    while (startOfCalendar <= endOfCalendar) {
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
    if (isAvailable === undefined) {
      return 'lightgray';
    }
    if (isAvailable) {
      return 'lightgreen';
    }
    return 'lightcoral';
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
    updateAvailabilityMonthAndYear(data);
  };

  const updateAvailabilityMonthAndYear = (data: any) => {
    setMonth(data.month);
    setYear(data.year);
  };

  const {
    register: registerAvailabilityDates,
    handleSubmit: handleSubmitAvailabilityDates,
    watch: watchAvailabilityDates,
    formState: { errors: errorsAvailabilityDates },
  } = useForm({
    defaultValues: {
      dateFrom: '',
      dateTo: '',
      price: 0,
    },
  });

  const onSubmitAvailabilityDates = (data: any) => {
    /*call backend*/
  };

  return (
    <div style={{ padding: '1rem' }}>
      <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
        <Grid item marginBottom={'1rem'}>
          <Typography variant="h4" align="left">
            Availability for accommodation: name
          </Typography>
        </Grid>
        <Grid item marginBottom={'1rem'}>
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

        <Paper elevation={6} className={styles.updateAvailabilityForm}>
          <Typography variant="h5" align="left">
            Update availability
          </Typography>
          <form onSubmit={handleSubmitAvailabilityDates(onSubmitAvailabilityDates)}>
            <div className={styles.inputContainer}>
              <input
                type="date"
                id="dateFrom"
                value={watchAvailabilityDates('dateFrom')}
                {...registerAvailabilityDates('dateFrom', {
                  required: 'This field is required.',
                  min: { value: new Date().toISOString(), message: 'Date must be in the future.' },
                })}
              />
              <label className={styles.label} htmlFor="dateFrom" id="label-dateFrom">
                <div className={styles.text}>From</div>
              </label>
              <label className={styles.errorLabel}>{errorsAvailabilityDates.dateFrom?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="date"
                id="dateTo"
                value={watchAvailabilityDates('dateTo')}
                {...registerAvailabilityDates('dateTo', {
                  required: 'This field is required.',
                  min: { value: watchAvailabilityDates('dateFrom'), message: 'Selected date is before starting date.' },
                })}
              />
              <label className={styles.label} htmlFor="dateTo" id="label-dateTo">
                <div className={styles.text}>To</div>
              </label>
              <label className={styles.errorLabel}>{errorsAvailabilityDates.dateTo?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="number"
                id="price"
                value={watchAvailabilityDates('price')}
                {...registerAvailabilityDates('price', {
                  required: 'This field is required.',
                  min: 0,
                })}
              />
              <label className={styles.label} htmlFor="price" id="label-price">
                <div className={styles.text}>Price</div>
              </label>
              <label className={styles.errorLabel}>{errorsAvailabilityDates.price?.message}</label>
            </div>
            <Button
              variant="contained"
              size="large"
              type="submit"
              sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
            >
              Update availability
            </Button>
          </form>
        </Paper>
      </div>
    </div>
  );
}

export default AvailabilityCalendar;
