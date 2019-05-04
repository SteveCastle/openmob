import React from 'react'
import PropTypes from 'prop-types'
import { Link } from 'gatsby'

import App from '@openmob/bluebird/src/components/layout/App'
import SideBar from '@openmob/bluebird/src/components/sidebar/SideBar'
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel'
import SideBarHeader from '@openmob/bluebird/src/components/sidebar/SideBarHeader'
import Menu from '@openmob/bluebird/src/components/menu/Menu'
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem'
import MyCauses from './MyCauses'
const Home = () => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/app">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem>
          <Link to="/app/new">Create Cause</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/plugins">Plugins</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/settings">Settings</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/admin">CRUD</Link>
        </MenuItem>
      </Menu>
    </SideBar>
    <ContentPanel>
      <MyCauses />
    </ContentPanel>
  </App>
)

Home.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Home
