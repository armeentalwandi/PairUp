import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { withStyles } from '@material-ui/core/styles'
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Rating from '@material-ui/lab/Rating'
import FiberManualRecordIcon from '@material-ui/icons/FiberManualRecord'
import { ThemeProvider, createTheme} from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';
import { CardHeader } from '@material-ui/core';

const theme = createTheme({
  typography: {
    fontFamily: [
      'Poppins',
    ]
  },

});

const useStyles = makeStyles((theme) => ({
  root: {
    minWidth: 180,
    minHeight: 300,
    backgroundColor: "#F9F9F9",
  },
  title: {
    fontSize: 27,
  },
  pos: {
    marginBottom: 12,
  },
  box: {
    alignItems: 'right',
  },
  large: {
    width: theme.spacing(7),
    height: theme.spacing(7),
  },
}));

const StyledRating = withStyles({
    iconFilled: {
        color: '#f6c6ea',
    },
})(Rating);

const StyledRating2 = withStyles({
    iconFilled: {
        color: '#CDF0EA',
    },
})(Rating);

export default function OutlinedCard({ props }) {
  const classes = useStyles();
  
  
  return (

    <ThemeProvider theme={theme}>
      <Card className={classes.root} variant="outlined">
      <CardHeader
        avatar={
          <Avatar className={classes.large}>N</Avatar>
        }
        titleTypographyProps={{variant:'h6' , fontSize: 20}}
        title="Natalie P."
        subheader="UW"
      />
      <CardContent>
        <Typography className={classes.pos} color="textSecondary">
          0.2 miles
          <br></br>
          2 mins
        </Typography>
        <StyledRating name="read-only" value={2} precision={0.5} icon={<FiberManualRecordIcon fontSize="/inherit"/>} readOnly/>
        <br></br>
        <StyledRating2 name="read-only" value={3} precision={0.5} icon={<FiberManualRecordIcon fontSize="/inherit"/>} readOnly/>
        
      </CardContent>


      <CardActions style={{justifyContent: 'center'}}>
        <Button variant="outlined" size="medium">Request</Button>
      </CardActions> 
    </Card>

    </ThemeProvider>
    
  );
}
