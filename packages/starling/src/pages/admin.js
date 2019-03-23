import React from "react"
import { Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import Layout from "../components/layout"
import SEO from "../components/seo"

const Admin = () => (
  <ThemeProvider theme={skyward}>
  <Layout title={"Admin"}>
    <SEO title="Admin" />
    <h1>Admin Page</h1>
    <p>Welcome to the Admin Page</p>
    <Link to="/">Go back to the homepage</Link>
  </Layout>
  </ThemeProvider>
)

export default Admin
