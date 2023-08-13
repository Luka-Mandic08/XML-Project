import { UpdatePersonalData, UpdateCredentials } from '@frontend/models';
import { Button, Grid, Typography } from '@mui/material';
import { set, useForm } from 'react-hook-form';
import styles from './profile-info.module.css';
import { useState, useEffect } from 'react';
import {
  DeleteAccount,
  GetAccountInformation,
  GetProfileInformation,
  UpdateAccountInformation,
  UpdateProfileInformation,
} from '@frontend/features/booking/profiles/data-access';

/* eslint-disable-next-line */
export interface ProfileInfoProps {}

export function ProfileInfo(props: ProfileInfoProps) {
  const [userInfo, setUserInfo] = useState<UpdatePersonalData>({
    name: '',
    surname: '',
    email: '',
    address: {
      street: '',
      city: '',
      country: '',
    },
  });
  const [accountInfo, setAccountInfo] = useState<UpdateCredentials>({
    username: '',
    password: '',
  });

  const [isDisabled, setIsDisabled] = useState<boolean>(true);

  useEffect(() => {
    GetAccountInformation().then((data) => {
      setAccountInfo(data);
    });

    GetProfileInformation().then((data) => {
      setUserInfo(data);
    });
  }, []);

  useEffect(() => {
    resetProfile(userInfo);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [userInfo]);

  useEffect(() => {
    resetAccount({
      username: accountInfo.username,
      password: '',
      confirmPassword: '',
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [accountInfo]);

  const {
    register: registerProfile,
    handleSubmit: handleSubmitProfile,
    watch: watchProfile,
    reset: resetProfile,
    formState: { errors: errorsProfile },
  } = useForm({
    defaultValues: {
      name: '',
      surname: '',
      email: '',
      address: {
        street: '',
        city: '',
        country: '',
      },
    },
  });

  const {
    register: registerAccount,
    getValues: getValuesAccount,
    handleSubmit: handleSubmitAccount,
    watch: watchAccount,
    reset: resetAccount,
    formState: { errors: errorsAccount },
  } = useForm({
    defaultValues: {
      username: '',
      password: '',
      confirmPassword: '',
    },
  });

  const onSubmitProfile = async (data: UpdatePersonalData) => {
    setUserInfo(await UpdateProfileInformation(data));
  };
  const onSubmitAccount = async (data: UpdateCredentials) => {
    const res: any = await UpdateAccountInformation(data);
    const updatedAccountInfo: UpdateCredentials = {
      username: res.username,
      password: '',
    };
    resetAccount(updatedAccountInfo);
    setAccountInfo(updatedAccountInfo);
  };

  const deleteAccount = async () => {
    await DeleteAccount();
    localStorage.clear();
    window.location.href = '/';
  };

  return (
    <Grid container marginY={'2rem'} justifyContent={'center'}>
      <Grid container direction={'row'} justifyContent={'space-between'} xs={10}>
        <Grid item sx={{ marginBottom: '1rem' }}>
          <Typography variant="h4">{accountInfo.username}'s Profile</Typography>
        </Grid>
        <div className={styles.buttonTopContainer}>
          <Grid item sx={{ marginBottom: '1rem' }}>
            <Button
              variant="contained"
              size="large"
              onClick={() => {
                setIsDisabled(!isDisabled);
              }}
              sx={{ color: 'white', background: '#212121', ':hover': { background: 'white', color: '#212121' } }}
            >
              {isDisabled ? 'Enable editing' : 'Disable editing'}
            </Button>
          </Grid>
          <Grid item sx={{ marginBottom: '1rem' }}>
            <Button
              variant="contained"
              size="large"
              onClick={() => {
                deleteAccount();
              }}
              sx={{ color: 'white', background: 'red', ':hover': { background: 'white', color: '#212121' } }}
            >
              Delete account
            </Button>
          </Grid>
        </div>
      </Grid>
      <Grid item xs={12} md={5} className={styles.formContainer}>
        <form onSubmit={handleSubmitProfile(onSubmitProfile)}>
          <fieldset disabled={isDisabled}>
            <Typography variant="h6" marginBottom={'1rem'}>
              Peronal information
            </Typography>
            <div className={styles.inputContainer}>
              <input
                type="text"
                id="name"
                value={watchProfile('name')}
                {...registerProfile('name', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="name" id="label-name">
                <div className={styles.text}>Name</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.name?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="text"
                id="surname"
                value={watchProfile('surname')}
                {...registerProfile('surname', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="surname" id="label-surname">
                <div className={styles.text}>Surname</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.surname?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="email"
                id="email"
                value={watchProfile('email')}
                {...registerProfile('email', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="email" id="label-email">
                <div className={styles.text}>Email</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.email?.message}</label>
            </div>
            <Typography variant="h6" marginBottom={'1rem'}>
              Address information
            </Typography>
            <div className={styles.inputContainer}>
              <input
                type="text"
                id="address.street"
                value={watchProfile('address.street')}
                {...registerProfile('address.street', { required: 'This field is required.' })}
              />
              <label className={styles.label} htmlFor="address.street" id="label-address.street">
                <div className={styles.text}>Street</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.address?.street?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="text"
                id="address.city"
                value={watchProfile('address.city')}
                {...registerProfile('address.city', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="address.city" id="label-praddress.city">
                <div className={styles.text}>City</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.address?.city?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="text"
                id="address.country"
                value={watchProfile('address.country')}
                {...registerProfile('address.country', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="address.country" id="label-address.country">
                <div className={styles.text}>Country</div>
              </label>
              <label className={styles.errorLabel}>{errorsProfile.address?.country?.message}</label>
            </div>
            <input style={{ width: '50%', marginLeft: '25%', marginRight: '25%' }} type="submit" value="Update Personal Information" />
          </fieldset>
        </form>
      </Grid>
      <Grid item xs={12} md={5} className={styles.formContainer}>
        <form onSubmit={handleSubmitAccount(onSubmitAccount)}>
          <fieldset disabled={isDisabled}>
            <Typography variant="h6" marginBottom={'1rem'}>
              Account information
            </Typography>
            <div className={styles.inputContainer}>
              <input
                type="text"
                id="username"
                value={watchAccount('username')}
                {...registerAccount('username', {
                  required: 'This field is required.',
                })}
              />
              <label className={styles.label} htmlFor="username" id="label-username">
                <div className={styles.text}>Username</div>
              </label>
              <label className={styles.errorLabel}>{errorsAccount.username?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input type="password" id="password" value={watchAccount('password')} {...registerAccount('password', {})} />
              <label className={styles.label} htmlFor="password" id="label-password">
                <div className={styles.text}>New password</div>
              </label>
              <label className={styles.errorLabel}>{errorsAccount.password?.message}</label>
            </div>

            <div className={styles.inputContainer}>
              <input
                type="password"
                id="confirmPassword"
                value={watchAccount('confirmPassword')}
                {...registerAccount('confirmPassword', {
                  validate: {
                    isSameAsPassword: (v) => v === getValuesAccount('password') || 'Passwords do not match',
                  },
                })}
              />
              <label className={styles.label} htmlFor="confirmPassword" id="label-confirmPassword">
                <div className={styles.text}>Confirm password</div>
              </label>
              <label className={styles.errorLabel}>{errorsAccount.confirmPassword?.message}</label>
            </div>

            <input style={{ width: '50%', marginLeft: '25%', marginRight: '25%' }} type="submit" value="Update Credentials" />
          </fieldset>
        </form>
      </Grid>
    </Grid>
  );
}

export default ProfileInfo;
