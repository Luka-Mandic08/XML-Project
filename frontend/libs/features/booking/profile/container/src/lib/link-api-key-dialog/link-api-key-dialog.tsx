import { Button, Dialog, DialogContent, DialogTitle, Divider, Typography } from '@mui/material';
import styles from './link-api-key-dialog.module.css';
import { ApiKey } from '@frontend/models';
import { useEffect, useState } from 'react';
import { CreateApiKey, GetApiKey, LinkToFlightsApp } from '@frontend/features/booking/profile/data-access';
import { useForm } from 'react-hook-form';

/* eslint-disable-next-line */
export interface LinkApiKeyDialogProps {
  open: boolean;
  onClose: () => void;
}

export function LinkApiKeyDialog(props: LinkApiKeyDialogProps) {
  const { open, onClose } = props;
  const [apiKey, setApiKey] = useState<ApiKey>();

  useEffect(() => {
    if (open && localStorage.getItem('role') === 'Guest') {
      getApiKey();
    }
  }, [open]);

  const getApiKey = async () => {
    const res = await GetApiKey();
    if (!res) {
      return;
    }
    const newApiKey = {
      apiKeyValue: res.apiKeyValue,
      validTo: new Date(res.validTo.seconds * 1000),
      isPermanent: res.isPermanent,
    };
    setApiKey(newApiKey);
  };

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    defaultValues: {
      username: '',
      password: '',
    },
  });

  const onSubmit = async (data: any) => {
    await LinkToFlightsApp(data.username, data.password, apiKey!);
    onClose();
  };

  const createPermanentApiKey = async () => {
    await CreateApiKey(true);
  };
  const createTemporaryApiKey = async () => {
    await CreateApiKey(false);
  };

  const handleClose = () => {
    onClose();
  };
  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle variant="h4">Connect to flights app</DialogTitle>
      <DialogContent className={styles.dialogContentContainer}>
        <div className={styles.infoContainer}>
          <Typography variant="h5" align="left">
            By connecting to flights app you will be able to purchase flight tickets from your reservation.
          </Typography>
          <Divider sx={{ backgroundColor: 'grey', width: '100%' }} />
          {apiKey && (
            <>
              <Typography variant="h5" align="left">
                You have an APIKey
              </Typography>
              {!apiKey.isPermanent && (
                <Typography variant="h5" align="left">
                  Valid to: {apiKey.validTo.toDateString()}
                </Typography>
              )}
              <Typography variant="h5" align="left">
                {apiKey.isPermanent ? 'Permanent' : 'Temporary'}
              </Typography>
              {apiKey.validTo < new Date() && (
                <div className={styles.lineContainer}>
                  <Button
                    variant="contained"
                    size="small"
                    onClick={createTemporaryApiKey}
                    sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
                  >
                    1 month
                  </Button>
                  <Button
                    variant="contained"
                    size="small"
                    onClick={createPermanentApiKey}
                    sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
                  >
                    Permanent
                  </Button>
                </div>
              )}
              {apiKey.validTo > new Date() && (
                <form onSubmit={handleSubmit(onSubmit)}>
                  <div className={styles.inputContainer}>
                    <input
                      type="text"
                      id="username"
                      value={watch('username')}
                      {...register('username', {
                        required: 'This field is required.',
                      })}
                    />
                    <label className={styles.label} htmlFor="username" id="label-username">
                      <div className={styles.text}>Flights App Username</div>
                    </label>
                    <label className={styles.errorLabel}>{errors.username?.message}</label>
                  </div>

                  <div className={styles.inputContainer}>
                    <input type="password" id="password" value={watch('password')} {...register('password', {})} />
                    <label className={styles.label} htmlFor="password" id="label-password">
                      <div className={styles.text}>Flights App Password</div>
                    </label>
                    <label className={styles.errorLabel}>{errors.password?.message}</label>
                  </div>

                  <Button
                    variant="contained"
                    size="small"
                    type="submit"
                    sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
                  >
                    Connect to flights app
                  </Button>
                </form>
              )}
            </>
          )}
          {!apiKey && (
            <>
              <Typography variant="h5" align="left">
                You do not have an APIKey. Create your APIKey to connect to flights app
              </Typography>
              <Typography variant="h6" align="left">
                It can be permanent or valid for 1 month
              </Typography>
              <div className={styles.lineContainer}>
                <Button
                  variant="contained"
                  size="small"
                  onClick={createTemporaryApiKey}
                  sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
                >
                  1 month
                </Button>
                <Button
                  variant="contained"
                  size="small"
                  onClick={createPermanentApiKey}
                  sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
                >
                  Permanent
                </Button>
              </div>
            </>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}

export default LinkApiKeyDialog;
