import { login } from "@frontend/features/flights/login/data-access";
import { AppRoutes } from "@frontend/models";
import { Grid, Typography, TextField, Button, Divider, Link  } from "@mui/material";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

/* eslint-disable-next-line */
export interface LoginPageProps {}

export function LoginPage(props: LoginPageProps) {

  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [error, setError] = useState<string>("");
  const navigate= useNavigate();

  const onUsernameChange = (e: any) => setUsername(e.target.value);
  const onPasswordChange = (e: any) => setPassword(e.target.value);

  const handleSubmit = async () => {
    let rsp;
    if(username !== '' && password !== ''){
      rsp = await login(username, password)
      if(rsp === undefined){
        console.log(rsp)
        setError("Wrong credentials")
      }
      navigate(AppRoutes.Home);
    }
    else
      setError("Please fill all fields")
  }

  return (
    <Grid container direction="row" justifyContent="center" alignItems="center" sx={{marginTop: '10%', width: '100%'}}>
      <Grid xs={4}>
      </Grid>
      <Grid container direction="column" justifyContent="start" alignItems="center" xs={3} sx={{ border: "3px solid #212121", height: 400}}>
        <Grid item sx={{mb: 4, mt: 6}}>
          <Typography variant="h4">Login</Typography>
        </Grid>
        <Grid item sx={{mb: 1}}>
          <TextField onChange={onUsernameChange} value={username} id="username" label="username" variant="outlined" sx={{ width: 290}} inputProps={{style: {height: 15}}}/>
        </Grid>
        <Grid item sx={{mb: 3}}>
          <TextField type="password" onChange={onPasswordChange} value={password}  id="password" label="password" variant="outlined" sx={{ width: 290}} inputProps={{style: {height: 15}}}/>
          <Typography color="red">{error}</Typography>
        </Grid>
        <Grid item sx={{mb: 1}}>
          <Button onClick={handleSubmit} variant="contained" sx={{width: 290,  backgroundColor: "#212121",  '&:hover': { backgroundColor: '#ffffff', color:  "#212121" }}}>Login</Button>
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
