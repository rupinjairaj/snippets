import React from 'react';
import { Route, Switch } from 'react-router-dom';

import Header from './Header';
import Sidebar from './Sidebar';
import SnippetView from './SnippetView';
import Tags from './Tags';

/**
 * The layout has been divided into 3 parts. The header which holds the 
 * app name (and the universal search bar later), the side navbar and the 
 * space which is the target of the route handler.
 */
function App() {
  return (
    <div>
      <Header />
      <div style={{ marginBottom: "2px" }}></div>
      <div style={{ marginLeft: "10px" }}>
        <div className="row">
          <div className="col-2 bg-light">
            <Sidebar />
          </div >
          <div className="col-10">
            <Switch>
              <Route path="/" component={Tags} exact />
              <Route path="/tags" component={Tags} />
              <Route path="/snippetView/:tagName" children={<SnippetView />} />
              {/* Uncomment these as you continue to build them one by one. */}
              {/* <Route path="/sent" component={Sent} /> */}
              {/* <Route path="/trending" component={Trending} /> */}
              {/* <Route path="/topics" component={Topics} /> */}
              {/* <Route path="/people" component={People} /> */}
            </Switch>
          </div>
        </div >
      </div >
    </div >
  );
}

export default App;
