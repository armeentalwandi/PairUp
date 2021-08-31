import React from 'react';
import MapWithADirectionsRenderer from './Homepage/Map';
import GoogleMaps from './Homepage/startingpoint';
import GoogleMaps1 from './Homepage/destination';
import {useState} from 'react';

function Home() {
    const [startlat, setStartlat] = useState(); 
    const [endlat, setEndlat] = useState(); // 3 more 
    const [startlong, setStartlong] = useState(); 
    const [endlong, setEndlong] = useState(); 

    return (
        <>
        console.log("hello"); 
        <GoogleMaps setStartlat = {setStartlat} setStartlong = {setStartlong}/>
        <GoogleMaps1 setEndlat = {setEndlat} setEndlong = {setEndlong}/>

      {/* <Button startlat = {startlat} startlong = {startlong} 
         endlat={endlat} endlong={endlong} onClick=}> Find Buddies</Button> */}

         <MapWithADirectionsRenderer startlat = {startlat} startlong = {startlong} 
         endlat={endlat} endlong={endlong}/>
     </>
    );
   
}

export default Home;