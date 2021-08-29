import './App.css';
import SideNav from './components/SideNav/SideNav';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './pages/Home';
import React from 'react';
import Invites from './pages/Invites';
import Sent from './pages/Sent';

function App() {
  return (
    <>
      <Router>
        <SideNav />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/invites' exact component={Invites} />
          <Route path='/sent' exact component={Sent} />
        </Switch>
      </Router>
    </>
  );
}

export default App;