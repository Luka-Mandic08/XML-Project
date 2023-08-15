import { AccommodationCreateUpdateDTO, BookingAppRoutes, BookingBaseURL } from '@frontend/models';
import styles from './create-update-accommodation.module.css';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { Grid, Paper, Typography, IconButton, Button } from '@mui/material';
import HighlightOffIcon from '@mui/icons-material/HighlightOff';
import { CreateUpdateAccommodationFunction } from '@frontend/features/booking/accomodation/data';

/* eslint-disable-next-line */
export interface CreateUpdateAccommodationProps {}

export function CreateUpdateAccommodation(props: CreateUpdateAccommodationProps) {
  const [amenities, setAmenities] = useState<string[]>([]);
  const [images, setImages] = useState<File[]>([]);
  const [imagesUrl, setImagesUrl] = useState<string[]>([]);
  const [accomodationDTO, setAccomodationDTO] = useState<AccommodationCreateUpdateDTO>({
    id: '',
    name: '',
    address: {
      street: '',
      city: '',
      country: '',
    },
    minGuests: 0,
    maxGuests: 0,
    priceIsPerGuest: false,
    hasAutomaticReservations: false,
    hostId: localStorage.getItem('userId')!,
  });

  const navigate = useNavigate();

  const {
    register: registerAmenities,
    handleSubmit: handleSubmitAmenities,
    watch: watchAmenities,
    reset: resetAmenities,
    formState: { errors: errorsAmenities },
  } = useForm({
    defaultValues: {
      amenity: '',
    },
  });

  const onSubmitAmenities = (data: any) => {
    amenities.push(data.amenity);
    setAmenities(amenities);
    resetAmenities();
  };

  const deleteAmenities = (id: number) => {
    amenities.splice(id, 1);
    setAmenities(amenities);
    resetAmenities();
  };

  const {
    register: registerImages,
    handleSubmit: handleSubmitImages,
    watch: watchImages,
    reset: resetImages,
    formState: { errors: errorsImages },
  } = useForm({
    defaultValues: {
      image: File,
    },
  });

  const onSubmitImages = (data: any) => {
    images.push(data);
    console.log(images);
    setImages(images);
    resetImages();
  };

  const deleteImages = (id: number) => {
    images.splice(id, 1);
    setImages(images);
    resetImages();
  };

  const {
    register,
    handleSubmit,
    watch,
    setValue,
    formState: { errors },
  } = useForm({
    defaultValues: {
      id: accomodationDTO.id,
      name: accomodationDTO.name,
      address: {
        street: accomodationDTO.address.street,
        city: accomodationDTO.address.city,
        country: accomodationDTO.address.country,
      },
      minGuests: accomodationDTO.minGuests,
      maxGuests: accomodationDTO.maxGuests,
      priceIsPerGuest: accomodationDTO.priceIsPerGuest,
      hasAutomaticReservations: accomodationDTO.hasAutomaticReservations,
      hostId: accomodationDTO.hostId,
    },
  });

  const onSubmit = (data: AccommodationCreateUpdateDTO) => {
    CreateUpdateAccommodationFunction(data, amenities, images);
    navigate(BookingAppRoutes.HomeHost);
  };

  return (
    <Paper elevation={3} sx={{ width: 'min(90%,1200px)', margin: '1rem auto', paddingX: '2.5rem', paddingY: '1rem' }}>
      <Grid container justifyContent={'center'} marginY={'1rem'} alignItems={'center'}>
        <Grid item paddingX={'1rem'}>
          <Typography variant="h2" align="center">
            Create Accommodation
          </Typography>
        </Grid>
      </Grid>

      <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
        <Grid item marginBottom={'1rem'}>
          <Typography variant="h4" align="left">
            Amenities
          </Typography>
        </Grid>
        <div className={styles.amenitiesContainer}>
          {amenities?.map((amenity, idx) => (
            <div className={styles.amenityCard}>
              <Typography>
                {idx + 1}. {amenity}
              </Typography>
              <IconButton onClick={() => deleteAmenities(idx)}>
                <HighlightOffIcon></HighlightOffIcon>
              </IconButton>
            </div>
          ))}
        </div>
      </Grid>
      <form onSubmit={handleSubmitAmenities(onSubmitAmenities)}>
        <div className={styles.lineContainer}>
          <div className={styles.inputContainer}>
            <input
              type="text"
              id="amenity"
              value={watchAmenities('amenity')}
              {...registerAmenities('amenity', {
                required: 'This field is required.',
              })}
            />
            <label className={styles.label} htmlFor="amenity" id="label-amenity">
              <div className={styles.text}>Amenity</div>
            </label>
            <label className={styles.errorLabel}>{errorsAmenities.amenity?.message}</label>
          </div>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Add Amenity
          </Button>
        </div>
      </form>

      <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
        <Grid item marginBottom={'1rem'}>
          <Typography variant="h4" align="left">
            Images
          </Typography>
        </Grid>
        <div className={styles.amenitiesContainer}>
          {imagesUrl?.map((image, idx) => (
            <div className={styles.amenityCard}>
              <Typography>Image {idx + 1}.</Typography>
              <IconButton onClick={() => deleteImages(idx)}>
                <HighlightOffIcon></HighlightOffIcon>
              </IconButton>
              <img src={image} alt="accommodation" />
            </div>
          ))}
        </div>
      </Grid>
      <form onSubmit={handleSubmitImages(onSubmitImages)}>
        <div className={styles.lineContainer}>
          <div className={styles.inputContainer}>
            <input
              type="file"
              id="image"
              {...registerImages('image', {
                required: 'This field is required.',
              })}
            />
            <label className={styles.label} htmlFor="image" id="label-image">
              <div className={styles.text}>Image</div>
            </label>
            <label className={styles.errorLabel}>{errorsImages.image?.message}</label>
          </div>
          <Button
            variant="contained"
            size="large"
            type="submit"
            sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
          >
            Add Image
          </Button>
        </div>
      </form>

      <Grid container marginY={'1rem'} alignItems={'left'} direction={'column'}>
        <Typography variant="h4" align="left">
          Accommodation information
        </Typography>
      </Grid>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className={styles.inputContainer}>
          <input
            type="text"
            id="name"
            value={watch('name')}
            {...register('name', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="name" id="label-name">
            <div className={styles.text}>Name</div>
          </label>
          <label className={styles.errorLabel}>{errors.name?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input type="text" id="address.street" value={watch('address.street')} {...register('address.street', { required: 'This field is required.' })} />
          <label className={styles.label} htmlFor="address.street" id="label-address.street">
            <div className={styles.text}>Street</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.street?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="address.city"
            value={watch('address.city')}
            {...register('address.city', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="address.city" id="label-praddress.city">
            <div className={styles.text}>City</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.city?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="text"
            id="address.country"
            value={watch('address.country')}
            {...register('address.country', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="address.country" id="label-address.country">
            <div className={styles.text}>Country</div>
          </label>
          <label className={styles.errorLabel}>{errors.address?.country?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="minGuests"
            value={watch('minGuests')}
            {...register('minGuests', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="minGuests" id="label-minGuests">
            <div className={styles.text}>Minimum number of guests</div>
          </label>
          <label className={styles.errorLabel}>{errors.minGuests?.message}</label>
        </div>

        <div className={styles.inputContainer}>
          <input
            type="number"
            id="maxGuests"
            value={watch('maxGuests')}
            {...register('maxGuests', {
              required: 'This field is required.',
            })}
          />
          <label className={styles.label} htmlFor="maxGuests" id="label-maxGuests">
            <div className={styles.text}>Maximum number of guests</div>
          </label>
          <label className={styles.errorLabel}>{errors.maxGuests?.message}</label>
        </div>
        <div className={styles.lineContainer}>
          <div className={styles.lineContainer} style={{ gap: '1rem' }}>
            <div>
              <Typography variant="h6">Price is per guest</Typography>
              <Typography variant="h6" color={'red'}>
                {errors.priceIsPerGuest?.message}
              </Typography>
            </div>
            <input style={{ width: '48px', height: '48px' }} type="checkbox" id="priceIsPerGuest" {...register('priceIsPerGuest')} />
          </div>
          <div className={styles.lineContainer} style={{ gap: '1rem' }}>
            <div>
              <Typography variant="h6">Has automatic reservations</Typography>
              <Typography variant="h6" color={'red'}>
                {errors.hasAutomaticReservations?.message}
              </Typography>
            </div>
            <input style={{ width: '48px', height: '48px' }} type="checkbox" id="hasAutomaticReservations" {...register('hasAutomaticReservations')} />
          </div>
        </div>

        <Button
          variant="contained"
          size="large"
          type="submit"
          sx={{ color: 'white', background: '#212121', height: '48px', width: '248px', ':hover': { background: 'white', color: '#212121' } }}
        >
          Create Accommodation
        </Button>
      </form>
    </Paper>
  );
}

export default CreateUpdateAccommodation;
