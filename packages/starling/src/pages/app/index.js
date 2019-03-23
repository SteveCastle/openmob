import React from "react"
import { Router } from "@reach/router"
import Admin from '../../components/admin/Admin'
import Home from '../../components/admin/Home'

import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import './reset.css';

function App() {
  return (
    <ThemeProvider theme={skyward}>
    <Router>
      <Home path="app" />
      <Admin path="app/admin/*" />
    </Router>
    </ThemeProvider>
  )
}

export default App
