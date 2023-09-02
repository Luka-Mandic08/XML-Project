import { BookingAppRoutes, SearchedAccommodationInfo } from '@frontend/models';
import styles from './make-reservation-dialog.module.css';
import { useForm } from 'react-hook-form';
import { useSearchParametersStore } from '@frontend/features/booking/store/container';
import { Dialog, DialogTitle, DialogContent, Button, Typography } from '@mui/material';
import { MakeReservationFunction } from '@frontend/features/booking/reservation/data-access';
import { useNavigate } from 'react-router-dom';

/* eslint-disable-next-line */
export interface MakeReservationDialogProps {
  open: boolean;
  selectedAccommodation: SearchedAccommodationInfo;
  onClose: () => void;
}

export function MakeReservationDialog(props: MakeReservationDialogProps) {
  const { open, selectedAccommodation, onClose } = props;
  const searchParameters = useSearchParametersStore((state) => state.searchParameters);

  const navigate = useNavigate();

  const {
    handleSubmit,
    formState: { errors },
  } = useForm({
    defaultValues: {
      startDate: searchParameters.dateFrom,
      endDate: searchParameters.dateTo,
      numberOfGuests: searchParameters.numberOfGuests,
      userId: localStorage.getItem('userId'),
      accommodationId: selectedAccommodation.id,
    },
  });

  const onSubmit = async (data: any) => {
    data.startDate = new Date(data.startDate);
    data.endDate = new Date(data.endDate);
    const res = await MakeReservationFunction(data);
    if (res) {
      navigate(BookingAppRoutes.GuestReservations);
    }
    onClose();
  };

  const handleClose = () => {
    onClose();
  };

  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle variant="h4">Make a reservation</DialogTitle>
      <DialogContent className={styles.dialogContentContainer}>
        <div className={styles.infoContainer}>
          <Typography variant="h5" align="left">
            Accommodation: {selectedAccommodation.name}
          </Typography>
          <Typography variant="h5" align="left">
            Check in date: {searchParameters.dateFrom}
          </Typography>
          <Typography variant="h5" align="left">
            Check out date: {searchParameters.dateTo}
          </Typography>
          <Typography variant="h5" align="left">
            Number of guests: {searchParameters.numberOfGuests}
          </Typography>
          <Typography variant="h5" align="left">
            Price per night: {selectedAccommodation.unitPrice}
          </Typography>
          <Typography variant="h5" align="left">
            Total price: {selectedAccommodation.totalPrice}
          </Typography>
        </div>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className={styles.lineContainer}>
            <Button
              variant="contained"
              size="small"
              onClick={handleClose}
              sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
            >
              Cancel
            </Button>
            <Button
              variant="contained"
              size="small"
              type="submit"
              sx={{ color: 'white', background: '#212121', height: '48px', minWidth: '200px', ':hover': { background: 'white', color: '#212121' } }}
            >
              Reserve
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default MakeReservationDialog;
