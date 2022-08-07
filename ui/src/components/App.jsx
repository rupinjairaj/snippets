import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Publish from './Publish';

import Navbar from './Navbar';
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
      <div style={{ marginBottom: "2px" }}></div>
      <div style={{ marginLeft: "10px" }}>
        <div style={{ marginBottom: "12px" }} className="row">
          <Navbar />
        </div>
        <div className="row">
          <div className="col-1">
          </div >
          <div className="col-10">
            <Switch>
              <Route path="/" component={Tags} exact />
              <Route path="/tags" component={Tags} />
              <Route path="/snippetView/:tagName" children={<SnippetView />} />
              <Route path="/publish" component={Publish} />
              {/* Uncomment these as you continue to build them one by one. */}
              {/* <Route path="/trending" component={Trending} /> */}
              {/* <Route path="/topics" component={Topics} /> */}
              {/* <Route path="/people" component={People} /> */}
            </Switch>
          </div>
          <div style={{ marginTop: "12px" }} className="row">
          </div>
        </div >
      </div >
    </div >
  );
}

export default App;
