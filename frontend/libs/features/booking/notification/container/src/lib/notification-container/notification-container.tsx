import { GetNotificationPreferences, GetNotifications, UpdateNotificationPreferences } from '@frontend/features/booking/notification/data-access';
import {
  Typography,
  Box,
  TableBody,
  TableCell,
  TableRow,
  Grid,
  styled,
  tableCellClasses,
  TableContainer,
  Paper,
  TableHead,
  Button,
  Checkbox,
  FormControlLabel,
} from '@mui/material';
import { useEffect, useState } from 'react';

/* eslint-disable-next-line */
export interface NotificationContainerProps {}

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  [`&.${tableCellClasses.head}`]: {
    backgroundColor: '#212121',
    color: theme.palette.common.white,
  },
  [`&.${tableCellClasses.body}`]: {
    fontSize: 15,
  },
}));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  '&:nth-of-type(odd)': {
    backgroundColor: theme.palette.action.hover,
  },
  // hide last border
  '&:last-child td, &:last-child th': {
    border: 0,
  },
}));

type NotificationPreferences = {
  ReservationCreated: boolean;
  ReservationCanceled: boolean;
  HostRated: boolean;
  AccommodationRated: boolean;
  OutstandingHostStatus: boolean;
  ReservationApprovedOrDenied: boolean;
};

export function NotificationContainer(props: NotificationContainerProps) {
  const [notifications, setNotifications] = useState<any[]>([]);
  const [notificationPreferences, setNotificationPreferences] = useState<NotificationPreferences>({
    ReservationCreated: false,
    ReservationCanceled: false,
    HostRated: false,
    AccommodationRated: false,
    OutstandingHostStatus: false,
    ReservationApprovedOrDenied: false,
  });

  const handleCheckboxChange = (event: { target: { name: string; checked: boolean } }) => {
    const { name, checked } = event.target;
    setNotificationPreferences({
      ...notificationPreferences,
      [name]: checked,
    });
  };

  const handleSave = () => {
    const truePreferences = Object.keys(notificationPreferences).filter((key) => (notificationPreferences as any)[key] === true);
    UpdateNotificationPreferences(truePreferences);
  };

  useEffect(() => {
    GetNotifications().then((data) => {
      setNotifications(data);
    });

    GetNotificationPreferences().then((data) => {
      setNotificationPreferences({
        ReservationCreated: data?.includes('ReservationCreated') ?? false,
        ReservationCanceled: data?.includes('ReservationCanceled') ?? false,
        HostRated: data?.includes('HostRated') ?? false,
        AccommodationRated: data?.includes('AccommodationRated') ?? false,
        OutstandingHostStatus: data?.includes('OutstandingHostStatus') ?? false,
        ReservationApprovedOrDenied: data?.includes('ReservationApprovedOrDenied') ?? false,
      });
    });
  }, []);

  return (
    <>
      <Box sx={{ mb: '40px' }}>
        <Typography variant="h4" sx={{ m: '35px' }}>
          Notification Preferences
        </Typography>
        {localStorage.getItem('role') === 'Host' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={<Checkbox name="ReservationCreated" checked={notificationPreferences.ReservationCreated} onChange={handleCheckboxChange} />}
            label="Reservation Created"
            style={{ color: 'black' }}
          />
        )}
        {localStorage.getItem('role') === 'Host' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={<Checkbox name="ReservationCanceled" checked={notificationPreferences.ReservationCanceled} onChange={handleCheckboxChange} />}
            label="Reservation Canceled"
          />
        )}
        {localStorage.getItem('role') === 'Host' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={<Checkbox name="HostRated" checked={notificationPreferences.HostRated} onChange={handleCheckboxChange} />}
            label="Host Rated"
          />
        )}
        {localStorage.getItem('role') === 'Host' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={<Checkbox name="AccommodationRated" checked={notificationPreferences.AccommodationRated} onChange={handleCheckboxChange} />}
            label="Accommodation Rated"
          />
        )}
        {localStorage.getItem('role') === 'Host' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={<Checkbox name="OutstandingHostStatus" checked={notificationPreferences.OutstandingHostStatus} onChange={handleCheckboxChange} />}
            label="Outstanding Host Status"
          />
        )}
        {localStorage.getItem('role') === 'Guest' && (
          <FormControlLabel
            sx={{ ml: '25px' }}
            control={
              <Checkbox name="ReservationApprovedOrDenied" checked={notificationPreferences.ReservationApprovedOrDenied} onChange={handleCheckboxChange} />
            }
            label="Reservation Approved Or Denied"
          />
        )}
        <Button variant="contained" sx={{ ml: '50px', background: 'black' }} onClick={handleSave}>
          Save
        </Button>
      </Box>
      <Box display="flex" justifyContent="center">
        <Box>
          <Grid sx={{ mt: 4 }}>
            <TableContainer component={Paper}>
              <TableHead>
                <TableRow sx={{ maxWidth: '900px' }}>
                  <StyledTableCell align="center" colSpan={3} sx={{ mt: 10, mb: 2, width: '1300px' }}>
                    <Typography variant="h4">Notifications</Typography>
                  </StyledTableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {notifications?.map((notification, idx) => (
                  <StyledTableRow key={idx} sx={{ width: '900px' }}>
                    <StyledTableCell sx={{ width: '700px' }}>{notification.notificationText}</StyledTableCell>
                    <StyledTableCell>{notification.type}</StyledTableCell>
                    <StyledTableCell>{new Date(notification.dateCreated.seconds * 1000).toDateString()}</StyledTableCell>
                  </StyledTableRow>
                ))}
              </TableBody>
            </TableContainer>
          </Grid>
        </Box>
      </Box>
    </>
  );
}

export default NotificationContainer;
