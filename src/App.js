
import './App.css';
import SideNav from './components/SideNav/SideNav';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import React from 'react';
import Home from './pages/Home';
import {useState} from 'react';
import GoogleMaps from './pages/Homepage/startingpoint';
import GoogleMaps1 from './pages/Homepage/destination';
import MapWithADirectionsRenderer from './pages/Homepage/Map';


//import { Button } from '@material-ui/core';




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
     
    
      
      <header className="content">
        <SideNav></SideNav>
        <Switch>
         <Route path= "/" exact component={Home}>
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