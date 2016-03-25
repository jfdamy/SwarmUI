import React from 'react';

import AppBar from 'material-ui/lib/app-bar';
import LeftNav from 'material-ui/lib/left-nav';
import MenuItem from 'material-ui/lib/menus/menu-item';

export default ({children}) => {
    return (
        <div id="container">
            <AppBar title="SwarmUI"/>
            {children}
        </div>
    );
}
