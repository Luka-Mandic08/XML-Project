import { useEffect, useState } from 'react';
import styles from './accomodation-comments.module.css';
import { Paper, Grid, Typography, CardHeader, CardContent, Avatar, Card, CardActions, IconButton } from '@mui/material';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import { GetAccomodationRatings } from '@frontend/features/booking/accomodation/data';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import PersonIcon from '@mui/icons-material/Person';

/* eslint-disable-next-line */
export interface AccomodationCommentsProps {}

interface AccommodationRating {
  id: string;
  guestId: string;
  accommodationId: string;
  score: number;
  comment: string;
  date: Date;
}

const commentsData: AccommodationRating[] = [
  {
    id: '1',
    guestId: '64d7c0c95f1270c79a9d96ca',
    accommodationId: 'House1',
    score: 5,
    comment: 'Great place to stay!',
    date: new Date('2023-08-26'),
  },
  {
    id: '2',
    guestId: 'guest2',
    accommodationId: 'House1',
    score: 4,
    comment: 'Nice view from the room.',
    date: new Date('2023-08-25'),
  },
];

export function AccomodationComments(props: AccomodationCommentsProps) {
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);
  const [comments, setComments] = useState<any[]>([]);
  const userRole = localStorage.getItem('role');
  const userID = localStorage.getItem('userId');

  useEffect(() => {
    GetAccomodationRatings(selectedAccommodation.id).then((data) => {
      setComments(commentsData);
    });
    console.log(userRole);
  }, []);

  return (
    <Paper elevation={6} sx={{ padding: '1.5rem 2rem 1.5rem 2rem', marginTop: '1.5rem' }}>
      <Grid container justifyContent={'start'}>
        <Grid container justifyContent={'center'} direction={'column'}>
          <Typography variant="h3" align="center" fontWeight={550}>
            Comments
          </Typography>
        </Grid>
        <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
          {comments?.map((comment, index) => (
            <Grid item xs={12} key={index} marginTop={'5px'}>
              <Card className={styles.card} sx={{ border: '0.5px solid #999', width: '100%' }}>
                <CardContent style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <div style={{ flex: 1 }}>
                    <CardHeader
                      avatar={
                        <Avatar>
                          <PersonIcon />
                        </Avatar>
                      }
                      title={'Guest'}
                      subheader={`Date: ${comment.date.toDateString()}`}
                    />
                    <Typography variant="body1" align="left">
                      {comment.comment}
                    </Typography>
                  </div>
                  <div>
                    <CardActions>
                      {userRole === 'Guest' && userID === comment.guestId && (
                        <>
                          <IconButton aria-label="Edit">
                            <EditIcon />
                          </IconButton>
                          <IconButton aria-label="Delete">
                            <DeleteIcon />
                          </IconButton>
                        </>
                      )}
                    </CardActions>
                  </div>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      </Grid>
    </Paper>
  );
}

export default AccomodationComments;
