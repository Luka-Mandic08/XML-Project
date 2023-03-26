import { Card, Grid, Typography, OutlinedInput, TextField, Button, Divider, Link  } from "@mui/material";

/* eslint-disable-next-line */
export interface LoginPageProps {}

export function LoginPage(props: LoginPageProps) {
  return (
    <Grid container direction="row" justifyContent="center" alignItems="center" sx={{marginTop: '10%', width: '100%'}}>
      <Grid xs={4}>
      </Grid>
      <Grid container direction="column" justifyContent="start" alignItems="center" xs={3} sx={{ border: "3px solid #212121", height: 400}}>
        <Grid item sx={{mb: 4, mt: 6}}>
          <Typography variant="h4">Login</Typography>
        </Grid>
        <Grid item sx={{mb: 1}}>
          <TextField id="username" label="username" variant="outlined" sx={{ width: 290}} inputProps={{style: {height: 15}}}/>
        </Grid>
        <Grid item sx={{mb: 3}}>
          <TextField id="password" label="password" variant="outlined" sx={{ width: 290}} inputProps={{style: {height: 15}}}/>
        </Grid>
        <Grid item sx={{mb: 1}}>
          <Button variant="contained" sx={{width: 290,  backgroundColor: "#212121",  '&:hover': { backgroundColor: '#ffffff', color:  "#212121" }}}>Login</Button>
        </Grid>
        <Divider sx={{backgroundColor: "#212121", width: 280, mb: 1}}/>
        <Grid item sx={{mb: 1}}>
         <Button variant="contained" sx={{width: 290, backgroundColor: "#212121", '&:hover': { backgroundColor: '#ffffff', color:  "#212121" }}}>Register</Button>
        </Grid>
        <Grid item>
          or
          <Link href="#" underline="hover" sx={{ color: "#212121"}}>
            {'  Enter as guest'}
          </Link>
        </Grid>
      </Grid>
      <Grid xs={4}>
      </Grid>
      </Grid>
  );
}

export default LoginPage;
