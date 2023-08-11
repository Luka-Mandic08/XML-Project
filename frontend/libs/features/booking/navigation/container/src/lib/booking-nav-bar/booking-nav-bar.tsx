import { AppRoutes, BookingAppRoutes, SharedRoutes } from '@frontend/models';
import { Box, AppBar, Toolbar, Typography, Button } from '@mui/material';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './booking-nav-bar.module.css';

/* eslint-disable-next-line */
export interface BookingNavBarProps {}

export function BookingNavBar(props: BookingNavBarProps) {
  const [role, setRole] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    setRole(localStorage.getItem('role')!);
  }, [role, navigate]);

  const logout = () => {
    localStorage.removeItem('role');
    localStorage.removeItem('userId');
    localStorage.removeItem('jwt');
    localStorage.removeItem('username');

    navigate(BookingAppRoutes.HomeGuest);
    window.location.reload();
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="absolute" sx={{ backgroundColor: '#212121', height: '70px' }}>
        <Toolbar>
          <Typography
            variant="h6"
            component="div"
            sx={{ flexGrow: 1, ml: 1, cursor: 'default', '&:hover': { cursor: 'pointer' } }}
            onClick={() => navigate(BookingAppRoutes.HomeGuest)}
          >
            eBooking
          </Typography>
          {role === null && (
            <>
              <Button onClick={() => navigate(SharedRoutes.Login)} color="inherit" sx={{ textTransform: 'none' }}>
                Login
              </Button>
              <Button onClick={() => navigate(BookingAppRoutes.Register)} color="inherit" sx={{ textTransform: 'none' }}>
                Register
              </Button>
            </>
          )}
          {role === 'Host' && (
            <>
              <Button onClick={() => navigate(BookingAppRoutes.HomeHost)} color="inherit" sx={{ textTransform: 'none', marginRight: '30px' }}>
                My Accomodation
              </Button>
              <Button onClick={() => navigate(BookingAppRoutes.Profile)} color="inherit" sx={{ textTransform: 'none', marginRight: '30px' }}>
                Profile
              </Button>
              <Button onClick={() => logout()} color="inherit" sx={{ textTransform: 'none' }}>
                Log out
              </Button>
            </>
          )}
          {role === 'Guest' && (
            <>
              <Button onClick={() => navigate(BookingAppRoutes.Profile)} color="inherit" sx={{ textTransform: 'none', marginRight: '30px' }}>
                Profile
              </Button>
              <Button onClick={() => logout()} color="inherit" sx={{ textTransform: 'none' }}>
                Log out
              </Button>
            </>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default BookingNavBar;
