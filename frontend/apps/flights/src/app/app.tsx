// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { Route, Routes } from 'react-router-dom';
import { AppRoutes } from '@frontend/models';
import { HomeContainer } from '@frontend/features/flights/home/container';
import { LoginPage } from '@frontend/features/flights/login/container';
import { NavBar } from '@frontend/features/flights/nav-bar/container';

export function App() {
  return (
    <>
      <NavBar></NavBar>
      <Routes>
            <Route path={AppRoutes.Login} element={<LoginPage />} />
            <Route path={AppRoutes.Home} element={<HomeContainer />} />
      </Routes>
    </>
  );
}

export default App;
