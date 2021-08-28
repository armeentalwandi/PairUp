import '../App.css'
import OutlinedCard from '../components/Card'
import Grid from '@material-ui/core/Grid'
import { makeStyles } from '@material-ui/core';



const useStyles = makeStyles({
  GridContainer: {
    paddingLeft: "20px",
    paddingRight: "20px",
    paddingTop: "20px",
  }

})

function Results() {
  const classes = useStyles();
  return (
  <Grid container spacing={2} className={classes.GridContainer} justify="center">
    <Grid item xs={12} sm={6} md={3}>
      <OutlinedCard/>
    </Grid>
    <Grid item xs={12} sm={6} md={3}>
      <OutlinedCard/>
    </Grid>
    <Grid item xs={12} sm={6} md={3}>
      <OutlinedCard/>
    </Grid>
    <Grid item xs={12} sm={6} md={3}>
      <OutlinedCard/>
    </Grid>
  </Grid>
  );
}

export default Results;