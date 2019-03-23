import React from "react"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import Layout from "./Layout"

const Admin = () => (
  <ThemeProvider theme={skyward}>
  <Layout title={"Admin"} />
  </ThemeProvider>
)

export default Admin
