import React from "react"
import { Link } from "gatsby"

import Layout from "../components/layout"
import SEO from "../components/seo"

const Admin = () => (
  <Layout title={"Admin"}>
    <SEO title="Admin" />
    <h1>Admin Page</h1>
    <p>Welcome to the Admin Page</p>
    <Link to="/">Go back to the homepage</Link>
  </Layout>
)

export default Admin
