
import './App.css';

import {Link} from 'react-router-dom';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'; 
import MapWithADirectionsRenderer from './Map';
import GoogleMaps from './startingpoint';
import GoogleMaps1 from './destination';

//import { Button } from '@material-ui/core';
 
import { useState } from 'react';



function App() {
//  const [startlat, setStartlat] = useState(43.4723); 
//  const [endlat, setEndlat] = useState(43.4739); // 3 more 
//  const [startlong, setStartlong] = useState(-80.5449); 
//  const [endlong, setEndlong] = useState(-80.5274);

 const [startlat, setStartlat] = useState(); 
 const [endlat, setEndlat] = useState(); // 3 more 
 const [startlong, setStartlong] = useState(); 
 const [endlong, setEndlong] = useState(); 
  return (
 
    <Router>
    <div className="App">
     
    
      <p><Link to="/homepage">Get Directions</Link></p> 
      <header className="content">
        <Switch>
         <Route path= "/homepage">
         <GoogleMaps setStartlat = {setStartlat} setStartlong = {setStartlong}/>
         <GoogleMaps1 setEndlat = {setEndlat} setEndlong = {setEndlong}/>

       {/* <Button startlat = {startlat} startlong = {startlong} 
          endlat={endlat} endlong={endlong} onClick=}> Find Buddies</Button> */}

          <MapWithADirectionsRenderer startlat = {startlat} startlong = {startlong} 
          endlat={endlat} endlong={endlong}/>
           
          </Route>
          </Switch>
        </header>
      </div> 

    </Router>
   
  );
}

export default App;
