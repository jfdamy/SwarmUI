import React from 'react';
import { Route, IndexRoute } from 'react-router';

import App from './App';
import ProjectPage from '../../pages/project/page';
import CreateProjectPage from '../../pages/create/page';
import EditProjectPage from '../../pages/edit/page';
import HomePage from '../../pages/home/page';


export default (
  <Route path="/" component={App}>
    <IndexRoute component={HomePage} />
    <Route path="project" component={HomePage} />
    <Route path="/project/:projectId" component={ProjectPage} />
    <Route path="/create/project" component={CreateProjectPage} />
    <Route path="/edit/:projectId" component={EditProjectPage} />
  </Route>
);
