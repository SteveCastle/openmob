import React from 'react'
import PropTypes from 'prop-types'
import { Link } from 'gatsby'

import App from '@openmob/bluebird/src/components/layout/App'
import SideBar from '@openmob/bluebird/src/components/sidebar/SideBar'
import SideBarHeader from '@openmob/bluebird/src/components/sidebar/SideBarHeader'
import Menu from '@openmob/bluebird/src/components/menu/Menu'
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem'

const Layout = ({ children, title, id, summary }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem>
          <Link to="/app/pages">Pages</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/contacts">Contacts</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/contacts">Volunteer Coordination</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/shop">Fundraising</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/events">Events</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/elections">Elections</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/boycotts">BoyCotts</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/elections">Map</Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/admin">CRUD</Link>
        </MenuItem>
      </Menu>
    </SideBar>
  </App>
)

Layout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Layout
