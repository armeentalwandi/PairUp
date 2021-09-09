import "./App.css";
import SideNav from "./components/SideNav/SideNav";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import React from "react";
import { useState } from "react";
import GoogleMaps from "./pages/Homepage/startingpoint";
import GoogleMaps1 from "./pages/Homepage/destination";
import MapWithADirectionsRenderer from "./pages/Homepage/Map";
import { Button } from "@material-ui/core";
import "./index.css";
import Results from "./pages/Results";
import Invites from "./pages/Invites";
import Sent from "./pages/Sent";
import axios from "axios";

function App() {
  const [startlat, setStartlat] = useState();
  const [endlat, setEndlat] = useState(); // 3 more
  const [startlong, setStartlong] = useState();
  const [endlong, setEndlong] = useState();

  function click() {
    axios
      .put("http://localhost:10000/location", {
        Id: "9",
        Startlat: startlat,
        StartLong: startlong,
        EndLat: endlat,
        EndLong: endlong,
      })
      .then(function (response) {
        console.log("here-now-really-really");
        console.log(response);
      });
  }

  return (
    <Router>
      <div className="App">
        <header className="content">
          <SideNav></SideNav>
          <Switch>
            <Route exact path="/">
              <GoogleMaps
                setStartlat={setStartlat}
                setStartlong={setStartlong}
              />
              <GoogleMaps1 setEndlat={setEndlat} setEndlong={setEndlong} />

              <Button
                startlat={startlat}
                startlong={startlong}
                endlat={endlat}
                endlong={endlong}
                onClick={click}
              >
                {" "}
                Find Buddies
              </Button>

              <MapWithADirectionsRenderer
                startlat={startlat}
                startlong={startlong}
                endlat={endlat}
                endlong={endlong}
              />
            </Route>
            <Route exact path="/results">
              {" "}
              <Results />{" "}
            </Route>
            <Route path="/invites" exact component={Invites} />
            <Route path="/sent" exact component={Sent} />
          </Switch>
        </header>
      </div>
    </Router>
  );
}

export default App;
