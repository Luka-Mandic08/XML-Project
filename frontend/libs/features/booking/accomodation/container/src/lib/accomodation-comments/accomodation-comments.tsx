import { useEffect, useState } from 'react';
import styles from './accomodation-comments.module.css';
import { Paper, Grid, Typography, CardHeader, CardContent, Avatar, Card, CardActions, IconButton, Rating, TextField } from '@mui/material';
import { useSelectedAccommodationStore } from '@frontend/features/booking/store/container';
import {
  DeleteAccommodationRating,
  DeleteHostRating,
  EditAccommodationRating,
  EditHostRating,
  GetAccomodationRatings,
  GetHostRatings,
} from '@frontend/features/booking/accomodation/data';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import PersonIcon from '@mui/icons-material/Person';

export interface AccomodationCommentsProps {
  showAccommodationComments: boolean;
  showHostComments: boolean;
}

interface AccommodationRating {
  id: string;
  guestId: string;
  accommodationId?: string | null;
  hostId?: string | null;
  score: number;
  comment: string;
  date: any;
}

export function AccomodationComments(props: AccomodationCommentsProps) {
  const selectedAccommodation = useSelectedAccommodationStore((state) => state.selectedAccommodation);
  const [comments, setComments] = useState<AccommodationRating[]>([]);
  const [userComments, setUserComments] = useState<AccommodationRating[]>([]);
  const userRole = localStorage.getItem('role');
  const userID = localStorage.getItem('userId');
  const [editedComment, setEditedComment] = useState('');
  const [editedIndex, setEditedIndex] = useState(0);
  const [editedRating, setEditedRating] = useState(0);

  useEffect(() => {
    if (props.showAccommodationComments) {
      GetAccomodationRatings(selectedAccommodation.id).then((data) => {
        if (userRole === 'Host') {
          setComments(data);
        } else {
          setUserComments(data?.filter((comment: any) => comment.guestId === userID));
          setComments(data?.filter((comment: any) => comment.guestId !== userID));
        }
      });
    }
    if (props.showHostComments) {
      GetHostRatings(selectedAccommodation.hostId).then((data) => {
        if (userRole === 'Host') {
          setComments(data);
        } else {
          setUserComments(data?.filter((comment: any) => comment.guestId === userID));
          setComments(data?.filter((comment: any) => comment.guestId !== userID));
        }
      });
    }
  }, []);

  function deleteComment(comment: any) {
    if (props.showAccommodationComments) {
      DeleteAccommodationRating(comment.id, comment.guestId).then(() => {
        const updatedUserComments = userComments.filter((com) => com.id !== comment.id);
        setUserComments(updatedUserComments);
      });
    } else {
      DeleteHostRating(comment.id, comment.guestId).then(() => {
        const updatedUserComments = userComments.filter((com) => com.id !== comment.id);
        setUserComments(updatedUserComments);
      });
    }
  }

  const handleCommentEdit = (comment: AccommodationRating, editedComment: string) => {
    comment.comment = editedComment;
    setEditedComment(editedComment);
  };

  const handleRatingEdit = (comment: AccommodationRating, editedScore: number) => {
    comment.score = editedScore;
    setEditedRating(editedScore);
  };

  function editComment(comment: AccommodationRating) {
    if (props.showAccommodationComments) {
      EditAccommodationRating(comment);
    } else {
      EditHostRating(comment);
    }
  }

  return (
    <Paper elevation={6} sx={{ padding: '1.5rem 2rem 1.5rem 2rem', margin: '2rem' }}>
      <Grid container justifyContent={'start'}>
        <Grid container justifyContent={'center'} direction={'column'}>
          <Typography variant="h4" align="center" fontWeight={550}>
            {props.showAccommodationComments && 'Accommodation comments'}
            {props.showHostComments && 'Host comments'}
          </Typography>
        </Grid>
        <Grid item direction={'row'} xs={12} marginTop={'1.25rem'}>
          {userComments?.map((comment, index) => (
            <Grid item xs={12} key={index} marginTop={'5px'}>
              <Card className={styles.card} sx={{ border: '0.5px solid #999', width: '100%' }}>
                <CardContent style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <div style={{ flex: 1 }}>
                    <div style={{ display: 'flex', alignItems: 'center' }}>
                      <CardHeader
                        avatar={
                          <Avatar>
                            <PersonIcon />
                          </Avatar>
                        }
                        title={'Guest'}
                        subheader={`Date: ${new Date(comment.date.seconds * 1000).toDateString()}`}
                      />
                    </div>
                    <TextField
                      label="Comment"
                      variant="outlined"
                      sx={{ margin: '3%', width: '90%' }}
                      value={comment.comment}
                      onChange={(event) => {
                        const editedComment = event.target.value;
                        handleCommentEdit(comment, editedComment);
                      }}
                    />
                    <TextField
                      label="Score"
                      variant="outlined"
                      sx={{ margin: '3%', width: '60%', marginTop: '0' }}
                      value={comment.score}
                      onChange={(event) => {
                        const editedRating = event.target.value;
                        handleRatingEdit(comment, Number(editedRating));
                      }}
                      inputProps={{
                        type: 'number',
                        min: 0,
                        max: 5,
                        step: 0.1,
                      }}
                    />
                    <Rating
                      value={comment.score}
                      onChange={(event, newValue) => {
                        handleRatingEdit(comment, Number(newValue));
                      }}
                      precision={1}
                      readOnly={false}
                      size="small"
                      sx={{ marginX: '20px', marginTop: '3%' }}
                    />
                  </div>
                  <div>
                    <CardActions>
                      {userRole === 'Guest' && userID === comment.guestId && (
                        <>
                          <IconButton aria-label="Edit" onClick={() => editComment(comment)}>
                            <EditIcon />
                          </IconButton>
                          <IconButton aria-label="Delete" onClick={() => deleteComment(comment)}>
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
          {comments?.map((comment, index) => (
            <Grid item xs={12} key={index} marginTop={'5px'}>
              <Card className={styles.card} sx={{ border: '0.5px solid #999', width: '100%' }}>
                <CardContent style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <div style={{ flex: 1 }}>
                    <div style={{ display: 'flex', alignItems: 'center' }}>
                      <CardHeader
                        avatar={
                          <Avatar>
                            <PersonIcon />
                          </Avatar>
                        }
                        title={'Guest'}
                        subheader={`Date: ${new Date(comment.date.seconds * 1000).toDateString()}`}
                      />
                      <Rating value={comment.score} precision={0.1} readOnly size="small" sx={{ marginX: '20px' }} />
                    </div>
                    <Typography variant="body1" align="left" sx={{ marginLeft: '3%' }}>
                      {comment.comment}
                    </Typography>
                  </div>
                </CardContent>
              </Card>
            </Grid>
          ))}
          {comments?.length === undefined && userComments?.length === undefined ? (
            <Typography variant="h6" align="center" sx={{ marginLeft: '3%', marginTop: '5%', width: '100%' }}>
              There are no comments yet
            </Typography>
          ) : null}
        </Grid>
      </Grid>
    </Paper>
  );
}

export default AccomodationComments;
