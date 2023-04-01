import { Box, AppBar, Toolbar, Typography, Button, IconButton } from '@mui/material';
import AirplaneTicketOutlinedIcon from '@mui/icons-material/AirplaneTicketOutlined';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AppRoutes } from '@frontend/models';

/* eslint-disable-next-line */
export interface NavBarProps {}

export function NavBar(props: NavBarProps) {
  const [role, setRole] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    setRole(localStorage.getItem('role')!);
    console.log(role);
  }, [role, navigate]);

  const logout = () => {
    localStorage.removeItem('role');
    localStorage.removeItem('id');
    navigate(AppRoutes.Home);
    window.location.reload();
  };

  return (
    <Box sx={{ flexGrow: 1, mb: 10 }}>
      <AppBar position="fixed" sx={{ backgroundColor: '#212121' }}>
        <Toolbar>
          <IconButton onClick={() => navigate(AppRoutes.Home)} color="inherit" aria-label="home">
            <AirplaneTicketOutlinedIcon />
          </IconButton>
          <Typography
            variant="h6"
            component="div"
            sx={{ flexGrow: 1, ml: 1, cursor: 'default', '&:hover': { cursor: 'pointer' } }}
            onClick={() => navigate(AppRoutes.Home)}
          >
            eFlight
          </Typography>
          {role === null && (
            <Button onClick={() => navigate(AppRoutes.Login)} color="inherit" sx={{ textTransform: 'none' }}>
              Login
            </Button>
          )}
          {role === 'USER' && (
            <>
              <Button onClick={() => navigate(AppRoutes.PurchasedTickets)} color="inherit" sx={{ textTransform: 'none', marginRight: '30px' }}>
                My tickets
              </Button>
              <Button onClick={() => logout()} color="inherit" sx={{ textTransform: 'none' }}>
                Log out
              </Button>
            </>
          )}
          {role === 'ADMIN' && (
            <Button onClick={() => logout()} color="inherit" sx={{ textTransform: 'none' }}>
              Log out
            </Button>
          )}
          {role === 'NONE' && (
            <Button onClick={() => navigate(AppRoutes.Login)} color="inherit" sx={{ textTransform: 'none' }}>
              Login
            </Button>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default NavBar;
