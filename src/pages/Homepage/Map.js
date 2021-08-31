

const { compose, withProps, lifecycle } = require("recompose");
const {
  withScriptjs,
  withGoogleMap,
  GoogleMap,
  DirectionsRenderer,
} = require("react-google-maps");


 const MapWithADirectionsRenderer = compose( 
  withProps({
    googleMapURL: "https://maps.googleapis.com/maps/api/js?key=AIzaSyDLF-yHHgLIQe8Z45WCTmMDm6Iwup_sHQs&v=3.exp&libraries=geometry,drawing,places",
    loadingElement: <div style={{ height: `100%` }} />,
    containerElement: <div style={{ height: `400px` }} />,
    mapElement: <div style={{ height: `100%`, width:'80%', marginLeft:'10%' }} />,
  }),
  withScriptjs,
  withGoogleMap,
  lifecycle({
    componentDidMount() {
      console.log(this.props.startlat)
      const DirectionsService = new window.google.maps.DirectionsService();
      
         
      DirectionsService.route({
      origin: new window.google.maps.LatLng(this.props.startlat, this.props.startlong), //starting point
      destination: new window.google.maps.LatLng(this.props.endlat, this.props.endlong), // destination
      travelMode: window.google.maps.TravelMode.WALKING,
   }, 
   
   (result, status) => {
     if (status === window.google.maps.DirectionsStatus.OK) {
       this.setState({
         directions: result,
       });
     } else {
       console.error(`error fetching directions ${result}`);
     }
   },
   
   );
    }, componentDidUpdate(prevProps, prevState) {
      

      if (prevProps.startlat !== this.props.startlat || prevProps.startlong !== this.props.startlong
        || prevProps.endlat !== this.props.endlat || prevProps.endlong !== this.props.endlong ) {
        console.log(this.props.startlat)
        const DirectionsService = new window.google.maps.DirectionsService();
      
        DirectionsService.route({
          origin: new window.google.maps.LatLng(this.props.startlat, this.props.startlong), //starting point
          destination: new window.google.maps.LatLng(this.props.endlat, this.props.endlong), // destination
          travelMode: window.google.maps.TravelMode.WALKING,
        },  
      
        (result, status) => {
          if (status === window.google.maps.DirectionsStatus.OK) {
            this.setState({
              directions: result,
            });
          } else {
            console.error(`error fetching directions ${result}`);
          }
        },
      );
      }
      
    }
 })
)(props =>
  <GoogleMap
    defaultZoom={7}
    defaultCenter={new window.google.maps.LatLng(41.8507300, -87.6512600)}
  >
    {props.directions && <DirectionsRenderer directions={props.directions} />}
  </GoogleMap>

);



export default MapWithADirectionsRenderer; 