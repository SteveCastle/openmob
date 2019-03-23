import React from "react"
import PropTypes from "prop-types"
import { Link } from "gatsby"

import App from "@openmob/bluebird/src/components/layout/App"
import SideBar from "@openmob/bluebird/src/components/sidebar/SideBar"
import SideBarHeader from "@openmob/bluebird/src/components/sidebar/SideBarHeader"
import Menu from "@openmob/bluebird/src/components/menu/Menu"
import MenuItem from "@openmob/bluebird/src/components/menu/MenuItem"

const Layout = ({ children, title, id, summary }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">
          Open Mob
        </Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem>
          <Link to="/app/admin">Advanced Admin</Link>
        </MenuItem>
      </Menu>
    </SideBar>
  </App>
)

Layout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Layout
