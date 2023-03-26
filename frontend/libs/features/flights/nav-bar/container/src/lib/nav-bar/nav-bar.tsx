import { Box, AppBar, Toolbar, Typography, Button, IconButton } from '@mui/material';
import AirplaneTicketOutlinedIcon from '@mui/icons-material/AirplaneTicketOutlined';

/* eslint-disable-next-line */
export interface NavBarProps {}

export function NavBar(props: NavBarProps) {
  return (
    <Box sx={{ flexGrow: 1, mb: 10 }}>
      <AppBar position="fixed" sx={{backgroundColor: '#212121'}}>
        <Toolbar>
          <AirplaneTicketOutlinedIcon />
          <Typography variant="h6" component="div" sx={{ flexGrow: 1, ml: 1 }}>
            eFlight
          </Typography>
          <Button color="inherit" sx={{textTransform: 'none'}}>Login</Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default NavBar;
