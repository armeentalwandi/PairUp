import React from 'react';
import './index.css';
import Results from './pages/Results';
import { BrowserRouter as Router, Route, Switch} from 'react-router-dom';

function App() {
  return (
    <Router>
      <div className="App">
        <div className="content">
          <Switch>
            <Route path="/results">
              <Results/>
            </Route>
          </Switch>
        </div>
      </div>
    </Router>
  )

}

export default App;
