import { Link } from "gatsby"
import PropTypes from "prop-types"
import React from "react"
import PageHeader from "@openmob/bluebird/src/components/type/Header"
import SubHeader from "@openmob/bluebird/src/components/type/SubHeader"

const Header = ({ siteTitle, id, summary }) => (
  <header
    style={{
      background: `rebeccapurple`,
      marginBottom: `1.45rem`,
    }}
  >
    <div
      style={{
        margin: `0 auto`,
        maxWidth: 960,
        padding: `1.45rem 1.0875rem`,
      }}
    >
      <PageHeader dark>
        <Link
          to="/"
          style={{
            color: `white`,
            textDecoration: `none`,
          }}
        >
          {siteTitle}
        </Link>
      </PageHeader>
      <Link
        to="/app"
        style={{
          position: "absolute",
          top: "0",
          right: "0",
          backgroundColor: "white",
          padding: "4px",
          margin: "4px",
          borderRadius: "4px"
        }}
      >
        Admin
      </Link>
      <SubHeader dark>{id}</SubHeader>
      <SubHeader dark>{summary}</SubHeader>
    </div>
  </header>
)

Header.propTypes = {
  siteTitle: PropTypes.string,
}

Header.defaultProps = {
  siteTitle: ``,
}

export default Header
